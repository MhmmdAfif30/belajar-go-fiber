package migration

import (
	"fmt"

	"tutor-go-fiber/database"
	"tutor-go-fiber/models/entity"
)

func RunMigrate() {
	err := database.DB.AutoMigrate(
		&entity.User{},
		&entity.Messages{},
		&entity.About{},
	)
	if err != nil {
		panic(err)
	}
	fmt.Println("Success to migrate")
}
