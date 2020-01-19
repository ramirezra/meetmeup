package graphql

import (
	"context"

	"github.com/ramirezra/meetmeup/models"
)

type userResolver struct{ *Resolver }

// User function defined
func (r *Resolver) User() UserResolver {
	return &userResolver{r}
}

func (u *userResolver) Meetups(ctx context.Context, obj *models.User) ([]*models.Meetup, error) {
	return u.MeetupsRepo.GetMeetupsForUser(obj)
}
