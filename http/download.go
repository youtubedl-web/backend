package http

import (
	"io"
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

	// open file
	f, err := os.Open(fixedFilename)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// get file size
	fi, err := os.Stat(fixedFilename)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	// convert size int64 to string
	size := strconv.FormatInt(fi.Size(), 10)

	// Set important headers
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", `attachment; filename="`+fixedFilename+`"`)
	w.Header().Set("Content-Length", size)

	// copy the file to the response
	io.Copy(w, f)

	// http.ServeFile(w, r, fixedFilename)

	return 0, nil
}
