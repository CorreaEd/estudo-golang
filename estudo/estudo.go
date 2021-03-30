package main

import (
	"fmt"
	"os"
)

func main() {

	exibirIntroducao()
	exibirMenu()

	comando := lerComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Exibindo Logs...")
	case 0:
		fmt.Println("Saindo do programa.")
		os.Exit(0)
	default:
		fmt.Println("Não conheço este comando")
		os.Exit(-1)
	}

}

func exibirIntroducao() {
	nome := "Douglas"
	versão := 1.1
	fmt.Println("Olá, sr.", nome)
	fmt.Println("Este programa está na versão", versão)
}

func exibirMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func lerComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido foi", comandoLido)

	return comandoLido
}
