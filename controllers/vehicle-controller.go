package controllers

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/models"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/services"
)

type VehicleController interface {
	FindAll() []models.Vehicle
	FindById(ctx *gin.Context) (models.Vehicle, error)
	Create(ctx *gin.Context) error
	Update(ctx *gin.Context) error
}

type controller struct {
	service services.VehicleService
}

func New(service services.VehicleService) VehicleController {
	return &controller{
		service: service,
	}
}

func (c *controller) FindAll() []models.Vehicle {
	return c.service.FindAll()
}

func (c *controller) FindById(ctx *gin.Context) (models.Vehicle, error) {
	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return models.Vehicle{}, err
	}
	return c.service.FindById(id), nil
}

func (c *controller) Create(ctx *gin.Context) error {
	var vehicle models.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)
	if err != nil {
		return err
	}
	c.service.Create(vehicle)
	return nil
}

func (c *controller) Update(ctx *gin.Context) error {
	var vehicle models.Vehicle
	err := ctx.ShouldBindJSON(&vehicle)
	if err != nil {
		return err
	}

	id, err := strconv.ParseUint(ctx.Param("id"), 0, 0)
	if err != nil {
		return err
	}
	vehicle.Id = id
	c.service.Update(vehicle)
	return nil
}
