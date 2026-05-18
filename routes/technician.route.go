package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TechnicianRoute(router fiber.Router, ca *controllers.TechnicianControllerImpl) {
	//----> Technician protected routes.
	//techProtectedRoute := router.Use(middleware.VerifyTokenJwt)

	//----> Admin protected routes.
	adminProtectedRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)
	adminProtectedRoute.Post("/", ca.CreateTechnicianController)
	adminProtectedRoute.Delete("/:id", ca.DeleteTechnicianByIdController)
	adminProtectedRoute.Patch("/:id", ca.EditTechnicianByIdController)
	adminProtectedRoute.Get("/", ca.GetAllTechniciansController)
	adminProtectedRoute.Get("/:id", ca.GetTechnicianByIdController)
	adminProtectedRoute.Get("/by-user-id/:userId", ca.GetTechnicianByUserIdController)
	adminProtectedRoute.Get("/by-specialty/:specialty", ca.GetTechniciansBySpecialtyController)

}
