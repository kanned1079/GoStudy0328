package main

// 1.Redis 是 NOSQL数据库， 不是传统的关系型数据库官网: https://redisio/和 http://www.redis.cn
// 2.Redis:Remote Dlctionary Server(远程字典服务器)，Redis性能非常高，单机能够达到15w qps，通常适合做缓存，也可以持久化。
// 3.是完全开源免费的，高性能的(key/value)分布式内存数据库，基于内存运行并支持持久化的NoSQL数据库，是最热门的NoSql数据库之一,也称为数据结构服务器

// Redis安装好后 有16个默认的数据库
// 添加key-val [set]
// 查看当前redis的 所有key [keys *]
// 获取key对应的值. [get key]
// 切换redis 数据库 [select index]
// 如何查看当前数据库的key-val数量[dbsize]
// 清空当前数据库的key-val和清空所有数据库的key-val [flushdb flushall]

// Redis支持的数据类型
// String字符串 Hash哈希 List列表 Set集合 Zset有序集合

// Crud操作
// string是最基本的数据类型 一个key对应一个value
// string是二进制安全的 因此可以存放图片等数据
// 字符串的value最大512M一个

// SET addreee 北京天安门
// DELETE address

// 值保留30s
//setex name 30 name

// 一次设置获得多个值
// MSET name1 tom name2 jerry
// MGET name1 name2 在golang中这个会返回一个切片

// Hash操作
// 通常可以用来存放姐沟通
// user1 name "Smith" age 30 job "coding"
// name 张三 age就是两对fieldMap
// HSET user1 name "kanna" age 30 job "engineer" 这个30存进去也是字符串了
// HGET user1 job
// HMSET user2 name jerry age 16 job "java coder"
// 判断是否存在
// HEXISTS user2 name
// 统计一个哈希有多少个键值
// HLEN user1

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

// Set操作
// 底层是HashTable数据结构 Set也是存放很多的String 但是类型默认是无序的 而且数据是不能重复的
// 比如用来存放邮箱地址
// SADD email 111@a.com 222@a.com 111@a.com 这样来添加元素
// SMEMBERS email 这是读取集合中的元素
// SISMEMBER email 111@a.com 判断集合中有没有这个
// SREM email 111@a.com 这
