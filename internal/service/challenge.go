package service

import (
	"52weeks/internal/models"
	"errors"
)

func ValidateChallengeTarget(challenge *models.Challenge) error {
	if challenge.TargetValue <= 0 {
		return errors.New("a meta não pode ser negativa ou igual a zero")
	}

	return nil
}
