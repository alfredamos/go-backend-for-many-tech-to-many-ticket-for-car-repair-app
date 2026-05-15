package models

import (
	"errors"

	"github.com/go-playground/validator/v10"
)

func ValidateLoginUser(loginUser *LoginUserRequest) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	userLogin := &LoginUserModel{
		Email:    loginUser.Email,
		Password: loginUser.Password,
	}

	//----> Check for validation error
	if err := validate.Struct(userLogin); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func ValidateChangeUserPassword(changeUserPassword ChangeUserPasswordRequest) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	userChangePassword := &ChangePasswordModel{
		Email:           changeUserPassword.Email,
		Password:        changeUserPassword.Password,
		ConfirmPassword: changeUserPassword.ConfirmPassword,
		NewPassword:     changeUserPassword.NewPassword,
	}

	//----> Check for validation error.
	if err := validate.Struct(userChangePassword); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response.
	return nil
}

func ValidateEditUserProfile(editUserProfile EditUserProfileRequest) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	userEditProfile := &EditUserProfileModel{
		Name:     editUserProfile.Name,
		Email:    editUserProfile.Email,
		Phone:    editUserProfile.Phone,
		Gender:   editUserProfile.Gender,
		Role:     editUserProfile.Role,
		Password: editUserProfile.Password,
		Image:    editUserProfile.Image,
	}

	//----> Check for validation error.
	if err := validate.Struct(userEditProfile); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}

func ValidateChangeUserRole(changeUserRole ChangeUserRoleRequest) error {
	validate := validator.New() //----> Initialize.

	//----> Validate data.
	userChangeRole := &ChangeUserRoleModel{
		Email: changeUserRole.Email,
	}

	//----> Check for validation error.
	if err := validate.Struct(userChangeRole); err != nil {
		return errors.New(err.Error())
	}

	return nil
}

func ValidateSignupUser(signupUser SignupUserRequest) error {
	validate := validator.New() //----> Initialize

	//----> Validate data.
	userSignup := &SignupUserModel{
		Email:           signupUser.Email,
		Password:        signupUser.Password,
		Name:            signupUser.Name,
		Phone:           signupUser.Phone,
		Gender:          signupUser.Gender,
		Role:            signupUser.Role,
		Image:           signupUser.Image,
		ConfirmPassword: signupUser.ConfirmPassword,
	}

	//----> Check for validation error.
	if err := validate.Struct(userSignup); err != nil {
		return errors.New(err.Error())
	}

	//----> Send back response
	return nil
}
