package student

// Student đại diện cho một sinh viên với danh sách nhiều điểm
type Student struct {
	Id     int
	Name   string
	Class  string
	Scores []float64 // Mảng động chứa các đầu điểm
}

// Method tính điểm trung bình để hiển thị
func (s Student) GetAverage() float64 {
	if len(s.Scores) == 0 {
		return 0
	}
	var total float64
	for _, score := range s.Scores {
		total += score
	}
	return total / float64(len(s.Scores))
}

func (s Student) GetId() int {
	return s.Id
}

func IsIdUniqueStudent(id int, list []Student) bool {
	for _, item := range list {
		if item.Id == id {
			return false
		}
	}

	return true
}
