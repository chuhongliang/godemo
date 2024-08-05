package model

type Seed struct {
	Id          int    `json:"id"          gorm:"primaryKey"`
	Level       int    `json:"level"`
	Title       string `json:"title"`
	Imgurl      string `json:"imgurl"`
	GrowTime    int    `json:"grow_time"`
	Desc        string `json:"desc"`
	PlantId     int    `json:"plant_id"`
	PlantNumber int    `json:"plant_number"`
}

func (Seed) TableName() string {
	return "seed"
}
