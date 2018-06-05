package store

import (
	mysqld "github.com/SamanShafigh/go-db-models/db"
	"github.com/SamanShafigh/go-db-models/model"
)

// UserStore implements models.UserStorer
type UserStore struct {
	db *mysqld.DB
}

// Get gets list of user entities based on filter handler
func (cs *UserStore) Get(fh ...model.UserFilterHandler) (*[]model.User, error) {
	var user model.User
	var users []model.User

	rows, err := cs.db.QueryWithFilter("SELECT * FROM _user", model.GetUserFilter(fh...))
	for rows.Next() {
		err := rows.StructScan(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return &users, err
}

// GetByID gets one user by ID
func (cs *UserStore) GetByID(ID string) (*model.User, error) {
	var user model.User
	err := cs.db.QueryRowx("SELECT * FROM _user WHERE id = ?", ID).StructScan(&user)

	return &user, err
}

// Add adds a user entity
func (cs *UserStore) Add(user model.User) error {
	_, err := cs.db.NamedExec(
		`INSERT INTO _user 
			(id, status, email, password, first_name, last_name) 
		VALUES 
			(:id, :status, :email, :password, :first_name, :last_name)`,
		user,
	)

	return err
}

// Update updates a user entity
func (cs *UserStore) Update(user model.User) error {
	return nil
}

// Delete deletes a user entity
func (cs *UserStore) Delete(fh ...model.UserFilterHandler) error {
	return nil
}

// NewUserStore returns a user store management
func NewUserStore(db *mysqld.DB) model.UserStorer {
	return &UserStore{db}
}
