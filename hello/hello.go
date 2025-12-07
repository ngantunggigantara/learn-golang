package main

import "fmt"

//declare variable
//var name string
// var a int
// var b int

// var age int = 20
// var name string = "Gigan"

var x float64

func main() {

	// age := true
	// name := "Gigan"
	// a = 10
	// b = 5

	// fmt.Println(a + b)

	// fmt.Println("Masukkan nama:")
	// fmt.Scanln(&name)
	// fmt.Printf("Hello %s, welcome to Go!", name)

	// fmt.Println(age)
	// fmt.Println(name)

	fmt.Print("Masukkan angka:")
	fmt.Scanln(&x)
	fmt.Printf("hasil float %f", x)

}
