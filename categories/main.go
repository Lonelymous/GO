package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type Category struct {
	Id       int
	Name     string
	Parent   *Category
	Children []*Category
	Depth    int
}

var (
	CategoryById map[int]*Category //map[CategoryId]
	Categories   *Category         //Categories root
)

func AddCategory(parentId int, id int, name string) error {
	parent, found := CategoryById[parentId]
	if !found {
		// panic("Parent not found")
		return errors.New("Parent not found")
	}
	category := Category{
		Id:       id,
		Name:     name,
		Parent:   parent,
		Children: []*Category{},
		Depth:    0,
	}
	if parent != nil {
		category.Depth = parent.Depth + 1
	}
	parent.Children = append(parent.Children, &category)
	CategoryById[id] = &category
	return nil
}

func PrintCategories() {
	var Print func(category *Category)
	Print = func(category *Category) {
		if category.Depth > 0 {
			fmt.Println("" + strings.Repeat("     ", category.Depth) + "└── " + category.Name + " (" + strconv.Itoa(category.Id) + ")")
		}
		for _, cat := range category.Children {
			Print(cat)
		}
	}
	Print(Categories)
}

// func DeferTest() (result int, err error) {

// 	defer func() {
// 		if err != nil {
// 			fmt.Println("Error volt a futás közben")
// 		} else {
// 			fmt.Println("Minden fasza")
// 		}
// 	}()

// 	defer func() {
// 		if r := recover(); r != nil {
// 			fmt.Println("Panic volt ", r)
// 		}
// 	}()

// 	i := 1
// 	fmt.Println("1.")
// 	i++
// 	defer fmt.Println(i)
// 	i++
// 	defer fmt.Println(i)
// 	fmt.Println("2.")
// 	return
// }

func CheckForError(err error) {
	if err != nil {
		fmt.Println("LOG", err)
		panic(err)
	}
}

func main() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("valami")
			//Do nothing
		}
	}()

	CategoryById = make(map[int]*Category)
	Categories = &Category{
		Id:       0,
		Name:     "",
		Parent:   &Category{},
		Children: []*Category{},
		Depth:    0,
	}

	CategoryById[0] = Categories

	CheckForError(AddCategory(0, 1, "Autó"))
	CheckForError(AddCategory(1, 2, "Gumik"))
	CheckForError(AddCategory(1, 3, "Alufelnik"))
	CheckForError(AddCategory(2222, 4, "Kaga"))
	// spew.Dump(Categories)

	PrintCategories()

}
