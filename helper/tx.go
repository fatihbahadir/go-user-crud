package helper

import "database/sql"

func CommitOrRollback(tx *sql.Tx) {
	err := recover()

	if err != nil {
		errRollBack := tx.Rollback()
		HandleError(errRollBack, "Failed to rollback transaction")
		panic(err)
	} else {
		errCommit := tx.Commit()
		HandleError(errCommit, "Failed to commit transaction")
	}

}