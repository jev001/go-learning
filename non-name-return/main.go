package main

import "fmt"

func main() {
	println("main start")
	var x = split(17, "name")
	fmt.Print(x)
}

func split(sum int, name string) (x string) {
	y := sum * 4 / 9
	println(y)
	// println(name)
	print(x)
	return
}
