package word

import (
	"database/sql"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/ninetypercentlanguage/entities/word"
)

func GetWordSubRoute(db *sql.DB, wordsHost string) func(chi.Router) {
	return func(router chi.Router) {
		router.Post("/", getAddKnownPartsOfSpeechRoute(db, wordsHost))
		router.Get("/", getGetKnownPartsOfSpeechRoute(db))
	}
}

func getAddKnownPartsOfSpeechRoute(db *sql.DB, wordsHost string) func(w http.ResponseWriter, r *http.Request) {
	// expect a body of { ids: [1, 2], add_tags: ['compound'] }
}

func getGetKnownPartsOfSpeechRoute(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	// return body of []PartOfSpeech (from word server)
	partsOfSpeech := make([]word.PartOfSpeech, 0)
	partsOfSpeech = append(partsOfSpeech word.PartOfSpeech{
		ID: 1,
		PartOfSpeech: "noun",
		Lemmas: make([]word.Lemma, 0),
	})
}
