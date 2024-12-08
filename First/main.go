package main

import "fmt"

func main() {
	var varString string = "I am a normal string"
	const constString string = "I am a constant string"
	var varInt int = 1337
	varSugar := "\"Look, I'm syntactic\""
	fmt.Printf("Hello there!\n")
	fmt.Printf("I have two strings. First one is \"%v\" and the other is \"%v\"\n", varString, constString)
	fmt.Printf("I have a number as well: %v\n", varInt)
	fmt.Print("It's easier to use Printf instead of normal print if you want to just write.\n")
	fmt.Print("Here is the first variable: ", varString, "\nAnd here is the second variable: ", constString, "\n")
	fmt.Print("This is also a string: ", varSugar)

	var firstName string

	fmt.Print("Input name: ")
	fmt.Scan(&firstName)
	fmt.Println(firstName)
}
