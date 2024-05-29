package utils

import (
	"encoding/json"
	"net/http"
)

func CreateAndWriteJSON(w http.ResponseWriter, v any) error {
	encoder := json.NewEncoder(w)
	return encoder.Encode(v)
}
