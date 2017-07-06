package mysql

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Info struct {
	Hostname string
	Database string
	User     string
	Password string
	Port     int
}

func (i *Info) dsn() string {
	// "user:password@/dbname?charset=utf8&parseTime=True&loc=Local"
	return i.User +
		":" +
		i.Password +
		"@/" +
		i.Database +
		"?charset=utf8&parseTime=True&loc=Local"
}

func (i *Info) Connect() (db *gorm.DB, err error) {
	db, err = gorm.Open("mysql", i.dsn())
	//defer db.Close()
	if err != nil {
		panic(err)
	}

	return
}
