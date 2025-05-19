package handler

import (
	"52weeks/internal/models"
	"52weeks/internal/service"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValideEmailIsEmpty(newUser.Email); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{
		"message": "User created successfully!",
		"user":    newUser,
	})

}

// Crie uma rota PUT que aceite um ID e atualize o nome de um usuário existente.
