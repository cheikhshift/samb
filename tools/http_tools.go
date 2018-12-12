package tools

import (
	"net/http"
	"strings"
)

// ShortenPath updates the path of a request
// by matching and replacing the string (prefix)
// supplied with nothing.
func ShortenPath(prefix string,r *http.Request){
	r.URL.Path = strings.Replace(r.URL.Path, prefix, "", 1)
}

