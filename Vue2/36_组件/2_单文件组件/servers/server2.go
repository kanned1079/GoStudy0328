package main

import "github.com/gin-gonic/gin"

type Car struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Price int    `json:"price"`
}

func main() {
	var carSlice []Car
	carSlice = append(carSlice, Car{"001", "奔驰", 199})
	carSlice = append(carSlice, Car{"002", "马自达", 109})
	carSlice = append(carSlice, Car{"003", "捷达", 120})

	r := gin.Default()

	r.GET("/cars", func(c *gin.Context) {
		c.JSON(200, carSlice)
	})

	r.Run("localhost:5002")

}
