package main

import "fmt"

func main() {
	fmt.Println("MAteri tentang slice")
	var numbers = []int{1, 2, 3, 4}
	fmt.Println(numbers[0:2])
	//perintah cap dan len
	fmt.Println("Perintah cap dan len")
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(len(fruits)) // len: 4
	fmt.Println(cap(fruits)) // cap: 4

	var aFruits = fruits[0:3]
	fmt.Println(len(aFruits)) // len: 3
	fmt.Println(cap(aFruits)) // cap: 4

	var bFruits = fruits[1:4]
	fmt.Println(len(bFruits)) // len: 3
	fmt.Println(cap(bFruits)) // cap: 3
	//fungsi append menambahkan yang ada ke dalam variable baru
	fmt.Println("Perintah append")
	var newFruits = append(fruits, "salak")
	fmt.Println(fruits)
	fmt.Println(newFruits)
	//perintah copy
	fmt.Println("Perintah copy")
	dst := []string{"potato", "ship", "nut"}
	src := []string{"gelas", "piring"}
	n := copy(dst, src)
	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	//pengaksesan slice dengan 3 indeks
	fruits = []string{"apple", "banana", "grape"}
	var cFruits = fruits[0:2]
	var dFruits = fruits[0:2:2]

	fmt.Println(fruits)      // ["apple", "grape", "banana"]
	fmt.Println(len(fruits)) // len: 3
	fmt.Println(cap(fruits)) // cap: 3

	fmt.Println(cFruits)      // ["apple", "grape"]
	fmt.Println(len(cFruits)) // len: 2
	fmt.Println(cap(cFruits)) // cap: 3

	fmt.Println(dFruits)      // ["apple", "grape"]
	fmt.Println(len(dFruits)) // len: 2
	fmt.Println(cap(dFruits)) // cap: 2
}
