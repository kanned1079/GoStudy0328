package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)
import "github.com/go-resty/resty/v2"

type Area struct {
	Code     string     `json:"code"`
	Location []Location `json:"location"`
	Refer    Refer      `json:"refer"`
}
type Location struct {
	Name      string `json:"name"`
	ID        string `json:"id"`
	Lat       string `json:"lat"`
	Lon       string `json:"lon"`
	Adm2      string `json:"adm2"`
	Adm1      string `json:"adm1"`
	Country   string `json:"country"`
	Tz        string `json:"tz"`
	UtcOffset string `json:"utcOffset"`
	IsDst     string `json:"isDst"`
	Type      string `json:"type"`
	Rank      string `json:"rank"`
	FxLink    string `json:"fxLink"`
}
type Refer struct {
	Sources []string `json:"sources"`
	License []string `json:"license"`
}

const key = "844af2d5f3f0481f9492e4f36519fa3b"

type NowWeather struct {
	Code       string `json:"code"`
	UpdateTime string `json:"updateTime"`
	FxLink     string `json:"fxLink"`
	Now        Now    `json:"now"`
	Refer      Refer  `json:"refer"`
}
type Now struct {
	ObsTime   string `json:"obsTime"`
	Temp      string `json:"temp"`
	FeelsLike string `json:"feelsLike"`
	Icon      string `json:"icon"`
	Text      string `json:"text"`
	Wind360   string `json:"wind360"`
	WindDir   string `json:"windDir"`
	WindScale string `json:"windScale"`
	WindSpeed string `json:"windSpeed"`
	Humidity  string `json:"humidity"`
	Precip    string `json:"precip"`
	Pressure  string `json:"pressure"`
	Vis       string `json:"vis"`
	Cloud     string `json:"cloud"`
	Dew       string `json:"dew"`
}

func main() {
	r := gin.Default()

	r.GET("/api", func(context *gin.Context) {
		postCity := context.PostForm("city")
		log.Println("city:", postCity)
		weatherApiUrl := fmt.Sprintf("https://geoapi.qweather.com/v2/city/lookup?location=%s&key=%s", postCity, key)
		log.Println("url:", weatherApiUrl)

		client := resty.New()
		resp, err := client.R().SetResult(&Area{}).Get(weatherApiUrl)
		if err != nil {
			log.Println(err)
		}
		apiResp := resp.Result().(*Area)
		areaId := apiResp.Location[0].ID
		url := fmt.Sprintf("https://devapi.qweather.com/v7/weather/now?location=%s&key=%s", areaId, key)
		log.Println("url:", url)
		resp, err = client.R().SetResult(&NowWeather{}).Get(url)
		nowWeather := resp.Result().(*NowWeather)
		if err != nil {
			log.Println(err)
		}
		log.Println(nowWeather)
		context.JSON(http.StatusOK, nowWeather)

		//context.JSON(http.StatusOK, gin.H{
		//	"code": apiResp.Code,
		//	"id":   apiResp.Location[0].ID,
		//	"name": apiResp.Location[0].Name,
		//	"tz":   apiResp.Location[0].Tz,
		//})
	})

	r.Run("localhost:8080")

}
