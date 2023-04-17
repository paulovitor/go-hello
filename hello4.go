package main

import "fmt"

type Carro struct {
	Nome string
}

func (c Carro) andar() {
	fmt.Println(c.Nome, "andou")
}

func main() {
	carro := Carro{
		Nome: "BMW",
	}

	carro.andar()

	resultado, str := soma(10, 20)

	fmt.Println(resultado, str)

	resultado2 := soma2(10, 20)

	fmt.Println(resultado2)

	resultado3 := somaTudo(10, 20, 30)

	fmt.Println(resultado3)

	resultado4 := func(x ...int) func() int {
		res := 0

		for _, v := range x {
			res += v
		}

		return func() int {
			return res * res
		}
	}

	fmt.Println(resultado4(1, 2, 3)())
}

func soma(a int, b int) (int, string) {
	return a + b, "somou!"
}

func soma2(a int, b int) (result int) {
	result = a + b
	return
}

func somaTudo(x ...int) int {
	resultado := 0

	for _, v := range x {
		resultado += v
	}
	return resultado
}
