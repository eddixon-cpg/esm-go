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

var handler controllers.DbHandler

func main() {
	configuration.SetPath("./configuration/")
	configuration := configuration.GetConfiguration()
	address := fmt.Sprintf("localhost:%d", configuration.Port)

	router := getRouter(configuration)

	router.Run(address)
}

func getRouter(configuration configuration.Config) *gin.Engine {
	DB := db.Init(configuration)
	handler = controllers.New(DB)

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:  []string{"*"},
		AllowMethods:  []string{"GET", "POST", "PUT", "DELETE", "HEAD", "OPTIONS", "PATCH"},
		AllowHeaders:  []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization", "Access-Control-Request-Headers", "Access-Control-Request-Method", "Connection", "Host", "Origin", "User-Agent", "Referer", "Cache-Control", "X-header", "X-Requested-With", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders: []string{"Content-Length"},
	}))

	root := router.Group("/v1")
	AddRoutes(root)

	return router
}

func authRoutes(superRoute *gin.RouterGroup) {
	authRouter := superRoute.Group("/auth")
	{
		authRouter.POST("/login", handler.Login)
		authRouter.POST("/signup", handler.Signup)
		authRouter.GET("/verify", handler.Verify)
	}
}

func apiRoutes(superRoute *gin.RouterGroup) {
	router := superRoute.Group("/api")
	{
		router.GET("/employee", handler.GetAllEmployees)
		router.GET("/employee/:id", handler.GetEmployee)
		router.Use(middlewares.Auth()).POST("/employee", handler.AddEmployee)
		router.PUT("/employee/:id", handler.UpdateEmployee)
		router.DELETE("/employee/:id", handler.DeleteEmployee)
	}
}

func AddRoutes(superRoute *gin.RouterGroup) {
	authRoutes(superRoute)
	apiRoutes(superRoute)
}
