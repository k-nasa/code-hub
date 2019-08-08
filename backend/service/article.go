package service

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"

	"github.com/voyagegroup/treasure-app/model"
	"github.com/voyagegroup/treasure-app/repository"
	"github.com/voyagegroup/treasure-app/util"
)

type ArticleService struct {
	db *sqlx.DB
}

func NewArticleService(db *sqlx.DB) *ArticleService {
	return &ArticleService{db}
}

func (a *ArticleService) Update(id int64, newArticle *model.Article) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := util.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.UpdateArticle(tx, id, newArticle)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article update transaction")
	}
	return nil
}

func (a *ArticleService) Destroy(id int64) error {
	_, err := repository.FindArticle(a.db, id)
	if err != nil {
		return errors.Wrap(err, "failed find article")
	}

	if err := util.TXHandler(a.db, func(tx *sqlx.Tx) error {
		_, err := repository.DestroyArticle(tx, id)
		if err != nil {
			return err
		}
		if err := tx.Commit(); err != nil {
			return err
		}
		return err
	}); err != nil {
		return errors.Wrap(err, "failed article delete transaction")
	}
	return nil
}

func (a *ArticleService) Create(newArticle *model.Article) (int64, error) {
	var createdId int64
	if err := util.TXHandler(a.db, func(tx *sqlx.Tx) error {
		result, err := repository.CreateArticle(tx, newArticle)
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
		createdId = id
		return err
	}); err != nil {
		return 0, errors.Wrap(err, "failed article insert transaction")
	}
	return createdId, nil
}
