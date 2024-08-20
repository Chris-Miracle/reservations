package main

import "fmt"

type Animal interface {
	Says() string
	NumberOfLegs() int
}

type Dog struct {
	Name  string
	Breed string
}

type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Spot",
		Breed: "Golden Retriever",
	}

	PrintInfo(&dog)

	gorilla := Gorilla{
		Name:  "Jack",
		Color: "Silver",
		NumberOfTeeth: 12,
	}

	PrintInfo(&gorilla)
}

func PrintInfo(animal Animal) {
	fmt.Println("This animal says", animal.Says(), "and has ", animal.NumberOfLegs(), "legs.")
}

func (dog *Dog) Says() string {
	return "woof"
}

func (dog *Dog) NumberOfLegs() int {
	return 4
}

func (dog *Gorilla) Says() string {
	return "Ugh"
}

func (dog *Gorilla) NumberOfLegs() int {
	return 2
}
