package main

import (
	"os"
	"os/signal"
	"fmt"
	"time"
)

func main() {
	/* liberando recursos do programa de forma apropriada!
		O programa é executado e quando o programa e cancelado com ctrl + c,
		a funcao abaixo executa liberando
		recursos de maquina.
	*/
	go func() {
		sinal := make(chan os.Signal, 1)
		signal.Notify(sinal, os.Interrupt)

		<-sinal

		fmt.Println("liberando recursos...")
		fmt.Println("liberacao finalizada!")

		os.Exit(0)
	}()

	fmt.Println("control + c para finalizar")

	for {
		time.Sleep(1 * time.Second)
		fmt.Println(".")
	}
}
