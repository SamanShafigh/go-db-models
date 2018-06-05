# go-db-models

Demonstrating database agnostic models in Golang with example. The main idea is based on https://github.com/steven-ferrer/go-db-models however I did some modifcation and refacotring

## Structure

Models (entities) are located in `model` folder and they only represent the structure of your model and also an interface for a model storer which handles the DB functionalities the model needs. In a model file we also define a model filter structure and some setter to set a filter which can be used to filter the result.

```go
package model

type User struct {
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
}

type UserStorer interface {
	Get(...UserFilterHandler) (*[]User, error)
	Add(User) error
	Update(User) error
	Delete(...UserFilterHandler) error
}
```

Stores are located in `store` folder. Note that a model is separated from its store so you can have multiple implementation of a store for a same model for different  storage drivers as long as they satisfy the model storer interface. 

```go
package store

type UserStore struct {
	db *sqly.DB
}

func (cs *UserStore) Get(fh ...model.UserFilterHandler) (*[]model.User, error) {
	return nil, nil
}

func (cs *UserStore) Add(user model.User) error {
	return nil
}

func (cs *UserStore) Update(user model.User) error {
	return nil
}

func (cs *UserStore) Delete(fh ...model.UserFilterHandler) error {
	return nil
}

func NewUserStore(db *sqly.DB) model.UserStorer {
	return &UserStore{db}
}
```

Then we can use this structure like this
```go
  userStore := store.NewUserStore(db)

  user := model.MakeUser(sqld.MakeUUID(), "Saman", "Shafigh", 1)
	userStore.Add(user)

	user, err := userStore.Get(
		model.UserStatusFilter(1),
		model.UserFirstNameFilter("Saman"),
	)
```

mysqld is just a simple wrapper around sqlx so you can implement some extra functions over there  

## Contributing

Feel free to improve this by creating a PR discussing your changes.


## DB 

```
DROP TABLE IF EXISTS `_user`;

CREATE TABLE `_user` (
  `id` varchar(255) NOT NULL DEFAULT '',
  `status` int(11) DEFAULT NULL,
  `email` varchar(255) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `first_name` varchar(255) DEFAULT NULL,
  `last_name` varchar(255) DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

INSERT INTO `_user` (`id`, `status`, `email`, `password`, `first_name`, `last_name`)
VALUES
	('2dc11aa2-c631-4509-b639-33dea9636fba',1,'','','Saman','Shafigh'),
	('981894c2-af67-4204-957c-11bfebe5d346',1,'','','Saman','Shafigh'),
	('92a31895-82e1-49cf-87da-351d771475f4',1,'','','Saman','Shafigh'),
	('5f3e462c-e632-4afa-9ac4-f666017d6e48',1,'','','Saman','Shafigh'),
	('76c0459b-16e7-420e-9b84-6ec06df64f6d',1,'','','Saman','Shafigh'),
	('2823787f-832e-4e79-b468-f1743ea7b8c5',2,'','','Saman','Shafigh'),
	('02cd9c7f-ed56-4090-b8e5-f080c3656078',2,'','','Saman','Shafigh'),
	('05f0b45e-b704-4e5a-a974-d9f18bb12127',2,'','','Saman','Shafigh'),
	('ecbc4a6e-79ea-496c-bf45-ad4ec6cfe16e',2,'','','Saman','Shafigh'),
	('840d6d7f-0b56-459d-93d1-078af2f106f4',3,'','','Saman','Shafigh'),
	('51ce831e-ef2a-48a0-89bf-db7bcdf67e4f',3,'','','Saman','Shafigh'),
	('f9de7415-13c0-43e3-945e-f4ce714065f8',3,'','','Saman','Shafigh');
```