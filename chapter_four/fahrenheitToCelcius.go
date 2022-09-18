package main

import "fmt"

func main() {
	fmt.Print("Enter your temperature in fahreinheit: ")
	var temp float64
	fmt.Scanf("%f", &temp)
	tempInCelcius := (temp - 32) * 5 / 9
	fmt.Printf("%.2f", tempInCelcius)
}
