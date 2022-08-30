package controllers

import "gorm.io/gorm"

type DbHandler struct {
	DB *gorm.DB
}

func New(db *gorm.DB) DbHandler {
	return DbHandler{db}
}
