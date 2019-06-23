package main

import(
	"fmt"
	log "github.com/sirupsen/logrus"
	"openshift-basic-identity-provider/db"
)

func main() {
	db.DB = &db.CRDB{DBLink: db.ConnectDB()}
	defer db.DB.Close()
	if err := db.DB.CreateTable(&db.User{}); err != nil {
                log.Error("create table user_ failed.", err)
                return
        }
	  // Create
	 if err:=db.DB.Save(&db.User{Username:"zxw",Password:"haha"}); err != nil{
	 log.Error("db save error",err)
	 return 
	 }

	  result := db.User{}
	  //query
	  if err:= db.DB.Find(&result, "username=?", "zxw"); err != nil && err.Error() == "record not found"{
		  log.Error("no record ",err)
		  return
	  }
	 fmt.Println("%v",result)

	  //update
	  if err := db.DB.Update(&db.User{Username: "zxw", Password:"hahaha", Email:"dfjdkf@djfkdj.com"}, "username=?", "zxw"); err != nil {
	  log.Error("db update error",err)
			return
		}
	// delete
	  if err:= db.DB.Delete(&db.User{}, "username=?", "zxw"); err != nil {
		  log.Error(err)
	  return
	  }
	  //fmt.Println("%v",result)
  }

