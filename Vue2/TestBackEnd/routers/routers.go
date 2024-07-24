package routers

import (
	"GoStudy0328/Vue2/TestBackEnd/users"
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

func StartServer() {
	r := gin.Default()

	// 中间件
	r.Use(func(context *gin.Context) {
		context.Header("Access-Control-Allow-Origin", "*")
		context.Header("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		context.Header("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if context.Request.Method == "OPTIONS" {
			context.AbortWithStatus(http.StatusOK)
			return
		}
		context.Next()
	})

	// 自定义日志中间件
	r.Use(func(c *gin.Context) {
		// 记录请求的方法和路径
		log.Printf("Method: %s, Path: %s", c.Request.Method, c.Request.URL.Path)

		// 记录查询参数
		query := c.Request.URL.Query()
		for key, values := range query {
			log.Printf("Query param: %s = %s", key, values)
		}

		// 记录请求头
		for key, values := range c.Request.Header {
			log.Printf("Header: %s = %s", key, values)
		}

		// 如果需要记录请求体
		if c.Request.Method == "POST" || c.Request.Method == "PUT" {
			bodyBytes, err := ioutil.ReadAll(c.Request.Body)
			if err == nil {
				log.Printf("Body: %s", string(bodyBytes))
			}
			// 重新设置请求体，因为读取后它会被消耗
			c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		}

		// 继续处理请求
		c.Next()
	})

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

	r.GET("/api/user/get", HandleGetAllUsers)
	r.Any("/api/user/add", HandleAddUser)
	r.Any("/api/user/del", HandleDelUser)
	r.Any("/api/user/edit", HandleEditUser)

	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "Hello World")
	})

	if err := r.Run("localhost:5001"); err != nil {
		log.Println(err)
	}
}

func handleTableData() []users.TableData {
	var tdSlice []users.TableData
	tdSlice = append(tdSlice, users.TableData{"Huawei", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"Xiaomi", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"Samsung", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"Apple", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"OnePlus", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"Google", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"SONY", getRand(), getRand(), getRand()})
	tdSlice = append(tdSlice, users.TableData{"Nokia", getRand(), getRand(), getRand()})
	return tdSlice
}

func handleOrderData() []users.CountData {
	var countSlice []users.CountData
	countSlice = append(countSlice, users.CountData{"今日支付订单", getRand() / 3, "success", "#2ec7c9"})
	countSlice = append(countSlice, users.CountData{"今日收藏订单", getRand() / 3, "star-on", "#ffb980"})
	countSlice = append(countSlice, users.CountData{"今日未支付订单", getRand() / 3, "s-goods", "#5ab1ef"})
	countSlice = append(countSlice, users.CountData{"本月支付订单", getRand(), "success", "#2ec7c9"})
	countSlice = append(countSlice, users.CountData{"本月收藏订单", getRand(), "star-on", "#ffb980"})
	countSlice = append(countSlice, users.CountData{"本月未支付订单", getRand(), "s-goods", "#5ab1ef"})
	return countSlice
}

func handleVideoData() users.VideoData {
	var videoData users.VideoData
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

func handleColumnData() []users.ColumnData {
	var columnSlice []users.ColumnData
	columnSlice = append(columnSlice, users.ColumnData{"周一", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周二", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周三", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周四", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周五", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周六", getRandN(300), getRandN(1000)})
	columnSlice = append(columnSlice, users.ColumnData{"周日", getRandN(300), getRandN(1000)})
	return columnSlice
}

func handlePiData() []users.PiData {
	var piSlice []users.PiData
	piSlice = append(piSlice, users.PiData{"Xiaomi", getRand()})
	piSlice = append(piSlice, users.PiData{"Huawei", getRand()})
	piSlice = append(piSlice, users.PiData{"Apple", getRand()})
	piSlice = append(piSlice, users.PiData{"Google", getRand()})
	piSlice = append(piSlice, users.PiData{"SONY", getRand()})
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
