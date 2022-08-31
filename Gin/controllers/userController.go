package controllers

import (
	"esm-backend/application/user"
	"esm-backend/models/in"

	"net/http"

	"github.com/gin-gonic/gin"
)

func (h DbHandler) Login(c *gin.Context) {
	credentials := in.CredentialsInput{}

	if err := c.ShouldBindJSON(&credentials); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	response, err := user.Login(credentials, h.DB)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, response)

}

func (h DbHandler) Signup(c *gin.Context) {
	_user := in.UserInput{}
	if err := c.ShouldBindJSON(&_user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err := user.Signup(_user, h.DB)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, "User created successfully")
}

func (h DbHandler) Verify(c *gin.Context) {
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.JSON(401, gin.H{"error": "request does not contain an access token"})
		c.Abort()
		return
	}
	result, err := user.Verify(tokenString, h.DB)

	if err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, result)
}
