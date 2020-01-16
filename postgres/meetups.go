package postgres

import (
	"github.com/go-pg/pg/v9"
	"github.com/ramirezra/meetmeup/models"
)

// MeetupsRepo struct defined to interact with postgres databases
type MeetupsRepo struct {
	DB *pg.DB
}

// GetMeetups retrieve all meetups from the database.
func (m *MeetupsRepo) GetMeetups() ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	err := m.DB.Model(&meetups).Select()
	if err != nil {
		return nil, err
	}
	return meetups, nil
}

// CreateMeetup inserts a new meetup into the database.
func (m *MeetupsRepo) CreateMeetup(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Returning("*").Insert()

	return meetup, err
}
