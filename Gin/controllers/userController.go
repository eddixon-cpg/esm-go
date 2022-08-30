package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	fmt.Println("Login")
}

func Signup(c *gin.Context) {
	fmt.Println("Signup")
}

func Verify(c *gin.Context) {
	fmt.Println("verify")
}
