//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
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
