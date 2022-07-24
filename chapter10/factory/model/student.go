package model

//定义一个学生结构体
type student struct {
	Name string
	score float64
}

func NewStudent(name string, score float64) *student {
	return &student{
		Name : name,
		score : score,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}