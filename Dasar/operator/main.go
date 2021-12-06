package main

import "fmt"

func main() {
	fmt.Println("Materi tentang operator")
	fmt.Println("Operator aritmatika")
	fmt.Printf("1 + 1 = %d\n", 1+1)
	fmt.Printf("1 * 1 = %d\n", 1*1)
	fmt.Printf("1 -1 = %d\n", 1-1)
	fmt.Printf("5 mod 2 = %d\n", 5%2)
	fmt.Printf("Operator perbandingan\n")
	var value = (((2+6)%3)*4 - 2) / 3
	var isEqual = (value == 2)
	fmt.Printf("nilai %d (%t) \n", value, isEqual)
	fmt.Println("Operator logika")
	var right = true
	var left = false
	fmt.Println(right && left)
	fmt.Println(right || left)

}
