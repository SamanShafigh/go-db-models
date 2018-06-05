package model

//User represents an application user
type User struct {
	ID        string `db:"id"`
	Status    int64  `db:"status"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// UserStorer represents a user store
// this can be used to store user on different
// data store
type UserStorer interface {
	Get(...UserFilterHandler) (*[]User, error)
	GetByID(string) (*User, error)
	Add(User) error
	Update(User) error
	Delete(...UserFilterHandler) error
}

// UserFilterHandler is used for filtering results
type UserFilterHandler func(*UserFilter) error

// UserFilter is the config for filtering results
type UserFilter struct {
	Limit     *int64  `db:"limit" json:"limit"`
	Status    *int64  `db:"status" json:"status"`
	Email     *string `db:"email" json:"email"`
	FirstName *string `db:"first_name" json:"first_name"`
	LastName  *string `db:"last_name" json:"last_name"`
}

// UserLimitFilter sets limit filter
func UserLimitFilter(limit int64) UserFilterHandler {
	return func(fc *UserFilter) error {
		fc.Limit = &limit
		return nil
	}
}

// UserStatusFilter sets status filter
func UserStatusFilter(status int64) UserFilterHandler {
	return func(fc *UserFilter) error {
		fc.Status = &status
		return nil
	}
}

// UserEmailFilter sets email filter
func UserEmailFilter(email string) UserFilterHandler {
	return func(fc *UserFilter) error {
		fc.Email = &email
		return nil
	}
}

// UserFirstNameFilter sets firstName filter
func UserFirstNameFilter(firstName string) UserFilterHandler {
	return func(fc *UserFilter) error {
		fc.FirstName = &firstName
		return nil
	}
}

// UserLastNameFilter sets lastName filter
func UserLastNameFilter(lastName string) UserFilterHandler {
	return func(fc *UserFilter) error {
		fc.LastName = &lastName
		return nil
	}
}

// GetUserFilter applays all filter handlers and returns a UserFilter object
func GetUserFilter(filters ...UserFilterHandler) UserFilter {
	fc := UserFilter{}
	for _, filter := range filters {
		filter(&fc)
	}

	return fc
}

// MakeUser makes a new user
func MakeUser(id, firstName, lastName string, status int64) User {
	return User{ID: id, FirstName: firstName, LastName: lastName, Status: status}
}
