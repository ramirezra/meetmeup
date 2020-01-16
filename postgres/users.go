package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/ramirezra/meetmeup/models"
)

// UsersRepo struct defined to interact with postgres databases
type UsersRepo struct {
	DB *pg.DB
}

// GetUserByID retrieve all meetups from the database.
func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where("id =?", id).First()
	if err != nil {
		return nil, err
	}
	return &user, nil
}
