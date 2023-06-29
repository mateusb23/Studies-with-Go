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

type quadrilatero struct {
	altura      float64
	comprimento float64
}

func (q quadrilatero) Area() float64 {
	return q.altura * q.comprimento
}

type TamanhoDaFigura interface {
	Area() float64
}

func PrintarAreaTotal(t TamanhoDaFigura) {
	fmt.Printf("A ÁREA DA FIGURA GEOMÉTRICA É: %.2f\n", t.Area())
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
