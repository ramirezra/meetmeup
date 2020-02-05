package graphql

import (
	"context"
	"errors"

	"github.com/ramirezra/meetmeup/models"
)

var (
	// ErrInput documents the bad credentials were passed by user during login process
	ErrInput = errors.New("input errors")
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
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.Register(ctx, input)
}

// Login method defined
func (m *mutationResolver) Login(ctx context.Context, input models.LoginInput) (*models.AuthResponse, error) {
	isValid := validation(ctx, input)
	if !isValid {
		return nil, ErrInput
	}
	return m.Domain.Login(ctx, input)
}

// CreateMeetup method checks if user is logged in, checks if meetup is valid, and then store meet up.
func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	return m.Domain.CreateMeetup(ctx, input)
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	return m.Domain.DeleteMeetup(ctx, id)
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	return m.Domain.UpdateMeetup(ctx, id, input)
}
