package utils

import (
	"GoStudy0328/Basic_Syntax/src/79_Chat/common/message"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"net"
)

// 将这些方法关联到结构体 中
type Transfer struct {
	// 分析需要那些字段
	Conn   net.Conn   // 连接
	Buffer [8096]byte // 缓冲
}

// readPkg 用于读取数据包
func (this *Transfer) ReadPkg() (mes message.Message, err error) {
	//buffer := make([]byte, 8096)
	fmt.Println("读取客户端发送的数据...")
	// 这里的conn只有在没有被关闭的情况下 才会阻塞
	// 如果客户端关闭了conn就不会阻塞了
	_, err = this.Conn.Read(this.Buffer[:4]) // 先读取前面四个字节
	if err != nil {
		fmt.Println("read err:", err) // 这里是读取头部长度出错
		//err = errors.New("read pkg header err")
		// 进行错误处理
		return
	}
	fmt.Println("读取到buf长度 = ", this.Buffer[:4])
	// 将这个长度转换为uint32
	var pkgLen uint32
	pkgLen = binary.BigEndian.Uint32(this.Buffer[0:4]) // 这里是Uint32转回来 前面是Putuint32
	// 根据pkgLen来读取内通
	n, err := this.Conn.Read(this.Buffer[:pkgLen]) // 从conn套接字中读取[:pkgLen]长度的内容到buffer里
	if n != int(pkgLen) || err != nil {            // 第一种可能读到了但是没有读全 还有一种就是对方关闭了管道
		fmt.Println("read err:", err) // 这里是读取内容出错
		//err = errors.New("read pkg body err")
		return
	} // 到这里buffer里面就有收到的消息了 但是还没完 这需要反序列化
	err = json.Unmarshal(this.Buffer[:pkgLen], &mes) // 注意这里要加上& 很重要 很重要 很重要
	if err != nil {
		fmt.Println("反序列化失败")
		return
	}
	return
}

// writePkg 用于向客户端再发送数据 包含与客户端发送相似逻辑
// 过来的信息已经是字节了就不需要再转换成字符串了
func (this *Transfer) WritePkg(data []byte) (err error) {
	// 1.先发送长度给客户端
	var pkgLen uint32
	pkgLen = uint32(len(data))
	//var buffer [4]byte
	binary.BigEndian.PutUint32(this.Buffer[0:4], pkgLen)
	n, err := this.Conn.Write(this.Buffer[:4])
	if n != 4 || err != nil {
		fmt.Println("server write err:", err)
		return
	}
	// 2. 发送data消息本身
	n, err = this.Conn.Write(data)
	if n != int(pkgLen) || err != nil {
		fmt.Println("server write err:", err)
		return
	}
	return
}
