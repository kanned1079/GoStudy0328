package main

import "os"

const filePath = "./log.txt"

type Customer struct {
	Name string
}

func (c *Customer) Store() bool {
	var err error
	file, err := os.OpenFile(filePath, os.O_CREATE)
}

func main() {

}
