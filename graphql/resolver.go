//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"github.com/ramirezra/meetmeup/postgres"
)

// Resolver struct defined. Over all struct that contains principal methods.
type Resolver struct {
	MeetupsRepo postgres.MeetupsRepo
	UsersRepo   postgres.UsersRepo
}
