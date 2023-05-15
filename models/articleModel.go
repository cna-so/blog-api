package models

import (
	"backend/initializer"
	"errors"
	"time"
)

type Article struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Creator     int       `json:"creator"`
	CreateAt    time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func (ar *Article) GetArticles() ([]Article, error) {
	rows, err := initializer.Db.Query("SELECT * FROM articles")
	if err != nil {
		return nil, err
	}
	var articles []Article

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Description, &article.Creator, &article.CreateAt, &article.UpdatedAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
		defer rows.Close()
	}
	return articles, nil
}
func (ar *Article) GetArticleWithId() (Article, error) {
	row := initializer.Db.QueryRow("SELECT * FROM articles WHERE id=$1", ar.ID)
	if row.Err() != nil {
		return Article{}, row.Err()
	}
	var article Article
	err := row.Scan(&article.ID, &article.Title, &article.Description, &article.Creator, &article.CreateAt, &article.UpdatedAt)
	if err != nil {
		return Article{}, err
	}
	if article.ID == ar.ID {
		return article, nil
	}
	return Article{}, errors.New("your article doesn't exist ")
}
