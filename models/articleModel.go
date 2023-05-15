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

// GetArticles TODO : enable pagination
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

func (ar *Article) DeleteArticle() (string, error) {
	row := initializer.Db.QueryRow("DELETE FROM articles WHERE id=$1 RETURNING id", ar.ID)
	if row.Err() != nil {
		return "0", row.Err()
	}

	var id string
	err := row.Scan(&id)
	if err != nil {
		return "0", err
	}
	return id, nil
}
