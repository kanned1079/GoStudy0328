package main

import (
	"GoStudy0328/src/p46_CustomerInfosManagementSystem/model"
	"GoStudy0328/src/p46_CustomerInfosManagementSystem/services"
	"fmt"
)

// 这里才是有一个切片的实例

type customerView struct {
	// 定义必要的字段
	key             string // 接受用户的输入
	loop            bool   // 表示是否循环显示主菜单
	customerService *services.CustomerService
}

func (this *customerView) mainMenu() {
	for {
		fmt.Println("------- 客户信息管理系统 ------")
		fmt.Println("             1. 添加一个新的客户")
		fmt.Println("             2. 修改客户")
		fmt.Println("             3. 删除客户")
		fmt.Println("             4. 列出所有客户")
		fmt.Println("             5. 退出")
		fmt.Printf("\n请选择(1-5): ")

		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			this.add()
		case "2":
			fmt.Println("修改客户")
		case "3":
			this.del()
		case "4":
			this.list()
		//	fmt.Println("客户列表")
		case "5":
			this.loop = false
		default:
			fmt.Println("输入有误 请重新输入")
		}

		if !this.loop {
			break
		}
	}
	fmt.Println("已经退出软件...")
}

// 显示所有客户的信息
func (this *customerView) list() {
	// 首先获取所有客户信息 信息在切片中
	cus := this.customerService.List()
	fmt.Println("\n------- Customer List -------")
	fmt.Println("编号\t姓名\t性别\t年龄\t电话\t\t邮箱")
	for index := 0; index < len(cus); index++ {
		//fmt.Println()
		fmt.Println(cus[index].GetInfo())
	}
	fmt.Println("------- All customer list ended -------\n")
}

// 得到用户的输入 构建新的客户并加入到切片
func (this *customerView) add() {
	fmt.Println("\n------- Add Customer -------")
	var name, gender, tele, mail string
	var age int
	fmt.Print("姓名：")
	fmt.Scanln(&name)
	fmt.Print("性别：")
	fmt.Scanln(&gender)
	fmt.Print("年龄：")
	fmt.Scanln(&age)
	fmt.Print("电话：")
	fmt.Scanln(&tele)
	fmt.Print("邮箱：")
	fmt.Scanln(&mail)
	// 构建一个customer实例
	// 这里没有ID让用户输入 因为ID是唯一的 让系统自己分配
	customer := model.NewCustomer2(name, gender, age, tele, mail)
	if this.customerService.AddCustomer(customer) {
		fmt.Println("------- Success Added Customer -------\n")
	} else {
		fmt.Println("添加失败")
	}
}

// 得到用户的输入id，删除该id对应的客户
func (this *customerView) del() {
	fmt.Println("\n------- Delete customer -------")
	fmt.Print("输入要删除客户的Id：")
	var selectedId int = -1
	fmt.Scanln(&selectedId)
	if selectedId == -1 {
		fmt.Println("取消了")
		return
	}
	// 调用customerService的delete()方法
	fmt.Print("你确定要删除吗(y/n)：")
	var choose string = ""
	if fmt.Scanln(&choose); choose == "y" || choose == "Y" {
		if this.customerService.DeleteCustomer(selectedId) {
			fmt.Println("------- Success Deleted Customer -------\n")
		} else {
			fmt.Println("DELETE FAILURE")
		}
	} else {
		fmt.Println("操作取消")
	}

}

func main() {
	var cv1 customerView = customerView{
		key:  "",
		loop: true,
	}
	cv1.customerService = services.NewCustomerService() // 完成初始化
	// 在这里对字段 customerService *services.CustomerService 的初始化

	cv1.mainMenu()
}
