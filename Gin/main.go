package main

import (
	"errors"
	"esm-backend/configuration"
	"esm-backend/controllers"
	"esm-backend/controllers/concepts"
	"esm-backend/db"
	"esm-backend/middlewares"
	"fmt"
	"html/template"
	"io"
	"os"
	"path/filepath"
	"time"

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
	fmt.Println("runing ", address)
	router.Run(address)
}

func getRouter(configuration configuration.Config) *gin.Engine {
	DB := db.Init(configuration)
	handler = controllers.New(DB)

	configurateLog(configuration)

	router := gin.Default()

	router.Delims("{{", "}}")
	router.SetFuncMap(template.FuncMap{
		"formatAsDate": formatAsDate,
	})

	router.LoadHTMLFiles("content/index.html")
	//router.LoadHTMLGlob("templates/**/*.tmpl") // TODO I shoultry this later!!!!
	//router.LoadHTMLGlob("templates/**")
	router.LoadHTMLGlob("templates/**/*")
	//router.LoadHTMLGlob("templates/*.*")
	router.StaticFile("favicon.ico", "./content/imgs/favicon.ico")
	router.Static("/content", "./content")
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

		router.Use(middlewares.Auth()).GET("/designations", handler.GetAllDesignations)

	}
}

func demoRoutes(superRoute *gin.RouterGroup) {
	router := superRoute.Group("/concepts")
	{
		router.GET("/index", concepts.Index)
		router.GET("/", concepts.Index)
		router.GET("/render/:name/:age/:pos/", concepts.RenderType)
		router.GET("/serving-external", concepts.ServingExternal)
		router.GET("/serving-from-file", concepts.ServingFromFile)
		router.POST("/login", concepts.LoginFromForm)
		router.GET("/map-from-query", concepts.MapFromQueryString)
		router.POST("/map-from-form", concepts.MapFromPostForm)
		router.POST("/upload-file", concepts.Upload)
		router.GET("/get-cookie", concepts.GetCookie)
		router.GET("/posts/index", concepts.PostIndex)
		router.GET("/users/index", concepts.UserIndex)
		router.POST("/redirect", concepts.Redirect)
		router.GET("/redirect", concepts.Redirect)
		router.GET("/custom-template-func", concepts.CustomTemplateFunc)
	}
}

func AddRoutes(superRoute *gin.RouterGroup) {
	authRoutes(superRoute)
	apiRoutes(superRoute)
	demoRoutes(superRoute)
}

func configurateLog(configuration configuration.Config) {
	if configuration.LogToFile {
		gin.DisableConsoleColor()

		if _, err := os.Stat(configuration.LogPath); err == nil {
		} else if errors.Is(err, os.ErrNotExist) {
			dir := filepath.Dir(configuration.LogPath)
			os.MkdirAll(dir, os.ModeDir)
		} else {
			fmt.Print("ERROR on logPAth ", err)
		}

		f, err := os.Create(configuration.LogPath)
		fmt.Println("creating gin.log was ", err)
		gin.DefaultWriter = io.MultiWriter(f)
	}
}

func formatAsDate(t time.Time) string {
	year, month, day := t.Date()
	hour, min, sec := t.Clock()
	date := fmt.Sprintf("%d/%02d/%02d %02d:%02d:%02d", year, month, day, hour, min, sec)
	fmt.Println("date is ", date)
	return date
}
