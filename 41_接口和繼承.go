package main

import "fmt"

type BirdAble interface {
	CanFly()
}

type FishAble interface {
	CanSwim()
}

type Monkey_41 struct {
	Name string
}

type LittleMonkey_41 struct {
	Monkey_41
}

func (this *Monkey_41) climbing() {
	fmt.Printf("monkey[%s] can climb\n", this.Name)
}

func (this *LittleMonkey_41) CanFly() {
	fmt.Printf("monkey[%s] can fly\n", this.Name)
}

func (this *LittleMonkey_41) CanSwim() {
	fmt.Printf("monkey[%s] can swim\n", this.Name)
}

func main() {
	var monk1 LittleMonkey_41 = LittleMonkey_41{
		Monkey_41{
			Name: "aa1",
		},
	}
	//monk1.Name = "A"
	monk1.climbing()
	monk1.CanFly()
	monk1.CanSwim()

}
