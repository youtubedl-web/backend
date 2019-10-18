package http

import (
	"io/ioutil"
	"net/http"
	"os/exec"
	"strings"

	"github.com/fatih/color"
	"github.com/gorilla/mux"
	"github.com/youtubedl-web/backend"
)

// GetAudioLink creates a link for the user to download the MP3 audio file from a youtube video
func GetAudioLink(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	vars := mux.Vars(r)

	videoURL := vars["url"]

	// log video download
	c.Logger.Infof("[%s] Fetching audio download link for %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL)

	// run youtube-dl
	cmd := exec.Command(c.ExecutablePath, "--get-url", "-x", "https://www.youtube.com/watch?v="+videoURL)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		color.Red("Couldn't connect stdout pipe")
	}

	stderr, err := cmd.StderrPipe()
	if err != nil {
		color.Red("Couldn't connect stderr pipe")
	}

	// start command
	err = cmd.Start()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get audio download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL, err.Error())
	}

	stdoutOutput, _ := ioutil.ReadAll(stdout)
	stderrOutput, _ := ioutil.ReadAll(stderr)

	// separate all linkss
	links := strings.Split(string(stdoutOutput), "\r")

	link := &backend.Link{
		URL: links[0],
	}

	err = cmd.Wait()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get audio download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL, err.Error())
		c.Logger.Warnf("[%s] Standard Output and Error:\n\tStdout %s\n\tStderr: %s", r.RemoteAddr, stdoutOutput, stderrOutput)
	}

	return jsonPrint(w, link)
}

// GetVideoLink creates a link for the user to download the MP3 video file from a youtube video
func GetVideoLink(w http.ResponseWriter, r *http.Request, c *backend.Config) (int, error) {
	vars := mux.Vars(r)

	videoURL := vars["url"]

	// log video download
	c.Logger.Infof("[%s] Fetching video download link for %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL)

	// run youtube-dl
	cmd := exec.Command(c.ExecutablePath, "--get-url", "https://www.youtube.com/watch?v="+videoURL)
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
		c.Logger.Errorf("[%s] Cannot get video download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL, err.Error())
	}

	stdoutOutput, _ := ioutil.ReadAll(stdout)
	stderrOutput, _ := ioutil.ReadAll(stderr)

	links := strings.Split(string(stdoutOutput), "\n")

	link := &backend.Link{
		URL: links[0],
	}

	err = cmd.Wait()
	if err != nil {
		c.Logger.Errorf("[%s] Cannot get video download link for %s. Error: %s", r.RemoteAddr, "http://www.youtube.com/watch? ="+videoURL, err.Error())
		c.Logger.Warnf("[%s] Standard Output and Error:\n\tStdout %s\n\tStderr: %s", r.RemoteAddr, stdoutOutput, stderrOutput)
	}

	return jsonPrint(w, link)
}
