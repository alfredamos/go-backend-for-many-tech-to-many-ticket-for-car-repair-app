package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"time"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func passwordNotMatch(passwordOne, passwordTwo string) bool {
	return passwordOne != passwordTwo
}

func passwordNotValid(rawPassword, encodedPassword string) error {
	//----> Compare the password with the hashed password
	if err := bcrypt.CompareHashAndPassword([]byte(encodedPassword), []byte(rawPassword)); err != nil {
		return errors.New("invalid password")
	}

	//----> No error
	return nil

}

func hashPassword(rawPassword string) (string, error) {
	//----> Hash password.
	passwd, err := bcrypt.GenerateFromPassword([]byte(rawPassword), 12)

	//----> Check for error.
	if err != nil {
		return "", errors.New(err.Error())
	}

	//----> Send back response
	return string(passwd), nil

}

func getUserByEmail(u UserAuthRepoImpl, email string) (models.User, error) {
	user := new(models.User)
	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return models.User{}, errors.New(err.Error())
	}

	//----> send back response
	return *user, nil
}

func generateTokensAndCookies(tokenJwt middleware.TokenJwt, ctx *fiber.Ctx, u UserAuthRepoImpl) (middleware.Session, error) {
	//----> Initialize token object.
	token := models.Token{}
	token.UserID = tokenJwt.UserId

	//----> Revoke all valid tokens for the user.
	tokenRepository := TokenRepositoryImpl{DB: u.DB}
	if err := tokenRepository.RevokeAllValidTokensByUserId(tokenJwt.UserId); err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Generate access-token and store it in cookies.
	accessToken, err := middleware.GenerateAccessToken(tokenJwt.Name, tokenJwt.Email, tokenJwt.UserId, tokenJwt.Role)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Set cookie.
	accessTokenParam := accessTokenParam()
	middleware.SetCookie(ctx, accessTokenParam.TokenPath, accessTokenParam.TokenName, accessToken, accessTokenParam.TokenExpiresIn)
	token.AccessToken = accessToken

	//----> Generate refresh-token and store it in cookies.
	refreshToken, err := middleware.GenerateRefreshToken(tokenJwt.Name, tokenJwt.Email, tokenJwt.UserId, tokenJwt.Role)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Set cookie.
	refreshTokenParam := refreshTokenParam()
	middleware.SetCookie(ctx, refreshTokenParam.TokenPath, refreshTokenParam.TokenName, refreshToken, refreshTokenParam.TokenExpiresIn)
	token.RefreshToken = refreshToken

	//----> Create token object.
	tokenObject := MakeNewToken(&token)
	if err := tokenRepository.CreateToken(tokenObject); err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Send back response.
	return makeSession(tokenJwt, accessToken), nil
}

func fromEditUserProfileToUser(editUserProfile *models.EditUserProfileRequest, user *models.User) *models.User {
	return &models.User{
		Name:     editUserProfile.Name,
		Email:    user.Email,
		Phone:    editUserProfile.Phone,
		Gender:   editUserProfile.Gender,
		Image:    editUserProfile.Image,
		Role:     user.Role,
		Password: user.Password,
		Type:     user.Type,
	}
}

func fromSignupUserToUser(signupUser *models.SignupUserRequest) *models.User {
	return &models.User{
		Name:     signupUser.Name,
		Email:    signupUser.Email,
		Phone:    signupUser.Phone,
		Gender:   signupUser.Gender,
		Image:    signupUser.Image,
		Password: signupUser.Password,
		Type:     signupUser.Type,
		Role:     models.UserRole,
	}
}

func makeSession(tokenJwt middleware.TokenJwt, accessToken string) middleware.Session {
	return middleware.Session{
		Email:       tokenJwt.Email,
		Name:        tokenJwt.Name,
		IsLoggedIn:  true,
		UserId:      tokenJwt.UserId,
		Role:        tokenJwt.Role,
		AccessToken: accessToken,
		IsAdmin:     tokenJwt.Role == string(models.AdminRole),
	}
}

func makeTokenJwt(user *models.User) middleware.TokenJwt {
	return middleware.TokenJwt{
		Email:  user.Email,
		Name:   user.Name,
		Role:   string(user.Role),
		UserId: user.ID,
	}
}

func accessTokenParam() TokenParam {
	accessTokenPath := "/"
	accessTokenName := "accessToken"
	accessTokenExpiresIn := time.Now().Add(1 * time.Hour)

	return TokenParam{TokenName: accessTokenName, TokenPath: accessTokenPath, TokenExpiresIn: accessTokenExpiresIn}
}

func refreshTokenParam() TokenParam {
	refreshTokenPath := "/api/auth/refresh"
	refreshTokenName := "refreshToken"
	refreshTokenExpiresIn := time.Now().Add(7 * 24 * time.Hour)

	return TokenParam{TokenName: refreshTokenName, TokenPath: refreshTokenPath, TokenExpiresIn: refreshTokenExpiresIn}
}

func initialUserSession() middleware.Session {
	return middleware.Session{
		UserId:      "",
		Name:        "",
		Email:       "",
		Role:        "",
		IsAdmin:     false,
		IsLoggedIn:  false,
		AccessToken: "",
	}
}
