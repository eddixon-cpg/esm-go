package concepts

import (
	"encoding/xml"
	"fmt"
	"net/http"
	"strconv"
	"strings"

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
