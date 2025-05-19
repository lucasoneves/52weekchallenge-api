package service

import (
	"52weeks/internal/models"
	"errors"
)

func ValidateReviewRating(comment models.Comment) error {
	if comment.Rating < 1 || comment.Rating > 5 {
		return errors.New("rating must be between 1 and 5")
	}

	return nil
}
