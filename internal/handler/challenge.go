package handler

import (
	"52weeks/internal/data"
	"52weeks/internal/models"
	"52weeks/internal/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetChallenges(c *gin.Context) {

	c.JSON(http.StatusOK, data.Challenges)
}

func CreateChallenge(c *gin.Context) {
	var newChallenge models.Challenge
	if err := c.ShouldBindJSON(&newChallenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateChallengeTarget(&newChallenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	idGenerated := strconv.Itoa(len(data.Challenges) + 1)

	id, err := strconv.Atoi(idGenerated)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate ID"})
	}
	newChallenge.ID = strconv.Itoa(id)

	data.Challenges = append(data.Challenges, newChallenge)

	data.SaveChallenge()
	c.JSON(http.StatusOK, newChallenge)
}

func GetChallengeByID(c *gin.Context) {
	id := c.Param("id")
	for _, challenge := range data.Challenges {
		if challenge.ID == id {
			c.JSON(http.StatusOK, challenge)
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
}

func DeleteChallenge(c *gin.Context) {
	id := c.Param("id")
	for i, challenge := range data.Challenges {
		if challenge.ID == id {
			data.Challenges = append(data.Challenges[:i], data.Challenges[i+1:]...)
			data.SaveChallenge()
			c.JSON(http.StatusOK, gin.H{"message": "Challenge deleted successfully"})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
}

func UpdateChallenge(c *gin.Context) {
	id := c.Param("id")
	var updatedChallenge models.Challenge
	if err := c.ShouldBindJSON(&updatedChallenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateChallengeTarget(&updatedChallenge); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if updatedChallenge.ID != "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "O id não deve ser fornecido no corpo da requisição para uma atualização."})
		return
	}
	for i, challenge := range data.Challenges {
		if challenge.ID == id {
			data.Challenges[i] = updatedChallenge
			data.Challenges[i].ID = id
			data.SaveChallenge()
			c.JSON(http.StatusOK, data.Challenges[i])
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})
}
