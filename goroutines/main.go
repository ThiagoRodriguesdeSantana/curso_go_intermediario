package main

import (
	"fmt"
	"sync"
)

var orquetrador sync.WaitGroup

func main() {

	//criar rotina concorrentes, threads

	orquetrador.Add(2)

	go contarAte100_1()
	go contarAte100_2()

	orquetrador.Wait()

}
func contarAte100_1() {
	for index := 0; index < 100; index++ {
		fmt.Println("metodo 01", index)
	}
	orquetrador.Done()

}

func contarAte100_2() {
	for index := 0; index < 100; index++ {
		fmt.Println("metodo 02", index)
	}
	orquetrador.Done()
}
