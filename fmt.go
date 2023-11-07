package main

import "fmt"

func main() {
	firstName := "Eko"
	lastName := "Khannedy"

	fmt.Println("Hello '", firstName, lastName, "'")

	fmt.Printf("Hello '%s %s'\n", firstName, lastName)
}
