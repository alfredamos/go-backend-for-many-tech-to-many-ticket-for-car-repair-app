package controllers

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/services"

	"github.com/gofiber/fiber/v2"
)

type CustomerControllerInt interface {
	ChangeCustomerStatusByIdController(ctx *fiber.Ctx) error
	CreateCustomerController(ctx *fiber.Ctx) error
	DeleteCustomerByIdController(ctx *fiber.Ctx) error
	EditCustomerByIdController(ctx *fiber.Ctx) error
	GetAllCustomersController(ctx *fiber.Ctx) error
	GetCustomerByIdController(ctx *fiber.Ctx) error
	GetActiveCustomersController(ctx *fiber.Ctx) error
	GetInactiveCustomersController(ctx *fiber.Ctx) error
	GetCustomerByUserIdController(ctx *fiber.Ctx) error
}

type CustomerControllerImpl struct {
	service services.CustomerServiceImpl
}

func NewCustomerControllerImpl(service services.CustomerServiceImpl) *CustomerControllerImpl {
	return &CustomerControllerImpl{service: service}
}

func (c *CustomerControllerImpl) ChangeCustomerStatusByIdController(ctx *fiber.Ctx) error {
	//----> Change customer status by id.
	return c.service.ChangeCustomerStatusById(ctx)
}

func (c *CustomerControllerImpl) CreateCustomerController(ctx *fiber.Ctx) error {
	//----> Create customer.
	return c.service.CreateCustomer(ctx)
}

func (c *CustomerControllerImpl) DeleteCustomerByIdController(ctx *fiber.Ctx) error {
	//----> Delete customer by id.
	return c.service.DeleteCustomerById(ctx)
}

func (c *CustomerControllerImpl) EditCustomerByIdController(ctx *fiber.Ctx) error {
	//----> Edit customer by id.
	return c.service.EditCustomerById(ctx)
}

func (c *CustomerControllerImpl) GetCustomerByIdController(ctx *fiber.Ctx) error {
	//----> Get customer by id.
	return c.service.GetCustomerById(ctx)
}

func (c *CustomerControllerImpl) GetActiveCustomersController(ctx *fiber.Ctx) error {
	//----> Get active customers.
	return c.service.GetActiveCustomers(ctx)
}

func (c *CustomerControllerImpl) GetAllCustomersController(ctx *fiber.Ctx) error {
	//----> Get all customers.
	return c.service.GetAllCustomers(ctx)
}

func (c *CustomerControllerImpl) GetCustomerByUserIdController(ctx *fiber.Ctx) error {
	//----> Get customer by user id.
	return c.service.GetCustomerByUserId(ctx)
}

func (c *CustomerControllerImpl) GetInactiveCustomersController(ctx *fiber.Ctx) error {
	//----> Get inactive customers.
	return c.service.GetInactiveCustomers(ctx)
}
