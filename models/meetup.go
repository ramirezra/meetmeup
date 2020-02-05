package models

// Meetup exported.
type Meetup struct {
	ID          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	UserID      string `json:"userId"`
}

// IsOwner chekcs to see if the user is the owner of the meeteup
func (m *Meetup) IsOwner(user *User) bool {
	return m.UserID == user.ID
}
