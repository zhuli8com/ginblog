package v1

import (
	"fmt"
	"ginblog/model"
	"ginblog/utils/errmsg"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var code int

// 添加用户
func AddUser(c *gin.Context) {
	var data model.User
	err := c.ShouldBindJSON(&data)
	if err != nil {
		fmt.Println(err)
	}
	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.CreateUser(&data)
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"message":errmsg.GetErrMsg(code),
	})
}

// 查询单个用户
func GetUserInfo(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	data, code := model.GetUser(id)
	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"data":    data,
			"total":   1,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 查询用户列表
func GetUsers(c *gin.Context) {
	pageSize, _ := strconv.Atoi(c.Query("pagesize"))
	pageNum, _ := strconv.Atoi(c.Query("pagenum"))
	username := c.Query("username")

	switch {
	case pageSize >= 100:
		pageSize = 100;
	case pageSize <=0:
		pageSize = 10
	}

	if pageNum == 0{
		 pageNum = -1
	}

	data, total := model.GetUsers(username, pageSize, pageNum)
	code = errmsg.SUCCESS
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"data":data,
		"total":total,
		"message":errmsg.GetErrMsg(code),
	})
}

// 编辑用户
func EditUser(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.CheckUser(data.Username)
	if code == errmsg.SUCCESS {
		model.EditUser(id,&data)
	}
	if code == errmsg.ERROR_USERNAME_USED {
		c.Abort()
	}

	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}
// 修改密码
func ChangeUserPassword(c *gin.Context) {
	var data model.User
	id, _ := strconv.Atoi(c.Param("id"))
	_ = c.ShouldBindJSON(&data)

	code = model.ChangePassword(id, &data)

	c.JSON(
		http.StatusOK, gin.H{
			"status":  code,
			"message": errmsg.GetErrMsg(code),
		},
	)
}

// 删除用户
func DeleteUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	code = model.DeleteUser(id)
	c.JSON(http.StatusOK,gin.H{
		"status":code,
		"message":errmsg.GetErrMsg(code),
	})
}