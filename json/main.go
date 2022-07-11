package main

import (
	"encoding/json"
	"fmt"
)

//Product is my first structure.
type Product struct {
	Id    int    `json:"-"`
	Name  string `json:"nameOfProduct"` //This is the name
	Stock int
}

func (p Product) GetName() string {
	return p.Name
}

func (p *Product) SetName(name string) {
	p.Name = name
}

func main() {
	prd := Product{
		Id:    5,
		Name:  "PC",
		Stock: 13,
	}

	data, _ := json.Marshal(prd)
	fmt.Println(string(data))

	prd.Name = "asdasgsdlgfnsdklgfnsdl"

	_ = json.Unmarshal(data, &prd)
	fmt.Println(prd)
}
