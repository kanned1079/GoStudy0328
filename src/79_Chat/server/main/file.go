package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
	"time"
)

type Conf struct {
	// 下面是服务端的启动配置
	ServerProtocol string `yaml:"server_protocol"`
	ServerIpAddr   string `yaml:"server_ip_addr"`
	ServerPort     int    `yaml:"server_port"`
	// 下面是Redis的连接信息和Redis配置
	RedisProtocol    string        `yaml:"redis_protocol"`
	RedisIpAddr      string        `yaml:"redis_ip_addr"`
	RedisPort        int           `yaml:"redis_port"`
	RedisMaxIdle     int           `yaml:"redis_max_idle"`
	RedisMaxActive   int           `yaml:"redis_max_active"`
	RedisIdleTimeOut time.Duration `yaml:"redis_idle_time_out"`
}

func (this *Conf) ReadConfig(path string) (err error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0400)
	if err != nil {
		log.Println("配置文件打开出错或不存在 err: ", err)
		return
	}
	defer func() {
		if err = file.Close(); err != nil {
			log.Println("关闭配置文件错误 err: ", err)
		}
	}()
	dec := yaml.NewDecoder(file)
	if err = dec.Decode(&this); err != nil && err != io.EOF {
		log.Println("err: ", err)
	}

	fmt.Println(this)

	return
}

func (c *Conf) ReadConfigFromFile(path string) error {
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		log.Println("配置文件打开出错或不存在 err: ", err)
		return err
	}
	defer file.Close()

	dec := yaml.NewDecoder(file)
	if err := dec.Decode(&c); err != nil {
		return err
	}
	fmt.Println(c)

	return nil
}
