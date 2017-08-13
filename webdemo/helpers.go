package main

import (
	"fmt"
	"net/http"
)

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

// does nothing, for sinking favicon or other common paths
// that we don't care about
func nothingHandler(w http.ResponseWriter, r *http.Request) {

}
