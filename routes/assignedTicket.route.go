package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"

	"github.com/gofiber/fiber/v2"
)

func AssignedTicketRoute(router fiber.Router, assignedTickController controllers.AssignedTicketControllerInt) {
	//----> Admin routes.
	adminAssignedTicketRoute := router.Use(middleware.VerifyTokenJwt, middleware.AdminRolePermission)

	adminAssignedTicketRoute.Post("/", assignedTickController.CreateAssignedTicket)
	adminAssignedTicketRoute.Patch("/change-status/:techId/:ticketId", assignedTickController.ChangeAssignedTicketStatus)
	adminAssignedTicketRoute.Get("/completed", assignedTickController.GetCompletedAssignedTickets)
	adminAssignedTicketRoute.Get("/incompleted", assignedTickController.GetInCompletedAssignedTickets)
	adminAssignedTicketRoute.Get("/by-tech-id/:techId", assignedTickController.GetAssignedTicketsByTechId)
	adminAssignedTicketRoute.Get("/by-ticket-id/:ticketId", assignedTickController.GetAssignedTicketsByTicketId)
	adminAssignedTicketRoute.Delete("/:techId/:ticketId", assignedTickController.DeleteAssignedTicketById)
	adminAssignedTicketRoute.Patch("/:techId/:ticketId", assignedTickController.EditAssignedTicketById)
	adminAssignedTicketRoute.Get("/:techId/:ticketId", assignedTickController.GetAssignedTicketById)
	adminAssignedTicketRoute.Get("/", assignedTickController.GetAllAssignedTickets)

}
