package model

type User struct {
	Account  string `gorm:"primary_key;not_null;" json:"account"`
	Password string `gorm:"not_null" json:"password"`
	Name     string `gorm:"not_null" json:"name"`
}
