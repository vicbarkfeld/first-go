package main

import "fmt"
import "os"
import "net/http"

func main() {

	exibeIntroducao()	
	for{
		exibeMenu()		
	/* 	exibeNomes() */
	
		comando := leComando()
		
		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Exibindo logs...")
		case 0:
			fmt.Println("Saindo do programa!")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando :/")
			os.Exit(-1)
		}
	}
	
}

func exibeIntroducao(){
	nome := "Victoria"
    idade := 25
    versao := 1.1
    fmt.Println("Olá", nome, "sua idade é", idade)
    fmt.Println("Este programa está na versão", versao)
}

func exibeMenu(){
	fmt.Println("1- Iniciar Monitoramento")
    fmt.Println("2- Exibir Logs")
    fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O comando escolhido é:", comandoLido)

	return comandoLido
}

func iniciarMonitoramento() {
    fmt.Println("Monitorando...")
	var sites [3]string
	sites[0] = "https://www.random-status-code.herokuapp.com/"
	sites[1] = "https://www.alura.com.br"
	sites[2] = "https://www.caelum.com.br"


	fmt.Println(sites)

    site := "https://www.alura.com.br"
    resp, _ := http.Get(site)

    if resp.StatusCode == 200 {
        fmt.Println("Site:", site, "foi carregado com sucesso!")
    }else{
        fmt.Println("Site:", site, "está com problemas. Status Code:", resp.StatusCode)
    }
} 