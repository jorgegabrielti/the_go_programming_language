// Exercicio 1.2: Modifique o programa echo par exibir o índice e o valor de cada um de seus argumentos, um por linha.
package main

import (
	"fmt"
	"os"
)

func main() {
	for index, arg := range os.Args {
		fmt.Println(index, arg)
	}
}
