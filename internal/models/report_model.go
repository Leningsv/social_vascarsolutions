package models

type CommentReport struct {
	Username string `json:"username"  validate:"required"`
	Reason   string `json:"reason"  validate:"required,min=2,max=500"`
	Comment  string `json:"comment" validate:"required,min=2,max=500"`
}

type PostReport struct {
	Username string `json:"username"  validate:"required"`
	Reason   string `json:"reason"  validate:"required,min=2,max=500"`
	Post     string `json:"post" validate:"required,min=2,max=500"`
}
