package model

import "time"

type Land struct {
	Id          int       `json:"id"            gorm:"primaryKey"`
	UserId      int       `json:"user_id"       gorm:"index"`
	Isunlock    bool      `json:"isunlock"      gorm:"type:boolean;default:false"`
	Position    int       `json:"position"      gorm:"default:1"`
	Free        bool      `json:"free"          gorm:"type:boolean;default:true"`
	Drought     bool      `json:"drought"       gorm:"type:boolean;default:false"`
	Level       int       `json:"level"         gorm:"default:1"`
	SeedId      int       `json:"seed_id"       gorm:"default:0"`
	SeedName    string    `json:"seed_name"     gorm:"default:''"`
	SeedGrowEnd time.Time `json:"seed_grow_end" gorm:"type:datetime; autoCreateTime"`
}

func (Land) TableName() string {
	return "land"
}
