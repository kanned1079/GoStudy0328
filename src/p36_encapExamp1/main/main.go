package main

import (
	"GoStudy0328/src/p36_encapExamp1/model"
	"fmt"
)

func main() {
	//var p1 = model.NewPerson("張三")
	//p1.SetAge(34)
	//p1.SetSalary(5200.0)
	//
	//fmt.Println(p1.Name)
	//fmt.Println(p1.GetAge())
	//fmt.Println(p1.GetSalary())

	operateAccount()
}

func operateAccount() {
	acc1, err := model.NewAccount("kanna123")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("账户合法 创建成功")
	}
	if err = acc1.SetPassword("123456"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("密码合法")
	}
	if err = acc1.SetBalance(120, "123456"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("设置余额成功")
	}
	var balance float64 = 0
	if balance, err = acc1.GetBalance("123456"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("余额 = ", balance)
	}
}
