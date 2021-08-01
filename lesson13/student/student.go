package student

import (
	"fmt"
)

type Student struct {
	FullName string
	AverageMark int
}


func (s *Student) getFullInfo() string {
	res := "FullName:%s; AverageMark: %d"
	result := fmt.Sprintf(res, s.FullName, s.AverageMark)
	return result
}


