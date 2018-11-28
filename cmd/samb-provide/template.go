package main

var providerTemplate = `

provider {
	name %s;
	type %s;
	return provider%s(w,r);
}`

var fnTemplate = `package api

import "net/http"

// %s does...
func provider%s(w http.ResponseWriter, r *http.Request) %s {

	return nil
}
`

var nameTemplate = "provider_%s.go"
