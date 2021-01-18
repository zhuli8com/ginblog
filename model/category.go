package model

import (
	"ginblog/utils/errmsg"
	"gorm.io/gorm"
)

type Category struct {
	ID uint `gorm:"primary_key;auto_increment" json:"id"`
	Name string `json:"name"`
}

//查询分类是否存在
func CheckCategory(name string) (code int) {
	var cate Category
	db.Select("id").Where("name = ?",name).First(&cate)
	if cate.ID > 0 {
		return errmsg.ERROR_CATENAME_USED
	}
	return errmsg.SUCCESS
}

//新增分类
func CreateCate(data *Category) int {
	err := db.Create(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

//查询单个分类信息
func GetCateInfo(id int) (Category,int) {
	var cate Category
	db.Where("id = ?",id).First(&cate)
	return cate,errmsg.SUCCESS
}

//查询分类列表
func GetCate(pageSize int, pageNum int) ([]Category,int64) {
	var cate []Category
	var total int64
	err := db.Find(&cate).Limit(pageSize).Offset((pageNum - 1) * pageSize).Error
	db.Model(&cate).Count(&total)
	if err != nil && err != gorm.ErrRecordNotFound{
		return nil, 0
	}
	return cate, total
}

//编辑分类
func EditCate(id int, data *Category) int {
	var maps = make(map[string]interface{})
	maps["name"] = data.Name
	db.Model(&Category{}).Where("id = ?",id).Updates(maps)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func DeleteCate(id int) int {
	err := db.Where("id = ?", id).Delete(&Category{}).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}