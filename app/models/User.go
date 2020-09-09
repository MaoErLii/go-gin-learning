package models

import (
	"errors"
	"fmt"
	"gin-demo/database"
	"math/rand"
	"strconv"
	"time"
)

// User is ...
type User struct {
	ID   uint   `json:"id" form:"id"`
	Name string `json:"name" form:"name"`
}

// TableName is ... 设置表名
// gorm 表名的默认形式是构造体的复数形式，所有默认为users
func (User) TableName() string {
	return "user"
}

// AddUser is ...
// 新增用户
func (u *User) AddUser() (id uint, err error) {

	// 先判断库中是否存在该用户
	row := database.DB.Raw("SELECT id, name FROM user WHERE name = ?", u.Name).Row()
	var tempUser User
	row.Scan(&tempUser.ID, &tempUser.Name)

	// 表中不存在该用户
	if tempUser.Name == "" {

		// 生成用户id
		year, month, day := int(time.Now().Year()), int(time.Now().Month()), int(time.Now().Day())
		hour, minute, second := int(time.Now().Hour()), int(time.Now().Minute()), int(time.Now().Second())
		left := strconv.Itoa(year + month + day + hour + minute + second)
		rand.Seed(time.Now().UnixNano())
		random := strconv.Itoa(rand.Intn(10))
		key, _ := strconv.Atoi(left + random)
		var user User
		user.Name = u.Name
		user.ID = uint(key)
		// user.ID = 114514

		// 向表中插入新用户
		result := database.DB.Create(&user)

		// 插入失败
		if result.Error != nil {
			err = result.Error
			return
		}

		id = uint(key)
	} else {
		err = errors.New("用户已存在")
		return
	}

	return
}

// GetUser is ...
// 获取用户信息
func (u *User) GetUser() (user User, err error) {
	row := database.DB.Raw("SELECT id, name FROM user WHERE id = ?", u.ID).Row()

	row.Scan(&user.ID, &user.Name)

	return
}

// DeleteUser is ...
// 删除用户
func (u *User) DeleteUser() (user User, err error) {
	fmt.Println("要删除的用户信息", u)

	user, err = u.GetUser()

	if err != nil {
		return
	}

	if user.ID == 0 {
		err = errors.New("没有该用户")
		return
	}

	fmt.Println("查找的用户", user)

	result := database.DB.Where("name = ?", u.Name).Delete(&u)

	fmt.Println("删除结果", result)

	// tempU, err := u.GetUser()

	// fmt.Println("从库里查询的结果", tempU)

	return
}
