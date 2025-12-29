package requests

import (
	"fmt"
	"net/http"
	"slices"
	"time"
)

func Solicitud(url string, linea string, timeout int) (string, int) {
	cliente := http.Client{Timeout: time.Second * time.Duration(timeout)}

	urlexp := fmt.Sprintf("%s/%s", url, linea) // url que se prueba

	req, reqerr := http.NewRequest(http.MethodGet, urlexp, nil)

	if reqerr == nil {

		req.Header.Set("user-agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/143.0.0.0 Safari/537.36 Edg/143.0.100.0")
		resp, resperr := cliente.Do(req)

		if resperr != nil {
			return "", 0
		}

		if !slices.Contains([]int{404, 500}, resp.StatusCode) {

			return urlexp, resp.StatusCode

		}
	}

	return "", 0
}
