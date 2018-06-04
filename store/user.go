package store

import (
	"github.com/SamanShafigh/go-db-models/model"
	"github.com/jmoiron/sqlx"
)

// UserStore implements models.UserStorer
type UserStore struct {
	db *sqlx.DB
}

func (cs *UserStore) List(filters ...model.UserFilter) ([]model.User, error) {

	userFilter := model.UserFilterConfig{}
	for _, filter := range filters {
		filter(&userFilter)
	}

	return nil, nil
}

func (cs *UserStore) Get(filters ...model.UserFilter) (model.User, error) {
	return model.User{}, nil
}

func (cs *UserStore) Create(User model.User) error {
	_, err := cs.db.NamedExec("INSERT INTO _user (id, email, password, first_name, last_name) VALUES (:id, :email, :password, :first_name, :last_name)", User)

	return err
}

func (cs *UserStore) Update(User model.User) error {
	return nil
}

func (cs *UserStore) Delete(filters ...model.UserFilter) error {
	return nil
}

func NewUserStore(db *sqlx.DB) model.UserStorer {
	return &UserStore{db}
}
