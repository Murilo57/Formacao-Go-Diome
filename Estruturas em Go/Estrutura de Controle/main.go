package main

import "fmt"

func main() {
	number := 1

	//1º forma
	for number <= 10 {
		fmt.Println(number)
		number = number + 1
	}

	//2º forma
	for number := 0; number <= 10; number++ {
		if number%2 == 0 {
			fmt.Println("Par")
		} else {
			fmt.Println("Impar")
		}
	}

	//Condição if para descobrir se é PAR ou IMPAR

}
