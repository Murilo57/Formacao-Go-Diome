package main

import "fmt"

const ebulicaoF float64 = 212

func main() {

	tempF := ebulicaoF
	tempC := (tempF - 32) * 5 / 9 //Conversão de F para C

	fmt.Printf("A temperatura de ebulição da água em ºF =   %g(%T).       A temperatura de ebulição da água em ºC =  %g.      ", tempF, tempF, tempC)

}
