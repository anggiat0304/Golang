package main

import "fmt"

func main() {
	fmt.Println("Materi tentang tipe data")
	fmt.Println("Jenis Jenis tipe data")
	var name string = "Anggiat Pangaribuan"
	fmt.Println("Tipe data string : ", name)
	var age int = 20
	fmt.Println("Tipe data integer : ", age)
	var decimalNumber = 2.342
	fmt.Printf("Decimal 2 dibelakang koma : %.2f\n", decimalNumber)
	fmt.Printf("Decimal 3 dibelakang koma : %.3f\n", decimalNumber)
	var boolean bool = true
	fmt.Println("tipe data boolean : ", boolean)
}
