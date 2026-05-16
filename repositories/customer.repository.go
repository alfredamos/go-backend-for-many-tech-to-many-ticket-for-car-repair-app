package repositories

import "go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

type CustomerRepository interface {
	createCustomer(customerCreate *models.CustomerCreate)
}
