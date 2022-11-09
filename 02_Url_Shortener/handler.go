package urlshort

import (
	"net/http"
)

func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		path := r.URL.Path
		if destination, ok := pathsToUrls[path]; ok { // se 'ok'
			http.Redirect(w, r, destination, http.StatusFound)
			return // deu bom, já redirecionamos
		}
		fallback.ServeHTTP(w, r)
	}
}
