package main

import (
	"database/sql"
	"log"

	"github.com/FamousLuisin/api-go/src/config/database/mysql"
	"github.com/FamousLuisin/api-go/src/config/logger"
	"github.com/FamousLuisin/api-go/src/controller"
	"github.com/FamousLuisin/api-go/src/controller/routes"
	"github.com/FamousLuisin/api-go/src/model/repository"
	"github.com/FamousLuisin/api-go/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	logger.Info("About to start application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	//Connection mysql
	db, err := mysql.NewMysqlConnection()

	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(db)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func initDependencies(db *sql.DB) controller.UserControllerInterface {
	//Inicializar dependencias
	repo := repository.NewUserRepository(db)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}
