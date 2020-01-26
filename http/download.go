package http

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

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

	// remove trailing whitespaces and the single quotes
	fixedFilename := strings.ReplaceAll(strings.Trim(files[0].Name(), "\t \n"), "'", "")

	// rename file to the correct format
	os.Rename(files[0].Name(), fixedFilename)

	// get file size
	fi, err := os.Stat(fixedFilename)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// file size
	size := strconv.FormatInt(fi.Size(), 10)

	// Set important headers
	w.Header().Set("Content-Disposition", "attachment; filename="+fixedFilename)
	// w.Header().Set("Content-Type", "audio/mpeg")
	w.Header().Set("Content-Length", size)
	http.ServeFile(w, r, fixedFilename)

	return 0, nil
}
