package main

import "fmt"

// baseado em: https://www.digitalocean.com/community/tutorials/understanding-boolean-logic-in-go-pt
func main() {
	//x := 5
	//y := 8

	fmt.Println("HELLO WORLD")

	OpEquals(2, 2)
	OpNotEquals(2, 3)
	OpLessOrEqual(2,2)

	OpGreater(4,3)
	OpGreaterOrEqual(100, 10)
	// adicione um exemplo de chamada de cada metodo

}

func OpEquals(x int, y int) bool {
	fmt.Println("x == y:", x == y)
	return x == y
}

func OpNotEquals(x int, y int) bool {
	fmt.Println("x != y:", x != y)
	return x != y
}

func OpLess(x int, y int) bool {
	fmt.Println("x < y:", x < y)
	return x < y
}

func OpGreater(x int, y int) bool {
	fmt.Println("x > y:", x > y)
	return x > y
}

func OpLessOrEqual(x int, y int) bool {
	fmt.Println("x <= y:", x <= y)
	return x <= y
}




func OpGreaterOrEqual(saldo int, saque int) bool {
	fmt.Println("saldo >= saque", saldo >= saque)

	// FIXME implementar logica

	return saldo >= saque
}

/*

Escreva nesta função os seguintes resultados para o calculo

	se o saldo for suficiente para o saque = SALDO_SUFICIENTE
	se o saldo for igual ao saque = SALDO_SUFICIENTE_IRA_ZERAR_O_SALDO
	se o saldo for menor que o saque = SEM SALDO

*/
func StatusSaque(saldo int, saque int) string {




	// FIXME implementar logica

	return "IMPOSSIVEL REALIZAR O SAQUE"
}


