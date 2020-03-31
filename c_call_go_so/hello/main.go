package main

import "C"

func main() {}

//export Hello
func Hello() string {
	return "Hello"
}

//export Test
func Test() {
	println("export Test")
}
