package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func TicketRoute(router fiber.Router, controller *controllers.TicketControllerImpl) {
	//----> Ticket routes.
	adminTicketRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)

	adminTicketRoute.Post("/", controller.CreateTicketController)
	adminTicketRoute.Delete("/:id", controller.DeleteTicketByIdController)
	adminTicketRoute.Patch("/:id", controller.EditTicketByIdController)
	adminTicketRoute.Get("/", controller.GetAllTicketsController)
	adminTicketRoute.Get("/:id", controller.GetTicketByIdController)
	adminTicketRoute.Get("/by-customer-id/:customerId", controller.GetTicketsByCustomerIdController)

}
