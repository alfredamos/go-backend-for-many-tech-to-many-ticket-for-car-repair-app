package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func UserRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize user repository.
	userRepo := repositories.NewUserRepositoryImpl(DB)

	//----> Initialize user service.
	userService := services.NewUserServiceImpl(userRepo)

	//----> Initialize user controller.
	userController := controllers.NewUserControllerImpl(userService)

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
