package repositories

import (
	"errors"
	"go-backend-for-many-tech-to-many-ticket-for-car-repair-app/models"

	"gorm.io/gorm"
)

type TechnicianRepositoryInt interface {
	CreateTechnician(technicianCreate *models.TechnicianCreate) (TechnicianResponse, error)
	DeleteTechnicianById(id string) (TechnicianResponse, error)
	EditTechnicianById(id string, technicianEdit *models.TechnicianEdit) (TechnicianResponse, error)
	GetTechnicianById(id string) (TechnicianResponse, error)
	GetAllTechnicians() ([]TechnicianResponse, error)
	GetTechnicianByUserId(userId string) (TechnicianResponse, error)
	GetTechnicianBySpecialty(specialty string) ([]TechnicianResponse, error)
}

type TechnicianRepositoryImpl struct {
	DB *gorm.DB
}

func NewTechnicianRepositoryImpl(DB *gorm.DB) *TechnicianRepositoryImpl {
	return &TechnicianRepositoryImpl{DB: DB}
}

func (t *TechnicianRepositoryImpl) CreateTechnician(technicianCreate *models.TechnicianCreate) (TechnicianResponse, error) {
	//----> Validate input.
	if err := models.ValidateTechnicianCreate(technicianCreate); err != nil {
		return TechnicianResponse{}, err
	}

	//----> Initialize technician.
	technician := &models.Technician{
		UserID:    technicianCreate.UserID,
		Specialty: technicianCreate.Specialty,
	}

	//----> Create technician.
	if err := t.DB.Create(&technician).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := t.DB.Where("id = ?", technicianCreate.UserID).First(user).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponse(*technician, user), nil
}

func (t *TechnicianRepositoryImpl) DeleteTechnicianById(id string) (TechnicianResponse, error) {
	//----> Check for existence of technician.
	technician, err := getOneTechnicianById(id, t)

	//----> Check for error.
	if err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Delete technician.
	if err := t.DB.Delete(technician).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := t.DB.Where("id = ?", technician.UserID).First(user).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponse(*technician, user), nil
}

func (t *TechnicianRepositoryImpl) EditTechnicianById(id string, technicianEdit *models.TechnicianEdit) (TechnicianResponse, error) {
	//----> Check for existence of technician.
	technician, err := getOneTechnicianById(id, t)

	//----> Check for error.
	if err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Validate input.
	if err := models.ValidateTechnicianEdit(technicianEdit); err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize and populate technician.
	technicianToEdit := &models.Technician{
		ID:        technician.ID,
		UserID:    technician.UserID,
		Specialty: technicianEdit.Specialty,
	}

	//----> Update technician.
	if err := t.DB.Model(&technician).Where("id = ?", id).Updates(technicianToEdit).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := t.DB.Where("id = ?", technician.UserID).First(user).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponse(*technician, user), nil
}

func (t *TechnicianRepositoryImpl) GetTechnicianById(id string) (TechnicianResponse, error) {
	//----> Fetch the technician with giving id.
	technician, err := getOneTechnicianById(id, t)

	//----> Check for error.
	if err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize user
	user := &models.User{}

	//----> Preload user.
	if err := t.DB.Where("id = ?", technician.UserID).First(user).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponse(*technician, user), nil
}

func (t *TechnicianRepositoryImpl) GetAllTechnicians() ([]TechnicianResponse, error) {
	//----> Initialize technicians.
	var technicians []models.Technician

	//----> Fetch all technicians.
	if err := t.DB.Find(&technicians).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponseList(technicians, t), nil
}

func (t *TechnicianRepositoryImpl) GetTechnicianBySpecialty(specialty string) ([]TechnicianResponse, error) {
	//----> Initialize technicians.
	var technicians []models.Technician

	//----> Fetch all technicians.
	if err := t.DB.Where("specialty = ?", specialty).Find(&technicians).Error; err != nil {
		return nil, errors.New(err.Error())
	}

	//----> Send back response.
	return toTechnicianResponseList(technicians, t), nil
}

func (t *TechnicianRepositoryImpl) GetTechnicianByUserId(userId string) (TechnicianResponse, error) {
	//----> Initialize technician.
	technician := models.Technician{}

	//----> Fetch technician by user id.
	if err := t.DB.Where("user_id = ?", userId).First(&technician).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//----> Initialize user.
	user := &models.User{}

	//----> Preload user.
	if err := t.DB.Where("id = ?", technician.UserID).First(user).Error; err != nil {
		return TechnicianResponse{}, errors.New(err.Error())
	}

	//-----> Send back response.
	return toTechnicianResponse(technician, user), nil
}
