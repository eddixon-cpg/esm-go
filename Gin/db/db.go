package db

import (
	"log"

	"esm-backend/models/domain"

	"esm-backend/configuration"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(configuration configuration.Config) *gorm.DB {
	dbURL := configuration.ConnectionString //"host=localhost user=postgres password=Monitor25* dbname=esm port=5432 sslmode=disable" //"postgres://pg:pass@localhost:5432/crud"
	//connStr := "user=postgres dbname=connect-db password=secure-password host=localhost sslmode=disable"
	log.Println("Opening!!")
	db, err := gorm.Open(postgres.Open(dbURL), &gorm.Config{})
	log.Println(err)
	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&domain.Employee{})
	db.AutoMigrate(&domain.Designation{})
	//db.AutoMigrate(&domain.Skill{})
	//db.AutoMigrate(&domain.EmployeeSkill{})
	//db.AutoMigrate(&domain.Level{})
	db.AutoMigrate(&domain.User{})
	//db.SetupJoinTable(&domain.Employee{}, "skills", &domain.EmployeesSkill{})

	log.Println("Migrated")
	return db
}
