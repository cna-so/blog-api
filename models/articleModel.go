package models

import (
	"backend/initializer"
	"time"
)

type Article struct {
	ID          uint      `json:"id,omitempty"`
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
