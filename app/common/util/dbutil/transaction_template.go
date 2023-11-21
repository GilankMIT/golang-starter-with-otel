package dbutil

import (
	"database/sql"
)

func ExecInTransactionWithNoCallback(db *sql.DB, proc func(tx *sql.Tx) error) (err error) {
	var tx *sql.Tx
	tx, err = db.Begin()
	if err != nil {
		return
	}

	err = proc(tx)
	if err != nil {
		tx.Rollback()
		return
	}

	err = tx.Commit()
	if err != nil {
		tx.Rollback()
	}

	return
}
