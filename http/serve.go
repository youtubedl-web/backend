package http

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youtubedl-web/backend"
)

// Serve starts the HTTP server
func Serve(c *backend.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/", Wrap(DownloadMP3, c)).Methods("GET")

	http.ListenAndServe(c.Host+":"+strconv.Itoa(c.Port), r)
}
