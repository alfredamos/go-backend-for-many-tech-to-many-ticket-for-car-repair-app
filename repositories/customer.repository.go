package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type CustomerRepository interface {
	ChangeCustomerStatusById(id string) (ResponseMessage, error)
	CreateCustomer(customerCreate *models.CustomerCreate) (ResponseMessage, error)
	DeleteCustomerById(id string) (ResponseMessage, error)
	EditCustomerById(id string, customerEdit *models.CustomerEdit) (ResponseMessage, error)
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

func (c *CustomerRepositoryImpl) ChangeCustomerStatusById(id string) (ResponseMessage, error) {
	//----> fetch the customer with the given id.
	customer, err := getOneCustomerById(id, c)

	//----> check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> change the status of the customer.
	customer.Active = !customer.Active
	if customer.Active {
		customer.Status = models.Valid
	} else if !customer.Active {
		customer.Status = models.Invalid
	}

	//----> update the customer.
	if err := c.DB.Model(&models.Customer{}).Where("id = ?", id).Updates(customer).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> send back response.
	return NewResponseMessage("Status changed successfully", 200, "Success"), nil

}

func (c *CustomerRepositoryImpl) CreateCustomer(customerCreate *models.CustomerCreate) (ResponseMessage, error) {
	//----> Validate input.
	if err := models.ValidateCustomerCreate(customerCreate); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize customer.
	customer := models.Customer{
		Address: customerCreate.Address,
		Notes:   customerCreate.Notes,
		UserID:  customerCreate.UserID,
		Active:  customerCreate.Active,
		Status:  customerCreate.Status,
	}

	//----> Create customer.
	if err := c.DB.Create(&customer).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Send back response.
	return NewResponseMessage("Customer created successfully", 201, "Success"), nil
}

func (c *CustomerRepositoryImpl) DeleteCustomerById(id string) (ResponseMessage, error) {
	//----> Check for existence of customer.
	customer, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Delete customer.
	if err := c.DB.Delete(customer).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Send back response.
	return NewResponseMessage("Customer deleted successfully", 200, "Success"), nil
}

func (c *CustomerRepositoryImpl) EditCustomerById(id string, customerEdit *models.CustomerEdit) (ResponseMessage, error) {
	//----> Check for existence of customer.
	_, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Validate input.
	if err := models.ValidateCustomerEdit(customerEdit); err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Initialize and populate customer.
	customer := models.Customer{
		ID:      id,
		Address: customerEdit.Address,
		Notes:   customerEdit.Notes,
		UserID:  customerEdit.UserID,
		Active:  customerEdit.Active,
		Status:  customerEdit.Status,
	}

	//----> Update customer.
	if err := c.DB.Updates(&customer).Error; err != nil {
		return ResponseMessage{}, errors.New(err.Error())
	}

	//----> Send back response.
	return NewResponseMessage("Customer updated successfully", 200, "Success"), nil
}

func (c *CustomerRepositoryImpl) GetCustomerById(id string) (CustomerResponse, error) {
	//----> Fetch user with the giving id.
	customer, err := getOneCustomerById(id, c)

	//----> Check for error.
	if err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(*customer), nil
}

func (c *CustomerRepositoryImpl) GetActiveCustomers() ([]CustomerResponse, error) {
	//----> Fetch all active customers.
	var customers []models.Customer
	if err := c.DB.Preload("User").Find(&customers, "active = ?", true).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(customers), nil
}

func (c *CustomerRepositoryImpl) GetAllCustomers() ([]CustomerResponse, error) {
	//----> Fetch all customers.
	var customers []models.Customer
	if err := c.DB.Preload("User").Find(&customers).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(customers), nil
}

func (c *CustomerRepositoryImpl) GetCustomerByUserId(userId string) (CustomerResponse, error) {
	//----> Initialize customer.
	customer := models.Customer{}

	//----> Fetch customer by user id.
	if err := c.DB.Preload("User").First(&customer, "user_id = ?", userId).Error; err != nil {
		return CustomerResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponse(customer), nil
}

func (c *CustomerRepositoryImpl) GetInactiveCustomers() ([]CustomerResponse, error) {
	//----> Fetch all inactive customers.
	customers := new([]models.Customer)
	if err := c.DB.Preload("User").Where("active = ? AND status = ?", false, "Invalid").Find(&customers).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toCustomerResponseList(*customers), nil
}
