package main

import (
	"fmt"
)

type Debugger interface {
	Debug() string
}

type Calculator struct {
	A int
	B int
}

func (c Calculator) Calc() int {
	return c.A + c.B
}

func (c Calculator) Debug() string {
	return fmt.Sprintf("A = %v, B = %v", c.A, c.B)
}

func Log(i Debugger) {
	fmt.Println("Log: ", i.Debug())
}

type Calculator2 struct {
	A int
	B int
}

func (c Calculator2) Calc() int {
	return c.A * c.B
}

func (c Calculator2) Debug() string {
	return fmt.Sprintf("A = %v, B = %v", c.A, c.B)
}

func main() {
	fmt.Println("Zsa")

	calc := Calculator{
		A: 10,
		B: 20,
	}
	calc2 := Calculator2{
		A: 100,
		B: -30,
	}

	fmt.Println(calc.Calc())
	fmt.Println(calc.Debug())

	Log(calc)
	Log(calc2)
}
