package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

// LIST操作
// 列表是简单的字符串列表 按照插入顺序排序
// LIST的本质是一个链表
// LIST的元素是有序的
// 元素的值可以重复
// LPUSH city 上海 北京 天津 [element ...]
// LRANGE citys 0 -1 这个代表从第一个到最后一个元素
// 但是上面这个不会推出list里边的元素
// lpop citys //这个是会推出列表中的元素的 向左侧推出
// rpop citys // 和上面一样 但是是从右侧推出
// del citys 删除整个链表 因为头是一个指针 执行后遍历就会返回nil
// List的使用细节
// 1)lindex 按照索引下标获取元素（从左往右 从0 开始）
// 2)LLEN key 返回列表key的长度 如果key不存在 那么就会被解析成一个空标 返回0
// 如果所有元素被删除 那么这个链表也就消失了

func main() {
	var err error
	conn, err := redis.Dial("tcp", "api.ikanned.com:16379")
	if err != nil {
		fmt.Println("redis dial err:", err)
		panic(err)
	}
	defer conn.Close()

	flush(conn)

	_, _ = conn.Do("lpush", "cities1", "a", "b", "c") // 顺序是cba 相当于前插法
	_, _ = conn.Do("rpush", "cities2", "a", "b", "c") // 顺序是abc 相当于尾插法

	// 使用range只是遍历 不会修改里面的数据
	strs1, err := redis.Strings(conn.Do("lrange", "cities1", 0, -1))
	fmt.Println(strs1)

	strs2, err := redis.Strings(conn.Do("lrange", "cities2", 0, -1))
	for i := range strs2 {
		fmt.Println("strs2[i]:", strs2[i])
	}

	_, _ = conn.Do("rpush", "cities1", "d", "e", "f", "g")
	// 将会在cba后加入defg

	fmt.Println("---------------------------")
	str3, err := redis.String(conn.Do("lpop", "cities1")) // 执行这个后将会从左侧弹出一个元素 c
	fmt.Println(str3)
	strs4, err := redis.Strings(conn.Do("lrange", "cities1", 0, -1))
	fmt.Println(strs4) // 上面弹出c后这个List中就只有[b a]

	// 下面是一些List的使用细节和注意事项
	// 1) Lindex 按照索引获得下标元素
	elem1, err := redis.String(conn.Do("Lindex", "cities1", 0))
	fmt.Println("elem1:", elem1) // 在这里将会输出第一个元素
	// 2) Llen key 返回列表key的长度 如果key不存在 那么key就会被解析成一个空列表 返回0
	lenOfcitites1, err := redis.Int(conn.Do("Llen", "cities1"))
	fmt.Println("len(cities1) = ", lenOfcitites1) // 在这里将会返回 6

	// List的数据可以从左或从右插入
	// 如果将List元素全部删除 那么这个List就被删除了

}

func flush(conn redis.Conn) {
	_, _ = conn.Do("FLUSHALL")
}

func deleteList(conn redis.Conn, name string) {
	_, _ = conn.Do("DEL", name)
}
