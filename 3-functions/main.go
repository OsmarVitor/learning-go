package main

import (
	"fmt"
	"os"
)

func main() {
	showIntroduction()
	option := inputOption()
	fmt.Println("Option ", option, ".")

	name := getName()
	fmt.Println(name)

	namee, age := getNameAndAge()
	fmt.Println(namee, age)

	_, agee := getNameAndAge()
	fmt.Println(agee)
}

func showIntroduction() {
	var nome string
	var versao float32 = 1.1

	fmt.Println("Digite seu nome:")
	fmt.Scanf("%s", &nome)
	fmt.Println("Olá ", nome)
	fmt.Println("Versão do sistema", versao)

}

func inputOption() int {
	var option int
	fmt.Println("Chose the option: ")
	fmt.Scanf("%d", &option)

	switch option {
	case 1:
		fmt.Println("Option 1")
	case 2:
		fmt.Println("option 2")
	case 3:
		fmt.Println("Option 3")
	default:
		fmt.Println("Undefined")
		os.Exit(-1)
	}

	return option

}

func getName() string {
	name := "vitor"
	return name
}

func getNameAndAge() (string, int) {
	name := "vitor"
	age := 24
	return name, age
}
