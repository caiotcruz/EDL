package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func checaSite() {
	var url string
	fmt.Println("\nInsira a url do site que deseja verificar: ")
	fmt.Scan(&url)

	response, err := http.Get(url)

	if err != nil {
		fmt.Println("Erro encontrado : ", err)
		fmt.Println("Possivelmente a url está incorreta, verifique se a url possuí http/https...")
	} else {
		fmt.Println("\nStatus do site ", url, ":  |", response.Status, "|")
	}
	menu()
}

func menu() {
	fmt.Println("\n1- Checar site.")
	fmt.Println("2- Sair")

	var comando int
	fmt.Scan(&comando)

	switch comando {
	case 1:
		checaSite()
	case 2:
		fmt.Println("Espero que tenha gostado. :)")
		time.Sleep(2 * time.Second)
		os.Exit(0)
	default:
		fmt.Println("Comando não reconhecido. Por favor, tente novamente.")
		menu()
	}

}

func main() {
	fmt.Println("--- Bem vindo!!! ---")
	menu()
}
