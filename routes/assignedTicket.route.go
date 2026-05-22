package routes

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/middleware"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func AssignedTicketRoute(router fiber.Router, DB *gorm.DB) {

	//----> Initialize assigned-ticket repository.
	assignedTicketRepo := repositories.NewAssignedTicketRepositoryImpl(DB)

	//----> Initialize assigned-ticket service.
	assignedTicketService := services.NewAssignedTicketServiceImpl(assignedTicketRepo)

	//----> Initialize assigned-ticket controller.
	assignedTickController := controllers.NewAssignedTicketControllerImpl(assignedTicketService)

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
