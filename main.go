package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Challenge struct {
	ID          string
	name        string
	description string
	targetValue float64
}

func main() {
	var challenges = []Challenge{
		{ID: "1", name: "Férias de verão", description: "Férias de verão com destino ao nordeste", targetValue: 10000}}
	r := gin.Default()
	r.GET("/challenges", func(c *gin.Context) {
		c.JSON(200, challenges)
	})
	fmt.Println(challenges)
	r.Run()
}
