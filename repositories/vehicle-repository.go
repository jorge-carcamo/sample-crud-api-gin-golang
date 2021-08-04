package repositories

import (
	"github.com/jinzhu/gorm"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/models"
)

type VehicleRepository interface {
	FindAll() []models.Vehicle
	FindById(uint64) models.Vehicle
	Create(models.Vehicle)
	Update(models.Vehicle)
	Delete(models.Vehicle)
}

type database struct {
	connection *gorm.DB
}

func New(conn *gorm.DB) VehicleRepository {
	return &database{
		connection: conn,
	}
}

func (db *database) FindAll() (vehicles []models.Vehicle) {
	db.connection.Find(&vehicles)
	return
}

func (db *database) FindById(id uint64) (vehicle models.Vehicle) {
	db.connection.First(&vehicle, id)
	return
}

func (db *database) Create(vehicle models.Vehicle) {
	db.connection.Create(&vehicle)
}

func (db *database) Update(vehicle models.Vehicle) {
	db.connection.Create(&vehicle)
}

func (db *database) Delete(vehicle models.Vehicle) {
	db.connection.Delete(&vehicle)
}
