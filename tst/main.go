package main

import (
	"errors"

	"github.com/sirupsen/logrus"
)

type Thang struct {
	Name string
}

// Error method returns the error message of a Thang
func (t Thang) Error() string {
	return t.Name
}

func main() {
	ErrorRun1()
	ErrorRun2()
}

func ErrorRun1() {
	t := &Thang{}
	var errThang *Thang
	if errors.As(t, &errThang) {
		logrus.Println("Thang is the error")
	} else {
		logrus.Println("Thang is not the error")
	}
}

func ErrorRun2() {
	t := &Thang{}
	var errThang *Thang
	if errors.As(t, &errThang) {
		logrus.Println("Thang is the error")
	} else {
		logrus.Println("Thang is not the error")
	}
}
