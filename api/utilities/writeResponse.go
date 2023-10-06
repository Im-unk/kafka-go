package utilities

import (
	"encoding/json"
	"net/http"
)

func writeResoponse(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
