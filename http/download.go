package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/youtubedl-web/backend"
)

// DownloadFile starts the download of the file requested
func DownloadFile(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	vars := mux.Vars(r)
	hash := vars["hash"]

	// change dir to file storage folder
	os.Chdir(filepath.Join(c.Storage, hash))

	// get a list of files - the only one file is the file the user wants to download
	files, err := ioutil.ReadDir("./")
	if err != nil {
		log.Fatal(err)
	}

	// file size
	size := strconv.Itoa(int(files[0].Size()))

	// Set important headers
	w.Header().Set("Content-Disposition", "attachment; filename="+files[0].Name())
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", size)
	http.ServeFile(w, r, files[0].Name())

	return 0, nil
}
