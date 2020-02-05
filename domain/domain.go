package domain

import (
	"github.com/ramirezra/meetmeup/models"
	"github.com/ramirezra/meetmeup/postgres"
)

// Domain struct defined
type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

// NewDomain links application to database
func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, MeetupsRepo: meetupsRepo}
}

// Ownable is an interface to see if a model can be owned by a specific user.
type Ownable interface {
	IsOwner(user *models.User) bool
}

func checkOwnership(o Ownable, user *models.User) bool {

	return o.IsOwner(user)
}
