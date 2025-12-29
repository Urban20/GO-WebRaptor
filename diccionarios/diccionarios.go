package diccionarios

import (
	"bufio"
	"os"
)

func Leer(diccionario string) chan string {

	linea := make(chan string)
	arch, archerr := os.Open(diccionario)

	if archerr != nil {

		panic(archerr)
	}

	go func() {
		defer close(linea)

		escaner := bufio.NewScanner(arch)
		for escaner.Scan() {

			linea <- escaner.Text()

		}
	}()

	return linea

}
