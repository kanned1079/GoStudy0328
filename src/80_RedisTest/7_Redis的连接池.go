package main

// Redis链接池
// 通过Golang 对Redis操作 还可以通过Redis链接池
// 流程如下
// 1)事先初始化一定数量的链接 放入到链接池公
// 2)当Go需要操作Redis时 直接从Redis链接池取出链接即可
// 3)这样可以节省临时获取Redis链接的时间 从而提高效率.

// 使用set
// 存放商品信息
// 商品名
// 价格
// 生产日期

func main() {
	//var err error
	// 连接池核心代码
	/*
		var pool *redis.Pool
		pool = &redis.Pool{
			MaxIdle:     8,
			MaxActive:   0,
			IdleTimeout: 100,
			Dial: func() (redis.Conn, error) {
				return redis.Dial("tcp", "api.ikanned.com:16379")
			},
		}
		//c := pool.Get()
		pool.Close()
	*/

}
