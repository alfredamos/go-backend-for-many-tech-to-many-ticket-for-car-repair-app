package repositories

import (
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"
)

func toTechnicianResponse(Tech models.Technician) TechnicianResponse {
	return TechnicianResponse{
		ID:        Tech.ID,
		Specialty: Tech.Specialty,
		Name:      Tech.User.Name,
		Email:     Tech.User.Email,
		Phone:     Tech.User.Phone,
		Gender:    Tech.User.Gender,
		Image:     Tech.User.Image,
		Type:      string(Tech.User.Type),
		UserId:    Tech.User.ID,
	}
}

func getOneTechnicianById(id string, u *TechnicianRepositoryImpl) (*models.Technician, error) {
	//----> Initialize technician.
	technician := new(models.Technician)

	//----> Fetch customer by id.
	if err := u.DB.Preload("User").Where("id = ?", id).First(&technician).Error; err != nil {
		return nil, err
	}

	return technician, nil
}

func toTechnicianResponseList(technicians []models.Technician) []TechnicianResponse {
	var technicianResponses []TechnicianResponse
	for _, technician := range technicians {
		technicianResponses = append(technicianResponses, toTechnicianResponse(technician))
	}
	return technicianResponses
}
