package test

import (
	"social_vascarsolutions/internal/entities"
	"social_vascarsolutions/internal/services"
	"testing"
)

func init() {
	entities.MigrateUser()
}

func TestCreateUser(t *testing.T) {
	userService := services.UserService{}
	var user = entities.UserEntity{Username: "Lenin"}
	userService.CreateUser(&user)
	if user.ID == 0 {
		t.Errorf("TestCreateUser invalid: %s", user.Username)
	}
}

func TestGetUserByUsername(t *testing.T) {
	userService := services.UserService{}
	username := "Lenin"
	var users = userService.GetUserByUsername(username)
	if users.ID == 0 {
		t.Errorf("TestGetUserByUsername username: %s", username)
	}
}
