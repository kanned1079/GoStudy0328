package models

type Nav struct {
	ID     uint
	Title  string
	Url    string
	Status int
	Sort   int
}

func (Nav) TableName() string {
	return "nav"
}
