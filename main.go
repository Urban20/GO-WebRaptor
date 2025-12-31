package main

import (
	archivo "GoWebRaptor/archivos"
	"GoWebRaptor/diccionarios"
	"GoWebRaptor/requests"
	"flag"
	"fmt"
	"sync"
)

const RUTA = "diccionario_txt/"

var usag_default string = "Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41"
var timeout_f = flag.Int("t", 5, "tiempo de espera de cada solicitud")
var hilos_f = flag.Int("hl", 500, "concurrencia")
var url_f = flag.String("url", "", "url del sitio a analizar")
var dic_f = flag.String("dic", "", "ruta del diccionario a utilizar en formato.txt")
var usrAG = flag.String("usr", usag_default, "user-agent a utilizar")
var subdom = flag.Bool("sd", false, "habilita la busqueda de subdominios")

const LOGO = ` 
                                                       
 _____         _ _ _     _   _____         _           
|   __|___ ___| | | |___| |_| __  |___ ___| |_ ___ ___ 
|  |  | . |___| | | | -_| . |    -| .'| . |  _| . |  _|
|_____|___|   |_____|___|___|__|__|__,|  _|_| |___|_|  
                                      |_|              
reescrito en Golang`

func main() {
	flag.Parse()
	fmt.Print("\033[0;35m" + LOGO + "\n\n" + "\033[0m")

	timeout := *timeout_f
	hilos := *hilos_f
	url := *url_f
	diccionario := *dic_f
	usr := *usrAG
	sd := *subdom

	limite := make(chan struct{}, hilos)
	wg := sync.WaitGroup{}

	dic, dicerr := diccionarios.Leer(diccionario)

	if dicerr != nil {
		fmt.Println(dicerr)
		return
	}

	for linea := range dic {
		limite <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() { <-limite }()
			defer wg.Done()
			url_encontrado, codigo := requests.Solicitud(url, linea, timeout, usr, sd)
			if codigo != 0 {
				formato := fmt.Sprintf("%s status >> %d\n", url_encontrado, codigo)
				fmt.Println(formato)
				archerr := archivo.Archivar(formato)
				if archerr != nil {
					fmt.Println(archerr)
					return
				}
			}

		}()
	}
	wg.Wait()
	fmt.Printf("urls guardadas en %s\n", archivo.DIRECTORIOS)
}
