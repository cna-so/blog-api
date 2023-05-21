package models

import (
	"backend/initializer"
	"errors"
	"github.com/lib/pq"
	"time"
)

type Article struct {
	ID          int       `json:"id,omitempty"`
	Title       string    `json:"title" binding:"required"`
	Description string    `json:"description" binding:"required"`
	CategoryID  int       `json:"category_id" binding:"required"`
	ImageId     []string  `json:"image_id"`
	Creator     int       `json:"creator" binding:"required"`
	CreateAt    time.Time `json:"create_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}

// GetArticles TODO : enable pagination
func (ar *Article) GetArticles() ([]Article, error) {
	rows, err := initializer.Db.Query("SELECT id , title,description,photos,creator,category_id , created_at FROM articles")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var articles []Article

	for rows.Next() {
		var article Article
		err = rows.Scan(&article.ID, &article.Title, &article.Description, (*pq.StringArray)(&article.ImageId), &article.Creator, &article.CategoryID, &article.CreateAt)
		if err != nil {
			return nil, err
		}
		articles = append(articles, article)
	}
	return articles, nil
}

func (ar *Article) GetArticleWithId() (Article, error) {
	row := initializer.Db.QueryRow("SELECT id, title, description, photos, creator, category_id, created_at, updated_at FROM articles WHERE id=$1", ar.ID)
	if row.Err() != nil {
		return Article{}, row.Err()
	}
	var article Article
	err := row.Scan(&article.ID, &article.Title, &article.Description, (*pq.StringArray)(&article.ImageId), &article.Creator, &article.CategoryID, &article.CreateAt, &article.UpdatedAt)
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

func (ar *Article) InsertArticle() (string, error) {
	row := initializer.Db.QueryRow(`INSERT INTO articles 
    (title, description, creator,photos, category_id , created_at ,updated_at) 
	VALUES ($1,$2,$3,$4,$5,$6,$7) RETURNING id`,
		ar.Title, ar.Description, ar.Creator, ar.ImageId, ar.CategoryID, time.Now(), time.Now(),
	)
	if row.Err() != nil {
		return "", row.Err()
	}
	var id string
	err := row.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}
