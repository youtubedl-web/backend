package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/youtubedl-web/backend"
	"github.com/youtubedl-web/backend/file"
)

// GetAudioLink creates a link for the user to download the audio file from a youtube video
func GetAudioLink(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	vars := mux.Vars(r)
	videoURL := vars["url"]

	// prepare storage dir
	hash, ok := file.GenerateHash(c.Storage)
	if ok != 1 {
		c.Logger.Errorf("Couldn't generate hash and storage folder")
	}

	// change dir to the storage one
	err := os.Chdir(filepath.Join(c.Storage, hash))
	if err != nil {
		c.Logger.Errorf("Couldn't find storage folder")
	}

	// log download
	c.Logger.Infof("[%s] Fetching audio download link for %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL)

	// prepare command
	cmd := exec.Command(c.ExecutablePath, "-f bestaudio", "--extract-audio", `-o '%(title)s.%(ext)s'`, "https://www.youtube.com/watch?v="+videoURL)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("Couldn't connect stdout pipe")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		color.Red("Couldn't connect stderr pipe")
	}

	// run command
	err = cmd.Start()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get audio download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL, err.Error())
	}

	stdoutOutput, _ := ioutil.ReadAll(stdout)
	stderrOutput, _ := ioutil.ReadAll(stderr)

	// generate download link
	link := &backend.Link{
		URL: "",
	}

	if link.URL = "http://"; c.Secure {
		link.URL = "https://"
	}

	link.URL += c.PublicHost + ":" + strconv.Itoa(c.Port) + "/dl/" + hash

	// watch command exit status
	err = cmd.Wait()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get audio download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL, err.Error())
		c.Logger.Warnf("[%s] Standard Output and Error:\n\tStdout %s\n\tStderr: %s", r.RemoteAddr, stdoutOutput, stderrOutput)

		if exiterr, ok := err.(*exec.ExitError); ok {
			fmt.Println(exiterr)
		}

		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, link)
}

// GetVideoLink creates a link for the user to download the video file from a youtube video
func GetVideoLink(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	vars := mux.Vars(r)
	videoURL := vars["url"]

	// prepare storage dir
	hash, ok := file.GenerateHash(c.Storage)
	if ok != 1 {
		c.Logger.Errorf("Couldn't generate hash and storage folder")
	}

	// change dir to the storage one
	err := os.Chdir(filepath.Join(c.Storage, hash))
	if err != nil {
		c.Logger.Errorf("Couldn't find storage folder")
	}

	// log video download
	c.Logger.Infof("[%s] Fetching video download link for %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL)

	// run youtube-dl
	cmd := exec.Command(c.ExecutablePath, "-f bestvideo+bestaudio", "https://www.youtube.com/watch?v="+videoURL)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("Couldn't connect stdout pipe")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		color.Red("Couldn't connect stderr pipe")
	}

	// start commandvideo
	err = cmd.Start()

	if err != nil {
		c.Logger.Errorf("[%s] Cannot get video download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL, err.Error())
	}

	stdoutOutput, _ := ioutil.ReadAll(stdout)
	stderrOutput, _ := ioutil.ReadAll(stderr)

	// generate download link
	link := &backend.Link{
		URL: c.PublicHost + "/dl/" + hash,
	}

	err = cmd.Wait()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get video download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch?v="+videoURL, err.Error())
		c.Logger.Warnf("[%s] Standard Output and Error:\n\tStdout %s\n\tStderr: %s", r.RemoteAddr, stdoutOutput, stderrOutput)

		if exiterr, ok := err.(*exec.ExitError); ok {
			fmt.Println(exiterr)
		}

		return http.StatusInternalServerError, err
	}

	return jsonPrint(w, link)
}
