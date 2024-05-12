package myUtils

import "fmt"

// 先聲明一個結構體
type FamilyAccount struct {
	// 聲明一個變量 保存用戶輸入的選項
	key string
	// 證明一個變量 控制是否要推出for循環
	loop    bool
	balance float64
	money   float64
	// 余额 每次收支金额
	note string
	// 定义变量是否有过收支
	flag bool
	//var details string = string("\n收支\t账户金额\t收支金额\t说明\t") // 有收支时候只需要对这个进行拼接即可
	details string
}

// 給結構體綁定相對應的方法
// 顯示主菜單
func (this *FamilyAccount) MainMenu() {
	//this.balance = 10000.0
	this.details = "\n收支\t账户金额\t收支金额\t说明\t"
	this.loop = false
	this.flag = false
	for {
		fmt.Println("\n----------家庭收支記帳軟件----------")
		fmt.Println("           1. 收支明細")
		fmt.Println("           2. 登記收入")
		fmt.Println("           3. 登記支出")
		fmt.Println("           4. 退出軟件")
		fmt.Println("----------------------------------")
		fmt.Print("請選擇（1-4）：")
		fmt.Scanln(&this.key)
		switch this.key {
		case "1":
			{
				this.ShowDetails()
			}
		case "2":
			{
				this.RegisterIncome()
			}
		case "3":
			{
				this.PayMoney()
			}
		case "4":
			{
				this.Exit()
			}
		default:
			fmt.Println("請輸入正確的選項")

		}
		if this.loop {
			fmt.Println("已退出...")
			//break
			return
		}
	}
}

// 顯示明細
func (this *FamilyAccount) ShowDetails() {
	if this.flag {
		fmt.Println("\n----------當前收支明細----------")
		fmt.Println(this.details)
	} else {
		fmt.Println("还没有收入或支出")
	}
}

// 登記收入
func (this *FamilyAccount) RegisterIncome() {
	fmt.Print("本次收入金额：")
	fmt.Scanln(&this.money)
	this.balance += this.money // 将money加入余额
	fmt.Print("本次收入说明：")
	fmt.Scanln(&this.note)
	// 进行拼接到details
	this.details += fmt.Sprintf("\n收入\t%v\t\t%v\t\t%s\t", this.balance, this.money, this.note)
	this.flag = true
}

// 支出
func (this *FamilyAccount) PayMoney() {
	fmt.Print("本次支出金额：")
	fmt.Scanln(&this.money)
	if this.money > this.balance {
		fmt.Println("余额不足")
		//break
		//return
	} else {
		this.balance -= this.money // 将money加入余额
		fmt.Print("本次支出说明：")
		fmt.Scanln(&this.note)
		// 进行拼接到details
		this.details += fmt.Sprintf("\n支出\t%v\t\t%v\t\t%s\t", this.balance, this.money, this.note)
		this.flag = true
	}
}

func (this *FamilyAccount) Exit() {
	sig := ""
	fmt.Print("你确定要退出吗？(y/n): ")
	fmt.Scanln(&sig)
	if sig == "y" {
		this.loop = true
	} else {

	}
}

// 编写工厂模式的构造方法 返回一个 *FamilyAccount
func NewAccount(bal float64) *FamilyAccount {
	return &FamilyAccount{
		balance: bal,
	}
}
