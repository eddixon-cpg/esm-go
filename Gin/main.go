package main

import (
	"esm-backend/configuration"
	"esm-backend/controllers"
	"esm-backend/db"
	"esm-backend/middlewares"
	"fmt"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	configuration.SetPath("./configuration/")
	configuration := configuration.GetConfiguration()
	address := fmt.Sprintf("localhost:%d", configuration.Port)

	router := getRouter(configuration) //.Methods(http.MethodPost)

	router.Run(address)
}

func getRouter(configuration configuration.Config) *gin.Engine {
	DB := db.Init(configuration)
	handler := controllers.New(DB)
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"},
		AllowHeaders:  []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Connection", "Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header", "X-Requested-With", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	router.POST("/login", handler.Login)
	router.POST("/signup", handler.Signup)
	router.GET("/verify", handler.Verify)

	router.GET("/employee", handler.GetAllEmployees)                      //.Methods(http.MethodGet)
	router.Use(middlewares.Auth()).POST("/employee", handler.AddEmployee) //.Methods(http.MethodPost)

	return router
}
