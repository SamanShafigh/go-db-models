package model

//User represents an application user
type User struct {
	ID        string `db:"id"`
	Email     string `db:"email"`
	Password  string `db:"password"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

// MakeUser makes a new user
func MakeUser(id, firstName, lastName string) User {
	return User{ID: id, FirstName: firstName, LastName: lastName}
}

// UserFilter is used for filtering results
type UserFilter func(*UserFilterConfig) error

// UserFilterConfig is the config for filtering results
type UserFilterConfig struct {
	Limit     int64  `json:"limit"`
	ID        string `json:"id"`
	Email     string `json:"email"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
}

func UserLimitFilter(limit int64) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Limit = limit
		return nil
	}
}

func UserIDFilter(id string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.ID = id
		return nil
	}
}

func UserEmailFilter(email string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.Email = email
		return nil
	}
}

func UserFirstNameFilter(firstName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.FirstName = firstName
		return nil
	}
}

func UserLastNameFilter(lastName string) UserFilter {
	return func(fc *UserFilterConfig) error {
		fc.LastName = lastName
		return nil
	}
}

// UserStorer represents a user store
// this can be used to store user on different
// data store
type UserStorer interface {
	List(...UserFilter) ([]User, error)
	Get(...UserFilter) (User, error)
	Create(User) error
	Update(User) error
	Delete(...UserFilter) error
}
