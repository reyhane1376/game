package userservice

import (
	"fmt"
	"game/entity"
	"game/pkg/phonenumber"
)


type Repository interface {
	IsPhoneNumberUnique(phonenumber string) (bool, error)
	Register(u entity.User) (entity.User, error)
}

type Service struct {
	repo Repository
}

type RegisterRequest struct {
	Name string
	PhoneNumber string
}

type RegisteResponse struct {
	User entity.User
}

func New(repo Repository) Service {
	return Service{repo: repo}
}

func (s *Service) Register(req RegisterRequest) (RegisteResponse, error) {
	// TODO - we should verify phone number by verification code
	//validate phone number

	if !phonenumber.IsValid(req.PhoneNumber) {
		return RegisteResponse{}, fmt.Errorf("Invalid phone number")
	}

	// check uniqueness of phone number
	if isUnique, err := s.repo.IsPhoneNumberUnique(req.PhoneNumber); err != nil || !isUnique {
		if err != nil {
			return RegisteResponse{}, fmt.Errorf("unexpected error : %w", err)
		}

		if !isUnique {
			return RegisteResponse{}, fmt.Errorf("phone number is not unique")
		}
	}

	//validate name

	if len(req.Name) < 3 {
		return RegisteResponse{}, fmt.Errorf("name length should be greater than 3")
	}

	user := entity.User{
		ID         : 0,
		Name       : req.Name,
		PhoneNumber: req.PhoneNumber,
	}
	// create new user in storage
	createdUser, error := s.repo.Register(user)

	if error != nil {
		return RegisteResponse{}, fmt.Errorf("unexpected error : %w", error)
	}


	// return created user
	return RegisteResponse{User: createdUser}, nil

}