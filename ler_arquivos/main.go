package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func main() {

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("erro ao abrir arquivo", err.Error())
		return
	}
	/*
		scanner := bufio.NewScanner(arquivo)
		for scanner.Scan() {
			linha := scanner.Text()

			println(linha)
		} */

	//pacoe especifico para leitura de csv
	leitorDeCsv := csv.NewReader(arquivo)
	conteudo, err := leitorDeCsv.ReadAll()

	if err != nil {
		fmt.Println("erro ao abrir arquivo", err.Error())
		return
	}

	for indice, linha := range conteudo {
		fmt.Printf("linha[%d] e %s\r\n", indice, linha)
		for indiceItem, item := range linha {
			fmt.Printf("item[%d] e %s\r\n", indiceItem, item)
		}
	}

	arquivo.Close()
}
