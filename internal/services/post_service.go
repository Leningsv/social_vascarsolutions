package services

import (
	"social_vascarsolutions/internal/db"
	"social_vascarsolutions/internal/entities"
)

type PostService struct {
}

func (postService *PostService) CreatePost(post *entities.PostEntity) {
	db.Database.Create(post)
}
