package db

import (
	"log"

	"ESM-backend-app/pkg/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	dbURL := "host=localhost user=postgres password=Monitor25* dbname=postgres port=5432 sslmode=disable" //"postgres://pg:pass@localhost:5432/crud"
	//connStr := "user=postgres dbname=connect-db password=secure-password host=localhost sslmode=disable"
	log.Println("Opening!!")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	log.Println(err)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println("Migrating")
	db.AutoMigrate(&models.Employee{})
	//db.AutoMigrate(&models.Employee{})

	//db.AutoMigrate(&models.Skill{})
	return db
}
