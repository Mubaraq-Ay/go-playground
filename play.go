package main

import "fmt"

func main () {
	name := "Mubaraq"
	age := 1000
	coffee := true

	fmt.Println("Welcome,", name)
	fmt.Println("Age:", age)
	fmt.Println("Likes coffee?:", coffee)

	if age >= 18 {
		fmt.Println("Access approved.")
	} else {
		fmt.Println("Access denied.")
	}

	for i := 1; i <= 5; i++ {
		fmt.Println("go says:", i)
	}
}