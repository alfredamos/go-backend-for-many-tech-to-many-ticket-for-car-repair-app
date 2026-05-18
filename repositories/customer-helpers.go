package repositories

import "go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

func toCustomerResponse(Customer models.Customer, user *models.User) CustomerResponse {
	return CustomerResponse{
		ID:      Customer.ID,
		Address: Customer.Address,
		Active:  Customer.Active,
		Notes:   Customer.Notes,
		Status:  Customer.Status,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		Gender:  user.Gender,
		Image:   user.Image,
		Type:    string(user.Type),
		UserId:  user.ID,
	}
}

func getOneCustomerById(id string, u *CustomerRepositoryImpl) (*models.Customer, error) {
	//----> Initialize customer.
	customer := new(models.Customer)

	//----> Fetch customer by id.
	if err := u.DB.Where("id = ?", id).First(&customer).Error; err != nil {
		return nil, err
	}

	return customer, nil
}

func toCustomerResponseList(customers []models.Customer, c *CustomerRepositoryImpl) []CustomerResponse {
	var customerResponses []CustomerResponse
	for _, customer := range customers {
		user, err := getOneUserByIdInCustomer(customer.UserID, c)

		//----> Check for error.
		if err != nil {
			continue
		}

		customerResponses = append(customerResponses, toCustomerResponse(customer, user))
	}
	return customerResponses
}

func getOneUserByIdInCustomer(id string, c *CustomerRepositoryImpl) (*models.User, error) {
	//----> Initialize user.
	user := new(models.User)

	//----> Fetch user by id.
	if err := c.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	//----> Send back response.
	return user, nil
}
