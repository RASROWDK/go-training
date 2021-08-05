package main

import (
	"fmt"
	"math"
)

func main() {

	fmt.Println("soma", soma(1,1))
	fmt.Println("subtraçao",subtraçao(1,1))
	fmt.Println("adicionarPercentual", adicionarPercentual(20, 100))
    fmt.Println("descobrirPercentual", descobrirPercentual(100,10))

}
 func soma(x float64,y float64) float64{
 	return x+y

 }

 func subtraçao(x float64,y float64) float64{
 	return x-y
}
 func multiplicaçao(x float64,y float64) float64{
 	return x * y
 }

 func divisao(x float64,y float64) float64{
 	return x/y
 }
 func adicionarPercentual(valor float64, porcentagem float64) float64 {
	 return valor + (valor * porcentagem / 100)
 }
 func descobrirPercentual(valor float64, valor2 float64) float64{
 	return(valor2/valor)*100

 }
func pow(valor float64, multiplo float64) float64{
	return math.Pow(valor, multiplo)
}






