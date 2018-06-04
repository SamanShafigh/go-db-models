package main

import (
	"fmt"

	"github.com/SamanShafigh/go-db-models/model"
	"github.com/SamanShafigh/go-db-models/store"
	"github.com/SamanShafigh/go-db-models/util"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

func main() {
	mysqlDb, _ := sqlx.Connect("mysql", "user:pass@tcp(127.0.0.1:port)/dbName")

	userStore := store.NewUserStore(mysqlDb)

	user := model.MakeUser(util.MakeUUID(), "Saman", "Shafigh")
	userStore.Create(user)

	users, _ := userStore.List(
		model.UserFirstNameFilter("Saman"),
		model.UserFirstNameFilter("Shafigh"),
	)

	fmt.Printf("%+v", users)
}
