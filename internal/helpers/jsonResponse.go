package helpers

import (
	"encoding/json"
	"net/http"
)

func JsonResponse(res *http.ResponseWriter, body any, status int) {
	// Set header
	(*res).Header().Set("Content-Type", "application/json")

	// Convert JSON body to bytes
	bytes, err1 := json.Marshal(body)
	if err1 != nil {
		panic(err1)
	}

	// Write response
	(*res).WriteHeader(status)
	_, err2 := (*res).Write(bytes)
	if err2 != nil {
		panic(err2)
	}
}
