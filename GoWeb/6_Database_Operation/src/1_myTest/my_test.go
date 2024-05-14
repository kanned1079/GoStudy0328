package main

import (
	"log"
	"strconv"
	"testing"
)
import _ "github.com/go-sql-driver/mysql"

func TestTruncateData(t *testing.T) {
	if err := TruncateData(); err != nil {
		log.Fatal(err)
	}
}

func TestMyMain(t *testing.T) {

	var db1 DbInfo
	db1.ReadConf("./dbConf.yaml")
	InitDb(&db1)

	for i := 1; i <= 100; i++ {
		log.Println("---------", i, "------------")
		if err := InsertData(&Student{
			Name:    "testUser" + strconv.Itoa(i),
			Age:     18 + i,
			Email:   "test@aaaa.com" + strconv.Itoa(i),
			Address: "ShangHai" + strconv.Itoa(i),
		}); err != nil {
			t.Error(err)
		}
	}

	//stu1 := Student{
	//	Name:    "kanna",
	//	Age:     18,
	//	Email:   "dewffw@edfev.com",
	//	Address: "shanghai",
	//}
	//if err := InsertData(&stu1); err != nil {
	//	t.Error(err)
	//}

}
