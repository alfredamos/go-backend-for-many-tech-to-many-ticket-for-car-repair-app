package main

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/db"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/initializers"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"log"
)

func main() {
	//----> Load environment variables
	if err := initializers.LoadEnvVariable(); err != nil {
		log.Fatal(err)
	}

	//----> Migrate the gorm models into mysql database.
	DB, err := db.ConnectDatabase()

	if err != nil {
		log.Fatal(err)
	}

	//----> Manually set up the join table for the Technician.Tickets association
	err = DB.SetupJoinTable(&models.Technician{}, "Tickets", &models.AssignedTicket{})
	if err != nil {
		log.Fatal(err)
	}

	//----> Also set up for the Ticket.Technicians association for bidirectional use
	err = DB.SetupJoinTable(&models.Ticket{}, "Technicians", &models.AssignedTicket{})
	if err != nil {
		log.Fatal(err)
	}

	//----> AutoMigrate all models, including the custom join table
	err = DB.AutoMigrate(&models.AssignedTicket{}, &models.Customer{}, &models.Technician{}, &models.Ticket{}, &models.Token{}, &models.User{})

	if err != nil {
		log.Fatal(err)
	}

}
