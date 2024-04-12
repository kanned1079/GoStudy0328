package model

import "fmt"

// 这里含有每个客户的信息
type Customer struct {
	Id        int
	Name      string
	Gender    string
	Age       int
	Phone_num string
	Email     string
}

// 工厂模式创建用户
func NewCustomer(id int, name, gender string, age int, phone_num, mail string) Customer {
	return Customer{
		Id:        id,
		Name:      name,
		Gender:    gender,
		Age:       age,
		Phone_num: phone_num,
		Email:     mail,
	}
}

// 这是第二个不带Id的
func NewCustomer2(name, gender string, age int, phone_num, mail string) Customer {
	return Customer{
		Name:      name,
		Gender:    gender,
		Age:       age,
		Phone_num: phone_num,
		Email:     mail,
	}
}

func (this *Customer) GetInfo() string {
	info := fmt.Sprintf("%v\t%v\t%v\t%v\t%v\t\t%v", this.Id, this.Name, this.Gender, this.Age, this.Phone_num, this.Email)
	return info
}
