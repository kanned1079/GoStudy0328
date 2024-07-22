package users

import "time"

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

type PiData struct {
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
