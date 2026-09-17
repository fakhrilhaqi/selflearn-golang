package main

import "fmt"

import "rsc.io/quote"

func main() {
	fmt.Println("Hello world")
	fmt.Println("wow")

	//number
	fmt.Println("Satu: ", 1)
	fmt.Println("Dua: ", 2)
	fmt.Println("Tiga koma lima: ", 3.5)

	//Boolean
	fmt.Println("Boolean: ", true)
	fmt.Println("Boolean: ", false)

	//string
	fmt.Println("String length: ", len("Hello world"))
	fmt.Println("hae antek aseng"[1])

	//variable
	name := "wowok"
	fmt.Println("Name: ", name)

	name = "hae"
	fmt.Println("Name: ", name)

	name = "wiwi"
	fmt.Println("Name: ", name)

	var (
		firstName  = "1"
		middleName = "2"
		lastName   = "3"
	)
	fmt.Println("First Name: ", firstName)
	fmt.Println("Middle Name: ", middleName)
	fmt.Println("Last Name: ", lastName)

	const (
		PI = 3.14
	)
	fmt.Println("PI: ", PI)

	//PI = 3.14159
	fmt.Println("PI: ", PI)

	//conversion
	var numberInt int = 10
	var numberFloat float64 = float64(numberInt)
	fmt.Println("Number Int: ", numberInt)
	fmt.Println("Number Float: ", numberFloat)

	fmt.Println(string("hae antek aseng"[1]))

	fmt.Println(quote.Go())
}
