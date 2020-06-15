package main

import (
	"fmt"
	// é possivel adicionar um "apelido" ao import
	m "math"
)

func main() {
	const PI float64 = 3.1415
	var raio = 3.8 // tipo (float64) inferido pelo compilador

	// forma reduzida de declarar variavel
	area := PI * m.Pow(raio, 2)

	fmt.Println("Valor da área é", area)

	const (
		a = 1
		b = 2
	)

	var (
		c = 3
		d = 4
	)

	fmt.Println(c, d)

	// declaração de mais de uma variavel na mesma linha
	var e, f bool = true, false

	fmt.Println(e, f)

	//delcaração de multiplas variaveis de forma reduzida
	g, h, i := 2, false, "Bruno"

	fmt.Println(g, h, i)
}
