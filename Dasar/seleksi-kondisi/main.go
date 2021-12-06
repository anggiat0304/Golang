package main

import "fmt"

func main() {
	fmt.Println("Materi tentang pemograman seleksi kondisi")

	//Penggunaan if else variable biasa
	fmt.Println("Penggunaan pada if else")
	var point = 9
	if point == 0 {
		fmt.Println("Point sama dengan 0")
	} else if point == 1 {
		fmt.Println("Point sama dengan 1")
	} else if point == 8 {
		fmt.Println("Point sama dengan 8")
	} else {
		fmt.Println("Gak ada satu pun yang benar adalah ", point)
	}

	//Penggunaan pada if else variable temporary
	fmt.Println("Penggunaan pada variable temporary")
	point = 8840.0
	if percent := point / 100; percent >= 100 {
		fmt.Println("perfect : ", percent, "%")
	} else if percent >= 70 {
		fmt.Println("Good : ", percent, "%")
	} else {
		fmt.Println("Bad : ", percent, "%")
	}

	//Penggunaan switch case
	fmt.Println("Penggunaan switch case. ")
	point = 20
	switch point {
	case 2:
		fmt.Println("Nilai nya sama dengan 2 ")
	case 3:
		fmt.Println("Nilai nya sama dengan 3 ")
	case 4:
		fmt.Println("Nilai nya sama dengan 4")
	default:
		fmt.Println("Salah semua nilai nya adalah ", point)
	}
	point = 6

	switch {
	case point == 8:
		fmt.Println("perfect")
	case (point < 8) && (point > 3):
		fmt.Println("awesome")
		fallthrough
	case point < 5:
		fmt.Println("you need to learn more")
	default:
		{
			fmt.Println("not bad")
			fmt.Println("you need to learn more")
		}
	}
}
