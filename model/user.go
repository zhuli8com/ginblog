package model

import (
	"fmt"
	"ginblog/utils/errmsg"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"log"
)

type User struct {
	ID uint
	Username string `json:"username"`
	Password string `json:"password"`
	Role int `json:"role"`
}

// 新增用户
func CreateUser(data *User) int {
	//data.Password = ScryptPw(data.Password) //->BeforeSave钩子函数实现
	err := db.Create(&data).Error
	if err != nil {
		fmt.Println(err.Error())
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 密码加密(钩子函数)https://gorm.io/zh_CN/docs/hooks.html
func (u *User) BeforeSave(_ *gorm.DB) (err error) {
	u.Password = ScryptPw(u.Password)
	return nil
}

// 查询用户是否存在
func CheckUser(name string) (code int) {
	var user User
	db.Select("id").Where("username = ?",name).First(&user)
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 更新查询
func CheckUpUser(id int, name string) (code int) {
	var user User
	db.Select("id,username").Where("username = ?",name).First(&user)
	if user.ID == uint(id) {
		return errmsg.SUCCESS
	}
	if user.ID > 0 {
		return errmsg.ERROR_USERNAME_USED
	}
	return errmsg.SUCCESS
}

// 查询用户
func GetUser(id int) (User,int) {
	var user User
	err := db.Where("id = ?", id).First(&user).Error
	if err != nil {
		return user,errmsg.ERROR
	}
	return user,errmsg.SUCCESS
}

// 查询用户列表
func GetUsers(username string,pageSize int,pageNum int) ([]User,int64) {
	var users []User
	var total int64

	if username != "" {
		db.Select("id,username,role").
			Where("username LIKE ?",username+"%").
			Limit(pageSize).
			Offset((pageNum - 1) *pageSize).
			Find(&users)
		db.Model(&users).Where("username LIKE ?",username+"%").Count(&total)
		return users,total
	}

	db.Select("id,username,role").Limit(pageSize).Offset((pageNum-1)*pageSize).Find(&users)
	db.Model(&users).Count(&total)

	if err != nil {
		return users,0
	}
	return users,total
}

// 编辑用户信息
func EditUser(id int, data *User) int {
	var user User
	var maps = make(map[string]interface{})
	maps["username"] = data.Username
	maps["role"] = data.Role//注意 当通过 struct 更新时，GORM 只会更新非零字段。 如果您想确保指定字段被更新，你应该使用 Select 更新选定字段，或使用 map 来完成更新操作
	err = db.Model(&user).Where("id = ?", id).Updates(maps).Error
	db.Where("id = ?",id)
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

func ChangePassword(id int, data *User) int {
	err := db.Select("password").Where("id = ?", id).Updates(&data).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 删除用户
func DeleteUser(id int) int {
	var user User
	err := db.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return errmsg.ERROR
	}
	return errmsg.SUCCESS
}

// 生成密码
func ScryptPw(password string) string {
	const cost = 10
	HashPw, err := bcrypt.GenerateFromPassword([]byte(password), cost)
	if err != nil{
		log.Fatal(err)
	}
	return string(HashPw)
}

// 后台登录验证

// 前台登录