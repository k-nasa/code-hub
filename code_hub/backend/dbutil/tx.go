package dbutil

import (
	"fmt"
	"log"
	"net/http"

	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
)

type DuplicationCodeError struct{}

func (err *DuplicationCodeError) Error() string { return "duplication" }

// TXHandler is handler for working with transaction.
// This is wrapper function for commit and rollback.
func TXHandler(db *sqlx.DB, f func(*sqlx.Tx) error) error {
	tx, err := db.Beginx()
	if err != nil {
		return errors.Wrap(err, "start transaction failed")
	}
	defer func() {
		if err := recover(); err != nil {
			tx.Rollback()
			log.Print("rollback operation.")
			return
		}
	}()
	if err := f(tx); err != nil {
		switch err.(type) {
		case *DuplicationCodeError:
			return err
		default:
			return errors.Wrap(err, "transaction: operation failed")
		}
	}
	return nil
}

func Error(w http.ResponseWriter, err error, code int) {
	http.Error(w, fmt.Sprintf("%s", err), code)
}
