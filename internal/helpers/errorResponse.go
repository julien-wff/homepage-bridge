package helpers

import "net/http"

func ErrorResponse(res *http.ResponseWriter, err error, statusCode int) {
	JsonResponse(res, map[string]string{"error": err.Error()}, statusCode)
}
