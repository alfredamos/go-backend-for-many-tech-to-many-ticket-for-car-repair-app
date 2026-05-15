package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

type QueryConditions struct {
	TokenStatus models.Status
	UserID      string
	accessToken string
}

func MakeNewToken(token *models.Token) *models.TokenRequest {
	return &models.TokenRequest{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		Expired:      false,
		Revoked:      false,
		TokenType:    models.Bearer,
		TokenStatus:  models.Valid,
		UserID:       token.UserID,
	}
}

func sliceOfTokenIds(tokens []models.Token) []models.Token {
	//----> Make slice of token ids.
	tokenIds := make([]models.Token, 0)
	for _, token := range tokens {
		tokenID := models.Token{ID: token.ID}
		tokenIds = append(tokenIds, tokenID)
	}

	//----> Send back the result.
	return tokenIds
}

func revokeValidTokens(tokens []models.Token, tokenRepository *TokenRepositoryImpl) error {
	revokedTokens := make([]models.Token, 0) //----> Initialize slice of token.

	//----> Revoke all tokens.
	for _, token := range tokens {
		token.Expired = true               //----> Token has expired.
		token.Revoked = true               //----> Token is revoked.
		token.TokenStatus = models.Invalid //----> Token is invalid.

		revokedTokens = append(revokedTokens, token) //----> Update slice of revoked tokens.
	}

	//----> Save all revoked tokens in the database.
	if err := saveAll(revokedTokens, tokenRepository); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil
}

func findValidOrInvalidTokens(queryConditions QueryConditions, tokenRepository *TokenRepositoryImpl) ([]models.Token, error) {
	//----> Initialize tokens.
	tokens := new([]models.Token)
	//----> Retrieve valid or invalid tokens from database.
	if err := tokenRepository.DB.Where(&queryConditions).Find(&tokens).Error; err != nil {
		return []models.Token{}, errors.New(err.Error())
	}

	//----> Send back results.
	return *tokens, nil
}

func deleteInvalidTokens(queryConditions QueryConditions, tokenRepository *TokenRepositoryImpl) error {
	invalidTokens, err := findValidOrInvalidTokens(queryConditions, tokenRepository)
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Collect ids of invalid tokens in a slice.
	tokenIds := sliceOfTokenIds(invalidTokens)

	//----> Delete all invalid tokens from token database.
	if err := tokenRepository.DB.Unscoped().Delete(&tokenIds, queryConditions).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back result.
	return nil
}

func saveAll(tokens []models.Token, tokenRepository *TokenRepositoryImpl) error {
	tx := tokenRepository.DB.Begin()
	if tx.Error != nil {
		return errors.New("error at the onset of saving all tokens")
	}

	for _, token := range tokens {
		err := tx.Model(&models.Token{}).Where("id = ?", token.ID).Updates(token).Error

		//----> Check for error
		if err != nil {
			tx.Rollback()
			// Handle error
			return errors.New("error updating token")
		}
	}

	err := tx.Commit().Error
	if err != nil {
		// Handle error
		return errors.New("unable to commit token to database")
	}
	return nil
}

func fromTokenRequestToToken(request *models.TokenRequest) *models.Token {
	return &models.Token{
		AccessToken:  request.AccessToken,
		RefreshToken: request.RefreshToken,
		Expired:      request.Expired,
		Revoked:      request.Revoked,
		TokenStatus:  request.TokenStatus,
		TokenType:    request.TokenType,
		UserID:       request.UserID,
	}
}
