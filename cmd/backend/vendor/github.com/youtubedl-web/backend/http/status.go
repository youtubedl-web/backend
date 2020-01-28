package http

import (
	"net/http"

	"github.com/youtubedl-web/backend"
)

// APIStatus simply writes 200 to the response body
// It should be used to verify if the app is OK
func APIStatus(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	w.Write([]byte("200"))

	return 0, nil
}
