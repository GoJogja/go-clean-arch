package httphandler

import (
	"encoding/json"
	"net/http"
)

func JsonOK(w http.ResponseWriter, payload *JsonSuccessResponse) (err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(
		payload,
	)
	return
}
