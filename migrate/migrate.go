package main

import (
	"fmt"
	"go-test/db"
	"go-test/model"
)

func main() {
	dbConn, err := db.NewDB()
	if err != nil {
		fmt.Println(err)
	}
	defer fmt.Println("Successfully Migrated")
	defer db.CloseDB(dbConn)
	dbConn.AutoMigrate(&model.User{}, &model.Detail{})
}
