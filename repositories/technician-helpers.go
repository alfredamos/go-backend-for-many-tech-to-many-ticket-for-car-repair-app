package repositories

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toTechnicianResponse(Tech models.Technician, user *models.User) TechnicianResponse {
	return TechnicianResponse{
		ID:        Tech.ID,
		Specialty: Tech.Specialty,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Gender:    user.Gender,
		Image:     user.Image,
		Type:      string(user.Type),
		UserId:    user.ID,
	}
}

func getOneTechnicianById(id string, u *TechnicianRepositoryImpl) (*models.Technician, error) {
	//----> Initialize technician.
	technician := new(models.Technician)

	//----> Fetch customer by id.
	if err := u.DB.Where("id = ?", id).First(&technician).Error; err != nil {
		return nil, err
	}

	return technician, nil
}

func toTechnicianResponseList(technicians []models.Technician, t *TechnicianRepositoryImpl) []TechnicianResponse {
	var technicianResponses []TechnicianResponse
	for _, technician := range technicians {
		user, err := getOneUserByIdInTechnician(technician.UserID, t)

		//----> Check for error.
		if err != nil {
			continue
		}

		technicianResponses = append(technicianResponses, toTechnicianResponse(technician, user))
	}
	return technicianResponses
}

func getOneUserByIdInTechnician(id string, c *TechnicianRepositoryImpl) (*models.User, error) {
	//----> Initialize user.
	user := new(models.User)

	//----> Fetch user by id.
	if err := c.DB.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, err
	}

	//----> Send back response.
	return user, nil
}
