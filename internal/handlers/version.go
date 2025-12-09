package internal

import (
	"encoding/json"
	"net/http"
)

func Version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := map[string]string{"version": "1.0.0"}
	json.NewEncoder(w).Encode(resp)
}
