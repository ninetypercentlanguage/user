package word

import (
	"database/sql"
	"fmt"
)

func addKnownPartsofSpeechToDB(db *sql.DB, userId int, request addRequest) error {
	valuesQuery := ""
	for _, id := range request.Ids {
		valuesQuery += fmt.Sprintf("(%v, %v)", userId, id)
	}

	_, err := db.Query(fmt.Sprintf(`INSERT INTO user_known_parts_of_speech ("user", part_of_speech) VALUES %v`, valuesQuery))
	if err != nil {
		panic(err)
	}
	return nil
}

func getKnownPartsofSpeechFromDB(db *sql.DB, userId int) ([]int, error) {
	rows, err := db.Query(`SELECT part_of_speech FROM user_known_parts_of_speech WHERE "user" = $1`, userId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	partsOfSpeech := make([]int, 0)
	for rows.Next() {
		var id int
		err = rows.Scan(&id)
		if err != nil {
			panic(err)
		}
		partsOfSpeech = append(partsOfSpeech, id)
	}
	return partsOfSpeech, nil
}
