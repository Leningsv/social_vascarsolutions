package db

import (
	"fmt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database = func() *gorm.DB {
	if auxDb, err := gorm.Open(sqlite.Open("social.db"), &gorm.Config{}); err != nil {
		fmt.Println("Connection error", err)
		panic(err)
	} else {
		fmt.Println("Success Connection")
		return auxDb
	}
}()
