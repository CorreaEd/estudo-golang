package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

const monitoramentos = 3
const delay = 5

func main() {

	exibirIntroducao()
	lerSites()
	for {

		exibirMenu()

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
	fmt.Println("")

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")

	/* sites := []string{"https://www.alura.com.br/", "https://random-status-code.herokuapp.com/", "https://www.caelum.com.br"} */

	sites := lerSites()

	for i := 0; i < monitoramentos; i++ {
		for i, site := range sites {
			fmt.Println("Testando site", i, ":", site)
			testaSites(site)
		}
		time.Sleep(delay * time.Second)
		fmt.Println("")
	}

	fmt.Println("")

}

func testaSites(site string) {
	resp, err := http.Get(site)

	if err != nil {
		fmt.Println("Ocorreu um erro:", err)
	}

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "foi carregado com sucesso!")
	} else {
		fmt.Println("Site", site, "está com problemas. Status Code", resp.StatusCode)
	}

}

func lerSites() []string {

	var sites []string

	/* arquivo, err := os.Open("sites_monitoramento.txt")*/
	arquivo, err := ioutil.ReadFile("sites_monitoramento.txt")

	if err != nil {
		fmt.Println("Ocorreu um erro", err)

	}

	fmt.Println(string(arquivo))
	return sites
}
