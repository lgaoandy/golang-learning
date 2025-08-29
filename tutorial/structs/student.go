package main

import "fmt"

type Student struct {
	name string
	grades []int
	age int
}

// methods
func (s *Student) setAge(age int) {
	s.age = age
}

func (s Student) getAverageGrade() float32 {
	sum := 0
	for _, v := range s.grades {
		sum += v
	}
	return float32(sum) / float32(len(s.grades))
}

func (s Student) getBestGrade() int {
	best := 0
	for _, v := range s.grades {
		best = max(best, v)
	}
	return best
}

func (s *Student) addGrade(newGrade int) {
	s.grades = append(s.grades, newGrade)
}

func main() {
	s1 := Student{"Tim", []int{70,90,80,85}, 19}
	fmt.Println(s1)
	s1.setAge(7)
	fmt.Println(s1)

	average := s1.getAverageGrade()
	fmt.Println(average)

	best := s1.getBestGrade()
	fmt.Println(best)

	s1.addGrade(95)
	average = s1.getAverageGrade()
	best = s1.getBestGrade()
	fmt.Println(average, best)
}