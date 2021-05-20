package main

import (
	"fmt"

	"banco/clientes"
	"banco/contas"
)

func main() {
	clienteEduardo := clientes.Titular{"Eduardo", "413.314.413.17", "Desenvolvedor Go"}
	contaDoEduardo := contas.ContaCorrente{clienteEduardo, 589, 123456, 500}

	fmt.Println(contaDoEduardo)
}
