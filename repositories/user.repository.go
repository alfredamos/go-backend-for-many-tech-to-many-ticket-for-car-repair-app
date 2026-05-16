package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	DeleteUserById(userID string) (*UserResponse, error)
	GetUserById(userID string) (*UserResponse, error)
	GetUserByEmail(email string) (*UserResponse, error)
	GetAllUsers() ([]*UserResponse, error)
}

type UserRepositoryImpl struct {
	DB *gorm.DB
}

func NewUserRepositoryImpl(DB *gorm.DB) *UserRepositoryImpl {
	return &UserRepositoryImpl{DB: DB}
}

func (u *UserRepositoryImpl) DeleteUserById(userID string) (*UserResponse, error) {
	//----> Check if user exists.
	user, err := getOneUserById(userID, u)

	//----> Check for error.
	if err != nil {
		return nil, err
	}

	//----> Delete user.
	if err = u.DB.Delete(&user).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToUserResponse(user), nil
}

func (u *UserRepositoryImpl) GetUserById(userID string) (*UserResponse, error) {
	//----> Fetch user by id.
	user, err := getOneUserById(userID, u)

	//----> Check for error.
	if err != nil {
		return nil, err
	}

	//----> Send back response.
	return ToUserResponse(user), nil
}

func (u *UserRepositoryImpl) GetUserByEmail(email string) (*UserResponse, error) {
	//----> Fetch user by email.
	user, err := getOneUserByEmail(email, u)

	//----> Check for error.
	if err != nil {
		return nil, err
	}

	//----> Send back response.
	return ToUserResponse(user), nil
}

func (u *UserRepositoryImpl) GetAllUsers() ([]*UserResponse, error) {
	//----> Fetch all users.
	var users []*models.User
	if err := u.DB.Find(&users).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return ToUsers(users), nil
}
