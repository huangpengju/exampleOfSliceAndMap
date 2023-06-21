// user 包用来初始化用户，以及用户的增删改查
package user

import (
	"exampleOfSliceAndMap/menu"
	"fmt"
)

// init 用来初始化用户默认的值，创建3个map
func init() {
	// 创建3个map
	client01 := map[string]string{"id": "1", "name": "张三", "sex": "男", "age": "21", "email": "zhangsan@qq.com"}
	client02 := map[string]string{"id": "2", "name": "李四", "sex": "女", "age": "22", "email": "lisi@qq.com"}
	client03 := map[string]string{"id": "3", "name": "王五", "sex": "男", "age": "21", "email": "wangwu@qq.com"}

	// 构建一个map的切片（map是切片的类型，key为一个string类型，value为一个string类型）
	// var a= []int{1,2,3}  // 字面量声明一个切片
	var clinets = []map[string]string{client01, client02, client03}
	ClinetFunc(menu.TopTitle, menu.AddClientStart, clinets, menu.Success)
}

// clientFunc  主页列表
// topTitle 是系统的菜单标题
// addClientStart 是添加客户开始时的小标题
// clients 是初始化的3个用户
// success 是添加客户完成时的小标题
func ClinetFunc(topTitle string, addClientStart string,
	clients []map[string]string, success string) {
	// 输出系统的菜单
	fmt.Println(topTitle)

	// 输出操作提示
	fmt.Print("请输入选择[1-5]：")
	// 定义个变量存放选择结果,进一步判断客户的操作
	var nums int

	// 定义的变量存放客户的信息数据
	var name, sex, age, email string
	fmt.Scan(&nums)

	// 添加客户操作
	switch nums {
	// 添加客户信息
	case 1:
		AddClient(topTitle, addClientStart, clients, success, name, sex, age, email)
	// 查看客户信息
	case 2:
		GetClientList(clients, topTitle, addClientStart, success, menu.ClientListStart, menu.ClientListEnd)
		// 更新客户信息
	case 3:
		UpdateClient(topTitle, addClientStart, success, clients, sex, age, email)
	// 删除客户信息
	case 4:
		DeleteClient(topTitle, addClientStart, success, clients)
	// 退出客户
	case 5:
		fmt.Println("谢谢您的使用！")
		// break
	}

}
