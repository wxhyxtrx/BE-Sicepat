package database

import (
	"fmt"
	"server/model"
	"server/pkg/connection"
)

func RunMigration() {
	err := connection.DB.AutoMigrate(
		&model.User{},
		&model.Chat{},
		&model.Room{},
		&model.Setting{},
		&model.Tiket{},
		&model.Order{},
	)

	if err != nil {
		fmt.Println(err)
		panic("Failed! Create Table to Database")
	}
	fmt.Println("Create Table Success")
}
