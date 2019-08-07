package service

import (
	"database/sql"
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/voyagegroup/treasure-app/domain/model"
)

type ArticleService struct {
	dbx *sqlx.DB
}

func NewArticleService(dbx *sqlx.DB) *ArticleService {
	return &ArticleService{dbx}
}

func (a *ArticleService) UpdateArticle(id int64, newArticle *model.Article) (sql.Result, error) {
	_, err := model.FindArticle(a.dbx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed article repository")
	}
	result, err := model.UpdateArticle(a.dbx, id, newArticle)
	if err != nil {
		return nil, errors.Wrap(err, "failed article repository")
	}
	return result, nil
}
