package main

import (
	"fmt"
	"net/http"
)

// inicia um server em localhost:8080
// que implementa um url shortener.
// fazendo localhost:8080/yaml-godoc por exemplo vai redirect pra pag corresp
// localhost:8080/invalid-path vai soh mostrar um hello world
func main() {
	// o mux serve de 'fallback', define o comportamento das urls
	// não cadastradas em pathsToUrls
	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := MapHandler(pathsToUrls, mux)

	// Build the YAMLHandler using the mapHandler as the
	// fallback

	yaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`

	yamlHandler, err := YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
}

// isso define que página é servida por mapHandler no path '/'
// e nos paths não definidos no pathsToUrls.
// no caso é a página definida por hello(w, r)
func defaultMux() *http.ServeMux {
	// ServeMux is an HTTP request multiplexer.
	// It matches the URL of each incoming request against a list of registered patterns,
	// and calls the handler for the pattern that most closely matches the URL
	mux := http.NewServeMux()
	//  the pattern "/" matches all paths not matched by other registered patterns
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Olá, enfermeira!")
}
