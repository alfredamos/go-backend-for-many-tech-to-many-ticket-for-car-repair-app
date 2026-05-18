package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func CustomerRoute(router fiber.Router, customerController controllers.CustomerControllerInt) {
	//----> Customer protected routes.
	customerProtecRoute := router.Use(middleware.VerifyTokenJwt)
	customerProtecRoute.Get("/:id", customerController.GetCustomerByIdController)

	//----> Admin protected routes.
	customerAdminRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)

	customerAdminRoute.Patch("/change-status/:id", customerController.ChangeCustomerStatusByIdController)
	customerAdminRoute.Post("/", customerController.CreateCustomerController)
	customerAdminRoute.Patch("/:id", customerController.EditCustomerByIdController)
	customerAdminRoute.Delete("/:id", customerController.DeleteCustomerByIdController)
	customerAdminRoute.Get("/", customerController.GetAllCustomersController)
	customerAdminRoute.Get("/all/active", customerController.GetActiveCustomersController)
	customerAdminRoute.Get("/all/inactive", customerController.GetInactiveCustomersController)
	customerAdminRoute.Get("/by-user-id/:userId", customerController.GetCustomerByUserIdController)
}
