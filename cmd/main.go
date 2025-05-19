package main

import (
	"52weeks/internal/data"
	"52weeks/internal/handler"

	"github.com/gin-gonic/gin"
)

func main() {
	data.LoadChallenges()

	router := gin.Default()
	router.GET("/challenges", handler.GetChallenges)
	router.GET("/challenges/:id", handler.GetChallengeByID)
	router.PUT("/challenges/:id", handler.UpdateChallenge)
	router.DELETE("/challenges/:id", handler.DeleteChallenge)
	router.POST("/challenges", handler.CreateChallenge)
	router.POST("/challenges/:id/comments", handler.PostComment)

	router.POST("/user/register", handler.CreateUser)

	router.Run()
}
