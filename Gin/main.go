package main

import (
	"esm-backend/configuration"
	"esm-backend/controllers"
	"esm-backend/controllers/concepts"
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

	concepts.SetLocalPath("./content")

	router := getRouter(configuration)

	router.Run(address)
}

func getRouter(configuration configuration.Config) *gin.Engine {
	DB := db.Init(configuration)
	handler = controllers.New(DB)

	router := gin.Default()
	//router.LoadHTMLGlob("content/**/*.html")
	router.LoadHTMLFiles("content/index.html")
	router.StaticFile("content/imgs/smiley.png", "./content/imgs/smiley.png")
	router.StaticFile("favicon.ico", "./content/imgs/favicon.ico")
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
		router.Use(middlewares.Auth()).GET("/employee", handler.GetAllEmployees)
		router.Use(middlewares.Auth()).GET("/employee/:id", handler.GetEmployee)
		router.Use(middlewares.Auth()).POST("/employee", handler.AddEmployee)
		router.Use(middlewares.Auth()).PUT("/employee/:id", handler.UpdateEmployee)
		router.Use(middlewares.Auth()).DELETE("/employee/:id", handler.DeleteEmployee)

		router.Use(middlewares.Auth()).GET("/skill", handler.GetAllSkills)
		router.Use(middlewares.Auth()).GET("/skill/:id", handler.GetSkill)
		router.Use(middlewares.Auth()).POST("/skill", handler.AddSkill)
		router.Use(middlewares.Auth()).DELETE("/skill/:id", handler.DeleteSkill)
		router.Use(middlewares.Auth()).POST("/assign-skill", handler.AssignSkill)
		router.Use(middlewares.Auth()).DELETE("/remove-skill/:employeeid/:skillid", handler.RemoveSkill)
		router.Use(middlewares.Auth()).GET("/employee-skills/:employeeid", handler.GetEmployeeSkills)
		router.Use(middlewares.Auth()).GET("/level", handler.SkillLevel)

	}
}

func demoRoutes(superRoute *gin.RouterGroup) {
	router := superRoute.Group("concepts")
	{
		router.GET("/index", concepts.Index)
		router.GET("/", concepts.Index)
		router.GET("/render/:name/:age/:pos/", concepts.RenderType)
		router.GET("/serving-external", concepts.ServingExternal)
	}
}

func AddRoutes(superRoute *gin.RouterGroup) {
	authRoutes(superRoute)
	apiRoutes(superRoute)
	demoRoutes(superRoute)
}
