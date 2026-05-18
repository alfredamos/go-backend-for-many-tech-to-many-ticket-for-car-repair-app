package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	ChangeCustomerStatusById(id string) (CustomerResponse, error)
	CreateCustomer(customerCreate *models.CustomerCreate) (CustomerResponse, error)
	DeleteCustomerById(id string) (CustomerResponse, error)
	EditCustomerById(id string, customerEdit *models.CustomerEdit) (CustomerResponse, error)
	GetCustomerById(id string) (CustomerResponse, error)
	GetActiveCustomers() ([]CustomerResponse, error)
	GetAllCustomers() ([]CustomerResponse, error)
	GetCustomerByUserId(userId string) (CustomerResponse, error)
	GetInactiveCustomers() ([]CustomerResponse, error)
}

type CustomerRepositoryImpl struct {
	DB *gorm.DB
}

func NewCustomerRepositoryImpl(DB *gorm.DB) *CustomerRepositoryImpl {
	return &CustomerRepositoryImpl{DB: DB}
}

func (c *CustomerRepositoryImpl) ChangeCustomerStatusById(id string) (CustomerResponse, error) {
	//----> fetch the customer with the given id.
	customer, err := getOneCustomerById(id, c)

	//----> check for error.
	if err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> change the status of the customer.
	customer.Active = !customer.Active
	if customer.Active {
		customer.Status = models.Valid
	} else {
		customer.Status = models.Invalid
	}

	//----> update the customer.
	if err := c.DB.Model(&customer).Where("id = ?", customer.ID).Updates(customer).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> initialize user
	user := &models.User{}

	//----> preload user.
	if err := c.DB.First(&user, "id = ?", customer.UserID).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> send back response.
	return toCustomerResponse(*customer, user), nil

}

func (c *CustomerRepositoryImpl) CreateCustomer(customerCreate *models.CustomerCreate) (CustomerResponse, error) {
	//----> Initialize customer.
	customer := &models.Customer{
		Address: customerCreate.Address,
		Notes:   customerCreate.Notes,
		UserID:  customerCreate.UserID,
		Active:  customerCreate.Active,
		Status:  customerCreate.Status,
	}

	//----> Validate input.
	if err := models.ValidateCustomerCreate(customerCreate); err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Create customer.
	if err := c.DB.Create(customer).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := c.DB.First(&user, "id = ?", customer.UserID).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer, user), nil

}

func (c *CustomerRepositoryImpl) DeleteCustomerById(id string) (CustomerResponse, error) {
	//----> Check for existence of customer.
	customer, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Delete customer.
	if err := c.DB.Delete(customer).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := c.DB.First(&user, "id = ?", customer.UserID).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer, user), nil
}

func (c *CustomerRepositoryImpl) EditCustomerById(id string, customerEdit *models.CustomerEdit) (CustomerResponse, error) {
	//----> Check for existence of customer.
	customer, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Validate input.
	if err := models.ValidateCustomerEdit(customerEdit); err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize and populate customer.
	customerToEdit := &models.Customer{
		ID:      customer.ID,
		Address: customerEdit.Address,
		Notes:   customerEdit.Notes,
		UserID:  customerEdit.UserID,
		Active:  customerEdit.Active,
		Status:  customerEdit.Status,
	}

	//----> Update customer.
	if err := c.DB.Model(customerEdit).Where("id = ?", id).Updates(&customerToEdit).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := c.DB.First(&user, "id = ?", customer.UserID).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer, user), nil
}

func (c *CustomerRepositoryImpl) GetCustomerById(id string) (CustomerResponse, error) {
	//----> Fetch user with the giving id.
	customer, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := c.DB.First(&user, "id = ?", customer.UserID).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer, user), nil
}

func (c *CustomerRepositoryImpl) GetActiveCustomers() ([]CustomerResponse, error) {
	//----> Fetch all active customers.
	var customers []models.Customer
	if err := c.DB.Find(&customers, "active = ?", true).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(customers, c), nil
}

func (c *CustomerRepositoryImpl) GetAllCustomers() ([]CustomerResponse, error) {
	//----> Fetch all customers.
	var customers []models.Customer
	if err := c.DB.Find(&customers).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(customers, c), nil
}

func (c *CustomerRepositoryImpl) GetCustomerByUserId(userId string) (CustomerResponse, error) {
	//----> Initialize customer.
	customer := &models.Customer{}

	//----> Fetch customer by user id.
	if err := c.DB.First(&customer, "user_id = ?", userId).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := c.DB.First(&user, "id = ?", userId).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer, user), nil
}

func (c *CustomerRepositoryImpl) GetInactiveCustomers() ([]CustomerResponse, error) {
	//----> Fetch all active customers.
	var customers []models.Customer
	if err := c.DB.Find(&customers, "active = ?", false).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(customers, c), nil
}
