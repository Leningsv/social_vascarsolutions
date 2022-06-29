package entities

import (
	"social_vascarsolutions/internal/db"
	"time"
)

type ReportEntity struct {
	ID        int64      `json:"id"`
	Reason    string     `json:"reason"`
	CreatedAt time.Time  `json:"createdAt"`
	UserId    int64      `json:"userId"`
	User      UserEntity `json:"user" gorm:"foreignKey:UserId"`
}

func MigrateReport() {
	db.Database.AutoMigrate(ReportEntity{})
}
