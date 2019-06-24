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
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	bodyString := string(bodyBytes)
	var userinfo db.User
	helper.UnmarshaUp(bodyString, &userinfo)
	//fmt.Println(userinfo)
	userinfo.Password=db.AesEncrypt(userinfo.Password, *db.SALT_KEY)
	if err:=db.DB.Save(&userinfo); err != nil{
		log.Printf("%s", err)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
		}
	w.WriteHeader(http.StatusOK)
	userinfo.Password=""
	execStatus, err := w.Write(helper.MarshaUp(userinfo))
	if err != nil {
		w.WriteHeader(execStatus)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	// defer func() {
	// 	err := recover()
	// 	if err != nil {
	// 		log.Printf("Internal error: %s", err)
	// 		var err_str string
	// 		var ok bool
	// 		if err_str, ok = err.(string); !ok {
	// 			err_str = "We encountered an internal error"
	// 		}
	// 		w.WriteHeader(http.StatusInternalServerError)
	// 		w.Write([]byte(err_str))
	// 	}
	// }()
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	temp := strings.Split(r.URL.Path, "/")
	username := temp[len(temp)-1]
	//fmt.Println(username)

	if err:= db.DB.Delete(&db.User{}, "username=?", username); err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	w.WriteHeader(http.StatusOK)
}


func ListUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")

	users := []db.User{}
	//query
	if err:= db.DB.Find(&users); err != nil && err.Error() != "record not found"{
		log.Printf("%s", err)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	for index := range users{
		users[index].Password=""
	}
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(users))
	if err != nil {
		w.WriteHeader(execStatus)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	// bodyBytes, err := ioutil.ReadAll(r.Body)
	// if err != nil {
	// 	log.Printf("%s", err)
	// 	w.WriteHeader(http.StatusForbidden)
	// 	_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
	// 	return
	// }
	// bodyString := string(bodyBytes)
	var userinfo db.User
	// helper.UnmarshaUp(bodyString, &userinfo)
	// userinfo.Password=db.AesEncrypt(userinfo.Password, *db.SALT_KEY)

	auth := strings.Replace(r.Header["Authorization"][0],"Basic ", "", 1)
	credential, _ := base64.StdEncoding.DecodeString(auth)
	userAndPassword := strings.Split(string(credential), ":")
	userAndPassword[1] = db.AesEncrypt(userAndPassword[1])


	// if err:= db.DB.Find(&userinfo, "username =? and password=?",userinfo.Username,userinfo.Password); err != nil{
	// 	log.Printf("%s", err)
	// 	w.WriteHeader(http.StatusUnauthorized)
	// 	_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
	// 	return
	// }
	if err:= db.DB.Find(&userinfo, "username =? and password=?",userAndPassword[0],userAndPassword[1]); err != nil{
		log.Printf("%s", err)
		w.WriteHeader(http.StatusUnauthorized)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	userinfo.Password=""
	w.WriteHeader(http.StatusOK)
	execStatus, err := w.Write(helper.MarshaUp(userinfo))
	if err != nil {
		w.WriteHeader(execStatus)
		// ignore error
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
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
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
	}
	bodyString := string(bodyBytes)
	//var result map[string]interface{}
	var userinfo db.User
	helper.UnmarshaUp(bodyString, &userinfo)
	//fmt.Println(userinfo)

	userinfo.Password=db.AesEncrypt(userinfo.Password, *db.SALT_KEY)
	if err := db.DB.Update(&userinfo, "username=?", username); err != nil {
		log.Printf("%s", err)
		w.WriteHeader(http.StatusForbidden)
		_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
		return
  }
	w.WriteHeader(http.StatusOK)
	// execStatus, err := w.Write(helper.MarshaUp(userinfo))
	// if err != nil {
	// 	w.WriteHeader(execStatus)
	// 	// ignore error
	// 	_, _ = w.Write(helper.MarshaUp(InlineResponse403{Error_: err.Error()}))
	// 	return
	// }
}







// func respSerialize(){
// 	vars := r.FormValue("id")
// 	if vars == ""{
// 		log.Printf("var is nil")
// 	}
// 	log.Printf("%v",vars)
// 	w.WriteHeader(http.StatusOK)
// }