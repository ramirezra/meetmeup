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

func (r *queryResolver) Meetups(ctx context.Context) ([]*models.Meetup, error) {

	return r.MeetupsRepo.GetMeetups()
}
