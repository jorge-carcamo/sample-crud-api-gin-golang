package services

import (
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/models"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/repositories"
)

type VehicleService interface {
	FindAll() []models.Vehicle
	FindById(uint64) models.Vehicle
	Create(models.Vehicle) error
	Update(models.Vehicle) error
	Delete(models.Vehicle) error
}

type vehicleService struct {
	repository repositories.VehicleRepository
}

func New(vehicleRepository repositories.VehicleRepository) VehicleService {
	return &vehicleService{
		repository: vehicleRepository,
	}
}

func (service *vehicleService) FindAll() []models.Vehicle {
	return service.repository.FindAll()
}

func (service *vehicleService) FindById(id uint64) models.Vehicle {
	return service.repository.FindById(id)
}

func (service *vehicleService) Create(vehicle models.Vehicle) error {
	service.repository.Create(vehicle)
	return nil
}

func (service *vehicleService) Update(vehicle models.Vehicle) error {
	service.repository.Update(vehicle)
	return nil
}

func (service *vehicleService) Delete(vehicle models.Vehicle) error {
	service.repository.Delete(vehicle)
	return nil
}
