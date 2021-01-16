package model

import "gorm.io/gorm"

type Article struct {
	gorm.Model
	Title string `json:"title"`
	Category Category `gorm:"foreignkey:Cid"`
	Cid int `json:"cid"`
	Desc string `json:"desc"`
	Content string `json:"content"`
	Img string `json:"img"`
}