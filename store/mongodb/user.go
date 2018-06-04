package store

import (
	"github.com/SamanShafigh/go-db-models/model"
	"gopkg.in/mgo.v2"
)

//UserStore implements models.UserStorer
type UserStore struct {
	session *mgo.Session
}

// Get gets list of user entities based on filter handler
func (cs *UserStore) Get(fh ...model.UserFilterHandler) (*[]model.User, error) {
	return nil, nil
}

// GetByID gets one user by ID
func (cs *UserStore) GetByID(ID string) (*model.User, error) {
	return nil, nil
}

// Add adds a user entity
func (cs *UserStore) Add(User model.User) error {
	return nil
}

// Update updates a user entity
func (cs *UserStore) Update(User model.User) error {
	return nil
}

// Delete deletes a user entity
func (cs *UserStore) Delete(fh ...model.UserFilterHandler) error {
	return nil
}

// NewUserStore returns a user store management
func NewUserStore(session *mgo.Session) model.UserStorer {
	return &UserStore{session}
}
