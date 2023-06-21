package user

import (
	"fmt"
	"strconv"
	"strings"
)

// 添加客户
func AddClient(topTitle string, addClient string, clients []map[string]string, success string,
	name, sex, age, email string) []map[string]string {
	// 显示添加客户时的小标题
	fmt.Println(addClient)
	fmt.Print("请输入客户的姓名：")
	fmt.Scan(&name)
	if len(name) != 0 {
		fmt.Print("请输入客户的性别：")
		fmt.Scan(&sex)
		if len(sex) != 0 {
			fmt.Print("请输入客户的年龄：")
			fmt.Scan(&age)
			if len(age) != 0 {
				fmt.Print("请输入客户的邮箱：")
				fmt.Scan(&email)
				// 获取切片的长度
				var incrementID = len(clients)
				// 新的id=原切片的长度+1
				id := strconv.Itoa(incrementID + 1)
				// 构建一个map用于添加新的数据
				newMap := map[string]string{"id": id, "name": name, "sex": sex, "age": age, "email": email}
				// 给切片添加元素
				clients = append(clients, newMap)
				//显示添加成功的小标题
				fmt.Println(success)
				// 显示所有的客户信息
				for _, clientMap := range clients {
					fmt.Printf("编号：%-8s姓名：%-8s性别：%-8s年龄：%-8s邮箱：%-8s", clientMap["id"],
						clientMap["name"], clientMap["sex"], clientMap["age"], clientMap["email"])
					if len(clientMap["email"]) != 0 {
						fmt.Println()
					}
				}
				// 返回到主页的操作
				homeClient(topTitle, addClient, success, clients)
			}
		}

	}
	return clients
}

// homeClient返回到主页[Y/N]
func homeClient(topTitle, addClient, success string, clients []map[string]string) {
	fmt.Println("请问是否返回上一次[Y/N]：")
	var isFlag string
	fmt.Scan(&isFlag)
	upper := strings.ToUpper(isFlag)
	switch upper {
	case "Y":
		fmt.Println("返回到主页")
		// 主页列表
		ClinetFunc(topTitle, addClient, clients, success)
	case "N":
		return
	default:
		fmt.Println("您的输入有误！请重新输入")
		homeClient(topTitle, addClient, success, clients)
	}
}

// 删除客户操作
func DeleteClient(topTitle, addClient, success string, clients []map[string]string) {
	fmt.Println("请输入你要删除的用户姓名")
	var name string
	fmt.Scan(&name)
	var deleteIndex = 0
	for index, clientMap := range clients {
		if clientMap["name"] == name {
			deleteIndex = index
		}
	}
	// 追加元素操作（删除）
	clients = append(clients[:deleteIndex], clients[1+deleteIndex:]...)
	fmt.Println("用户信息已删除！")
	// 主页列表
	ClinetFunc(topTitle, addClient, clients, success)
}

// 更新客户信息
func UpdateClient(topTitle, addClient, success string, clients []map[string]string, sex, age, email string) []map[string]string {
	fmt.Print("请输入需要更新的客户姓名：")
	var name, id string
	fmt.Scan(&name)
	var deleteIndex = 0
	if len(name) != 0 {
		for index, clientMap := range clients {
			if clientMap["name"] == name {
				deleteIndex = index
				// 删除元素操作
				clients = append(clients[:deleteIndex], clients[1+deleteIndex:]...)
				fmt.Println("需要更新用的id:", deleteIndex+1)
				id = strconv.Itoa(deleteIndex + 1)
			}
		}
		fmt.Print("请输入新的姓名:")
		fmt.Scan(&name)
		if len(name) != 0 {
			fmt.Print("请输入新的客户性别：")
			fmt.Scan(&sex)
			if len(sex) != 0 {
				fmt.Print("请输出新的客户年龄：")
				fmt.Scan(&age)
				if len(age) != 0 {
					fmt.Print("请输入新的客户邮箱：")
					fmt.Scan(&email)
					// 构建一个新的切片元素
					newMap := map[string]string{"id": id, "name": name, "sex": sex, "age": age, "email": email}
					clients = append(clients, newMap)
				}
			}
		}
		fmt.Println("更新成功!")
		// 显示主页列表
		ClinetFunc(topTitle, addClient, clients, success)
	}
	return clients
}

// 查询用户操作
func GetClientList(clients []map[string]string, topTile, addClient, success, clientListStart, clientListEnd string) {
	fmt.Println(clientListStart)
	for _, clientMap := range clients {
		fmt.Printf("编号：%-8s姓名：%-8s性别：%-8s年龄：%-8s邮箱：%-8s", clientMap["id"], clientMap["name"], clientMap["sex"], clientMap["age"], clientMap["email"])
		if len(clientMap["email"]) != 0 {
			fmt.Println()
		}
	}
	fmt.Println(clientListEnd)

	// homeClient返回到主页[Y/N]
	homeClient(topTile, addClient, success, clients)
}
