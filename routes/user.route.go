package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router, userController controllers.UserControllerInt) {
	//----> User protected routes.
	userProtecRoute := router.Use(middleware.VerifyTokenJwt)
	userProtecRoute.Delete("/:id", userController.DeleteUserByIdController)
	userProtecRoute.Get("/:id", userController.GetUserByIdController)
	userProtecRoute.Get("/by-email/:email", userController.GetUserByEmailController)

	//----> Admin protected routes.
	userAdminRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	userAdminRoute.Get("/", userController.GetAllUsersController)

}
