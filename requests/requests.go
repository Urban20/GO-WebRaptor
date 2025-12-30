package requests

import (
	"fmt"
	v2 "math/rand/v2"
	"net/http"
	"regexp"
	"slices"
	"strings"
	"time"
)

// crea un url nuevo con el subdominio a probar y la url que se le pasa
func subdominios(url, linea string) (string, error) {
	exp, experr := regexp.Compile(`https?\://(\w+)\.`)
	if experr != nil {
		return "", experr
	}
	sd := exp.FindStringSubmatch(url)[1]
	return strings.Replace(url, sd, linea, 1), nil

}

func nueva_solicitud(cliente http.Client, urlexp string, usr string) (string, int) {

	req, reqerr := http.NewRequest(http.MethodGet, urlexp, nil)
	suspendido := v2.Float32()

	if reqerr == nil {

		req.Header.Set("user-agent", usr)
		time.Sleep(time.Duration(suspendido) * time.Second)
		resp, resperr := cliente.Do(req)

		if resperr != nil {
			return "", 0
		}

		if !slices.Contains([]int{404, 500, 502}, resp.StatusCode) {

			return urlexp, resp.StatusCode

		}
	}
	return "", 0
}

func Solicitud(url string, linea string, timeout int, usr string, sd bool) (string, int) {
	cliente := http.Client{Timeout: time.Second * time.Duration(timeout)}

	if sd {

		urlexp, sderr := subdominios(url, linea)
		if sderr != nil {
			return "", 0
		}

		url, codigo := nueva_solicitud(cliente, urlexp, usr)
		return url, codigo

	}

	urlexp := fmt.Sprintf("%s/%s", url, linea)
	url, codigo := nueva_solicitud(cliente, urlexp, usr)
	return url, codigo

}
