package postgres

import (
	"fmt"

	"github.com/go-pg/pg/v9"
	"github.com/ramirezra/meetmeup/models"
)

// UsersRepo struct defined to interact with postgres databases
type UsersRepo struct {
	DB *pg.DB
}

// GetUserByField retrieve all users from the database.
func (u *UsersRepo) GetUserByField(field, value string) (*models.User, error) {
	var user models.User
	err := u.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

// GetUserByID retrieve all users from the database.
func (u *UsersRepo) GetUserByID(id string) (*models.User, error) {
	return u.GetUserByField("id", id)
}

// GetUserByEmail retrieve all users from the database.
func (u *UsersRepo) GetUserByEmail(email string) (*models.User, error) {
	return u.GetUserByField("email", email)
}

// GetUserByUsername retrieve all users from the database.
func (u *UsersRepo) GetUserByUsername(username string) (*models.User, error) {
	return u.GetUserByField("username", username)
}

// CreateUser insert a new user into database
func (u *UsersRepo) CreateUser(tx *pg.Tx, user *models.User) (*models.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err

}
