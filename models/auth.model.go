package models

type ChangePasswordModel struct {
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required"`
	NewPassword     string `json:"newPassword" validate:"required"`
	ConfirmPassword string `json:"confirmPassword" validate:"required"`
}

type ChangeUserRoleModel struct {
	Email string `validate:"required,email"`
}

type EditUserProfileModel struct {
	Name     string `validate:"required" json:"name"`
	Password string `validate:"required" json:"password"`
	Phone    string `validate:"required" json:"phone"`
	Email    string `validate:"required,email" json:"email"`
	Role     Role   `json:"role" default:"User"`
	Gender   Gender `validate:"required,oneof=Female Male" json:"gender"`
	Image    string `validate:"required" json:"image"`
}

type LoginUserModel struct {
	Email    string `validate:"required,email"`
	Password string `validate:"required"`
}

type SignupUserModel struct {
	Name            string `validate:"required" json:"name"`
	Email           string `validate:"required,email"`
	Password        string `validate:"required" json:"password"`
	ConfirmPassword string `validate:"required" json:"confirmPassword"`
	Phone           string `validate:"required" json:"phone"`
	Role            Role   `json:"role"`
	Gender          Gender `validate:"required,oneof=Female Male" json:"gender"`
	Image           string `validate:"required" json:"image"`
}
