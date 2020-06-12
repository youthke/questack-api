package conf

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	db *gorm.DB
	err error
)


func Init(){
	db, err = gorm.Open("mysql", conf.DB.User+":"+conf.DB.Password+"@tcp("+conf.DB.IP+")/questack?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db.BlockGlobalUpdate(true)
}

func GetDB() *gorm.DB{
	return db
}

