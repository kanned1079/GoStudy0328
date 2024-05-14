package main

import (
	"bufio"
	"gopkg.in/yaml.v2"
	"log"
	"os"
)

type Node79 struct {
	SSL    bool   `yaml:"ssl"`
	IpAddr string `yaml:"ipAddr"`
	Port   int    `yaml:"port"`
	Path   string `yaml:"path"`
}

type Server79 struct {
	ServerIpAddr string `yaml:"serverIpAddr"`
	ServerPort   int    `yaml:"serverPort"`
	Timeout      int    `yaml:"timeout"`
}

type Conf79 struct {
	Node79
	Server79
}

func main() {

	testNode1 := Node79{
		SSL:    true,
		IpAddr: "192.168.0.1",
		Port:   80,
		Path:   "/home/www",
	}

	server1 := Server79{
		ServerIpAddr: "192.168.0.1",
		ServerPort:   80,
		Timeout:      10,
	}

	conf1 := Conf79{
		testNode1,
		server1,
	}

	data, err := yaml.Marshal(conf1)
	if err != nil {
		log.Println("error:", err)
		return
	}
	log.Println(string(data))

	file, err := os.OpenFile("./testYaml.yaml", os.O_CREATE|os.O_RDWR, 0777)
	if err != nil {
		log.Println("error:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString(string(data))

	writer.Flush()
}
