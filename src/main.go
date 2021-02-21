package main

import (
	"log"
	"net/http"

	"github.com/ninetypercentlanguage/user-server/db"

	"github.com/ninetypercentlanguage/user-server/getflags"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/ninetypercentlanguage/user-server/word"
)

func main() {
	flagVals := getflags.GetFlags()

	r := chi.NewRouter()
	r.Use(middleware.Logger)

	d := db.InitDB(flagVals.DB)

	r.Route("/{userId}", func(subrouter chi.Router) {
		// get words you know
		// add words you know
		subrouter.Route("/words", word.GetWordSubRoute(d, flagVals.Words))
	})

	log.Println("Serving from 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
