package config

import (
	"gin_go_learn/internal/models"
	"os"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	URI_DB := os.Getenv("DSN")
	db, err := gorm.Open(postgres.Open(URI_DB), &gorm.Config{})
	if err != nil {
		defer logrus.Error("Connection to DB Failed")
		logrus.Fatal(err.Error())
	}
	logrus.Info("Success Connect To DB")
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		defer logrus.Error("Auto Migrate DB Failed")
		logrus.Fatal(err.Error())
	}
	return db
}
