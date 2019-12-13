package model
import (
	_"fmt"	
)

type student struct{
	Name string
	score float64
}

func StudentFactory(n string, s float64) *student {
	return &student{
		Name: n,
		score: s,
	}
}

func (stu *student) GetScore() float64 {
	return stu.score
}