package test

import (
	"social_vascarsolutions/internal/entities"
	"social_vascarsolutions/internal/models"
	"social_vascarsolutions/internal/services"
	"testing"
)

func init() {
	entities.MigrateEntities()
}

func TestCreatePost(t *testing.T) {
	reportService := services.ReportService{}
	post := models.PostReport{
		Post:     "post 1",
		Reason:   "reason",
		Username: "Lenin1",
	}
	postEntity := reportService.CreatePost(post)
	if postEntity.ID == 0 {
		t.Errorf("TestCreatePost can not create post: %s", post)
	}
}

func TestCreateComment(t *testing.T) {
	reportService := services.ReportService{}
	comment := models.CommentReport{
		Comment:  "comment",
		Reason:   "reason",
		Username: "Lenin1",
	}
	postEntity := reportService.CreateComment(comment)
	if postEntity.ID == 0 {
		t.Errorf("TestCreatePost can not create comment: %s", comment)
	}
}
