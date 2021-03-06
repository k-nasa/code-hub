package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/k-nasa/code-hub/dbutil"
	"github.com/k-nasa/code-hub/model"
	"github.com/k-nasa/code-hub/repository"
	"github.com/pkg/errors"
)

type Code struct {
	db *sqlx.DB
}

func NewCodeService(db *sqlx.DB) *Code {
	return &Code{db}
}

func (c *Code) Create(newCode *model.Code) (*model.Code, error) {
	code := &model.Code{}

	if err := dbutil.TXHandler(c.db, func(tx *sqlx.Tx) error {
		// if foudCode, _ := repository.FindCodeByUserIdAndTitle(tx, newCode.UserID, newCode.Title); foudCode != nil {
		//
		// 	return &dbutil.DuplicationCodeError{}
		// }
		//
		result, err := repository.CreateCode(tx, newCode)
		if err != nil {
			return err
		}

		if err := tx.Commit(); err != nil {
			return err
		}

		id, err := result.LastInsertId()
		if err != nil {
			return err
		}

		code, err = repository.FindCode(c.db, id)

		return err
	}); err != nil {
		switch err.(type) {
		case *dbutil.DuplicationCodeError:
			return nil, err
		default:
			return nil, errors.Wrap(err, "failed code insert transaction")
		}
	}

	return code, nil
}

func (c *Code) FindUserCode(user_id int64) (*model.UserCodes, error) {
	userCodes := &model.UserCodes{}

	var err error
	u, err := repository.GetUserById(c.db, user_id)
	if err != nil {
		return nil, err
	}
	userCodes.User = *u

	userCodes.Codes, err = repository.FindUserCodes(c.db, user_id)
	if err != nil {
		return nil, err
	}

	return userCodes, nil
}
