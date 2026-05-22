package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func TicketRoute(router fiber.Router, DB *gorm.DB) {
	//----> Initialize ticket repository.
	ticketRepo := repositories.NewTicketRepositoryImpl(DB)

	//----> Initialize ticket service.
	ticketService := services.NewTicketServiceImpl(ticketRepo)

	//----> Initialize ticket controller.
	ticketController := controllers.NewTicketControllerImpl(ticketService)

	//----> Ticket routes.
	adminTicketRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)

	adminTicketRoute.Post("/", ticketController.CreateTicketController)
	adminTicketRoute.Delete("/:id", ticketController.DeleteTicketByIdController)
	adminTicketRoute.Patch("/:id", ticketController.EditTicketByIdController)
	adminTicketRoute.Get("/", ticketController.GetAllTicketsController)
	adminTicketRoute.Get("/:id", ticketController.GetTicketByIdController)
	adminTicketRoute.Get("/by-customer-id/:customerId", ticketController.GetTicketsByCustomerIdController)

}
