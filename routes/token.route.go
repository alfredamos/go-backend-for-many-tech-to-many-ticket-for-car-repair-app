package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TokenRoute(router fiber.Router, tokenController controllers.TokenControllerInt) {
	//----> Delete invalid tokens by user id.
	tokenAuthenticated := router.Use(middleware.VerifyTokenJwt)
	tokenAuthenticated.Delete("/delete-by-user-id/:userId", tokenController.DeleteInvalidTokensByUserIdController)

	//----> Delete all invalid tokens, admin only.
	tokenAdminAuthorized := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	tokenAdminAuthorized.Delete("/all/delete-all", tokenController.DeleteAllInvalidTokensController)
}
