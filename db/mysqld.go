package util

import (
	"encoding/json"
	"fmt"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	uuid "github.com/satori/go.uuid"
)

type DB struct {
	*sqlx.DB
}

type Tx struct {
	*sqlx.Tx
}

// Connect makes a connection to DB
func Connect(dataSourceName string) (*DB, error) {
	db, _ := sqlx.Connect("mysql", dataSourceName)

	return &DB{db}, nil
}

// Query makes a DB query
func (db *DB) Query(query string, filter interface{}) (*sqlx.Rows, error) {
	f, _ := getQueryString(filter)
	if f != "" {
		query = fmt.Sprintf("%s WHERE %s", query, f)
	}

	return db.NamedQuery(query, filter)
}

// MakeUUID returns a random generated UUID
func MakeUUID() string {
	uuid := uuid.Must(uuid.NewV4())

	return uuid.String()
}

// getQueryString generates a Mysql query string
func getQueryString(filter interface{}) (string, error) {
	b, err := json.Marshal(filter)
	if err != nil {
		return "", err
	}

	m := make(map[string]interface{})
	json.Unmarshal(b, &m)

	query := []string{}
	for key, val := range m {
		if key != "limit" && val != nil {
			query = append(query, fmt.Sprintf("%s = :%s", key, key))
		}
	}

	return strings.Join(query, " AND "), nil
}
