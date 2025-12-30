# GO-WebRaptor

Reescritura en Golang de mi herramienta de fuzzing web "WebRaptor" como ejercicio personal

para mas info ver: https://github.com/Urban20/WebRaptor

NOTA IMPORTANTE: utilizar con cuidado, el uso inadecuado puede resultar en el bloqueo de tu ip por parte del sitio

## Objetivo:
mayor estabilidad, rendimiento y velocidad 

## Flags imprementadas:
```
-dic string
        diccionario a utilizar en formato.txt
  -hl int
        concurrencia (default 500)
  -sd
        habilita la busqueda de subdominios
  -t int
        tiempo de espera de cada solicitud (default 5)
  -url string
        url del sitio a analizar
  -usr string
        user-agent a utilizar (default "Chrome/51.0.2704.106 Safari/537.36 OPR/38.0.2220.41")
  ```      
