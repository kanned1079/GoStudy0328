package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
)

type Monster3 struct {
	Name  string
	Age   int
	Skill string
}

func main() {
	var err error
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("redis dial fail")
		return
	}
	fmt.Println("redis dial ok")
	defer func() {
		if err := conn.Close(); err != nil {
			fmt.Println("conn close fail")
		}
	}()

	//flushRedis(conn)
	//// 要求从键盘输入三个Monster信息
	sliceMyMonster := make([]Monster3, 3)
	for i := 0; i < 3; i++ {
		var name, skill string
		var age int
		fmt.Print("输入Name  > ")
		fmt.Scanln(&name)
		fmt.Print("输入Age   > ")
		fmt.Scanln(&age)
		fmt.Print("输入Skill > ")
		fmt.Scanln(&skill)
		sliceMyMonster[i] = NewMonster3(name, age, skill)
		sliceMyMonster[i].writeToRedis(conn)
		readFormRedis(conn, sliceMyMonster[i].Name)
	}
}

func NewMonster3(name string, age int, skill string) Monster3 {
	return Monster3{
		Name:  name,
		Age:   age,
		Skill: skill,
	}
}

func (m *Monster3) writeToRedis(conn redis.Conn) {
	// 使用Hash进行存储
	for i := 0; i < reflect.TypeOf(m).Elem().NumField(); i++ {
		structName := reflect.TypeOf(m).Elem().Name()
		fieldName := reflect.TypeOf(m).Elem().Field(i).Name
		fieldValue := reflect.ValueOf(m).Elem().Field(i).Interface()

		_, err := conn.Do("Hset", structName+m.Name, fieldName, fmt.Sprint(fieldValue))
		if err != nil {
			fmt.Println("执行HMGET失败 ", err)
		}
	}
}

func readFormRedis(conn redis.Conn, key string) {
	var err error
	slices := make([]string, 0)
	slices, err = redis.Strings(conn.Do("HMget", "Monster3"+key, "Name", "Age", "Skill"))
	if err != nil {
		fmt.Println("执行HMGET失败 ", err)
	}
	fmt.Println("执行HMGET成功")
	for i := range slices {
		fmt.Println(slices[i])
	}
}

func flushRedis(conn redis.Conn) {
	if _, err := conn.Do("FLUSHALL"); err != nil {
		fmt.Println("刷新失败 ", err)
	} else {
		fmt.Println("刷新成功")
	}
}
