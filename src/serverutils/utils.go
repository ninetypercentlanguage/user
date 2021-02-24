package serverutils

import (
	"net/http"
)

func HandleError(w http.ResponseWriter, err error, status int) {
	http.Error(w, err.Error(), status)
}

func DefaultHandleOptions(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
}
