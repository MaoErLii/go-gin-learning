package controllers

import (
	"fmt"
	"gin-demo/app/models"
	"log"
	"net/http"
	"regexp"
	"strconv"

	"github.com/gin-gonic/gin"
)

// UserInfo is ...
// 根据id获取用户信息
func UserInfo(c *gin.Context) {
	// 获取query中的id
	id := c.Query("id")
	// 判断是否有传id
	// 有传id
	if id != "" {
		// 判断字符串是否是数字
		pattern := "\\d+"

		result, _ := regexp.MatchString(pattern, id)

		// 字符串为数字
		if result {
			// string 转换为 uint
			intID, err := strconv.Atoi(id)
			uintID := uint(intID)

			// 类型转换失败
			if err != nil {
				c.JSON(http.StatusForbidden, gin.H{
					"code":    http.StatusForbidden,
					"message": err.Error(),
				})
			}

			// 从库中查找用户
			u := models.User{ID: uintID}

			res, err := u.GetUser()

			// 没有找到用户
			if err != nil {
				fmt.Println(err)
				log.Fatalln(err)
			}

			// 用户名不为空，有该id的用户
			if res.Name != "" {
				// fmt.Println(res)
				c.JSON(http.StatusOK, gin.H{
					"data": res,
				})
			} else {
				//	库中用户名为空，则没有该用户
				// fmt.Println(res)
				c.JSON(http.StatusForbidden, gin.H{
					"code":    http.StatusForbidden,
					"message": "用户不存在",
				})
			}
		} else {
			// id 为非数字
			c.JSON(http.StatusForbidden, gin.H{
				"code":    http.StatusForbidden,
				"message": "id 格式错误",
			})
		}
		// fmt.Println(result)

	} else {
		// 没有传id
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "no id",
		})
	}
}

// AddUser is function to add user in database which table' name is user
// 新增用户
func AddUser(c *gin.Context) {

	// 利用bind绑定json数据
	var form models.User

	err := c.BindJSON(&form)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
		})
		return
	}

	if form.Name == "" {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": "缺少用户名",
		})
		return
	}

	// 向数据库中新增用户
	id, err := form.AddUser()

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": err.Error(),
		})
		return
	}

	if id != 0 {
		fmt.Println("id", id)
		c.JSON(http.StatusOK, gin.H{
			"code":    http.StatusOK,
			"message": "新增成功",
		})
	}

	// 获取body中的数据
	// data, _ := ioutil.ReadAll(c.Request.Body)

	// var resp models.User
	// err := json.Unmarshal(data, &resp)

	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"code":    http.StatusForbidden,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// name := resp.Name

	// if name == "" {
	// 	fmt.Println("没有用户名")
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"code":    http.StatusForbidden,
	// 		"message": "缺少用户名",
	// 	})
	// } else {
	// 	// fmt.Println("要新增的用户是 \n", "name: ", name)
	// 	// 向数据库中新增用户
	// 	id, err := resp.AddUser()

	// 	if err != nil {
	// 		fmt.Println(err)
	// 		c.JSON(http.StatusForbidden, gin.H{
	// 			"code":    http.StatusForbidden,
	// 			"message": err.Error(),
	// 		})
	// 		return
	// 	}

	// 	if id != 0 {
	// 		c.JSON(http.StatusOK, gin.H{
	// 			"code":    http.StatusOK,
	// 			"message": "新增用户成功",
	// 		})
	// 	}
	// }
}

// DeleteUser is function to delete user in database which table name is user
func DeleteUser(c *gin.Context) {

	var form models.User

	err := c.BindJSON(&form)

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": err.Error(),
		})
		return
	}

	// fmt.Println("请求数据", "\n name: ", form.Name, "\n id:", form.ID)

	user, err := form.DeleteUser()

	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{
			"code":    http.StatusForbidden,
			"message": err.Error(),
		})
		return
	}

	fmt.Println("删除用户成功", user)

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "删除成功",
	})

	// fmt.Println("删除用户操作")
	// // 获取body数据
	// body, _ := ioutil.ReadAll(c.Request.Body)

	// fmt.Println("请求body", body)

	// var resp models.User
	// // 读取json数据
	// err := json.Unmarshal(body, &resp)

	// if err != nil {
	// 	fmt.Println(err)
	// 	c.JSON(http.StatusForbidden, gin.H{
	// 		"code":    http.StatusForbidden,
	// 		"message": err.Error(),
	// 	})
	// 	return
	// }

	// fmt.Println("请求数据", "\n name: ", resp.Name, "\n id:", resp.ID)
}
