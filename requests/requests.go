package requests

import (
	"fmt"
	v2 "math/rand/v2"
	"net/http"
	"slices"
	"time"
)

func Solicitud(url string, linea string, timeout int, usr string) (string, int) {
	cliente := http.Client{Timeout: time.Second * time.Duration(timeout)}

	urlexp := fmt.Sprintf("%s/%s", url, linea) // url que se prueba

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
