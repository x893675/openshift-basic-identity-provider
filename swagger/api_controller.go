/*
 * openshift basic identity
 *
 * openshift basic identity provider
 *
 * API version: 1.0.0
 * Contact: zhu.xiaowei@99cloud.net
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"openshift-basic-identity-provider/db"
	"openshift-basic-identity-provider/helper"
	"strings"
)

func CreateUser(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bodyString := string(bodyBytes)
	//var result map[string]interface{}
	var userinfo db.User
	helper.UnmarshaUp(bodyString, &userinfo)
	fmt.Println(userinfo)
	err = db.Insert(userinfo)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(userinfo))
	if err != nil {
		w.WriteHeader(execStatus)
		// ignore error
		_, _ = w.Write(helper.MarshaUp(err))
		return
	}
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Internal error:", err)
			var err_str string
			var ok bool
			if err_str, ok = err.(string); !ok {
				err_str = "We encountered an internal error"
			}
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err_str))
		}
	}()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	temp := strings.Split(r.URL.Path, "/")
	username := temp[len(temp)-1]
	fmt.Println(username)
	err := db.Delete(username)
	if err != nil {
		fmt.Println("todo delete db error")
	}
	w.WriteHeader(http.StatusOK)
}

// func GetUserByName(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
// 	w.WriteHeader(http.StatusOK)
// }

func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	users, err := db.Query()
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(users))
	if err != nil {
		w.WriteHeader(execStatus)
		// ignore error
		_, _ = w.Write(helper.MarshaUp(err))
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bodyString := string(bodyBytes)
	var userinfo db.User
	helper.UnmarshaUp(bodyString, &userinfo)
	user, err := db.ValidatePassword(userinfo.Username, userinfo.Password)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		_, _ = w.Write(helper.MarshaUp(err))
		return
	}
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(user))
	if err != nil {
		w.WriteHeader(execStatus)
		// ignore error
		_, _ = w.Write(helper.MarshaUp(err))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	temp := strings.Split(r.URL.Path, "/")
	username := temp[len(temp)-1]
	fmt.Println(username)
	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	bodyString := string(bodyBytes)
	//var result map[string]interface{}
	var userinfo db.User
	helper.UnmarshaUp(bodyString, &userinfo)
	fmt.Println(userinfo)
	err = db.Update(username, userinfo)
	if err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(userinfo))
	if err != nil {
		w.WriteHeader(execStatus)
		// ignore error
		_, _ = w.Write(helper.MarshaUp(err))
		return
	}
}
