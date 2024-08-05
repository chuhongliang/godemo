package model

type Admin struct {
	Id       int    `json:"id"         gorm:"primaryKey" `
	Username string `json:"username"   gorm:"unique" `
	Password string `json:"password"   gorm:"-" `
}

func (Admin) TableName() string {
	return "admin"
}
