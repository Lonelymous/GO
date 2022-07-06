package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

// func sum(num1, num2 int) int {
// 	return num1 + num2
// }

type Game struct {
	Title    string
	Platform string
}

type User struct {
	Id           int
	Name         string
	FavoriteGame *Game
}

type Numbers interface {
	int | float64
}

func sum[T Numbers](num1, num2 T) T {
	return num1 + num2
}

func sum2[T constraints.Ordered](num1, num2 T) T {
	return num1 + num2
}

func main() {
	// var i int32
	// var f float64
	// var b bool = true
	// var s string = "asd"

	// s = `c
	// c
	// c
	// c
	// c
	// c`

	// i = 20
	// f = 8.0

	// for index := 0; index < 10000; index++ {
	// 	f += 0.15
	// }

	// fmt.Println("Hello World!!!!", i, f, b, s)

	// var itomb []int
	// itomb = append(itomb, 1, 2, 3, 4, 5, 6)
	// itomb = append(itomb, itomb...)

	// fmt.Println("Hello World!!!!", itomb)

	// var i int
	// j := &i

	// fmt.Println(i, j, *j)

	// *j = 13
	// j = nil

	// fmt.Println(i)
	// fmt.Println(sum(2, 3))

	fmt.Println(sum(2, 3))
	fmt.Println(sum2(2, 3))

	gta := Game{
		Title:    "Grand Theft Auto V",
		Platform: "Windows",
	}

	user := User{
		Id:           10,
		Name:         "Ferike",
		FavoriteGame: &gta,
	}
	fmt.Println(user)
	user.Name = "Sanyika"
	fmt.Println(user)

	witcher3 := Game{
		Title:    "The Witcher 3",
		Platform: "Windows",
	}

	_ = witcher3

	fmt.Println(witcher3)
	fmt.Println(gta)
}
