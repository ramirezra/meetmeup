package domain

import "github.com/ramirezra/meetmeup/postgres"

// Domain struct defined
type Domain struct {
	UsersRepo   postgres.UsersRepo
	MeetupsRepo postgres.MeetupsRepo
}

// NewDomain links application to database
func NewDomain(usersRepo postgres.UsersRepo, meetupsRepo postgres.MeetupsRepo) *Domain {
	return &Domain{UsersRepo: usersRepo, MeetupsRepo: meetupsRepo}
}
