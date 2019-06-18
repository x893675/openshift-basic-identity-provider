package db

import (
	"database/sql"
	"log"
	"openshift-basic-identity-provider/helper"

	_ "github.com/mattn/go-sqlite3"
)

var user_cols = []string{"username", "password", "email", "name"}

var db_path = new(string)

var db_driver *sql.DB

const createTable string = `CREATE TABLE IF NOT EXISTS user(
	id integer not null primary key, 
	username text not null unique,
	password text not null,
	email text,
	name text
	);`

func init() {
	helper.SetLocalVar("DB_PATH", db_path, "./user.db")
	log.Printf(*db_path)
}

func InitDB() {
	var err error
	db_driver, err = sql.Open("sqlite3", "test.db")
	if err != nil {
		log.Fatal(err)
	}
	_, err = db_driver.Exec(createTable)
	if err != nil {
		log.Fatal(err)
	}
	// sqlStmt := `Describe user;`
	// _, err = db_driver.Exec(sqlStmt)
	// if err != nil {
	// 	log.Fatal(err)
	// }
}

func Insert(userinfo map[string]string) error {
	for _, v := range user_cols {
		if _, ok := userinfo[v]; ok {

		}
	}
	return nil
}

func Update(userinfo map[string]string) error {
	return nil
}

func Delete(user string) error {
	return nil
}

func CloseDB() {
	db_driver.Close()
}
