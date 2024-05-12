package main

import "fmt"

// 我们在前面去定义一个结构体时候，实际上就是把一类事物的共有的属性(字段)和行为仿法)提取出来
// 形成一个物理模型(结构体)。这种研究问题的方法称为抽象。

type BankAccount struct {
	AccountNumber string
	Password      string
	Balance       float64
}

// 方法
// 1. 可以存款
func (account *BankAccount) Deposit(money float64, pwd string) {
	// 驗證密碼正確
	if pwd != account.Password {
		fmt.Println("密碼錯誤")
		return
	}
	// 驗證金額是否正確
	if money <= 0 {
		fmt.Println("存款不能為0")
		return
	}
	account.Balance += money
	fmt.Println("存款成功, 餘額為", account.Balance)
}

// 2. 可以取款
func (account *BankAccount) Withdraw(money float64, pwd string) {
	// 驗證密碼正確
	if pwd != account.Password {
		fmt.Println("密碼錯誤")
		return
	}
	// 驗證金額是否正確
	if money <= 0 || money > account.Balance {
		fmt.Println("輸入金額不正確")
		return
	}
	account.Balance -= money
	fmt.Println("取款成功, 餘額為", account.Balance)
}

// 3. 查詢餘額
func (account *BankAccount) QueryBalance(pwd string) {
	// 驗證密碼正確
	if pwd != account.Password {
		fmt.Println("密碼錯誤")
		return
	}
	//
	fmt.Printf("你的帳號是%s, 餘額為%v\n", account.AccountNumber, account.Balance)
}

func main() {
	account1 := &BankAccount{
		AccountNumber: "GS123456789123456",
		Password:      "666666",
		Balance:       100.0,
	}
	account1.QueryBalance("666666")
	account1.Deposit(800, "666666")
	account1.Withdraw(150, "666666")

}
