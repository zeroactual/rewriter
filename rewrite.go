package rewriter

import (
	"net/http"
	"strings"
)

func Rewrite(h http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		if r.RequestURI != "/" {
			if r.RequestURI[len(r.RequestURI)-1:] == "/" {
				s := strings.TrimRight(r.URL.Path, "/")
				http.Redirect(w, r, s, http.StatusFound)
				return
			}
		}
		h.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
