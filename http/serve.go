package http

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/youtubedl-web/backend"
)

// Serve starts the HTTP server
func Serve(c *backend.Config) {
	r := mux.NewRouter()

	r.HandleFunc("/link/{url}/audio", Wrap(GetAudioLink, c)).Methods("GET", "OPTIONS")
	r.HandleFunc("/link/{url}/video", Wrap(GetVideoLink, c)).Methods("GET", "OPTIONS")

	r.HandleFunc("/dl/{hash}", Wrap(DownloadFile, c)).Methods("GET", "OPTIONS")

	// application status check
	r.HandleFunc("/ping", Wrap(APIStatus, c)).Methods("GET", "OPTIONS")

	fmt.Printf("Server running on port ")
	color.Green(strconv.Itoa(c.Port))
	http.ListenAndServe(":"+strconv.Itoa(c.Port), r)
}
