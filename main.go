package main

import (
	"log"

	"github.com/chris-miracle/reservations/helpers"
)

const numPool = 10000

func CalcVal(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)

	intChan <- randomNumber
}

func main() {
	intChan := make(chan int)
	defer close(intChan)

	go CalcVal(intChan)

	num := <-intChan

	log.Println(num)
}
