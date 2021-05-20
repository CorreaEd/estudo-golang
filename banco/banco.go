package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"
)

func main() {
	contaDoEduardo := contas.ContaCorrente{Titular: clientes.Titular{
		Nome:      "Eduardo",
		CPF:       "413.314.413.17",
		Profissao: "Desenvolvedor Go"},
		NumeroAgencia: 589, NumeroConta: 123456, Saldo: 500}

	fmt.Println(contaDoEduardo)
}
