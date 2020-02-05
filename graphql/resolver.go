//go:generate go run github.com/99designs/gqlgen -v

package graphql

import (
	"github.com/ramirezra/meetmeup/domain"
)

// Resolver struct defined. Over all struct that contains principal methods.
type Resolver struct {
	Domain *domain.Domain
}
