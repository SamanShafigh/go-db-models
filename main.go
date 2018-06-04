package main

import (
	"fmt"
	"log"

	mysqld "github.com/SamanShafigh/go-db-models/db"
	"github.com/SamanShafigh/go-db-models/model"
	"github.com/SamanShafigh/go-db-models/store"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, _ := mysqld.Connect("user:pass@tcp(127.0.0.1:port)/dbName")
	userStore := store.NewUserStore(db)

	user1 := model.MakeUser(mysqld.MakeUUID(), "Saman", "Shafigh", 1)
	userStore.Add(user1)

	user2, err := userStore.GetByID("981894c2-af67-4204-957c-11bfebe5d346")
	fmt.Printf("%+v\n\n", user2)
	if err != nil {
		log.Fatal(err)
	}

	users, err := userStore.Get(
		model.UserStatusFilter(1),
		model.UserFirstNameFilter("Saman"),
		model.UserLastNameFilter("Shafigh"),
	)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%+v", users)
}
