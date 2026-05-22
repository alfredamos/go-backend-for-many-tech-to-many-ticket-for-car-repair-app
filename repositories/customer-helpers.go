package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toCustomerResponse(customer models.Customer) CustomerResponse {
	//----> Check for existence of user.
	if customer.User == nil {
		return CustomerResponse{}
	}

	//----> Send back response.
	return CustomerResponse{
		ID:      customer.ID,
		Address: customer.Address,
		Active:  customer.Active,
		Notes:   customer.Notes,
		Status:  customer.Status,
		Name:    customer.User.Name,
		Email:   customer.User.Email,
		Phone:   customer.User.Phone,
		Gender:  customer.User.Gender,
		Image:   customer.User.Image,
		Type:    string(customer.User.Type),
		UserId:  customer.User.ID,
	}
}

func getOneCustomerById(id string, u *CustomerRepositoryImpl) (*models.Customer, error) {
	//----> Initialize customer.
	customer := new(models.Customer)

	//----> Fetch customer by id.
	if err := u.DB.Preload("User").Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	return customer, nil
}

func toCustomerResponseList(customers []models.Customer) []CustomerResponse {
	//----> Check for empty customers.
	if len(customers) == 0 {
		return []CustomerResponse{}
	}

	var customerResponses []CustomerResponse

	//----> Iterate through customers.
	for _, customer := range customers {

		if customer.User != nil {
			customerResponses = append(customerResponses, toCustomerResponse(customer))
		}

	}
	return customerResponses
}
