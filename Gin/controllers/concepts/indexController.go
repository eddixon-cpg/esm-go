package concepts

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

var rootPath string

func SetLocalPath(path string) {
	rootPath = path
}

func Index(c *gin.Context) {
	fmt.Println(rootPath)

	c.HTML(http.StatusOK, "index.html", nil)
}

func RenderType(c *gin.Context) {

	var msg struct {
		Type string
		Id   int
	}

	_type := strings.ToLower(c.Param("type"))
	id, _ := strconv.Atoi(c.Param("id"))

	msg.Type = _type
	msg.Id = id
	if _type == "json" {
		c.JSON(http.StatusOK, msg)
	} else if _type == "xml" {

		c.XML(http.StatusOK, msg) //gin.H{"_type": _type, "id": id}
	} else {
		c.JSON(http.StatusBadRequest, "Invalid format")
	}
}
