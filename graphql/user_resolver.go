package graphql

import (
	"context"

	"github.com/ramirezra/meetmeup/models"
)

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

// User function defined
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}
