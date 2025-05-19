package service

import "errors"

func ValideEmailIsEmpty(email string) error {
	if email == "" {
		return errors.New("email is empty")
	}
	return nil
}
