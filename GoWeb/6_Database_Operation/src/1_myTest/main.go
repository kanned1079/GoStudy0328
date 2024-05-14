package main

import (
	"database/sql"
	"fmt"
	"gopkg.in/yaml.v2"
	"io"
	"log"
	"os"
)

type Student struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Email   string `json:"email"`
	Address string `json:"address"`
}

type DbInfo struct {
	DbProto string `yaml:"db_proto"`
	DbAddr  string `yaml:"db_addr"`
	DbPort  string `yaml:"db_port"`
	DbUser  string `yaml:"db_user"`
	DbPwd   string `yaml:"db_pwd"`
	DbName  string `yaml:"db_name"`
}

var (
	Db  *sql.DB
	err error
)

func (d *DbInfo) ReadConf(path string) {
	var err error
	file, err := os.OpenFile(path, os.O_RDONLY, 0444)
	if err != nil {
		log.Println("打开文件错误 err: ", err)
		return
	}
	defer func() {
		if err := file.Close(); err != nil {
			log.Println("关闭文件错误 err: ", err)
			return
		}
	}()
	decoder := yaml.NewDecoder(file)
	if err = decoder.Decode(&d); err != nil || err == io.EOF {
		log.Println("Read Ended. err: ", err)
	}
	log.Println(d)
}

func InitDb(db *DbInfo) {
	// 格式
	Db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@%s(%s:%s)/%s", db.DbUser, db.DbPwd, db.DbProto, db.DbAddr, db.DbPort, db.DbName))
	if err != nil {
		log.Println("打开数据库失败 err: ", err)
	} else {
		log.Println("打开数据库成功")

	}
}

func InsertData(s *Student) (err error) {
	sqlStr := "insert into students(name, age, email, address) values (?, ?, ?, ?)"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		log.Println("Db.Prepare err: ", err)
	}
	lines, err := stmt.Exec(s.Name, s.Age, s.Email, s.Address)
	log.Println("lines = ", lines)
	if err != nil {
		log.Println("Db.Exec err: ", err)
		return err
	}
	log.Println("Exec Success.")
	return nil
}

func TruncateData() (err error) {
	sqlStr := "truncate table students"
	stmt, err := Db.Prepare(sqlStr)
	if err != nil {
		log.Println("Db.Prepare err: ", err)
	}
	_, err = stmt.Exec()
	if err != nil {
		log.Println("Db.Exec err: ", err)
	}
	log.Println("Exec Success.")
	return nil
}

func GetSingleRow(fromId int) *Student {

	return nil
}
