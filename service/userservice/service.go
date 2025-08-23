package userservice

import (
	"GameApp/entity"
	"GameApp/pkg/phonenumber"
	"errors"
	"fmt"
)

type Repository interface {
	isPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}
type Service struct {
	repo Repository
}

type RegisterRequest struct {
	PhoneNumber string
	Name        string
}
type RegisterResponse struct {
	entity.User
}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//TODO - we should verify phone number with verification code

	//validate phonenumber
	if _, err := phonenumber.IsValid(req.PhoneNumber); err != nil {
		return RegisterResponse{}, errors.New(err.Error())
	}

	//check uniqueness
	if isUnique, err := s.repo.isPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, errors.New("phone number is already used")
		}
	}

	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, errors.New("name is too short")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
	}

	createdUser, err := s.repo.Register(user)
	if err != nil {
		return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	// todo - you must change it
	return RegisterResponse{
		User: createdUser,
	}, nil
}
