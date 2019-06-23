package db

import (
	"time"
	"openshift-basic-identity-provider/helper"
log "github.com/sirupsen/logrus"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
  )
  
var db_path = new(string)

  // key的长度为16,24,32位字符串
var salt_key string = "1234567887654321"
var DB Store

func init() {
	helper.SetLocalVar("DB_PATH", db_path, "./user.db")
	//log.Printf(*db_path)
}


type Store interface {
	SetLink(*gorm.DB)
	GetLink() (*gorm.DB)
	Save(interface{}) error
	Load(interface{}, int, int, ...interface{}) error
	Find(out interface{}, where ...interface{}) error
	Delete(in interface{}, where ...interface{}) error
	Update(in interface{}, query string, where ...interface{}) error
	CreateTable(interface{}) error
	Close()
}


func ConnectDB() (db *gorm.DB) {
	for {
		var err error
		db, err = gorm.Open("sqlite3", *db_path)
		db.DB().SetMaxOpenConns(100)
		db.DB().SetMaxIdleConns(10)
		db.DB().SetConnMaxLifetime(20*time.Second)
		if err != nil {
			log.Error("initDB error:", err)
			time.Sleep(10 * time.Second)
			continue
		}
		break
	}
	return db
}



//   func main() {
// 	db, err := gorm.Open("sqlite3", "/tmp/gorm.db")
// 	defer db.Close()
//   }
