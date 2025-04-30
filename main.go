package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Challenge struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	TargetValue float64 `json:"targetValue"`
}

func main() {
	var challenges = []Challenge{
		{ID: "1", Name: "Férias de verão", Description: "Férias de verão com destino ao nordeste", TargetValue: 10000},
		{ID: "2", Name: "Playstation 5", Description: "Console novo 2025", TargetValue: 5000},
	}
	r := gin.Default()
	r.GET("/challenges", func(c *gin.Context) {
		c.JSON(200, challenges)
	})
	fmt.Println(challenges)
	r.Run()
}
