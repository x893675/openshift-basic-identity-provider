package db

import (
	"errors"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	log "github.com/sirupsen/logrus"
)

//CDB is
type CRDB struct {
	DBLink *gorm.DB
}

func (crdb *CRDB) SetLink(link *gorm.DB) {
	crdb.DBLink = link
}

func (crdb *CRDB) GetLink() (*gorm.DB) {
	return crdb.DBLink
}

func (crdb *CRDB) Save(in interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	//create table and store
	if err := crdb.DBLink.Create(in).Error; err != nil {
		log.Error("store into db error:", err)
		return err
	}
	return nil
}

func (crdb *CRDB) Load(out interface{}, offset int, limit int, where ...interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	//load all data by time DESC
	//create_at is the timestamp
	if err := crdb.DBLink.Order("create_at DESC").Limit(limit).Offset(offset).Find(out, where...).Error; err != nil {
		log.Error("load from db error:", err)
		return err
	}
	return nil
}

//need to optimize
func (crdb *CRDB) Find(out interface{}, where ...interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	//authentication table
	err := crdb.DBLink.Find(out, where...).Error
	if err != nil {
		return err
	}

	return nil
}

func (crdb *CRDB) Delete(in interface{}, where ...interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	result := crdb.DBLink.Delete(in, where...)
	if result.Error != nil {
		log.Errorf("delete db record error: %v", result.Error)
		return result.Error
	}else if result.RowsAffected == 0 {
		//log.Errorf("record is not exist ",result.RowsAffected)
		return errors.New("record is not exist")
	}
	return nil
}

func (crdb *CRDB) Update(in interface{}, query string, where ...interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	result := crdb.DBLink.Model(in).Where(query, where...).Update(in)
	if result.Error != nil {
		log.Error("update db record failed.", err)
		return err
	}else if result.RowsAffected == 0 {
		return errors.New("record is not exist")
	}
	return nil
}

func (crdb *CRDB) CreateTable(in interface{}) error {

	if err := crdb.DBLink.DB().Ping(); err != nil {
		return err
	}

	if err := crdb.DBLink.AutoMigrate(in).Error; err != nil {
		log.Error("create table failed.", err)
		return err
	}
	return nil
}


func (crdb *CRDB) Close() {

	crdb.DBLink.DB().Ping()

	if err := crdb.DBLink.Close(); err != nil {
		log.Error("close datebase failed")
	}
}

