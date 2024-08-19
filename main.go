package main

import "fmt"

func main() {
	fmt.Println("Hello, World")

	var whatToSay string
	var i int

	whatToSay = "Goodbye, World"
	fmt.Println(whatToSay)

	i = 7

	fmt.Println("i is set to", i)

	whatWasSaid, theOtherthing := saySomething()

	fmt.Println("The function returned", whatWasSaid, theOtherthing)
}

func saySomething() (string, string) {
	return "Saying something", "Else"
}
