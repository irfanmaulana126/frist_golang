package models

type User struct {
	Id      uint   `form:"id" json:"id"`
	Name    string `form:"name" json:"name" binding:"required"`
	Email   string `form:"email" json:"email" binding:"required"`
	Phone   string `form:"phone" json:"phone" binding:"required"`
	Address string `form:"address" json:"address" binding:"required"`
}

func (b *User) TableName() string {
	return "user"
}
