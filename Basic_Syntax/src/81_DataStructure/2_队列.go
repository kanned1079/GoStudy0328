package main

import (
	"errors"
	"fmt"
	"log"
)

// 使用数组实现队列的思路
// 1)创健一个数组 arrary,是作为队列的一个字段
// 2)front初始化为-1
// 3)real,表示队列尾部，初始化为-1
// 4)完成队列的基本查找

const maxSize = 4

type SeqQueue struct {
	maxSize int
	array   [4]int //使用数组模拟队列
	front   int    // 表示指向队列的前端
	rear    int    // 表示指向队列的尾部
}

// AddQueue 添加数据到队列
func (this *SeqQueue) AddQueue(val int) (err error) {
	// 先判断队列是否为满
	if this.rear == maxSize-1 { // 含有最后一个元素
		log.Println("Queue is full.")
		return errors.New("队列满")
	}
	this.rear++
	this.array[this.rear] = val
	return nil
}

// ShowQueue 显示队列
func (this *SeqQueue) ShowQueue() {
	// 先找到队首 再遍历到队列尾
	// this.front 不含有队列首
	fmt.Println("队列元素为：")
	for i := this.front + 1; i <= this.rear; i++ {
		fmt.Printf("array[%d] = %d\n", i, this.array[i])
	}
}

// 从队列中取出数据
func (this *SeqQueue) GetQueue() (val int, err error) {
	// 先判断队列是否为空
	if this.front == this.rear {
		log.Println("Queue is empty..")
		return -1, errors.New("队列空")
	}
	this.front++
	return this.array[this.front], nil
}

func main() {
	var Q SeqQueue
	Q.maxSize = maxSize
	Q.AddQueue(1)
	Q.AddQueue(2)
	Q.AddQueue(3)
	Q.ShowQueue()
	fmt.Println("-----------")
	u1, _ := Q.GetQueue()
	fmt.Println(u1)
	u1, _ = Q.GetQueue()
	fmt.Println(u1)
	u1, _ = Q.GetQueue()
	fmt.Println(u1)
	u1, _ = Q.GetQueue()
	fmt.Println(u1)
}
