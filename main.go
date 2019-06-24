/*
 * openshift basic identity
 *
 * openshift basic identity provider
 *
 * API version: 1.0.0
 * Contact: zhu.xiaowei@99cloud.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package main

import (
	"net/http"

	// WARNING!
	// Change this to a fully-qualified import path
	// once you place this file into your project.
	// For example,
	//
	//    sw "github.com/myname/myrepo/go"
	//
	"openshift-basic-identity-provider/db"
	sw "openshift-basic-identity-provider/swagger"
	log "github.com/sirupsen/logrus"
)

func main() {
	log.Printf("Server started")

	router := sw.NewRouter()

	db.DB = &db.CRDB{DBLink: db.ConnectDB()}
	defer db.DB.Close()
	if err := db.DB.CreateTable(&db.User{}); err != nil {
                log.Error("create table user_ failed.", err)
                return
        }

	initTable()

	log.Fatal(http.ListenAndServe(":8080", router))
}


func initTable(){
	admin := db.User{
		Username: "admin",
		Password: "admin",
		Email: "admin@admin.com",
		Name: "admin"
	}
	developer := db.User{
		Username: "developer",
		Password: "developer",
		Email: "developer@developer.com",
		Name: "developer"
	}
	operator := db.User{
		Username: "operator",
		Password: "operator",
		Email: "operator@operator.com",
		Name: "operator"
	}
	_ := db.DB.Save(&admin)
	_ := db.DB.Save(&developer)
	_ := db.DB.Save(&operator)
}