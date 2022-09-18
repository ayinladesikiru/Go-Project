package main

import "fmt"

func main() {
	fmt.Println("Enter your length in feet: ")
	var lengthInFeet float64
	fmt.Scanf("%d", lengthInFeet)
	lengthInMeter := lengthInFeet * 0.3048
	fmt.Println(lengthInMeter)
}
