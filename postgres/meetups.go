package postgres

import (
	"fmt"
	"github.com/go-pg/pg/v9"
	"github.com/ramirezra/meetmeup/models"
)

// MeetupsRepo struct defined to interact with postgres databases
type MeetupsRepo struct {
	DB *pg.DB
}

// GetMeetups retrieve all meetups from the database.
func (m *MeetupsRepo) GetMeetups(filter *models.MeetupFilter, limit, offset *int) ([]*models.Meetup, error) {
	var meetups []*models.Meetup

	query := m.DB.Model(&meetups).Order("id")
	if filter != nil {
		if filter.Name != nil && *filter.Name != "" {
			// "%jon%"
			//
			query.Where("name ILIKE ?", fmt.Sprintf("%%%s%%", *filter.Name))
		}
	}

	if limit != nil {
		query.Limit(*limit)
	}

	if offset != nil {
		query.Offset(*offset)
	}

	err := query.Select()
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

// GetByID ...
func (m *MeetupsRepo) GetByID(id string) (*models.Meetup, error) {
	var meetup models.Meetup
	err := m.DB.Model(&meetup).Where("id = ?", id).First()
	return &meetup, err
}

// Update ...
func (m *MeetupsRepo) Update(meetup *models.Meetup) (*models.Meetup, error) {
	_, err := m.DB.Model(meetup).Where("Id= ?", meetup.ID).Update()
	return meetup, err
}

// Delete ..
func (m *MeetupsRepo) Delete(meetup *models.Meetup) error {
	_, err := m.DB.Model(meetup).Where("Id= ?", meetup.ID).Delete()
	return err
}

// GetMeetupsForUser ...
func (m *MeetupsRepo) GetMeetupsForUser(user *models.User) ([]*models.Meetup, error) {
	var meetups []*models.Meetup
	err := m.DB.Model(&meetups).Where("user_id = ?", user.ID).Order("id").Select()
	return meetups, err
}
