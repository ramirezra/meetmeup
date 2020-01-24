package graphql

import (
	"context"
	"errors"
	"fmt"
	"log"

	customMW "github.com/ramirezra/meetmeup/middleware"
	"github.com/ramirezra/meetmeup/models"
)

var (
	// ErrBadCredentials documents the bad credentials were passed by user during login process
	ErrBadCredentials = errors.New("email/password combination do not match")
	// ErrUnauthenticated documents that user has not been authenticated
	ErrUnauthenticated = errors.New("user not authenticated")
)

// Mutation function defined
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

// Register method defined
func (m *mutationResolver) Register(ctx context.Context, input models.RegisterInput) (*models.AuthResponse, error) {
	// check if email already exists in database
	_, err := m.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("email already in use")
	}

	// check if username already exists in database
	_, err = m.UsersRepo.GetUserByUsername(input.Username)
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

	tx, err := m.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("error creating a transaction: %v", err)
		return nil, errors.New("something went wrong")

	}
	defer tx.Rollback()

	if _, err := m.UsersRepo.CreateUser(tx, user); err != nil {
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
func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	user, err := m.UsersRepo.GetUserByEmail(input.Email)
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

// CreateMeetup method checks if user is logged in, checks if meetup is valid, and then store meet up.
func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	currentUser, err := customMW.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      currentUser.ID,
	}
	return m.MeetupsRepo.CreateMeetup(meetup)
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exist")
	}
	err = m.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup")
	}

	return true, nil
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exist")
	}

	didUpdate := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Name = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("No update done")
	}
	meetup, err = m.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updatein meetup: %v", err)
	}
	return meetup, nil
}
