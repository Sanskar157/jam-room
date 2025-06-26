package controllers

import (
	"jam-room/models"
	"net/http"
	"github.com/gin-gonic/gin"
)

type registerInput struct { 
	Password string `json:"password" binding:"required"`
	Email string `json:"email" binding:"required"`
	Name string `json:"name" binding:"required"`
}

type loginInput struct {
	Email string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func Register(c *gin.Context){
	var input registerInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	u := models.User{}

	u.Email = input.Email
	u.Password = input.Password
	u.Name = input.Name
	
	_,err := u.SaveUser()

	if err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message":"registration success"})
}

func Login(c *gin.Context) {
	var input loginInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	u := models.User{}
	u.Email = input.Email
	u.Password = input.Password

	token, err := models.LoginCheck(u.Email,u.Password)
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}