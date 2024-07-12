package main

import "github.com/gin-gonic/gin"

type Student struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	var studentSlice []Student
	studentSlice = append(studentSlice, Student{"001", "tom", 18})
	studentSlice = append(studentSlice, Student{"002", "jerry", 19})
	studentSlice = append(studentSlice, Student{"003", "tony", 120})

	r := gin.Default()

	r.Any("/students", func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
		c.JSON(200, studentSlice)
	})

	r.Run("localhost:5001")

}
