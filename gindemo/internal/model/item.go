package model

type Item struct {
	Id         int    `json:"id"            gorm:"primaryKey" `
	UserId     int    `json:"user_id"       gorm:"default:0"`
	ItemType   int    `json:"item_type"     gorm:"default:0"`
	ItemId     int    `json:"item_id"       gorm:"default:0"`
	ItemTitle  string `json:"item_title"    gorm:"default:''"`
	ItemNumber int    `json:"item_number"   gorm:"default:0"`
	ItemImgurl string `json:"item_imgurl"   gorm:"default:''"`
	ItemDesc   string `json:"item_desc"     gorm:"default:''"`
}

func (Item) TableName() string {
	return "item"
}
