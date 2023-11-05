package routes

import (
	"fmt"
	"net/http"
)

func Router() {
	// Router using standard library
	http.HandleFunc("/api/v1/minio", HandleMinio)

	fmt.Println("Listening on port 8080")

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
