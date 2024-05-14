package main

import (
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"reflect"
	"time"
)

type ServerConfig struct {
	ServerProtocol string `yaml:"server_protocol"`
	ServerIpAddr   string `yaml:"server_ip_addr"`
	ServerPort     int    `yaml:"server_port"`
}

type RedisServerConfig struct {
	RedisProtocol    string        `yaml:"redis_protocol"`
	RedisIpAddr      string        `yaml:"redis_ip_addr"`
	RedisPort        int           `yaml:"redis_port"`
	RedisMaxIdle     int           `yaml:"redis_max_idle"`
	RedisMaxActive   int           `yaml:"redis_max_active"`
	RedisIdleTimeOut time.Duration `yaml:"redis_idle_time_out"`
}

type Conf struct {
	// 下面是服务端的启动配置
	ServerConfig `yaml:"server_config"`
	// 下面是Redis的连接信息和Redis配置
	RedisServerConfig `yaml:"redis_server_config"`
}

// ReadConfig 用于读取配置文件内容
func (this *Conf) ReadConfig(path string) (err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		log.Println("打开配置文件出错 err: ", err)
		return
	}
	defer func() { // 延迟关闭文件
		if err = file.Close(); err != nil {
			log.Println("关闭配置文件错误 err: ", err)
			return
		}
	}()
	dec := yaml.NewDecoder(file)                              // 先新建NewDecoder
	if err = dec.Decode(&this); err != nil && err != io.EOF { // 再创建实例 就像Unmarshel一样
		log.Println("解码文件错误 err: ", err)
		return
	}
	this.ShowConfDetails()
	return nil
}

// ShowConfDetails 用于显示具体信息
func (this *Conf) ShowConfDetails() {
	log.Println("需要的配置个数 =", reflect.TypeOf(this).Elem().NumField())
	log.Println("配置文件信息 =", this)
}
