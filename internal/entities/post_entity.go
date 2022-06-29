package entities

import (
	"social_vascarsolutions/internal/db"
	"time"
)

type PostEntity struct {
	ID        int64        `json:"id"`
	Text      string       `json:"text"`
	CreatedAt time.Time    `json:"createdAt"`
	ReportId  int64        `json:"reportId"`
	Report    ReportEntity `json:"report" gorm:"foreignKey:ReportId"`
}

func MigratePost() {
	db.Database.AutoMigrate(PostEntity{})
}
