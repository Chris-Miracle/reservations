package main

import "log"


type class struct{
	StudentName string
}

func (c *class) showFirstName() string {
	return c.StudentName
}

func main() {
	var myVar class
	myVar.StudentName = "John"

	myVar2 := class{
		StudentName: "Mary",
	}

	log.Println(myVar.showFirstName())
	log.Println(myVar2.showFirstName())
}
