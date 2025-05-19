package handler

import (
	"52weeks/internal/data"
	"52weeks/internal/models"
	"52weeks/internal/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func PostComment(c *gin.Context) {
	// Implemente a lógica para adicionar um comentário a um desafio
	challengeID := c.Param("id")

	var newComment models.Comment

	if err := c.ShouldBindJSON(&newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := service.ValidateReviewRating(newComment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	for i, challenge := range data.Challenges {
		if challenge.ID == challengeID {
			challenge.Comment = append(challenge.Comment, newComment)
			data.Challenges[i] = challenge
			data.SaveChallenge()
			c.JSON(http.StatusCreated, gin.H{
				// Return a data array with json objects
				"message": "Comment added successfully",
				"data":    challenge,
				"status":  http.StatusCreated,
			})
			return
		}
	}

	c.JSON(http.StatusNotFound, gin.H{"error": "Challenge not found"})

}
