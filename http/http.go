package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/fatih/color"
	"github.com/youtubedl-web/backend"
)

// RequestHandler is a special type of function that handles HTTP requests and has the application config in its scope
type RequestHandler func(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error)

type response struct {
	Code int    `json:"code"`
	URL  string `json:"url,omitempty"`
}

// Wrap is a handler for the API methods which converts them into standard HandlerFunc
func Wrap(h RequestHandler, c *backend.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// LOG REQUESTS
		LogRequest(r.Method, r.URL.Path, r.RemoteAddr)

		allowedOrigins := c.AllowedOrigins
		if c.Development {
			allowedOrigins = c.AllowedOriginsDevelopment
		}

		// handle pre-flight requests
		w.Header().Set("Access-Control-Allow-Headers", "content-type")
		w.Header().Set("Access-Control-Allow-Origin", allowedOrigins)
		w.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS")

		// don't waste computing efforts when the request method is OPTIONS
		if r.Method == "OPTIONS" {
			return
		}

		var (
			code int
			err  error
		)

		// this piece of code is specialised in handling the results
		// from the execution of the RequestHandler functions
		defer func() {
			// whenever a function writes anything into the response
			// don't treat the response body as usual
			if code == 0 && err == nil {
				return
			}

			res := response{
				Code: code,
			}

			// marshal struct into json and write it into the response
			// set content-type
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(code)

			// marshal and write
			buf, err := json.MarshalIndent(res, "", "  ")
			if err != nil {
				color.Red(err.Error())
				return
			}
			w.Write(buf)

			return
		}()

		code, err = h(w, r, c)
	}
}

// LogRequest prints every request info (method, url and request remote address)
func LogRequest(method string, url string, reqAddr string) {
	fmt.Printf("[")
	color.New(color.FgGreen, color.Bold).Printf("INFO")
	fmt.Printf("]")

	fmt.Printf(" %s request to %s by %s\n", strings.ToUpper(method), url, reqAddr)
}
