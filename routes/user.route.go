package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func UserRoute(router fiber.Router, userController controllers.UserControllerInt) {
	//----> User protected routes.
	userProtecRoute := router.Use(middleware.VerifyTokenJwt)
	userProtecRoute.Get("/by-email/:email", userController.GetUserByEmailController)

	//----> Owner protected routes.
	userOwnerRoute := router.Use(middleware.VerifyTokenJwt, middleware.OwnerByIdOrAdmin)
	userOwnerRoute.Delete("/:id", userController.DeleteUserByIdController)
	userOwnerRoute.Get("/:id", userController.GetUserByIdController)

	//----> Admin protected routes.
	userAdminRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	userAdminRoute.Get("/", userController.GetAllUsersController)

}
