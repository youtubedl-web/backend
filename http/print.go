package http

import (
	"encoding/json"
	"net/http"
)

func jsonPrint(w http.ResponseWriter, v interface{}) (int, error) {
	buf, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		return http.StatusInternalServerError, err
	}

	w.Write(buf)

	return 0, nil
}
