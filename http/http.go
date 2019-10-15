package http

import (
	"net/http"

	"github.com/youtubedl-web/backend"
)

// RequestHandler is a special type of function that handles HTTP requests and has the application config in its scope
type RequestHandler func(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error)

func Wrap(h RequestHandler, c *backend.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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

			return
		}()

		code, err = h(w, r, c)
	}
}
