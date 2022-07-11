package main

import "fmt"

//Product is my first structure.
type Product struct {
	id    int
	Name  string //This is the name
	Stock int
}

func (p Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

type Empty struct{}

func (e Empty) Calc() int {
	return 10
}

func main() {
	prd := Product{
		Name:  "PC",
		Stock: 2,
	}

	prd.id = 5

	e := Empty{}
	fmt.Println(e.Calc())

	fmt.Println(prd)
	fmt.Println(prd.GetName())

	prd.SetName("asdasdasd")

	fmt.Println(prd.GetName())

	if (prd.id == 5 && prd.Name == "asd") || prd.id == 4 {

	} else if prd.Name == "Lajos" {
		fmt.Println("Lajos")
	} else {
		fmt.Errorf("Error")
	}

	// result := e.Calc()
	if result := e.Calc(); result > 120 {

	} else if e.Calc() > 1000 {

	}

	arr := make([]int, 0)
	arr2 := []int{1, 2, 3}

	fmt.Println(arr, arr2)

	for i := 0; i < 32; i++ {
		arr = append(arr, i*2)
		fmt.Println(arr[i])
	}

	fmt.Println(arr, arr2)

	for i := range arr {
		fmt.Println(arr[i])
	}

	fmt.Println(arr, arr2)

	for i, elem := range arr {
		fmt.Println(arr[i], elem)
	}

	fmt.Println(arr, arr2)

	for _, elem := range arr {
		fmt.Println(elem)
		elem = 6
	}

	fmt.Println(arr, arr2)

	switch prd.id {
	case 1:
		fmt.Println("1")
		break
	case 2:
		fmt.Println("2")
		break
	case 3, 4:
		fmt.Println("3 & 4")
		break
	default:
		fmt.Println("default")
		break
	}

	switch {
	case prd.id > 10:

	case prd.Name == "asdasd":

	}

}
