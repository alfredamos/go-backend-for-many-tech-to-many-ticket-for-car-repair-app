package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func getOneUserById(id string, u *UserRepositoryImpl) (*models.User, error) {
	//----> Initialize user.
	user := new(models.User)

	if err := u.DB.First(&user, "id = ?", id).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return user, nil
}

func getOneUserByEmail(email string, u *UserRepositoryImpl) (*models.User, error) {
	user := new(models.User)
	if err := u.DB.First(&user, "email = ?", email).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> send back response
	return user, nil
}
