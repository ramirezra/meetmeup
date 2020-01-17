package graphql

import (
	"context"
	"errors"
	"fmt"

	"github.com/ramirezra/meetmeup/models"
)

// Mutation function defined
func (r *Resolver) Mutation() MutationResolver {
	return &mutationResolver{r}
}

type mutationResolver struct{ *Resolver }

func (m *mutationResolver) CreateMeetup(ctx context.Context, input models.NewMeetup) (*models.Meetup, error) {
	if len(input.Name) < 3 {
		return nil, errors.New("name not long enough")
	}
	if len(input.Description) < 3 {
		return nil, errors.New("description not long enough")
	}

	meetup := &models.Meetup{
		Name:        input.Name,
		Description: input.Description,
		UserID:      "1",
	}
	return m.MeetupsRepo.CreateMeetup(meetup)
}

func (m *mutationResolver) DeleteMeetup(ctx context.Context, id string) (bool, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return false, errors.New("meetup not exist")
	}
	err = m.MeetupsRepo.Delete(meetup)
	if err != nil {
		return false, fmt.Errorf("error while deleting meetup")
	}

	return true, nil
}

func (m *mutationResolver) UpdateMeetup(ctx context.Context, id string, input models.UpdateMeetup) (*models.Meetup, error) {
	meetup, err := m.MeetupsRepo.GetByID(id)
	if err != nil || meetup == nil {
		return nil, errors.New("meetup not exist")
	}

	didUpdate := false
	if input.Name != nil {
		if len(*input.Name) < 3 {
			return nil, errors.New("name is not long enough")
		}
		meetup.Name = *input.Name
		didUpdate = true
	}
	if input.Description != nil {
		if len(*input.Description) < 3 {
			return nil, errors.New("description is not long enough")
		}
		meetup.Name = *input.Description
		didUpdate = true
	}

	if !didUpdate {
		return nil, errors.New("No update done")
	}
	meetup, err = m.MeetupsRepo.Update(meetup)
	if err != nil {
		return nil, fmt.Errorf("error while updatein meetup: %v", err)
	}
	return meetup, nil
}
