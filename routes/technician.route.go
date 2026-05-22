package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TechnicianRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize technician repository.
	technicianRepo := repositories.NewTechnicianRepositoryImpl(DB)

	//----> Initialize technician service.
	technicianService := services.NewTechnicianServiceImpl(technicianRepo)

	//----> Initialize technician controller.
	technicianController := controllers.NewTechnicianControllerImpl(technicianService)

	//----> Owner protected routes.
	technicianOwnerRoute := router.Use(middleware.VerifyTokenJwt, middleware.OwnerByUserIdOrAdmin)
	technicianOwnerRoute.Get("/by-user-id/:userId", technicianController.GetTechnicianByUserIdController)

	//----> Admin protected routes.
	adminProtectedRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	adminProtectedRoute.Post("/", technicianController.CreateTechnicianController)
	adminProtectedRoute.Delete("/:id", technicianController.DeleteTechnicianByIdController)
	adminProtectedRoute.Patch("/:id", technicianController.EditTechnicianByIdController)
	adminProtectedRoute.Get("/", technicianController.GetAllTechniciansController)
	adminProtectedRoute.Get("/:id", technicianController.GetTechnicianByIdController)
	adminProtectedRoute.Get("/by-specialty/:specialty", technicianController.GetTechniciansBySpecialtyController)

}
