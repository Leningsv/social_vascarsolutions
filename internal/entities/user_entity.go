package entities

import (
	"social_vascarsolutions/internal/db"
)

type UserEntity struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

type Users []UserEntity

func MigrateUser() {
	db.Database.AutoMigrate(UserEntity{})
}
