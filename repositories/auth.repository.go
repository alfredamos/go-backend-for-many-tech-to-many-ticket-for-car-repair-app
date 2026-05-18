package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type UserAuthRepoInt interface {
	ChangeUserPassword(request models.ChangeUserPasswordRequest) error
	ChangeUserRole(ctx *fiber.Ctx, request models.ChangeUserRoleRequest) (*UserResponse, error)
	EditUserProfile(request models.EditUserProfileRequest) error
	GetCurrentUser(ctx *fiber.Ctx) (*UserResponse, error)
	GetUserSession(ctx *fiber.Ctx) (middleware.Session, error)
	LoginUser(ctx *fiber.Ctx, request models.LoginUserRequest) (middleware.Session, error)
	LogoutUser(ctx *fiber.Ctx) (middleware.Session, error)
	RefreshUserToken(ctx *fiber.Ctx) (middleware.Session, error)
	SignupUser(request models.SignupUserRequest) error
}

type UserAuthRepoImpl struct {
	DB *gorm.DB
}

func NewUserAuthRepoImpl(DB *gorm.DB) *UserAuthRepoImpl {
	return &UserAuthRepoImpl{DB: DB}
}

func (u UserAuthRepoImpl) ChangeUserPassword(request models.ChangeUserPasswordRequest) error {
	//----> Validate input.
	err := models.ValidateChangeUserPassword(request)

	//----> Check for error.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Check for password match.
	if passwordNotMatch(request.NewPassword, request.ConfirmPassword) {
		return errors.New("password does not match")
	}

	//----> Check for existing user.
	user, err := getUserByEmail(u, request.Email)

	//----> Check for error.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Check for password validity.
	if err := passwordNotValid(request.Password, user.Password); err != nil {
		return errors.New(err.Error())
	}

	//----> Hash new password.
	hashedPassword, err := hashPassword(request.NewPassword)

	//----> Check for error.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Update user password.
	err = u.DB.Model(&user).Update("password", hashedPassword).Error

	//----> Check for error.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil

}

func (u UserAuthRepoImpl) ChangeUserRole(ctx *fiber.Ctx, request models.ChangeUserRoleRequest) (*UserResponse, error) {
	//----> Validate input.
	if err := models.ValidateChangeUserRole(request); err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Get the user session.
	session, err := u.GetUserSession(ctx)

	//----> Check for error.
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Check for admin privilege.
	if !session.IsAdmin {
		return nil, errors.New("you are not allowed to perform this action")
	}

	//----> Check for existing user.
	user, err := getUserByEmail(u, request.Email)

	//----> Check for error.
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Change user role.
	if user.Role == models.AdminRole {
		user.Role = models.UserRole
	} else {
		user.Role = models.AdminRole
	}

	//----> Update user role.
	if err := u.DB.Save(user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToUserResponse(&user), nil
}

func (u UserAuthRepoImpl) EditUserProfile(request models.EditUserProfileRequest) error {
	//----> Validate input.
	if err := models.ValidateEditUserProfile(request); err != nil {
		return errors.New(err.Error())
	}

	//----> Check for existing user.
	user, err := getUserByEmail(u, request.Email)

	//----> Check for error.
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Check for password validity.
	if err := passwordNotValid(request.Password, user.Password); err != nil {
		return errors.New(err.Error())
	}

	//----> Map edit user profile to user model.
	userToEdit := fromEditUserProfileToUser(&request, &user)

	//----> Update user profile.
	userToEdit.ID = user.ID
	userToEdit.CreatedAt = user.CreatedAt
	if err := u.DB.Save(userToEdit).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func (u UserAuthRepoImpl) GetCurrentUser(ctx *fiber.Ctx) (*UserResponse, error) {
	//----> Get the user session.
	session, err := u.GetUserSession(ctx)

	//----> Check for error.
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Get the user.
	user, err := getUserByEmail(u, session.Email)

	//----> Check for error.
	if err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToUserResponse(&user), nil
}

func (u UserAuthRepoImpl) GetUserSession(ctx *fiber.Ctx) (middleware.Session, error) {
	//----> Get the access-token from cookie.
	accessTokenPar := accessTokenParam()
	accessToken := middleware.GetCookie(ctx, accessTokenPar.TokenName)

	//----> Validate the access-token.
	tokenJwt, err := middleware.ValidateToken(accessToken, ctx)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Make user session.
	return makeSession(tokenJwt, accessToken), nil
}

func (u UserAuthRepoImpl) LoginUser(ctx *fiber.Ctx, request models.LoginUserRequest) (middleware.Session, error) {
	//----> Validate login input.
	if err := models.ValidateLoginUser(&request); err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Check for existence of user.
	user, err := getUserByEmail(u, request.Email)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Check for password validity.
	if err := passwordNotValid(request.Password, user.Password); err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> make token-jwt.
	tokenJwt := makeTokenJwt(&user)

	//----> Generate tokens and store them in cookies.
	return generateTokensAndCookies(tokenJwt, ctx, u)

}

func (u UserAuthRepoImpl) LogoutUser(ctx *fiber.Ctx) (middleware.Session, error) {
	//----> Initialize token repository.
	tokenRepository := TokenRepositoryImpl{DB: u.DB}

	//----> delete all cookies.
	accessTokenPar := accessTokenParam()
	middleware.DeleteCookie(ctx, accessTokenPar.TokenName, accessTokenPar.TokenPath)
	RefreshTokenPar := refreshTokenParam()
	middleware.DeleteCookie(ctx, RefreshTokenPar.TokenName, RefreshTokenPar.TokenPath)

	//----> Get the user session.
	session, err := u.GetUserSession(ctx)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Revoke valid tokens.
	if err := tokenRepository.RevokeAllValidTokensByUserId(session.UserId); err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Send back response.
	return initialUserSession(), nil

}

func (u UserAuthRepoImpl) RefreshUserToken(ctx *fiber.Ctx) (middleware.Session, error) {
	//----> Get the refresh-token from cookie.
	refreshTokenPar := refreshTokenParam()
	refreshToken := middleware.GetCookie(ctx, refreshTokenPar.TokenName)

	//----> Validate token.
	tokenJwt, err := middleware.ValidateToken(refreshToken, ctx)

	//----> Check for error.
	if err != nil {
		return middleware.Session{}, errors.New(err.Error())
	}

	//----> Generate tokens and store them in cookies.
	return makeSession(tokenJwt, refreshToken), nil
}

func (u UserAuthRepoImpl) SignupUser(request models.SignupUserRequest) error {
	//----> Validate input.
	if err := models.ValidateSignupUser(request); err != nil {
		return errors.New(err.Error())
	}

	//----> Check for password match.
	if passwordNotMatch(request.Password, request.ConfirmPassword) {
		return errors.New("password does not match")
	}

	//----> Check for existing user.
	user := new(models.User)
	if err := u.DB.Where("email = ?", request.Email).First(&user).Error; err == nil {
		return errors.New("user already exists")
	}

	//----> Hash password.
	encodedPassword, err := hashPassword(request.Password)
	if err != nil {
		return errors.New(err.Error())
	}

	//----> Map signup user to user model.
	request.Password = encodedPassword
	user = fromSignupUserToUser(&request)

	//----> Save user.
	if err := u.DB.Create(user).Error; err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil

}
