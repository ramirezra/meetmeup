package graphql

import (
	"context"

	"github.com/ramirezra/meetmeup/models"
)

// Query function defined
func (r *Resolver) Query() QueryResolver {
	return &queryResolver{r}
}

type queryResolver struct{ *Resolver }

func (r *queryResolver) Meetups(ctx context.Context, filter *models.MeetupFilter, limit *int, offset *int) ([]*models.Meetup, error) {

	return r.Domain.MeetupsRepo.GetMeetups(filter, limit, offset)
}

func (r *queryResolver) User(ctx context.Context, id string) (*models.User, error) {

	return r.Domain.UsersRepo.GetUserByID(id)
}
