package main

import "fmt"

func main() {
	var nums []int
	fmt.Println(nums, len(nums), cap(nums))

	nums = make([]int, 5)
	fmt.Println(nums, len(nums), cap(nums))

	capitais := []string{"Lisboa"}
	fmt.Println(capitais, len(capitais), cap(capitais))

	//o append  limilar ao 'lista.Add()' do C#
	capitais = append(capitais, "brasilia")
	fmt.Println(capitais, len(capitais), cap(capitais))

	cidade := make([]string, 5)

	cidade[0] = "nova york"
	cidade[1] = "londre"
	cidade[2] = "toquio"
	cidade[3] = "singapura"
	cidade[4] = "brasília"

	fmt.Println(cidade, len(cidade), cap(cidade))

	for index, c := range cidade {
		fmt.Printf("%v %v \n", index, c)
	}

	//busca dados do slice pelo intervalo de index informado,
	//considerando que na busca o primeiro index comeca com 0 e segundo começa nno 1
	capitaisAsia := cidade[2:4]
	fmt.Println(capitaisAsia, len(capitaisAsia), cap(capitaisAsia))

	//buscar apenas os 2 ultimos itens do slice
	doisultimos := cidade[len(cidade)-2:]
	fmt.Println(doisultimos)

	//Remover um item do slice

}
