package main

import "fmt"

func soma(numberA, numberB int) int {
	return numberA + numberB
}

func subtracao(numberA, numberB int) int {
	return numberA - numberB
}

func multiplicacao(numberA, numberB int) int {
	return numberA * numberB
}

func divisao(numberA, numberB int) int {
	if numberB == 0 {
		return 0
	} else {
		return numberA / numberB
	}
}

func main() {
	var numberA int
	fmt.Print("Digite um número inteiro: ")
	fmt.Scan(&numberA)
	fmt.Println("O número digitado foi: ", numberA)
	var numberB int
	fmt.Print("Digite outro número inteiro: ")
	fmt.Scan(&numberB)
	fmt.Print("Você digitou o número: ", numberB)

	resposta := soma(numberA, numberB)
	fmt.Printf("\nA soma dos dois números que você escolheu é de: %d\n", resposta)
	resposta = subtracao(numberA, numberB)
	fmt.Printf("\nA substração dos dois números que você escolheu é de: %d\n", resposta)
	resposta = multiplicacao(numberA, numberB)
	fmt.Printf("\nA multiplicação dos dois números que você escolheu é de: %d\n", resposta)
	resposta = divisao(numberA, numberB)
	fmt.Printf("\nA divisão dos dois números que você escolheu é de: %d\n", resposta)
}
