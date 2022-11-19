package main

import (
	"fmt"
	"math"
	"os"
	"os/exec"
	"runtime"
	"time"
)

//As funções ClearTerminar e runCMD eu encontrei no google. Estou utilizando apenas para deixar o programa mais bonitinho :) //

func ClearTerminal() {
	switch runtime.GOOS {
	case "darwin":
		runCmd("clear")
	case "linux":
		runCmd("clear")
	case "windows":
		runCmd("cmd", "/c", "cls")
	default:
		runCmd("clear")
	}
}

func runCmd(name string, arg ...string) {
	cmd := exec.Command(name, arg...)
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func menu() {
	var comando int
	fmt.Println("Selecione o que deseja fazer : ")
	fmt.Println("1- Resolver uma equação.")
	fmt.Println("2- Sair do programa.")
	fmt.Scan(&comando)

	if comando == 1 {
		ResolveEquacao()
	} else if comando == 2 {
		fmt.Println("Espero que tenha gostado ;)")
		os.Exit(0)
	} else {
		ClearTerminal()
		fmt.Println("Esse comando não existe. Por favor tente novamente")
		menu()
	}

}

func ResolveEquacao() {
	var a int
	var b int
	var c int
	fmt.Println("Insira o valor de a:")
	fmt.Scan(&a)
	fmt.Println("Insira o valor de b:")
	fmt.Scan(&b)
	fmt.Println("Insira o valor de c:")
	fmt.Scan(&c)

	delta := (b * b) - (4 * a * c)

	if a <= 0 {
		fmt.Print("Para ser uma equação de segundo grau, 'a' deve ser maior que 0.")
	} else if delta < 0 {
		fmt.Println("A função não possui raízes reais.")
	} else if delta == 0 {
		raiz := (-b / (2 * a))
		fmt.Println("A função possui uma única raiz : ", raiz)
	} else {
		deltaRaiz := math.Sqrt(float64(delta))
		raiz1 := (((float64(-b)) + (deltaRaiz)) / float64(2*a))
		raiz2 := (((float64(-b)) - (deltaRaiz)) / float64(2*a))
		fmt.Println("A função possui duas raízes :")
		fmt.Println("x1 : ", raiz1)
		fmt.Println("x2 : ", raiz2)
	}
	time.Sleep(3 * time.Second)
	menu()
}

func main() {
	fmt.Println("Olá, bem vindo a Calculadora de equação de segundo grau.")
	menu()
}
