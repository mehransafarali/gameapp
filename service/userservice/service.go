package userservice

import (
	"GameApp/entity"
	"GameApp/pkg/hashing"
	"GameApp/pkg/phonenumber"
	"errors"
	"fmt"
)

type Repository interface {
	IsPhoneNumberUnique(phoneNumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
	GetUserByPhoneNumber(phoneNumber string) (entity.User, bool, error)
}
type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name        string `json:"name"`
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type RegisterResponse struct {
	entity.User
}

type LoginRequest struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}
type LoginResponse struct{}

func (s *Service) Register(req RegisterRequest) (RegisterResponse, error) {
	//TODO - we should verify phone number with verification code

	//validate phonenumber
	if _, err := phonenumber.IsValid(req.PhoneNumber); err != nil {
		return RegisterResponse{}, errors.New(err.Error())
	}

	//check uniqueness
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisterResponse{}, fmt.Errorf("unexpected error: %w", err)
		}

		if !isUnique {
			return RegisterResponse{}, errors.New("phone number is already used")
		}
	}

	//validate name
	if len(req.Name) < 3 {
		return RegisterResponse{}, errors.New("name length must be at least 3")
	}

	//TODO- check the password with regex pattern
	//validate password
	if len(req.Password) < 8 {
		return RegisterResponse{}, errors.New("password length must be at least 8")
	}

	user := entity.User{
		ID:          0,
		PhoneNumber: req.PhoneNumber,
		Name:        req.Name,
		Password:    hashing.Hash(req.Password),
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

func (s *Service) Login(req LoginRequest) (LoginResponse, error) {
	user, exists, err := s.repo.GetUserByPhoneNumber(req.PhoneNumber)
	if err != nil {
		return LoginResponse{}, fmt.Errorf("unexpected error: %w", err)
	}

	if !exists {
		return LoginResponse{}, fmt.Errorf("username or password is wrong")
	}

	if user.Password != hashing.Hash(req.Password) {
		return LoginResponse{}, fmt.Errorf("username or password is wrong")
	}

	return LoginResponse{}, nil
}
