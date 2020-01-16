//go:generate go run github.com/99designs/gqlgen -v

package meetmeup

import (
	"context"
	"errors"

	"github.com/ramirezra/meetmeup/models"
	"github.com/ramirezra/meetmeup/postgres"
)

// Define data storage (in memory for now)
var meetups = []*models.Meetup{
	{
		ID:          "1",
		Name:        "A meetup",
		Description: "A description",
		UserID:      "1",
	},
	{
		ID:          "2",
		Name:        "Second meetup",
		Description: "Another description",
		UserID:      "2",
	},
}

var users = []*models.User{
	{
		ID:       "1",
		Username: "Bob",
		Email:    "bob@test.com",
	},
	{
		ID:       "2",
		Username: "Jon",
		Email:    "jon@test.com",
	},
}

// Resolver struct defined. Over all struct that contains principal methods.
type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}

// Meetup function defined
func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}

// Mutation function defined
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

// Query function defined
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

// User function defined
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type meetupResolver struct{ *Resolver }

func (m *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	return m.UsersRepo.GetUserByID(obj.UserID)
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}
	return m.MeetupsRepo.CreateMeetup(meetup)
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {

	return r.MeetupsRepo.GetMeetups()
}

type userResolver struct{ *Resolver }

func (r *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	var userMeetups []*models.Meetup

	for _, meetup := range meetups {
		if meetup.UserID == obj.ID {
			userMeetups = append(userMeetups, meetup)
		}
	}

	return userMeetups, nil

}
