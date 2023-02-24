package service

import (
	"fmt"
	"ginchat/models"
	"ginchat/utils"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/asaskevich/govalidator"
	"github.com/gin-gonic/gin"
)

type User struct{}

// GetUserList
// @Tags 用户接口
// @Summary 用户列表
// @Success 200 {json} {"code", "message"}
// @Router /user/getUserList [get]
func (u User) GetUserList(c *gin.Context) {
	data := models.GetUserList()
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "",
		"data":    data,
	})
}

// GetUser
// @Tags 用户接口
// @Summary 用户信息
// @Accept  application/x-www-form-urlencoded
// @Product application/json
// @Param name formData string true "用户姓名"
// @Param password formData string true "输入密码"
// @Success 200 {json} {"code", "message"}
// @Router /user/getUserInfo [post]
func (u User) GetUser(c *gin.Context) {
	name := c.PostForm("name")
	password := c.PostForm("password")

	findUserByName := models.FindUserByname(name)
	validPassword := utils.ValidPassword(password, findUserByName.Salt, findUserByName.PassWord)
	if findUserByName.ID < 1 || !validPassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "账号密码错误",
			"data":    "",
		})
		return
	}
	userInfo := models.FindUserBynameAndPwd(name, findUserByName.PassWord)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "用户信息",
		"data":    userInfo,
	})
}

// CreateUser
// @Tags 用户接口
// @Summary 创建用户
// @Accept  application/x-www-form-urlencoded
// @Product application/json
// @Param name formData string true "用户姓名"
// @Param password formData string true "输入密码"
// @Param repassword formData string true "再次输入密码"
// @Param phone formData string true "电话号码"
// @Param email formData string true "邮箱"
// @Success 200 {json} {"code", "message"}
// @Router /user/createUser [post]
func (u User) CreateUser(c *gin.Context) {
	user := models.UserBasic{}
	if models.FindUserByname(c.PostForm("name")).Name != "" {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "用户已存在",
			"data":    user,
		})
		return
	}
	password := c.PostForm("password")
	repassword := c.PostForm("repassword")
	if password != repassword {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": "两次密码输入不一致",
			"data":    user,
		})
		return
	}
	rand.Seed(time.Now().UnixMicro())
	salt := fmt.Sprintf("%d", rand.Int31())
	user.Salt = salt
	user.PassWord = utils.MakePassword(password, salt)
	user.Name = c.PostForm("name")
	user.Phone = c.PostForm("phone")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
			"data":    user,
		})
		return
	}
	models.CreateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "创建用户成功",
		"data":    user,
	})
}

// UpdateUser
// @Tags 用户接口
// @Summary 修改用户信息
// @Accept  application/x-www-form-urlencoded
// @Product application/json
// @Param id formData uint true "用户ID"
// @Param name formData string true "用户姓名"
// @Param newPassword formData string true "新密码"
// @Param phone formData string true "电话号码"
// @Param email formData string true "邮箱"
// @Success 200 {json} {"code", "message"}
// @Router /user/updateUser [put]
func (u User) UpdateUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(id)
	user.Name = c.PostForm("name")
	user.Phone = c.PostForm("phone")
	user.PassWord = c.PostForm("newPassword")
	user.Email = c.PostForm("email")
	_, err := govalidator.ValidateStruct(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"code":    -1,
			"message": err,
			"data":    user,
		})
		return
	}
	models.UpdateUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "修改用户信息成功",
		"data":    user,
	})
}

// DeleteUser
// @Tags 用户接口
// @Summary 删除用户
// @Param id query uint true "用户ID"
// @Success 200 {json} {"code", "message"}
// @Router /user/deleteUser [delete]
func (u User) DeleteUser(c *gin.Context) {
	user := models.UserBasic{}
	id, _ := strconv.Atoi(c.Query("id"))
	user.ID = uint(id)
	models.DeleteUser(&user)
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "删除用户成功",
		"data":    user,
	})
}
