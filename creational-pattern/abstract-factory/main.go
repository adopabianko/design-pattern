package main

import "fmt"

func main() {
	adidasFactory, _ := getSportsFactory("adidas")
	nikeFactory, _ := getSportsFactory("nike")

	nikeShoe := nikeFactory.makeShoe()
	nikeShort := nikeFactory.makeShort()

	adidasShoe := adidasFactory.makeShoe()
	adidasShort := adidasFactory.makeShort()

	printShoeDetails(nikeShoe)
	printShirtDetails(nikeShort)

	printShoeDetails(adidasShoe)
	printShoeDetails(adidasShort)
}

func printShoeDetails(s iShoe) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}

func printShirtDetails(s iShort) {
	fmt.Printf("Logo: %s", s.getLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.getSize())
	fmt.Println()
}
