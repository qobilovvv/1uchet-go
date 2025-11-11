package models

type User struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	PhoneNumber string `json:"phone_number" gorm:"phone_number"`
	Password    string `json:"password"`
}
