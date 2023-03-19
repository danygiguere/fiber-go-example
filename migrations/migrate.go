package migrations

import (
	"go-fiber-example/m/v2/configs"
	"go-fiber-example/m/v2/models"
	"log"
)

// to execute this file, run: go run migrations/migrate.go

func init() {
	configs.ConnectToDB()
}

func main() {
	err := configs.DBConn.AutoMigrate(&models.Product{})
	if err != nil {
		log.Fatal("Failed to AutoMigrate. \n", err)
		return
	}
}
