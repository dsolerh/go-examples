package main

import "fmt"

type s1 struct{}

func (s s1) Meth1() string {
	return "meth 1"
}

func (s s1) Meth2() string {
	return "meth 2"
}

type I1 interface {
	Meth1() string
}

type I2 interface {
	Meth2() string
}

func newS1() any {
	return s1{}
}

func main() {
	val1 := newS1()

	ival1, ok := val1.(I1)
	if !ok {
		fmt.Println("val 1 is not interface I1")
	} else {
		fmt.Println("val 1 is interface I2")
	}

	_, ok = ival1.(I2)
	if !ok {
		fmt.Println("val 1 (as I1) is not interface I2")
	} else {
		fmt.Println("val 1 (as I1) is interface I2")
	}

	_, ok = val1.(I2)
	if !ok {
		fmt.Println("val 1 is not interface I2")
	} else {
		fmt.Println("val 1 is interface I2")
	}
}
