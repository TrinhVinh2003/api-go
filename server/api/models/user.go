package models

type User struct {
	Id       uint   `json:"id" gorm:"primaryKey"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password []byte `json:"password"`
}
