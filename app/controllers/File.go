// 文件操作相关
package controllers

import "os"

// IsDirExist is 文件目录是否存在
func IsDirExist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			// fmt.Println("文件夹存在")
			return true
		}
		if os.IsNotExist(err) {
			// fmt.Println("文件夹不存在")
			return false
		}
		// fmt.Println(err)
		// fmt.Println("需要创建文件")
		return false
	}
	return true
}
