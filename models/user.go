package models

type User struct {
	Id       uint
	Name     string
	Login    string `gorm:"unique"`
	Password []byte
}
