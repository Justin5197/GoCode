package main

import (
	"fmt"
	"math"
	"github.com/Justin5197/string"
)

var myNumber = 1.23

func main() {
	roundedUp := math.Ceil(myNumber)
	roundedDown := math.Floor(myNumber)
	fmt.Println(roundedUp, roundedDown)
	fmt.Println(string.Reverse("hello Justin! You're a Go developer!"))
}
