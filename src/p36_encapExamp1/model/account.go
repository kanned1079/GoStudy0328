package model

import "errors"

type account struct {
	accountName string
	balance     float64
	password    string
}

func NewAccount(name string) (acc *account, err error) {
	if len(name) >= 6 && len(name) <= 10 {
		return &account{
			accountName: name,
		}, nil
	} else {
		err = errors.New("非法账号名长度")
		return nil, err
	}
}

func (a *account) SetBalance(money float64, pwd string) (err error) {
	if a == nil {
		err = errors.New("还未设置用户 不能设置余额")
		return
	}
	if pwd == "" {
		err = errors.New("还未设置密码 不能设置余额")
		return
	}
	if pwd != a.password {
		err = errors.New("密码错误")
		return err
	}
	if money < 20 {
		err = errors.New("金额至少20")
		return
	}
	a.balance = money
	err = nil
	return err
}

func (a *account) GetBalance(pwd string) (res float64, err error) {
	if a == nil {
		err = errors.New("还未设置用户 不能查询余额")
		return
	}
	if pwd != a.password {
		err = errors.New("密码错误")
		return
	}
	res = a.balance
	return
}

func (a *account) SetPassword(password string) (err error) {
	if a == nil {
		err = errors.New("还未设置用户 不能设置密码")
		return
	}
	if len(password) != 6 {
		err = errors.New("密码长度需要六位")
		return err
	}
	a.password = password
	err = nil
	return
}
