package main

import (
	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"paysee2/constants"
	"paysee2/handlers/userHandlers"
	"paysee2/internalFunc/configs"
	"paysee2/layers/models"
	"paysee2/layers/repository"
	"paysee2/layers/services"
	"paysee2/routers"
)

var db *gorm.DB

func main() {

	dsn := configs.CreateDsn()

	// اتصال db
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(constants.ErrorConnectDB)
		return
	}

	err = db.AutoMigrate(
		&models.User{},
		&models.KYC{},
		&models.LoginHistory{},
		&models.TransAction{},
	)
	if err != nil {
		log.Fatalln(constants.ErrorCreateTable)
		return
	}

	userRepo := repository.NewGormDB(db)
	userServ := services.NewUserService(userRepo)
	uH := userHandlers.NewUserHandlers(*userServ, db)

	r := gin.Default()
	routers.Router(r, uH)
	err = r.Run(":3000")
	if err != nil {
		panic(err)
	}
}
