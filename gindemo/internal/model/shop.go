package model

type Shop struct {
	Id         int    `json:"id"          gorm:"primaryKey"`
	ItemType   int    `json:"item_type"`
	ItemId     int    `json:"item_id"`
	ItemTitle  string `json:"item_title"`
	ItemImgurl string `json:"item_imgurl"`
	ItemDesc   string `json:"item_desc"`
	ItemPrice  int    `json:"item_price"`
}

func (Shop) TableName() string {
	return "shop"
}
