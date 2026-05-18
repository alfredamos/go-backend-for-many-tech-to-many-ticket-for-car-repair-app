package services

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/repositories"

	"github.com/gofiber/fiber/v2"
)

type CustomerServiceInt interface {
	ChangeCustomerStatusById(ctx *fiber.Ctx) error
	CreateCustomer(ctx *fiber.Ctx) error
	DeleteCustomerById(ctx *fiber.Ctx) error
	EditCustomerById(ctx *fiber.Ctx) error
	GetAllCustomers(ctx *fiber.Ctx) error
	GetCustomerById(ctx *fiber.Ctx) error
	GetActiveCustomers(ctx *fiber.Ctx) error
	GetInactiveCustomers(ctx *fiber.Ctx) error
	GetCustomerByUserId(ctx *fiber.Ctx) error
}

type CustomerServiceImpl struct {
	repo repositories.CustomerRepositoryImpl
}

func NewCustomerServiceImpl(repo repositories.CustomerRepositoryImpl) *CustomerServiceImpl {
	return &CustomerServiceImpl{repo: repo}
}

func (c *CustomerServiceImpl) ChangeCustomerStatusById(ctx *fiber.Ctx) error {
	//----> Get the customer id from payload.
	id := ctx.Params("id")

	//----> Change the customer status.
	customer, err := c.repo.ChangeCustomerStatusById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func (c *CustomerServiceImpl) CreateCustomer(ctx *fiber.Ctx) error {
	//----> Initialize customerCreate.
	customerCreate := &models.CustomerCreate{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&customerCreate); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Create customer.
	customer, err := c.repo.CreateCustomer(customerCreate)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusCreated).JSON(customer)

}

func (c *CustomerServiceImpl) DeleteCustomerById(ctx *fiber.Ctx) error {
	//----> Get the customer id from payload.
	id := ctx.Params("id")

	//----> Delete the customer with the giving id.
	customer, err := c.repo.DeleteCustomerById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func (c *CustomerServiceImpl) EditCustomerById(ctx *fiber.Ctx) error {
	//----> Get the customer id from payload.
	id := ctx.Params("id")

	//----> Initialize customerEdit.
	customerEdit := &models.CustomerEdit{}

	//----> Get the request payload.
	if err := ctx.BodyParser(&customerEdit); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(err.Error())
	}

	//----> Edit the customer with the giving id.
	customer, err := c.repo.EditCustomerById(id, customerEdit)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func (c *CustomerServiceImpl) GetCustomerById(ctx *fiber.Ctx) error {
	//----> Get the customer id from payload.
	id := ctx.Params("id")

	//----> Fetch the customer with the giving id.
	customer, err := c.repo.GetCustomerById(id)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func (c *CustomerServiceImpl) GetActiveCustomers(ctx *fiber.Ctx) error {
	//----> Fetch all active customers.
	customers, err := c.repo.GetActiveCustomers()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customers)
}

func (c *CustomerServiceImpl) GetAllCustomers(ctx *fiber.Ctx) error {
	//----> Fetch all customers.
	customers, err := c.repo.GetAllCustomers()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customers)
}

func (c *CustomerServiceImpl) GetCustomerByUserId(ctx *fiber.Ctx) error {
	//----> Get the user id from payload.
	userId := ctx.Params("userId")

	//----> Fetch the customer with the giving user id.
	customer, err := c.repo.GetCustomerByUserId(userId)

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customer)
}

func (c *CustomerServiceImpl) GetInactiveCustomers(ctx *fiber.Ctx) error {
	//----> Fetch all inactive customers.
	customers, err := c.repo.GetInactiveCustomers()

	//----> Check for error.
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(err.Error())
	}

	//----> Send back response.
	return ctx.Status(fiber.StatusOK).JSON(customers)
}
