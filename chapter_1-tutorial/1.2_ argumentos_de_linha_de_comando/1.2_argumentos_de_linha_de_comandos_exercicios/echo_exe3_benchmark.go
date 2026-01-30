// Exercicio 1.3: Benchmark melhorado para medir diferença de desempenho
// Este programa repete cada operação muitas vezes para ver a diferença real

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Criar um slice grande de strings para testar
	args := make([]string, 1000)
	for i := 0; i < 1000; i++ {
		args[i] = fmt.Sprintf("argumento_%d", i)
	}

	iterations := 1000 // Repetir 1000 vezes

	// Versão 1: Concatenação com for indexado
	start1 := time.Now()
	for n := 0; n < iterations; n++ {
		var s, sep string
		for i := 0; i < len(args); i++ {
			s += sep + args[i]
			sep = " "
		}
		_ = s // Usar a variável para evitar otimização do compilador
	}
	duration1 := time.Since(start1)

	// Versão 2: Concatenação com range
	start2 := time.Now()
	for n := 0; n < iterations; n++ {
		var s, sep string
		for _, arg := range args {
			s += sep + arg
			sep = " "
		}
		_ = s
	}
	duration2 := time.Since(start2)

	// Versão 3: strings.Join (mais eficiente)
	start3 := time.Now()
	for n := 0; n < iterations; n++ {
		s := strings.Join(args, " ")
		_ = s
	}
	duration3 := time.Since(start3)

	// Resultados
	fmt.Println("=== Benchmark: 1000 strings, 1000 iterações ===")
	fmt.Printf("Versão 1 (for indexado):     %v\n", duration1)
	fmt.Printf("Versão 2 (range):            %v\n", duration2)
	fmt.Printf("Versão 3 (strings.Join):     %v\n", duration3)
	fmt.Println("\n=== Comparação ===")
	fmt.Printf("Versão 1 é %.2fx mais lenta que strings.Join\n", float64(duration1)/float64(duration3))
	fmt.Printf("Versão 2 é %.2fx mais lenta que strings.Join\n", float64(duration2)/float64(duration3))
}
