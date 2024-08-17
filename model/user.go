package model

type User struct {
	ID       uint   `json:"id" gorm:"primaryKey"`
	Password string `json:"password"`
}
