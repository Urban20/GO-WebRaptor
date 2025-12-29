package archivo

import "os"

const DIRECTORIOS = "directorios.txt"

func Archivar(url string) error {
	archivo, archerr := os.OpenFile(DIRECTORIOS, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if archerr != nil {
		return archerr
	}
	_, error_ := archivo.Write([]byte(url))
	if error_ != nil {
		return error_
	}
	return nil
}
