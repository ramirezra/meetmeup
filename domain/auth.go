package domain

import (
	"context"
	"errors"
	"log"

	"github.com/ramirezra/meetmeup/models"
)

var (
	// ErrBadCredentials documents the bad credentials were passed by user during login process
	ErrBadCredentials = errors.New("email/password combination do not match")
	// ErrUnauthenticated documents that user has not been authenticated
	ErrUnauthenticated = errors.New("user not authenticated")
)

// Register method defined
func (d *Domain) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	// check if email already exists in database
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// check if username already exists in database
	_, err = d.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("username already in use")
	}

	// Now register new User.
	user := &models.User{
		Username:  input.Username,
		Email:     input.Email,
		FirstName: input.FirstName,
		LastName:  input.LastName,
	}

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("error while hasing password: %v", err)
		return nil, errors.New("something went wrong")
	}
	// TODO: send verification code

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")

	}
	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("error creating a user: %v", err)
		return nil, errors.New("something went wrong")
	}

	if err := tx.Commit(); err != nil {
		log.Printf("error while commiting: %v", err)
		return nil, errors.New("something went wrong")
	}

	token, err := user.GenToken()
	if err != nil {
		log.Printf("error while generating token: %v", err)
		return nil, errors.New("something went wrong")
	}

	authResponse := &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}

	return authResponse, nil
}

// Login method defined
func (d *Domain) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err != nil {
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {

		return nil, ErrBadCredentials
	}

	token, err := user.GenToken()
	if err != nil {
		return nil, errors.New("somthing went wrong")
	}

	return &models.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

}
