package main

import "fmt"

// Bird é uma interface para pássaros.
type Bird interface {
    Fly() string
}

// Duck é uma estrutura que representa um pato.
type Duck struct{}

// Fly implementa o método Fly para Duck.
func (d Duck) Fly() string {
    return "Duck flying"
}

// Ostrich é uma estrutura que representa um avestruz.
type Ostrich struct{}

// Fly implementa o método Fly para Ostrich.
func (o Ostrich) Fly() string {
    return "Ostrich cannot fly"
}

// MakeBirdFly faz um pássaro voar.
func MakeBirdFly(b Bird) {
    fmt.Println(b.Fly())
}

func main() {
    duck := Duck{}
    ostrich := Ostrich{}

    MakeBirdFly(duck)
    MakeBirdFly(ostrich) // Esse não voará.
}
