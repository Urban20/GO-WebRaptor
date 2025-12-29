package diccionarios

import (
	"bufio"
	"errors"
	"os"
	"regexp"
)

func Leer(diccionario string) (chan string, error) {
	linea := make(chan string)

	match, marcherr := regexp.Match(`\w+\.txt$`, []byte(diccionario))

	if marcherr != nil || !match {
		return linea, errors.New("diccionario invalido")
	}

	arch, archerr := os.Open(diccionario)

	if archerr != nil {

		return linea, archerr
	}

	go func() {
		defer close(linea)

		escaner := bufio.NewScanner(arch)
		for escaner.Scan() {

			linea <- escaner.Text()

		}
	}()

	return linea, nil

}
