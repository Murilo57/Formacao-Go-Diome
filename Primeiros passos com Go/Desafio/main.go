/*
Desafio

Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da água de Kelvin para graus Celsius.

Dica: C = K - 273
*/
package main

import "fmt"

func main() {
	var temperatura float64 = 273.15
	var Kelvin float64 = 373.15

	conversaoKC(Kelvin, temperatura)
}

func conversaoKC(K float64, T float64) {
	var Celsius float64 = K - T
	fmt.Printf("A conversão de Kelvin para Celsius é de : %g", Celsius)
}
