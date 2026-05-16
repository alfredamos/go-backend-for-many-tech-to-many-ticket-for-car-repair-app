package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AuthRoute(router fiber.Router, authController controllers.AuthControllerInt) {
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
