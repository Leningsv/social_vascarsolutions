package services

import (
	"social_vascarsolutions/internal/db"
	"social_vascarsolutions/internal/entities"
)

type UserService struct {
}

func (userService *UserService) CreateUser(user *entities.UserEntity) {
	db.Database.Create(user)
}

func (userService *UserService) GetUserByUsername(username string) entities.UserEntity {
	var user entities.UserEntity
	db.Database.Where(&entities.UserEntity{Username: username}).First(&user)
	return user
}
