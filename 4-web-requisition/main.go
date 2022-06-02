package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	fmt.Println("Iniciando servico de requisicao")
	url := "https://www.google.com"
	resp, _ := http.Get(url)

	if resp.StatusCode == 200 {
		fmt.Println("SUCESS")
		fmt.Println(resp.Status)
	} else {
		fmt.Println("FAIL")
		os.Exit(-1)
	}

}
