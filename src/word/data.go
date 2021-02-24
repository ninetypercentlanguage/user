package word

import "database/sql"

type addRequest struct {
	Ids []int `json:"ids"`
}

func getAdder(db *sql.DB, wordsHost string) func(int, addRequest) error {
	return func(userId int, a addRequest) error {
		// todo validate that the ids are real against the word host
		return addKnownPartsofSpeechToDB(db, userId, a)
	}
}

func getGetter(db *sql.DB) func(int) ([]int, error) {
	return func(userId int) ([]int, error) {
		return getKnownPartsofSpeechFromDB(db, userId)
	}
}
