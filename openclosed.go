package main

import "fmt"

// Shape é uma interface para formas geométricas.
type Shape interface {
    Area() float64
}

// Rectangle é uma estrutura que representa um retângulo.
type Rectangle struct {
    Width  float64
    Height float64
}

// Area calcula a área do retângulo.
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}

// Circle é uma estrutura que representa um círculo.
type Circle struct {
    Radius float64
}

// Area calcula a área do círculo.
func (c Circle) Area() float64 {
    return 3.14 * c.Radius * c.Radius
}

// CalculateArea calcula a área total de várias formas geométricas.
func CalculateArea(shapes []Shape) float64 {
    totalArea := 0.0
    for _, shape := range shapes {
        totalArea += shape.Area()
    }
    return totalArea
}

func main() {
    shapes := []Shape{
        Rectangle{Width: 5, Height: 10},
        Circle{Radius: 3},
    }
    totalArea := CalculateArea(shapes)
    fmt.Printf("Total area: %f\n", totalArea)
}
