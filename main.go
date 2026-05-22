package main

import (
	"fmt"
	_ "fmt"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/controllers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/db"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/initializers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/routes"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

// TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>
func main() {
	//TIP <p>Press <shortcut actionId="ShowIntentionActions"/> when your caret is at the underlined text
	// to see how GoLand suggests fixing the warning.</p><p>Alternatively, if available, click the lightbulb to view possible fixes.</p>
	//----> Load environment variables
	if err := initializers.LoadEnvVariable(); err != nil {
		log.Fatal(err)
	}

	//----> Connect to database
	DB, err := db.ConnectDatabase()

	//----> Check for error
	if err != nil {
		log.Fatal(err)
	}

	//----> Initialize fiber app.
	app := fiber.New()

	//----> Custom CORS configuration
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:4200, http://localhost:5174", // Comma-separated list of origins
		AllowHeaders:     "Origin, Content-Type, Accept",                 // Allowed headers
		AllowMethods:     "GET, POST, HEAD, PUT, DELETE, PATCH",          // Allowed methods
		AllowCredentials: true,                                           // Set to true if you need to handle cookies/credentials
	}))

	//----> Auth routes.
	authRoutes := app.Group("/api/auth")
	routes.AuthRoute(authRoutes, DB)

	//----> Assigned-ticket routes.
	assignedTicketRoutes := app.Group("/api/assign-tickets")
	routes.AssignedTicketRoute(assignedTicketRoutes, DB)

	//----> Customer routes.
	CustomerRoutes := app.Group("/api/customers")
	routes.CustomerRoute(CustomerRoutes, DB)

	//----> Technician routes.
	TechnicianRoutes := app.Group("/api/technicians")
	routes.TechnicianRoute(TechnicianRoutes, DB)

	//----> Ticket routes.
	ticketRoutes := app.Group("/api/tickets")
	routes.TicketRoute(ticketRoutes, DB)

	//----> Token routes.
	tokenRoutes := app.Group("/api/tokens")
	routes.TokenRoute(tokenRoutes, DB)

	//----> Initialize user repository.
	userRepo := repositories.NewUserRepositoryImpl(DB)

	//----> Initialize user service.
	userService := services.NewUserServiceImpl(userRepo)

	//----> Initialize user controller.
	UserController := controllers.NewUserControllerImpl(userService)

	//----> User routes.
	userRoutes := app.Group("/api/users")
	routes.UserRoute(userRoutes, UserController)

	//----> Start server
	err = app.Listen(":5000")
	if err != nil {
		log.Fatal(err)

	}

	fmt.Println("Server is running on port 5000")

}
