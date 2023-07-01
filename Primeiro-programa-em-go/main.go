package main

import "fmt"

func soma(a, b int) int {
	return a + b
}

func dividir(a, b int) int {
	if b == 0 {
		return 0
	} else {
		return a / b
	}

}

func subtrair(a, b int) int {
	return a - b
}

func multiplicar(a, b int) int {
	return a * b
}

func main() {
	for i := 0; i < 100; i++ {
		for j := 0; j < 100; j++ {
			fmt.Println(i)
			fmt.Println(j)

			resposta := soma(i, j)
			fmt.Printf("Soma: %d\n", resposta)
			resposta = dividir(i, j)
			fmt.Printf("Divisão: %d\n ", resposta)
			resposta = subtrair(i, j)
			fmt.Printf("Substração: %d\n ", resposta)
			resposta = multiplicar(i, j)
			fmt.Printf("Multiplicação:%d\n ", resposta)
		}
	}
}
