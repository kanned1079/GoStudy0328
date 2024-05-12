package model

import "fmt"

type person struct {
	Name   string
	age    int
	salary float64
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {
	if age > 0 && age < 150 {
		p.age = age
	} else {
		fmt.Println("年齡範圍不正確")
		// 也可以給默認值
	}
}

func (p *person) GetAge() int {
	return p.age
}

func (p *person) SetSalary(sal float64) {
	if sal >= 3000 && sal <= 30000 {
		p.salary = sal
	} else {
		fmt.Println("薪水範圍不正確")
	}
}

func (p *person) GetSalary() float64 {
	return p.salary
}
