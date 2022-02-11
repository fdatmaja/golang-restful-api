package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()
	if err != nil {
		errRoll := tx.Rollback()
		if errRoll != nil {
			panic(errRoll)
		}
		panic(err)
	} else {
		errComm := tx.Commit()
		if errComm != nil {
			panic(errComm)
		}
	}
}
