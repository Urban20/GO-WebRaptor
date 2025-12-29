package main

import (
	"GoWebRaptor/diccionarios"
	"GoWebRaptor/requests"
	"flag"
	"fmt"
	"sync"
)

const RUTA = "diccionario_txt/"

var timeout_f = flag.Float64("t", 5, "tiempo de espera de cada solicitud")
var hilos_f = flag.Int("hl", 500, "concurrencia")
var url_f = flag.String("url", "", "url del sitio a analizar")
var dic_f = flag.String("dic", "", "diccionario a utilizar en formato.txt")

var timeout = *timeout_f
var hilos = *hilos_f
var url = *url_f
var diccionario = *dic_f

func main() {

	flag.Parse()

	limite := make(chan struct{}, hilos)
	wg := sync.WaitGroup{}
	dic_ruta := fmt.Sprintf("%s/%s", RUTA, diccionario)
	diccionario := diccionarios.Leer(dic_ruta)

	for linea := range diccionario {
		limite <- struct{}{}
		wg.Add(1)

		go func() {
			defer func() { <-limite }()
			defer wg.Done()
			url_encontrado, codigo := requests.Solicitud(url, linea, timeout)
			if codigo != 0 {

				fmt.Printf("%s status >> %d\n", url_encontrado, codigo)
			}

		}()
	}
	wg.Wait()
}
