package main

import (
	"github.com/gin-gonic/gin"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

type TableData struct {
	Name     string `json:"name"`
	TodayBuy int    `json:"todayBuy"`
	MonthBuy int    `json:"monthBuy"`
	TotalBuy int    `json:"totalBuy"`
}

type CountData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
	Icon  string `json:"icon"`
	Color string `json:"color"`
}

type VideoData struct {
	Data []map[string]int `json:"data"`
	Date []string         `json:"date"`
}

type piData struct {
	Name  string `json:"name"`
	Value int    `json:"value"`
}

//
//type Data struct {
//	Items map[string]int `json:"items"`
//}

type ColumnData struct {
	Date   string `json:"date"`
	New    int    `json:"new"`
	Active int    `json:"active"`
}

type User struct {
	Name  string    `json:"name"`
	Age   int       `json:"age"`
	Sex   int       `json:"sex"` // 1男 0女
	Birth time.Time `json:"birth"`
	Addr  string    `json:"addr"`
}

func main() {
	r := gin.Default()

	r.GET("/api/home/getData", func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.JSON(http.StatusOK, gin.H{
			"tableData":  handleTableData(),
			"orderData":  handleOrderData(),
			"videoData":  handleVideoData(),
			"columnData": handleColumnData(),
			"piData":     handlePiData(),
		})
	})

	r.Run(":5001")
}

func handleTableData() []TableData {
	var tdSlice []TableData
	tdSlice = append(tdSlice, TableData{"Huawei", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"Xiaomi", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"Samsung", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"Apple", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"OnePlus", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"Google", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"SONY", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, TableData{"Nokia", getRand(), getRand(), getRand()})
	return tdSlice
}

func handleOrderData() []CountData {
	var countSlice []CountData
	countSlice = append(countSlice, CountData{"今日支付订单", getRand() / 3, "success", "#2ec7c9"})
	countSlice = append(countSlice, CountData{"今日收藏订单", getRand() / 3, "star-on", "#ffb980"})
	countSlice = append(countSlice, CountData{"今日未支付订单", getRand() / 3, "s-goods", "#5ab1ef"})
	countSlice = append(countSlice, CountData{"本月支付订单", getRand(), "success", "#2ec7c9"})
	countSlice = append(countSlice, CountData{"本月收藏订单", getRand(), "star-on", "#ffb980"})
	countSlice = append(countSlice, CountData{"本月未支付订单", getRand(), "s-goods", "#5ab1ef"})
	return countSlice
}

func handleVideoData() VideoData {
	var videoData VideoData
	for i := 0; i < 7; i++ {
		var dataMap = make(map[string]int)
		dataMap["Apple"] = getRand()
		dataMap["Huawei"] = getRand()
		dataMap["Xiaomi"] = getRand()
		dataMap["SONY"] = getRand()
		dataMap["Samsung"] = getRand()
		dataMap["Google"] = getRand()
		dataMap["OnePlus"] = getRand()
		videoData.Data = append(videoData.Data, dataMap)

	}
	for i := 1; i <= 7; i++ {
		str := "2019100" + strconv.Itoa(i)
		videoData.Date = append(videoData.Date, str)
	}
	return videoData
}

func handleColumnData() []ColumnData {
	var columnSlice []ColumnData
	columnSlice = append(columnSlice, ColumnData{"周一", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周二", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周三", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周四", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周五", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周六", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, ColumnData{"周日", getRandN(300), getRandN(1000)})
	return columnSlice
}

func handlePiData() []piData {
	var piSlice []piData
	piSlice = append(piSlice, piData{"Xiaomi", getRand()})
	piSlice = append(piSlice, piData{"Huawei", getRand()})
	piSlice = append(piSlice, piData{"Apple", getRand()})
	piSlice = append(piSlice, piData{"Google", getRand()})
	piSlice = append(piSlice, piData{"SONY", getRand()})
	return piSlice
}

func getRand() int {
	return rand.Intn(10000)
}

func getRandN(x int) int {
	return rand.Intn(x)
}

func getStatisticalData() {

}
