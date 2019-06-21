package main

import(
	"openshift-basic-identity-provider/db"
)

func main() {
	db.DB = &db.CRDB{DBLink: db.ConnectDB()}
	defer db.DB.Close()
	
	admin db.User = {Username: "admin"}
	db.DB.InitTableData(&admin, db.User{Username:"admin"})
}