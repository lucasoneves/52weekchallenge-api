package handler

import (
	"52weeks/internal/data"
	"52weeks/internal/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func GetChallenges(c *gin.Context) {

	c.JSON(200, data.Challenges)
}

func CreateChallenge(c *gin.Context) {
	var newChallenge models.Challenge
	if err := c.ShouldBindJSON(&newChallenge); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idGenerated := strconv.Itoa(len(data.Challenges) + 1)

	id, err := strconv.Atoi(idGenerated)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate ID"})
	}
	newChallenge.ID = strconv.Itoa(id)

	data.Challenges = append(data.Challenges, newChallenge)

	data.SaveChallenge()
	c.JSON(200, newChallenge)
}

func GetChallengeByID(c *gin.Context) {
	id := c.Param("id")
	for _, challenge := range data.Challenges {
		if challenge.ID == id {
			c.JSON(200, challenge)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Challenge not found"})
}

func DeleteChallenge(c *gin.Context) {
	id := c.Param("id")
	for i, challenge := range data.Challenges {
		if challenge.ID == id {
			data.Challenges = append(data.Challenges[:i], data.Challenges[i+1:]...)
			data.SaveChallenge()
			c.JSON(200, gin.H{"message": "Challenge deleted successfully"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Challenge not found"})
}

func UpdateChallenge(c *gin.Context) {
	id := c.Param("id")
	var updatedChallenge models.Challenge
	if err := c.ShouldBindJSON(&updatedChallenge); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	if updatedChallenge.ID != "" {
		c.JSON(400, gin.H{"error": "O ID não deve ser fornecido no corpo da requisição para uma atualização. O ID da URL é utilizado."})
		return
	}
	for i, challenge := range data.Challenges {
		if challenge.ID == id {
			data.Challenges[i] = updatedChallenge
			data.Challenges[i].ID = id
			data.SaveChallenge()
			c.JSON(200, data.Challenges[i])
			return
		}
	}

	c.JSON(404, gin.H{"error": "Challenge not found"})
}
