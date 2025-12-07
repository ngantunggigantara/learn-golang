package main

import "fmt"

func main() {
	var a, b float64
	var op string

	fmt.Print("Masukkan angka pertama: ")
	fmt.Scanln(&a)

	fmt.Print("Masukkan angka kedua: ")
	fmt.Scanln(&b)

	fmt.Print("Pilih operasi (+ - * /): ")
	fmt.Scanln(&op)

	var result float64

	if op == "+" {
		result = a + b
	} else if op == "-" {
		result = a - b
	} else if op == "*" {
		result = a * b
	} else if op == "/" {
		if b == 0 {
			fmt.Println("Error: pembagian dengan nol!")
			return
		}
		result = a / b
	} else {
		fmt.Println("Operasi tidak dikenal!")
		return
	}

	fmt.Println("Hasil:", result)
}
