package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"reflect"
	"strconv"
)

type Monster3 struct {
	Name  string
	Age   int
	Skill string
}

func main() {
	var monst1 Monster3
	monst1.Name = "aaaaaa"
	monst1.Age = 650
	monst1.Skill = "aabcewvergeverefger"
	monst1.writeToRedis()
	monst1.readFormRedis()
}

func (m *Monster3) writeToRedis() {
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
	// 使用Hash进行存储
	for i := 0; i < reflect.TypeOf(m).Elem().NumField(); i++ {
		structName := reflect.TypeOf(m).Elem().Name()
		fieldName := reflect.TypeOf(m).Elem().Field(i).Name
		fieldValue := reflect.ValueOf(m).Elem().Field(i).Interface()

		_, err := conn.Do("Hset", structName, fieldName, fmt.Sprint(fieldValue))
		if err != nil {
			fmt.Println("redis hset fail")
		}
	}
}

func (m *Monster3) readFormRedis() {
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
	slices := make([]string, 0)
	slices, err = redis.Strings(conn.Do("HMget", "Monster3", "Name", "Age", "Skill"))
	if err != nil {
		fmt.Println("redis hget fail")
	}
	fmt.Println("redis hget ok")
	for i := range slices {
		fmt.Println(slices[i])
	}
	num, err := strconv.Atoi(slices[1])
	fmt.Println(num)
}
