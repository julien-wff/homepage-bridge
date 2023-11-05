package routes

import (
	"github.com/julien-wff/homepage-bridge/internal/collectors"
	"github.com/julien-wff/homepage-bridge/internal/helpers"
	"net/http"
)

func HandleMinio(res http.ResponseWriter, _ *http.Request) {
	collected, err := collectors.MinioCollector()
	if err != nil {
		helpers.ErrorResponse(&res, err, http.StatusInternalServerError)
	} else {
		helpers.JsonResponse(&res, collected, http.StatusOK)
	}
}
