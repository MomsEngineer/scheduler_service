package handlers

import "net/http"

func Healthz(res http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodGet {
		http.Error(res, "Only GET requests are allowed!", http.StatusMethodNotAllowed)
		return
	}

	res.WriteHeader(http.StatusOK)
}
