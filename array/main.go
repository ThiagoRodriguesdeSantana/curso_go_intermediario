package main

import "fmt"

func main() {

	var teste [3]int
	teste[0] = 3
	teste[1] = 2
	teste[2] = 1

	fmt.Println("Qual o tamanho de array?", len(teste))

	cantores := [2]string{"thisgo", "jhon lenon"}

	fmt.Printf("O que á nesse array\n\r%v\r\n", cantores)

	//array nem tamanho definido

	capitais := [...]string{"lisboa", "brasilia", "luanda", "maputo"}

	for index, cidade := range capitais {
		fmt.Printf("Capital [%v] %s\r\n", index, cidade)

	}
}
