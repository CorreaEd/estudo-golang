package main

import "fmt"

func main() {
	contaDoEduardo := ContaCorrente{Titular: "Eduardo", Saldo: 500}
	contaDaKelly := ContaCorrente{Titular: "Kelly", Saldo: 300}

	status := contaDoEduardo.Transferir(200, &contaDaKelly)

	fmt.Println(status)
	fmt.Println(contaDoEduardo)
	fmt.Println(contaDaKelly)

}
