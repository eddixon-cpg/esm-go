package concepts

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

var rootPath string

type ExampleXml struct {
	XMLName  xml.Name `xml:"example"`
	Name     string   `xml:"name"`
	Age      int      `xml:"age"`
	Position string   `xml:"position"`
}

func SetLocalPath(path string) {
	rootPath = path
}

func Index(c *gin.Context) {
	fmt.Println(rootPath)

	c.HTML(http.StatusOK, "index.html", nil)
}

func RenderType(c *gin.Context) {
	_type := strings.ToLower(c.Param("name"))
	id, _ := strconv.Atoi(c.Param("age"))
	pos := c.Param("pos")
	format := c.Query("format")

	msg := ExampleXml{}
	msg.Name = _type
	msg.Age = id
	msg.Position = pos

	if format == "json" {
		c.JSON(http.StatusOK, msg)
	} else if format == "xml" {

		c.XML(http.StatusOK, msg) //gin.H{"_type": _type, "id": id}
	} else {
		c.JSON(http.StatusBadRequest, "Invalid format")
	}
}

func ServingExternal(c *gin.Context) {
	response, err := http.Get("https://es.wikipedia.org/wiki/Alemania") //"https://raw.githubusercontent.com/gin-gonic/logo/master/color.png"
	if err != nil || response.StatusCode != http.StatusOK {
		c.Status(http.StatusServiceUnavailable)
		return
	}

	reader := response.Body
	contentLength := response.ContentLength
	contentType := response.Header.Get("Content-Type")

	extraHeaders := map[string]string{}

	c.DataFromReader(http.StatusOK, contentLength, contentType, reader, extraHeaders)
}

func ServingFromFile(c *gin.Context) {
	c.File("./content/imgs/smiley.png")
}

type LoginForm struct {
	UserName string `form:"userName" binding:"required"`
	Password string `form:"password" binding:"required"`
	IsAdmin  int    `form:"isAdmin" binding:"required"`
}

func LoginFromForm(c *gin.Context) {
	fmt.Println("Authenticating trying to validate model")

	var form LoginForm
	if err := c.ShouldBind(&form); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if form.UserName == "mbartra" && form.Password == "!1234Qqwer" {
		c.JSON(200, gin.H{"status": "you are   authorized"})
	} else {
		c.JSON(401, gin.H{"status": "you are unauthorized"})
	}
}

func MapFromQueryString(c *gin.Context) {
	ids := c.QueryMap("ids")
	fmt.Printf("ids: %v; ", ids)
}

func MapFromPostForm(c *gin.Context) {

	names := c.PostFormMap("names")
	fmt.Printf("names: %v", names)
}

func Upload(c *gin.Context) {
	name := c.PostForm("name")
	email := c.PostForm("email")

	root := "C:\\Files\\Job\\uploads\\"
	// Source
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "get form err: %s", err.Error())
		return
	}

	filename := filepath.Base(file.Filename)
	fmt.Printf("Storing @ %v\n", root+filename)
	if err := c.SaveUploadedFile(file, root+filename); err != nil {
		c.String(http.StatusBadRequest, "upload file err: %s", err.Error())
		return
	}

	c.String(http.StatusOK, "File %s uploaded successfully with fields name=%s and email=%s.", file.Filename, name, email)
}

func GetCookie(c *gin.Context) {
	cookie, err := c.Cookie("gin_cookie")

	if err != nil {
		fmt.Printf("Error getting cookie %v\n", err)
		cookie = "NotSet"
		c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
	}

	fmt.Printf("Cookie value: %s \n", cookie)
}

func PostIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "posts/index.tmpl", gin.H{
		"title": "Posts",
	})
}

func UserIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "users/index.tmpl", gin.H{
		"title": "Users of ESM",
	})
}

type Mail struct {
	Message string `json:"message"`
	Email   string `json:"email"`
}

func Redirect(c *gin.Context) {
	redirect := c.Query("must")
	mustRedirect, err := strconv.ParseBool(redirect)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if mustRedirect {
		c.Redirect(http.StatusMovedPermanently, "https://developer.mozilla.org/en-US/docs/Web/HTTP/Status/301")
		return
	}

	c.JSON(200, "Your message has been sent")
}

func CustomTemplateFunc(c *gin.Context) {
	c.HTML(http.StatusOK, "func/raw.tmpl", map[string]interface{}{
		"title": "ESM users",
		"now":   time.Now(), //Date(2017, 12, 23, 10, 11, 12, 13, time.UTC),
		"name":  "Lazlo",
		"count": 7,
	})
}
