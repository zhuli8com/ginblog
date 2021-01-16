package model

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}