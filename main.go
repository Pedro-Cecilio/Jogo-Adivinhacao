package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

func main() {
	limparTerminal()
	for {
		exibeIntroducao()
		var Qt_tentativas, tentativa int
		numAleatorio := gerarNumAleatorio()
		pegarTentativa(numAleatorio, &tentativa, &Qt_tentativas)
		exibeOpcoes()
		opcao := pegarOpcao()
		if opcao == 1 {
			limparTerminal()
		}
		if opcao == 0 {
			limparTerminal()
			break
		}
	}
}
func exibeIntroducao() {
	fmt.Println("==== Jogo de Adivinhação ====")
	fmt.Println("Seu objetivo é adivinhar o número que nós sorteamos.")
	fmt.Println("Valor sorteado vai de 1 a 1000.")
	fmt.Println("Para desistir, basta inserir 0 como resposta. Boa sorte!!")
}
func gerarNumAleatorio() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(1000) + 1
}
func pegarTentativa(resposta int, tent, contagemTentativas *int) {
	for {
		fmt.Println("Insira sua", *contagemTentativas+1, "tentativa:")
		fmt.Scan(tent)
		*contagemTentativas++
		acertou, desistiu := confereTentativa(resposta, *tent, *contagemTentativas)
		if acertou || desistiu == 1 {
			break
		}
	}
}
func confereTentativa(resposta, tent, qt_tentativas int) (bool, int) {
	if tent == resposta {
		fmt.Println("Você acertou!! Parabéns!")
		fmt.Println("Quantidade de Tentativas:", qt_tentativas)
		return true, 0
	} else {
		if tent == 0 {
			fmt.Println("Uma pena sua desistência... Nos vemos na próxima!")
			return false, 1
		}
		if tent > resposta {
			fmt.Println("Chutou alto...")
		} else {
			fmt.Println("Chutou Baixo...")
		}
		return false, 0
	}
}
func exibeOpcoes() {
	fmt.Println("1 - Jogar novamente.", "\n"+"0 - Sair.")
}
func pegarOpcao() int {
	var opcao int
	fmt.Scan(&opcao)
	return opcao
}
func limparTerminal() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}
