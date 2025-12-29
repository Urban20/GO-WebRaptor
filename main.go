package main

import (
	"GoWebRaptor/diccionarios"
	"GoWebRaptor/requests"
	"flag"
	"fmt"
	"sync"
)

const RUTA = "diccionario_txt/"

var usag_default string = "Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41 "
var timeout_f = flag.Int("t", 5, "tiempo de espera de cada solicitud")
var hilos_f = flag.Int("hl", 500, "concurrencia")
var url_f = flag.String("url", "", "url del sitio a analizar")
var dic_f = flag.String("dic", "", "diccionario a utilizar en formato.txt")
var usrAG = flag.String("usr", usag_default, "user-agent a utilizar")

const LOGO = ` 
··············
:GO-WebRaptor:
··············
reescrito en Golang`

func main() {
	flag.Parse()
	fmt.Print(LOGO + "\n\n")

	timeout := *timeout_f
	hilos := *hilos_f
	url := *url_f
	diccionario := *dic_f
	usr := *usrAG

	limite := make(chan struct{}, hilos)
	wg := sync.WaitGroup{}
	dic_ruta := fmt.Sprintf("%s/%s", RUTA, diccionario)
	dic := diccionarios.Leer(dic_ruta)

	for linea := range dic {
		limite <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() { <-limite }()
			defer wg.Done()
			url_encontrado, codigo := requests.Solicitud(url, linea, timeout, usr)
			if codigo != 0 {

				fmt.Printf("%s status >> %d\n", url_encontrado, codigo)
			}

		}()
	}
	wg.Wait()
}
