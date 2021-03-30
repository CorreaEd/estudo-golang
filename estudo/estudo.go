package main

import (
	"fmt"
	"net/http"
	"os"
	"reflect"
)

func main() {

	var sites [4]string
	sites[0] = "https://www.alura.com.br/"
	sites[1] = "https://www.alura.com.br/dsadaga"
	sites[2] = "https://random-status-code.herokuapp.com/"
	sites[3] = "https://www.alura.com.br/xsafgsa"
	/* 	sites[4] = "https://random.com/"
	Como o array criado tem [4] posições, o que exceder as 4 posições,
	por exemplo, será entendido como invalid array index */

	fmt.Println(reflect.TypeOf(sites))
	fmt.Println(sites)
	exibeNomes()
	/* exibirIntroducao() */
	for {
		/* 	exibirMenu() */

		comando := lerComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
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

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	site := "https://www.alura.com.br/"
	resp, _ := http.Get(site)
	/* fmt.Println(resp) */

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code", resp.StatusCode)
	}

}

func exibeNomes() {
	nomes := []string{"Douglas", "Daniel", "Denis"}
	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

	nomes = append(nomes, "Eduardo", "Kelly")

	fmt.Println(nomes)
	fmt.Println(reflect.TypeOf(nomes))
	fmt.Println("O meu slice tem", len(nomes), "itens")
	fmt.Println("O meu slice tem capacidade para", cap(nomes), "itens")

}

/*Array X Slice - Na prática, não aparenta diferença alguma, mas têm suas diferenças:
Array por exemplo, é limitado à quantidade de posições informada inicialmente

	var exemplo [4]string

Nesse caso, é um array de 4 posições:

	exemplo[1] = "string1"
	exemplo[2] = "string2"
	exemplo[3] = "string3"
	exemplo[4] = "string4"

Em Go, os slices são arrays, são feitos em cima de um array,
porém o slice facilita o seu uso, por se "ajustar" exatamente na quantidade de posições necessárias:

	exemplo := []string{"string1", "string2", "string3", "string4", "string5", "string6"}


Quando é feito um append >> o slice dobra a quantidade de posições iniciais!
Ou seja, neste exemplo tinhamos 6 posições; Ao fazer um apend,
ele passará a ter 12 posições, apesar de estar usando, neste exemplo, apenas 9 delas:

	exemplo := []string{"string1", "string2", "string3", "string4", "string5", "string6"}
	exemplo = append(exemplo, "string7", "string8", "string9")

*/
