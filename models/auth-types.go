package models

import "time"

type Role string

const (
	AdminRole Role = "Admin"
	UserRole  Role = "User"
)

type Gender string

const (
	Female Gender = "Female"
	Male   Gender = "Male"
)

type TokenParam struct {
	TokenName      string
	TokenPath      string
	TokenExpiresIn time.Time
}

type TokenJwt struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  Role   `json:"role"`
}

type UserDto struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Phone  string `json:"phone"`
	Email  string `json:"email"`
	Role   Role   `json:"role"`
	Gender Gender `json:"gender"`
	Image  string `json:"image"`
}

type Session struct {
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Email       string `json:"email"`
	Role        string `json:"role"`
	AccessToken string `json:"accessToken"`
	IsLoggedIn  bool   `json:"isLoggedIn"`
	IsAdmin     bool   `json:"isAdmin"`
}

type ResponseMessage struct {
	Message       string `json:"message"`
	StatusMessage string `json:status`
	StatusCode    int    `json:"statusCode"`
}

type ChangeUserPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
}

type ChangeUserRoleRequest struct {
	Email string `json:"email"`
}

type ChangeUserRoleResponse struct {
	Email string `json:"email"`
}

type LoginUserRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

type EditUserProfileRequest struct {
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Phone    string `validate:"required" json:"phone"`
	Email    string `validate:"required,email" json:"email"`
	Role     Role   `json:"role" default:"User"`
	Gender   Gender `validate:"required,oneof=Female Male" json:"gender"`
	Image    string `validate:"required" json:"image"`
}

type SignupUserRequest struct {
	Name            string `validate:"required" json:"name"`
	Email           string `validate:"required,email" json:"email"`
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirmPassword"`
	Phone           string `validate:"required" json:"phone"`
	Role            Role   `json:"role"`
	Gender          Gender `validate:"required,oneof=Female Male" json:"gender"`
	Image           string `validate:"required" json:"image"`
}
