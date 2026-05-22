package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TokenRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize token repository.
	tokenRepo := repositories.NewTokenRepositoryImpl(DB)

	//----> Initialize token service.
	tokenService := services.NewTokenServiceImpl(tokenRepo)

	//----> Initialize token controller.
	tokenController := controllers.NewTokenControllerImpl(tokenService)

	//----> Delete invalid tokens by user id.
	tokenOwnershipRoute := router.Use(middleware.VerifyTokenJwt, middleware.OwnerByUserIdOrAdmin)
	tokenOwnershipRoute.Delete("/delete-by-user-id/:userId", tokenController.DeleteInvalidTokensByUserIdController)

	//----> Delete all invalid tokens, admin only.
	tokenAdminAuthorized := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	tokenAdminAuthorized.Delete("/all/delete-all", tokenController.DeleteAllInvalidTokensController)
}
