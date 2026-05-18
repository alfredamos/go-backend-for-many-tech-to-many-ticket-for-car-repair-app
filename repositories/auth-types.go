package repositories

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"time"
)

type ChangeUserPasswordRequest struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"oldPassword" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type ChangeUserRoleRequest struct {
	Email string `json:"email" validate:"required,email"`
}

type EditUserProfileRequest struct {
	Name     string        `validate:"required" json:"name"`
	Password string        `validate:"required" json:"password"`
	Phone    string        `validate:"required" json:"phone"`
	Email    string        `validate:"required,email" json:"email"`
	Role     models.Role   `json:"role" default:"User"`
	Gender   models.Gender `validate:"required,oneof=Female Male" json:"gender"`
	Image    string        `validate:"required" json:"image"`
}

type LoginUserRequest struct {
	Email    string `validate:"required,email" json:"email"`
	Password string `validate:"required" json:"password"`
}

type SignupUserRequest struct {
	Name            string          `validate:"required" json:"name"`
	Email           string          `validate:"required,email" json:"email"`
	Password        string          `validate:"required" json:"password"`
	ConfirmPassword string          `validate:"required" json:"confirmPassword"`
	Phone           string          `validate:"required" json:"phone"`
	Role            models.Role     `json:"role"`
	Gender          models.Gender   `validate:"required,oneof=Female Male" json:"gender"`
	Image           string          `validate:"required" json:"image"`
	Type            models.UserType `json:"type" validate:"required,oneof=Customer Technician" json:"type"`
}

type UserResponse struct {
	ID     string `json:"id"`
	Name   string `json:"name"`
	Email  string `json:"email"`
	Phone  string `json:"phone"`
	Role   string `json:"role"`
	Gender string `json:"gender"`
	Image  string `json:"image"`
	Type   models.UserType
}

type TokenParam struct {
	TokenName      string
	TokenPath      string
	TokenExpiresIn time.Time
}

func ToUserResponse(user *models.User) *UserResponse {
	return &UserResponse{
		ID:     user.ID,
		Name:   user.Name,
		Email:  user.Email,
		Phone:  user.Phone,
		Role:   string(user.Role),
		Gender: string(user.Gender),
		Image:  user.Image,
		Type:   user.Type,
	}
}

func ToUsers(users []*models.User) []*UserResponse {
	var userResponses []*UserResponse
	for _, user := range users {
		userResponses = append(userResponses, ToUserResponse(user))
	}
	return userResponses
}
