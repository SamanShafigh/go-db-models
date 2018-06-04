package store

import (
	"github.com/SamanShafigh/go-db-models/models"
	"gopkg.in/mgo.v2"
)

//UserStore implements models.UserStorer
type UserStore struct {
	session *mgo.Session
}

func (cs *UserStore) List(filters ...models.UserFilter) ([]models.User, error) {
	return nil, nil
}

func (cs *UserStore) Get(filters ...models.UserFilter) (models.User, error) {
	return models.User{}, nil
}

func (cs *UserStore) Create(User models.User) error {
	return nil
}

func (cs *UserStore) Update(User models.User) error {
	return nil
}

func (cs *UserStore) Delete(filters ...models.UserFilter) error {
	return nil
}

func NewUserStore(session *mgo.Session) models.UserStorer {
	return &UserStore{session}
}
