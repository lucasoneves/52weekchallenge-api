package main

import (
	"52weeks/models"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var challenges []models.Challenge

func main() {
	loadChallenges()

	router := gin.Default()
	router.GET("/challenges", getChallenges)
	router.GET("/challenges/:id", getChallengeByID)
	router.PUT("/challenges/:id", updateChallenge)
	router.DELETE("/challenges/:id", deleteChallenge)
	router.POST("/challenges", createChallenge)
	router.GET("/welcome", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Olá Mundo"})
	})

	router.Run()
}

func getChallenges(c *gin.Context) {

	c.JSON(200, challenges)
}

func createChallenge(c *gin.Context) {
	var newChallenge models.Challenge
	if err := c.ShouldBindJSON(&newChallenge); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}

	idGenerated := strconv.Itoa(len(challenges) + 1)

	id, err := strconv.Atoi(idGenerated)
	if err != nil {
		c.JSON(500, gin.H{"error": "Failed to generate ID"})
	}
	newChallenge.ID = strconv.Itoa(id)

	challenges = append(challenges, newChallenge)

	saveChallenge()
	c.JSON(200, newChallenge)
}

func getChallengeByID(c *gin.Context) {
	id := c.Param("id")
	for _, challenge := range challenges {
		if challenge.ID == id {
			c.JSON(200, challenge)
			return
		}
	}

	c.JSON(404, gin.H{"error": "Challenge not found"})
}

func loadChallenges() {
	file, err := os.Open("dados/challenge.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&challenges); err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}
}

func saveChallenge() {
	file, err := os.Create("dados/challenge.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(challenges); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}

func updateChallenge(c *gin.Context) {
	c.JSON(200, gin.H{"message": "Updated deleted successfully"})
}

func deleteChallenge(c *gin.Context) {
	id := c.Param("id")
	for i, challenge := range challenges {
		if challenge.ID == id {
			challenges = append(challenges[:i], challenges[i+1:]...)
			saveChallenge()
			c.JSON(200, gin.H{"message": "Challenge deleted successfully"})
			return
		}
	}

	c.JSON(404, gin.H{"error": "Challenge not found"})
}
