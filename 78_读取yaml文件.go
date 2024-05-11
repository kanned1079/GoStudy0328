package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Conf struct {
	Protocol string `yaml:"protocol"`
	IpAddr   string `yaml:"ipaddr"`
	Port     int    `yaml:"port"`
}

type Config struct {
	Conf Conf `yaml:"conf"`
}

func ReadConfigFromFile(path string) (*Config, error) {
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		log.Println("配置文件打开出错或不存在 err: ", err)
		return nil, err
	}
	defer file.Close()

	var config Config
	dec := yaml.NewDecoder(file)
	if err := dec.Decode(&config); err != nil {
		log.Println("解析配置文件出错: ", err)
		return nil, err
	}

	return &config, nil
}

func main() {
	config, err := ReadConfigFromFile("./conf.yaml")
	if err != nil {
		log.Fatal("读取配置文件出错: ", err)
	}

	fmt.Println("配置信息：", config.Conf)
}
