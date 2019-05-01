package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"
	"strings"

	"github.com/ThiagoRodriguesdeSantana/curso_go_intermediario/escrever_arquivos/model"
)

func main() {

	arquivoJson, err := os.Create("cidades.json")

	if err != nil {
		fmt.Println("erro ao gerar aquivo", err.Error())
	}

	arquivo, err := os.Open("cidades.csv")
	if err != nil {
		fmt.Println("erro ao abrir arquivo", err.Error())
		return
	}
	leitorDeCsv := csv.NewReader(arquivo)
	conteudo, err := leitorDeCsv.ReadAll()

	if err != nil {
		fmt.Println("erro ao abrir arquivo", err.Error())
		return
	}

	escritor := bufio.NewWriter(arquivoJson)
	escritor.WriteString("[\r\n")

	for _, linha := range conteudo {
		for indiceItem, item := range linha {
			dados := strings.Split(item, "/")
			cidade := model.Cidade{}

			cidade.Nome = dados[0]
			cidade.Estado = dados[1]

			cidadeJson, err := json.Marshal(cidade)

			if err != nil {
				fmt.Println("erro ao gerar json", err.Error())
			}
			escritor.WriteString(" " + string(cidadeJson))

			if (indiceItem + 1) < len(linha) {
				escritor.WriteString(",\r\n")
			}
		}
	}
	escritor.WriteString("\r\n]")
	escritor.Flush()
	arquivoJson.Close()
	arquivo.Close()

}
