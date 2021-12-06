package main

import "fmt"

func main() {
	fmt.Println("Materi tentang perulangan")
	//keyword for
	for i := 1; i <= 5; i++ {
		fmt.Println(i)
	}
	//penggunaan for dengan argumen
	fmt.Println("Perulangan dengan argumen")
	var i = 0
	for i < 5 {
		fmt.Println(i)
		i++
	}

	//penggunaan for tanpa argumen
	i = 0
	fmt.Println("Perulangan tanpa argumen")
	for {
		fmt.Println(i)
		i++
		if i > 5 {
			break
		}
	}

	//penggunaan keyword & break
	fmt.Println("Perulangan dengan continue dan break")
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue
		} else if i > 10 {
			break
		}
		fmt.Println(i)
	}
	//perulangan bersarang
	fmt.Println("Perulangan bersarang (nested for)")
	for i := 1; i <= 10; i++ {
		for j := i; j < 10; j++ {
			print(j)
		}
		fmt.Println()
	}
}
