package data

import (
	"52weeks/internal/models"
	"encoding/json"
	"fmt"
	"os"
)

var Challenges []models.Challenge

func LoadChallenges() {
	file, err := os.Open("dados/challenge.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		panic(err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)

	if err := decoder.Decode(&Challenges); err != nil {
		fmt.Println("Error decoding JSON:", err)
		panic(err)
	}
}

func SaveChallenge() {
	file, err := os.Create("dados/challenge.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()

	encoder := json.NewEncoder(file)
	if err := encoder.Encode(Challenges); err != nil {
		fmt.Println("Error encoding JSON:", err)
	}
}
