package model

type Task struct {
	Id     int    `json:"id"          gorm:"primaryKey"`
	Title  string `json:"title"`
	Desc   string `json:"desc"`
	Step   int    `json:"step"`
	Type   int    `json:"type"`
	Reward string `json:"reward" gorm:"type:json"`
}
