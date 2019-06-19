package db

import (
	"database/sql"
	"log"
	"openshift-basic-identity-provider/helper"

	_ "github.com/mattn/go-sqlite3"
)

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

func Insert(userinfo User) error {
	tx, err := db_driver.Begin()
	if err != nil {
		return err
		//log.Fatal(err)
	}
	stmt, err := tx.Prepare("insert into user(username,password,email,name) values(?, ?, ?, ?)")
	if err != nil {
		return err
		//log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userinfo.Username, userinfo.Password, userinfo.Email, userinfo.Name)
	if err != nil {
		return err
	}

	tx.Commit()
	return nil
}

func Update(userinfo User) error {
	stmt, err := db_driver.Prepare("update user set (username,password,email,name) values(?,?,?,?) where username = ?")
	if err != nil {
		return err
		//log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(userinfo.Username, userinfo.Password, userinfo.Email, userinfo.Name, userinfo.Username)
	if err != nil {
		return err
	}
	return nil
}

func Delete(username string) error {
	stmt, err := db_driver.Prepare("delete from user where username = ?")
	if err != nil {
		return err
		//log.Fatal(err)
	}
	defer stmt.Close()

	_, err = stmt.Exec(username)
	if err != nil {
		return err
	}
	return nil
}

func Query() error {
	return nil
}

func CloseDB() {
	db_driver.Close()
}
