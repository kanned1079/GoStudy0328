package myUtils

type admins struct {
	name     string
	uuid     string
	password string
}

func NewAdmin(n, pwd string) *admins {
	if len(n) >= 4 && len(pwd) >= 6 {
		return &admins{
			name:     n,
			password: pwd,
		}
	} else {
		return nil
	}
}

func (this *admins) GetName() (n string) {
	if this != nil {
		n = this.name
	}
	return
}

func (this *admins) GetPwd() (p string) {
	if this != nil {
		p = this.password
	}
	return
}
