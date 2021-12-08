package main

import (
	"fmt"
)

func main() {
	fmt.Println("Materi tentang fungsi")
	//function rand.send
	fmt.Println(calculate(10, 6))
}

//function for plus biasa
func plus(a int, b int) int {
	return a + b
}

//function disperency dengan type data param sama dengan cara 2
func disperency(a, b int) int {
	return b - a
}

func div(a, b int) int {
	return b / a
}

//function multiple return
func calculate(x, y int, bangun string) (int, int, int) {
	plus := plus(x, y)
	dis := disperency(x, y)
	div := div(x, y)
	return plus, dis, div
}
