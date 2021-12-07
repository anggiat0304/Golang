package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "Anggiat"
	names[1] = "Coki"
	names[2] = "Si tampan"
	names[3] = "Yang Malang"
	fmt.Println(names)

	//array tanpa initialisasi jumplah index
	var numbers = [...]int{1, 3, 5, 6, 7, 4, 3}
	fmt.Println(numbers)
	//array multidimensi
	var number1 = [2][3]int{{3, 2, 1}, {25, 67, 43}}
	fmt.Println(number1)
	//perulangan element array dengan menggunakan keyword for
	var fruits = [3]string{"Anggur", "Apel", "Pepaya"}
	for i := 0; i < len(fruits); i++ {
		fmt.Println("Index ", ":", i, " ", fruits[i])
	}
	fmt.Println("Dengan keyword range")
	//perulangan element array dengan menggunakan keyword range
	for i, fruit := range fruits {
		fmt.Println("index : ", i, " ", fruit)
	}
	//alokasi elemen array dengan keyword make
	fmt.Println("Alokasi element array dengan keyword make")
	var cars = make([]string, 2)
	cars[0] = "Ferrari"
	cars[1] = "Lamborghini"
	fmt.Println(cars)
}
