package main

import "fmt"

func main() {
	nome := "Vitor"
	var versao float32 = 1.1
	var opcao int

	fmt.Println("Olá ", nome)
	fmt.Println("Versão do sistema", versao)

	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")

	fmt.Println("Digite a opção:")
	fmt.Scanf("%d", &opcao)
	fmt.Println(opcao)
}
