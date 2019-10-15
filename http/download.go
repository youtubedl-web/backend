package http

import (
	"net/http"

	"github.com/youtubedl-web/backend"
)

// DownloadMP3 creates a link for the user to download the MP3 audio file from a youtube video
func DownloadMP3(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	return http.StatusOK, nil
}
