package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/controllers"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/models"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/persistence"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/repositories"
	"github.com/jorge-carcamo/sample-crud-api-gin-golang/services"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/swag/example/basic/docs"
)

// @title Sample CRUD API GIN GOLANG
// @version 1.0
// @description #

// @contact.name Jorge Cárcamo
// @contact.email j.carcamo.bustamante@gmail.com

// @host localhost:9000
// @BasePath /api/v1
func main() {

	db, err := persistence.ConnectDB()
	if err != nil {
		fmt.Println(err)
	}
	db.AutoMigrate(&models.Vehicle{})
	defer db.Close()

	var (
		vehicleRepository repositories.VehicleRepository = repositories.New(db)
		vehicleService    services.VehicleService        = services.New(vehicleRepository)
		vehicleController controllers.VehicleController  = controllers.New(vehicleService)
	)

	server := gin.New()
	server.Use(gin.Recovery(), gin.Logger(), cors.Default())

	apiRoutes := server.Group(docs.SwaggerInfo.BasePath)
	{

		/*brands := apiRoutes.Group("/brands")
		{
			brands.GET(":id", auptopressAPI.GetMarcaById)
		}

		models := apiRoutes.Group("/models")
		{
			models.GET(":id", auptopressAPI.GetModeloById)
			models.GET(":id/versions", auptopressAPI.GetVersionByModeloId)
		}*/

		vehicles := apiRoutes.Group("/vehicle")
		{
			vehicles.GET("", func(ctx *gin.Context) {
				ctx.JSON(http.StatusOK, vehicleController.FindAll())
			})

			vehicles.GET(":id", func(ctx *gin.Context) {

				vehicle, err := vehicleController.FindById(ctx)
				if err != nil {
					ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
				} else {
					ctx.JSON(http.StatusOK, vehicle)
				}

			})
		}

	}

	server.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := os.Getenv("PORT")
	if port == "" {
		port = "9000"
	}
	server.Run(":" + port)

}
