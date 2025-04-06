package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Hola classe")
	//Haz una calculadora sencilla
	var num1 float64
	var num2 float64
	var operacion string
	var resultado float64

	fmt.Println("Introduce el primer número:")
	fmt.Scan(&num1)
	fmt.Println("Introduce el segundo número:")
	fmt.Scan(&num2)
	fmt.Println("Introduce la operación:")
	fmt.Scan(&operacion)

	switch operacion {
	case "+":
		resultado = num1 + num2
	case "-":
		resultado = num1 - num2
	case "*":
		resultado = num1 * num2
	case "/":
		resultado = num1 / num2
	default:
		fmt.Println("Operación no válida")
		return
	}

	fmt.Println("El resultado es:", resultado)
	//guarda el resultado tambe en un arxiu anomenat resultat.txt
	salida := fmt.Sprintf("El resultado es: %.2f\n", resultado)
	err := os.WriteFile("resultat.txt", []byte(salida), 0644)
	if err != nil {
		fmt.Println("Error al guardar el resultado:", err)
		return
	}
	fmt.Println("El resultat s'ha guardat en l'archivo resultat.txt")
}
