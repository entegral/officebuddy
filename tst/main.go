package main

import "github.com/sirupsen/logrus"

type Thing struct {
	*Thang
}

type Thang struct {
	Name string
}

func (t *Thang) Tst() string {
	if t == nil {
		t = &Thang{Name: "im not nil anymore!"}
		return t.Name
	}
	return "its not nil"
}

func main() {
	t := Thing{}
	o := t.Tst()
	logrus.Println(o)
}
