//go:generate go run github.com/99designs/gqlgen -v

package meetmeup

import (
	"context"
	"errors"
	"log"
	"strconv"

	"github.com/ramirezra/meetmeup/models"
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

// Generated using gqlgen.
type Resolver struct{}

func (r *Resolver) Meetup() MeetupResolver {
	return &meetupResolver{r}
}
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

type meetupResolver struct{ *Resolver }

func (r *meetupResolver) User(ctx context.Context, obj *models.Meetup) (*models.User, error) {
	user := new(models.User)

	for _, u := range users {
		if u.ID == obj.UserID {
			user = u
			break
		}
	}
	if user == nil {
		return nil, errors.New("text: User with id does not exist")
	}
	return user, nil
}

type mutationResolver struct{ *Resolver }

func (r *mutationResolver) CreateMeetup(ctx context.Context, input NewMeetup) (*models.Meetup, error) {

	size := len(meetups)
	previousID, err := strconv.Atoi(meetups[size-1].ID)
	if err != nil {
		log.Fatalln(err)
	}
	newID := strconv.Itoa(previousID + 1)

	var newMeetup = models.Meetup{
		ID:          newID,
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}

	meetups = append(meetups, &newMeetup)

	return &newMeetup, nil
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {
	return meetups, nil
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
