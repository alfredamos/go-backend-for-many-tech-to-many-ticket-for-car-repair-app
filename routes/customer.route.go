package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func CustomerRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize customer repository.
	customerRepo := repositories.NewCustomerRepositoryImpl(DB)

	//----> Initialize customer service.
	customerService := services.NewCustomerServiceImpl(customerRepo)

	//----> Initialize customer controller.
	customerController := controllers.NewCustomerControllerImpl(customerService)

	//----> Customer protected routes.
	customerProtecRoute := router.Use(middleware.VerifyTokenJwt)
	customerProtecRoute.Get("/:id", customerController.GetCustomerByIdController)

	//----> Owner protected routes.
	customerOwnerRoute := router.Use(middleware.VerifyTokenJwt, middleware.OwnerByUserIdOrAdmin)
	customerOwnerRoute.Get("/by-user-id/:userId", customerController.GetCustomerByUserIdController)

	//----> Admin protected routes.
	customerAdminRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)

	customerAdminRoute.Patch("/change-status/:id", customerController.ChangeCustomerStatusByIdController)
	customerAdminRoute.Post("/", customerController.CreateCustomerController)
	customerAdminRoute.Patch("/:id", customerController.EditCustomerByIdController)
	customerAdminRoute.Delete("/:id", customerController.DeleteCustomerByIdController)
	customerAdminRoute.Get("/", customerController.GetAllCustomersController)
	customerAdminRoute.Get("/all/active", customerController.GetActiveCustomersController)
	customerAdminRoute.Get("/all/inactive", customerController.GetInactiveCustomersController)

}
