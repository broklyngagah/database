package database

import (
	"fmt"
	as "github.com/aerospike/aerospike-client-go"
	"github.com/broklyngagah/database/aerospike"
	"github.com/broklyngagah/database/mongo"
	"github.com/broklyngagah/database/mysql"
	"github.com/jinzhu/gorm"
	"gopkg.in/mgo.v2"
)

const (
	mongoDbDriver string = "mongodb"
	mysqlDbDriver string = "mysql"
	aeroDbDriver  string = "aero"
)

type ProviderInterface interface {
	Connect() error
}

type (
	DbInfo struct {
		MongoInfo mongo.Info
		MysqlInfo mysql.Info
		AeroInfo  aerospike.Info
	}

	Engine struct {
		MongoEngine *mgo.Session
		Mongo       *mgo.Database
		Mysql       *gorm.DB
		AeroEngine  *as.Client
	}
)

func (cfg *DbInfo) Handler(driver string) (db Engine, err error) {
	switch driver {
	case mongoDbDriver:
		dbInfo := cfg.MongoInfo
		db.MongoEngine, err = dbInfo.Connect()
		defer db.MongoEngine.Close()
		if err != nil {
			fmt.Println("Can not connect to MongoDB")
			return
		}

		warmUp := db.MongoEngine.Copy()
		db.Mongo = warmUp.DB(dbInfo.Database)

		return
	case mysqlDbDriver:
		dbInfo := cfg.MysqlInfo
		db.Mysql, err = dbInfo.Connect()
		if err != nil {
			fmt.Println("Cannot connect to mysql database")
			return
		}

		return

	case aeroDbDriver:
		dbInfo := cfg.AeroInfo
		db.AeroEngine, err = dbInfo.Connect()
		if err != nil {
			fmt.Println("Cannot connect to aero database")
			return
		}

		return
	}

	return
}
