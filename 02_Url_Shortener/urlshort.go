package main

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		// declarar uma variável 'destination' com um map lookup
		if destination, ok := pathsToUrls[path]; ok { // se 'ok'
			// redirecione URL
			http.Redirect(w, r, destination, http.StatusFound)
			// deu bom, já redirecionamos, não queremos chegar no fallback
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

func YAMLHandler(yamlBytes []byte, fallback http.Handler) (http.HandlerFunc, error) {

	// vamos ler os yamlBytes, e extrair dele pares {Path, URL}
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yamlBytes, &pathUrls)

	if err != nil {
		fmt.Println("deu ruim no unmarshal")
		return nil, err
	}

	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		pathsToUrls[pu.Path] = pu.URL
	}

	return MapHandler(pathsToUrls, fallback), nil
}

type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
