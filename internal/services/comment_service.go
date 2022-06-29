package services

import (
	"social_vascarsolutions/internal/db"
	"social_vascarsolutions/internal/entities"
)

type CommentService struct {
}

func (commentService *CommentService) CreateComment(comment *entities.CommentEntity) {
	db.Database.Create(comment)
}
