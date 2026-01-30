// Exercicio 1.3: Experimente mdedir a diferença de tempo de execução entre nossas versões potencialmente ineficientes e a versão que usa strings.Join. (A seção 1.6 modstra parte do pacote time, e a seção 11.4 mostra como escrever testes comparativos para uma avaliação sistemática de desempenho.)

package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	//versão 1: Ineficiente (concatenação com for indexado)
	start1 := time.Now()
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	duration1 := time.Since(start1)
	fmt.Printf("Versão 1 (for indexado):     %v (%d ns)\n", duration1, duration1.Nanoseconds())

	//versão 2: (concatenação com range)
	start2 := time.Now()
	s, sep = "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	duration2 := time.Since(start2)
	fmt.Printf("Versão 2 (range):            %v (%d ns)\n", duration2, duration2.Nanoseconds())

	//versão 3: (strings.Join - mais eficiente)
	start3 := time.Now()
	s = strings.Join(os.Args[1:], " ")
	duration3 := time.Since(start3)
	fmt.Printf("Versão 3 (strings.Join):     %v (%d ns)\n", duration3, duration3.Nanoseconds())
}
