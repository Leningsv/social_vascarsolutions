package services

import (
	"social_vascarsolutions/internal/db"
	"social_vascarsolutions/internal/entities"
	"social_vascarsolutions/internal/models"
)

type ReportService struct {
}

func (reportService *ReportService) CreatePost(post models.PostReport) entities.PostEntity {
	postService := PostService{}
	report := reportService.createReport(post.Username, post.Reason)
	postEntity := entities.PostEntity{
		Text:   post.Post,
		Report: report,
	}
	postService.CreatePost(&postEntity)
	return postEntity
}

func (reportService *ReportService) CreateComment(comment models.CommentReport) entities.CommentEntity {
	commentService := CommentService{}
	report := reportService.createReport(comment.Username, comment.Reason)
	commentEntity := entities.CommentEntity{
		Text:   comment.Comment,
		Report: report,
	}
	commentService.CreateComment(&commentEntity)
	return commentEntity
}

func (reportService *ReportService) createReport(username string, reason string) entities.ReportEntity {
	userService := UserService{}
	user := userService.GetUserByUsername(username)
	if user.ID == 0 {
		user.Username = username
		userService.CreateUser(&user)
	}
	report := entities.ReportEntity{
		Reason: reason,
		User:   user,
	}
	db.Database.Create(&report)
	return report
}
