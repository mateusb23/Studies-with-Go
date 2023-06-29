package main

import (
	"fmt"
	"math"
)

type circulo struct {
	raio float64
}

func (c circulo) Area() float64 {
	return math.Pow(c.raio, 2) * 3.4
}

func (c circulo) TipoFigura() (string, string) {
	nome := "círculo"
	tipoDimensional := "bidimensional"
	return nome, tipoDimensional
}

type quadrilatero struct {
	altura      float64
	comprimento float64
}

func (q quadrilatero) Area() float64 {
	return q.altura * q.comprimento
}

func (q quadrilatero) TipoFigura() (string, string) {
	var nome string
	tipoDimensional := "bidimensional"
	if q.altura == q.comprimento {
		nome = "quadrado"
		return nome, tipoDimensional
	} else {
		nome = "retângulo"
		return nome, tipoDimensional
	}
}

type DadosDaFigura interface {
	Area() float64
	TipoFigura() (string, string)
}

func PrintarAreaTotal(d DadosDaFigura) {
	fmt.Println(d.TipoFigura())
	fmt.Printf("A ÁREA É: %.2f\n\n", d.Area())
}

func main() {
	c1 := circulo{
		raio: 4.3,
	}

	q1 := quadrilatero{
		altura:      12,
		comprimento: 6,
	}

	var q2 quadrilatero
	q2 = quadrilatero{
		altura:      5,
		comprimento: 5,
	}

	PrintarAreaTotal(c1)
	PrintarAreaTotal(q1)
	PrintarAreaTotal(q2)
}
