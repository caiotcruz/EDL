package main

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"
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

func checaComando(casa string, comando string, l1 [3]string, l2 [3]string, l3 [3]string) string {
	if casa == "" {
		return comando
	} else {
		imprimeTabuleiro(l1, l2, l3)
		fmt.Println("Esse espaço já foi preenchido, escolha outro por favor.")
		return jogada(l1, l2, l3)
	}
}

func imprimeTabuleiro(l1 [3]string, l2 [3]string, l3 [3]string) {
	ClearTerminal()
	fmt.Println("A   ", l1[0], " | ", l1[1], " | ", l1[2], " ")
	fmt.Println("   ---------------")
	fmt.Println("B   ", l2[0], " | ", l2[1], " | ", l2[2], " ")
	fmt.Println("   ---------------")
	fmt.Println("C   ", l3[0], " | ", l3[1], " | ", l3[2], " ")
	fmt.Println("")
	fmt.Println("    1   2    3")
}

func quemGanhou(desenho string) {
	if desenho == "X" {
		fmt.Println("\nJogador 1 venceu a partida. Parabéns!!!")
		time.Sleep(3 * time.Second)
		ClearTerminal()
		menu()
	} else {
		fmt.Println("\nJogador 2 venceu a partida. Parabéns!!!")
		time.Sleep(3 * time.Second)
		ClearTerminal()
		menu()
	}
}

func checkVitoria(l1 [3]string, l2 [3]string, l3 [3]string) {

	if l1[0] != "" && l1[0] == l1[1] && l1[1] == l1[2] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[0])
	} else if l2[0] != "" && l2[0] == l2[1] && l2[1] == l2[2] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l2[0])
	} else if l3[0] != "" && l3[0] == l3[1] && l3[1] == l3[2] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l3[0])
	} else if l1[0] != "" && l1[0] == l2[0] && l2[0] == l3[0] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[0])
	} else if l1[1] != "" && l1[1] == l2[1] && l2[1] == l3[1] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[1])
	} else if l1[2] != "" && l1[2] == l2[2] && l2[2] == l3[2] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[2])
	} else if l1[0] != "" && l1[0] == l2[1] && l2[1] == l3[2] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[0])
	} else if l1[2] != "" && l1[2] == l2[1] && l2[1] == l3[0] {
		imprimeTabuleiro(l1, l2, l3)
		quemGanhou(l1[2])
	} else if l1[0] != "" && l1[1] != "" && l1[2] != "" && l2[0] != "" && l2[1] != "" && l2[2] != "" && l3[0] != "" && l3[1] != "" && l3[2] != "" {
		imprimeTabuleiro(l1, l2, l3)
		fmt.Println("\nDeu velha... :/")
		time.Sleep(3 * time.Second)
		ClearTerminal()
		menu()
	}
}

func jogada(l1 [3]string, l2 [3]string, l3 [3]string) string {
	fmt.Println("Sua jogada :")
	var comando string
	fmt.Scan(&comando)
	comando = strings.ToUpper(comando)

	switch comando {
	case "A1":
		return checaComando(l1[0], comando, l1, l2, l3)
	case "A2":
		return checaComando(l1[1], comando, l1, l2, l3)
	case "A3":
		return checaComando(l1[2], comando, l1, l2, l3)
	case "B1":
		return checaComando(l2[0], comando, l1, l2, l3)
	case "B2":
		return checaComando(l2[1], comando, l1, l2, l3)
	case "B3":
		return checaComando(l2[2], comando, l1, l2, l3)
	case "C1":
		return checaComando(l3[0], comando, l1, l2, l3)
	case "C2":
		return checaComando(l3[1], comando, l1, l2, l3)
	case "C3":
		return checaComando(l3[2], comando, l1, l2, l3)
	default:
		imprimeTabuleiro(l1, l2, l3)
		fmt.Println("Comando invalido, por favor tente novamente.")
		return jogada(l1, l2, l3)
	}
}

func game() {
	jogadorAtual := 1
	desenho := "X"

	var A1 string
	var A2 string
	var A3 string
	var B1 string
	var B2 string
	var B3 string
	var C1 string
	var C2 string
	var C3 string

	linha1 := [3]string{A1, A2, A3}
	linha2 := [3]string{B1, B2, B3}
	linha3 := [3]string{C1, C2, C3}

	for {
		imprimeTabuleiro(linha1, linha2, linha3)

		fmt.Println("\njogador ", jogadorAtual, ", sua vez.\n ")

		comando := jogada(linha1, linha2, linha3)

		switch comando {
		case "A1":
			linha1[0] = desenho
		case "A2":
			linha1[1] = desenho
		case "A3":
			linha1[2] = desenho
		case "B1":
			linha2[0] = desenho
		case "B2":
			linha2[1] = desenho
		case "B3":
			linha2[2] = desenho
		case "C1":
			linha3[0] = desenho
		case "C2":
			linha3[1] = desenho
		case "C3":
			linha3[2] = desenho
		}

		if jogadorAtual == 1 {
			jogadorAtual = 2
		} else {
			jogadorAtual = 1
		}

		if desenho == "X" {
			desenho = "O"
		} else {
			desenho = "X"
		}

		checkVitoria(linha1, linha2, linha3)

	}

}

func menu() {
	fmt.Println("1- Começar novo jogo.")
	fmt.Println("2- Sair do programa")

	var comando int
	fmt.Scan(&comando)

	if comando == 1 {
		ClearTerminal()
		game()
	} else if comando == 2 {
		fmt.Println("Saindo do programa, espero que tenha gostado.")
		os.Exit(0)
	} else {
		fmt.Println("Comando não reconhecido, tente novamente por favor.")
		time.Sleep(3 * time.Second)
		ClearTerminal()
		menu()
	}
}

func main() {
	fmt.Println("Bem vindo ao jogo da velha.")
	menu()
}
