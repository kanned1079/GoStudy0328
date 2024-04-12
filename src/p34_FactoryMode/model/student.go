package model

type student struct {
	name  string
	score float64
}

func NewStudent(n string, s float64) *student {
	return &student{
		name:  n,
		score: s,
	}
}

func (s *student) GetName() string {
	return s.name
}

func (s *student) SetName(n string) {
	s.name = n
}
