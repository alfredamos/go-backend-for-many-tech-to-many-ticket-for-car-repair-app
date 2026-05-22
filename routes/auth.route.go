package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AuthRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize auth repository
	authRepo := repositories.NewUserAuthRepoImpl(DB)

	//----> Initialize auth service.
	authService := services.NewAuthServiceImpl(authRepo)

	//----> Initialize auth controller.
	authController := controllers.NewAuthController(authService)

	//----> Auth public routes.
	router.Post("/login", authController.LoginUserController)          //----> Login user.
	router.Post("/logout", authController.LogoutUserController)        //----> Logout user.
	router.Post("/refresh", authController.RefreshUserTokenController) //----> Refresh user token.
	router.Post("/signup", authController.SignupUserController)

	//----> Auth protected routes.
	protectedRouter := router.Use(middleware.VerifyTokenJwt)
	protectedRouter.Patch("/change-password", authController.ChangeUserPasswordController)
	protectedRouter.Patch("/change-role", authController.ChangeUserRoleController)
	protectedRouter.Patch("/edit-profile", authController.EditUserProfileController)
	protectedRouter.Get("/me", authController.GetCurrentUserController)

}
