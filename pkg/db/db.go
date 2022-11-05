package db

import (
	"github.com/vandong9/go-grpc-product-svc/pkg/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
)

type Handler struct {
	DB *gorm.DB
}

func Init(url string) Handler {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(models.Product{})          // why need &
	db.AutoMigrate(models.StockDecreaseLog{}) // why need &

	return Handler{db}
}
