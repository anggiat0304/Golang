package main

import "fmt"

func main() {
	fmt.Println("Materi tentang map.")
	var chicken map[string]int
	chicken = map[string]int{}
	chicken["januari"] = 30
	chicken["februari"] = 29
	fmt.Println("Januari: ", chicken["januari"])

	//mengahupus item map dengan perintah delete(var , key)
	//sebelum dihapus
	fmt.Println("Sebelum dihapus: ", chicken)
	delete(chicken, "februari")
	fmt.Println("Setelah dihapus: ", chicken)
}
