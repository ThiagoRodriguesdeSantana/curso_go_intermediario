package model

//Galinha Cacareja
type Galinha interface {
	Cacareja() string
}

//Pata faz quack
type Pata interface {
	Quack() string
}

//Ave : uma ave qualquer
type Ave struct {
	Nome string
}

//Cacareja : retorna o cacarejo de uma ave
func (a Ave) Cacareja() string {
	return a.Nome + " cocorico..."
}

// Quack  : Retorana o som dos patos
func (a Ave) Quack() string {
	return a.Nome + " quack,quack..."
}
