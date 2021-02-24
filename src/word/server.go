package word

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/ninetypercentlanguage/user-server/serverutils"

	"github.com/go-chi/chi"
)

func GetWordSubRoute(db *sql.DB, wordsHost string) func(chi.Router) {
	return func(router chi.Router) {
		router.Put("/", getAddKnownPartsOfSpeechRoute(db, wordsHost))
		router.Get("/", getGetKnownPartsOfSpeechRoute(db))
	}
}

func getAddKnownPartsOfSpeechRoute(db *sql.DB, wordsHost string) func(http.ResponseWriter, *http.Request) {
	add := getAdder(db, wordsHost)

	return func(writer http.ResponseWriter, r *http.Request) {
		var toAdd addRequest
		dec := json.NewDecoder(r.Body)
		dec.DisallowUnknownFields()
		err := dec.Decode(&toAdd)
		if err != nil {
			serverutils.HandleError(writer, fmt.Errorf("body is not an add request"), http.StatusBadRequest)
			return
		}
		userIDParam := chi.URLParam(r, "userId")
		userID, err := strconv.Atoi(userIDParam)
		if err != nil {
			serverutils.HandleError(writer, fmt.Errorf("userID %v is not integer", userID), http.StatusNotFound)
			return
		}
		err = add(userID, toAdd)
		if err != nil {
			serverutils.HandleError(writer, err, http.StatusInternalServerError)
			return
		}

		writer.WriteHeader(200)
	}
}

func getGetKnownPartsOfSpeechRoute(db *sql.DB) func(http.ResponseWriter, *http.Request) {
	get := getGetter(db)

	return func(writer http.ResponseWriter, r *http.Request) {
		rawID := chi.URLParam(r, "userId")
		userID, err := strconv.Atoi(rawID)
		if err != nil {
			serverutils.HandleError(writer, fmt.Errorf("userID %v is not integer", userID), http.StatusNotFound)
			return
		}
		res, err := get(userID)
		if err != nil {
			serverutils.HandleError(writer, err, http.StatusInternalServerError)
			return
		}
		responseJson, err := json.Marshal(res)
		if err != nil {
			serverutils.HandleError(writer, err, http.StatusInternalServerError)
			return
		}
		writer.Write([]byte(responseJson))
	}
}
