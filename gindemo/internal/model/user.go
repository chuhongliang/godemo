package model

import (
	"time"
)

type Extra struct {
	Guide int `json:"guide"`
}

type User struct {
	Id        int       `json:"id"          gorm:"primaryKey"`
	Username  string    `json:"username"    gorm:"unique"`
	Password  string    `json:"password"    gorm:"-"`
	Nickname  string    `json:"nickname"    gorm:"default:''"`
	Avatar    string    `json:"avatar"      gorm:"default:''"`
	Level     int       `json:"level"       gorm:"default:1"`
	Exp       int       `json:"exp"         gorm:"default:0"`
	Gold      int       `json:"gold"        gorm:"default:0"`
	Diamond   int       `json:"diamond"     gorm:"default:0"`
	LandLevel int       `json:"land_level"  gorm:"default:1"`
	LoginAt   time.Time `json:"login_at"    gorm:"autoCreateTime"`
	Extra     string    `json:"extra"       gorm:"type:json;not null"`
}

func (User) TableName() string {
	return "user"
}
