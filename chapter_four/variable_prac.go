package main

import "fmt"

var x string = "hello Asa!"
var y string = "The lion is here"

func main() {
	fmt.Println(y == x)
}

func f() {
	fmt.Println(x + y)
}
