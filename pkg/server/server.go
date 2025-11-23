package server

import (
	"diploma/pkg/api"
	"log"
	"net/http"
)

func Run(p string, wD string) error {
	log.Printf("Starting web server on port %s, serving files from %s", p, wD)
	api.Init()
	http.Handle("/", http.FileServer(http.Dir(wD)))

	return http.ListenAndServe(":"+p, nil)

}
