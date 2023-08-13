package handler

import (
	"log"
	"sync"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

const dsn = "host=localhost user=arvan password=arvan dbname=arvan port=5432 sslmode=disable"

var doOnce sync.Once
var singleconn *gorm.DB

func GetConnection(destination string) *gorm.DB {
	doOnce.Do(func() {
		db, err := gorm.Open(postgres.Open(destination), &gorm.Config{NamingStrategy: schema.NamingStrategy{SingularTable: true}, NowFunc: func() time.Time {
			return time.Now().Local()
		}})
		if err != nil {
			log.Fatalln("GetConnection Error is: ", err)
		}

		singleconn = db
	})
	return singleconn
}
