package main

import (
	"curso_go_intermediario/interfaces/model"
	"fmt"
)

func main() {
	jojo := model.Ave{}
	jojo.Nome = "Jojo da silva"

	cacarejo(jojo)

	quackDosPatos(jojo)
}

func cacarejo(g model.Galinha) {
	fmt.Println(g.Cacareja())
}

func quackDosPatos(p model.Pata) {
	fmt.Println(p.Quack())
}
