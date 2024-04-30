package main

import (
	"fmt"
	"strconv"
)

// channle本质就是数据是先进先出
// 1)Channel的本质上就是一个数据结构 -> 队列
// 2)数据是先进先出 [FIFO]
// 3)线程安全，多goroutine访问时，不需要加锁，就是说channel本身就是线程安全的
// 4)channel是有类型的，一个string的channel只能存放string类型数据

// Channel的声明/定义
// channel是引用类型
// channel必须初始化后才能使用
// channel是有类型的

func main() {
	// 创建一个可以存放3个int类型的管道
	var intChan chan int
	intChan = make(chan int, 3)

	fmt.Printf("intChan = %v, &intChan = %v\n", intChan, &intChan)
	// intChan = 0x140000b8000, &intChan = 0x140000aa018

	// 向管道写入数据
	intChan <- 10
	numb := 20
	intChan <- numb

	intChan <- 30
	num2 := <-intChan
	fmt.Println("num2 = ", num2) // 先进先出 那么这里出来的就是10
	intChan <- 40
	// 这里会报错 fatal error: all goroutines are asleep - deadlock!
	// 但是只要取了一个就又可以装了

	// 从Channel中读取数据
	// 输出intChan的length和cap
	fmt.Println("len(intChan) = ", len(intChan))
	fmt.Println("cap(intChan) = ", cap(intChan))
	// 在没有使用协程的情况下 如果Channel的数据已经全部取出 再取就会报告deadlock

	fmt.Println(<-intChan)
	fmt.Println(<-intChan)
	fmt.Println(<-intChan) // 到这里都是可以取到的
	//fmt.Println(<-intChan) // 到这里会报错因为已经取完了数据

	//test63_1()
	//test63_2()
	//test63_3()
	//test63_4()

	afterClassExam1()

}

func test63_1() {
	// 创建一个Channel 最多存放3个int 演示存入和读取
	var intChannel1 chan int
	intChannel1 = make(chan int, 3)
	for i := 1; i <= cap(intChannel1); i++ {
		intChannel1 <- i * 10
		fmt.Println("data = ", i*10)
	}
	fmt.Println("len(intChannel1) = ", len(intChannel1))
	fmt.Println("cap(intChannel1) = ", cap(intChannel1))
	for i := 1; i <= cap(intChannel1); i++ {
		fmt.Printf("intChannel1[%d] = %d\n", i, <-intChannel1)
	}
}

func test63_2() {
	fmt.Println("-----------------------------")
	// 创建一个mapChan 最多可以存放10个map[string]string的kyy-value
	var mapChan1 chan map[string]string
	mapChan1 = make(chan map[string]string, 10)
	m1 := make(map[string]string, 1)
	for i := 1; i <= cap(mapChan1); i++ {
		m1[fmt.Sprintf("%s", "city")] = "上海"
		mapChan1 <- m1
	}
	fmt.Println(len(m1))
	fmt.Println("len(mapChan1) = ", len(mapChan1))
	fmt.Println("cap(mapChan1) = ", cap(mapChan1))
	for i := 1; i <= cap(mapChan1); i++ {
		fmt.Printf("mapChan1[%02d] = %v\n", i, <-mapChan1)
	}
}

type Person63 struct {
	Name string
	Age  int
}

func test63_3() {
	var personChan1 chan Person63
	personChan1 = make(chan Person63, 3)
	p1 := Person63{
		Name: "kanna",
		Age:  20,
	}
	p2 := Person63{
		Name: "kinggyo",
		Age:  16,
	}
	personChan1 <- p1
	personChan1 <- p2
	fmt.Printf("personChan1.len = %d, personChan1.cap = %d\n", len(personChan1), cap(personChan1))

	fmt.Println("person = ", <-personChan1)
	fmt.Println("person = ", <-personChan1)
	fmt.Printf("personChan1.len = %d, personChan1.cap = %d\n", len(personChan1), cap(personChan1))

}

type Cat63 struct {
	Name string
	Age  int
}

func test63_4() {
	fmt.Println("--------------------------")
	var allChan chan interface{}
	allChan = make(chan interface{}, 10)

	a := 10
	d := "Hello"

	allChan <- a
	allChan <- d
	allChan <- Cat63{"neko1", 1}
	allChan <- Cat63{"neko2", 3}
	fmt.Printf("allChan1.len = %d, allChan1.cap = %d\n", len(allChan), cap(allChan))
	//var nums int = len(allChan)
	//for i := 0; i < nums; i++ {
	//	fmt.Println("allchan = ", <-allChan)
	//}
	fmt.Printf("allChan1.len = %d, allChan1.cap = %d\n", len(allChan), cap(allChan))
	// 如果需要获取第三个元素 那么就先要推出前两个
	<-allChan
	<-allChan
	//newCat1 := (<-allChan).(Cat63)
	//newCat1 := <-allChan
	//fmt.Println("newCat1 = ", newCat1) // 这样能用
	// 但是不能取出字段
	// 因为这个时候newCat1还只是一个interface{}
	//fmt.Println("newCat1's Name = ", newCat1.Name)
	// 那么就需要进行类型断言
	newCat2 := (<-allChan).(Cat63)
	fmt.Printf("newCat2's Name = %s, Age = %d\n", newCat2.Name, newCat2.Age)

}

type Person63_2 struct {
	Name    string
	Age     int
	Address string
}

func afterClassExam1() {
	// 创建一个 Person结构体 [Name, Age,Address]
	// 使用rand方法配合随机创建10个Person实例，并放入到channel中3)遍历channel
	// 将各个Person实例的信息显示在终端
	var personChannel chan Person63_2
	personChannel = make(chan Person63_2, 10)
	for i := 0; i < 10; i++ {
		personChannel <- Person63_2{"p" + strconv.Itoa(i), 12, "addr" + strconv.Itoa(i)}
	}
	fmt.Println("len = ", len(personChannel))
	fmt.Println("cap = ", cap(personChannel))
	for i := 0; i < cap(personChannel); i++ {
		fmt.Println("personChannel[", i, "] = ", <-personChannel)
	}
	close(personChannel)

}
