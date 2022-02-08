package main

import "fmt"

func main() {
	emoji := "🐱"
	switch emoji {
	//Podemos agrupar caso en Go
	// case "🐱", "🐶":
	// y se imprime "es un animal por ejemplo"
	// o mediante comparativos && o ||
	case "🐱":
		fmt.Println("Es un pinche gato")

	case "🐶":
		fmt.Println("Es un pinche perro")

	default:
		fmt.Println("No tengo ni puta idea :v")
	}
}
