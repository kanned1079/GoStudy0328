package services

import (
	"GoStudy0328/src/p46_CustomerInfosManagementSystem/model"
)

// 这里有一个客户的切片
type CustomerService struct {
	// 这里有增删改查操作
	customers []model.Customer
	// 声明一个字段表示当前切片有多少个客户
	customerNumber int // 这个也可以作为新客户的编号
}

func NewCustomerService() *CustomerService {
	// 先初始化一个客户
	customerService := &CustomerService{}
	customerService.customerNumber = 1
	customer := model.NewCustomer(1, "张三", "男", 20, "112234", "zhangsan@xxx.com")
	customerService.customers = append(customerService.customers, customer)
	return customerService
}

// 编写List() 返回客户信息 其实就是切片
func (this *CustomerService) List() []model.Customer {
	return this.customers
}

// 添加客户到切片中
// 这里一定要使用 this *CustomerService
// 这里如果没有使用指针 那么每次新建客户的时候 老的客户就看不见了
func (this *CustomerService) AddCustomer(customer model.Customer) bool {
	// 确定一个规则 分配Id
	this.customerNumber++
	customer.Id = this.customerNumber
	this.customers = append(this.customers, customer)
	return true
}

func (this *CustomerService) FindTargetId(index int) (i int) {
	i = -1
	// 遍历切片 查找Id
	for a := 0; a < len(this.customers); a++ {
		if this.customers[a].Id == index {
			// 找到了
			i = a
		}
	}
	return
}

func (this *CustomerService) DeleteCustomer(id int) bool {
	// 这里的删除是从切片中删除
	index := this.FindTargetId(id)
	if index == -1 {
		// -1表示没有这个客户
		return false
	} else {
		// 从切片中删除一个元素
		// 这里的[:index]表示从0开时到index-1
		this.customers = append(this.customers[:index], this.customers[index+1:]...)
		return true
	}
}
