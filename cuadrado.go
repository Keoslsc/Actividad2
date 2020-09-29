package main

import "fmt"

func main() {
	var lado float32
	var area float32

	fmt.Println("Calcular área de un cuadrado.")
	fmt.Println("-----------------------------")
	fmt.Printf("Digite el valor del lado:")
	fmt.Scan(&lado)
	area = lado * lado
	fmt.Printf("El area del cuadrado es: %.2f.", area)
	fmt.Printf("Fin")
}
