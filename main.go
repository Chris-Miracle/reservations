package main

import "log"


func main() {
	var color string
	color = "Green"

	log.Println("Color is", color)

	changeColor(&color)

	log.Println("New Color is", color)
}

func changeColor(s *string) {
	newColor := "Blue"
	*s = newColor
}
