package main

import (
	"fmt"
)

// 功能1 先完成可以顯示主菜單 並且可以輸入4退出程序
// 功能2 完成可以显示明细和登记收入的功能
//
//	因为需要显示明细 需要定义变量detail string来记录
//	定义变量来记录余额balance和每次收支金额monry每次收支说明note
//
// 功能3 完成登记支出的功能
func main() {
	// 聲明一個變量 保存用戶輸入的選項
	key := ""
	// 證明一個變量 控制是否要推出for循環
	loop := true
	balance, money := 10000.0, 0.0
	// 余额 每次收支金额
	var note string = string("")
	// 定义变量是否有过收支
	var flag bool = false
	var details string = string("\n收支\t账户金额\t收支金额\t说明\t") // 有收支时候只需要对这个进行拼接即可
	// 顯示主菜單
	for {
		fmt.Println("\n----------家庭收支記帳軟件----------")
		fmt.Println("           1. 收支明細")
		fmt.Println("           2. 登記收入")
		fmt.Println("           3. 登記支出")
		fmt.Println("           4. 退出軟件")
		fmt.Println("----------------------------------")
		fmt.Print("請選擇（1-4）：")
		fmt.Scanln(&key)
		switch key {
		case "1":
			{
				if flag == true {
					fmt.Println("\n----------當前收支明細----------")
					fmt.Println(details)
				} else {
					fmt.Println("还没有收入或支出")
				}

			}
		case "2":
			{
				fmt.Print("本次收入金额：")
				fmt.Scanln(&money)
				balance += money // 将money加入余额
				fmt.Print("本次收入说明：")
				fmt.Scanln(&note)
				// 进行拼接到details
				details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%s\t", balance, money, note)
				flag = true
			}
		case "3":
			{
				fmt.Print("本次支出金额：")
				fmt.Scanln(&money)
				if money > balance {
					fmt.Println("余额不足")
					break
				} else {
					balance -= money // 将money加入余额
					fmt.Print("本次支出说明：")
					fmt.Scanln(&note)
					// 进行拼接到details
					details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%s\t", balance, money, note)
					flag = true
				}

			}
		case "4":
			sig := ""
			fmt.Print("你确定要退出吗？(y/n): ")
			fmt.Scanln(&sig)
			if sig == "y" {
				loop = false
			} else {

			}

		default:
			fmt.Println("請輸入正確的選項")

		}
		if !loop {
			fmt.Println("已退出...")
			break
		}
	}

}
