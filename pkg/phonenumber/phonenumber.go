package phonenumber

import (
	"errors"
	"strconv"
)

func IsValid(phoneNumber string) (bool, error) {
	if len(phoneNumber) != 11 {
		return false, errors.New("phone number must be 11 characters long")
	}

	if phoneNumber[0:2] != "09" {
		return false, errors.New("phone number must begin with '09'")
	}

	if _, err := strconv.Atoi(phoneNumber[2:]); err != nil {
		return false, errors.New("phone number must be numeric")
	}

	return true, nil
}
