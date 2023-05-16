package models

import (
	"backend/initializer"
	"time"
)

type CategoryModel struct {
	ID        int       `json:"id" binding:"required"`
	Title     string    `json:"title"  binding:"required"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (m *CategoryModel) GetCategories() ([]CategoryModel, error) {
	rows, err := initializer.Db.Query("SELECT id, title, created_at, updated_at FROM categories ")
	if err != nil {
		return nil, err
	}
	var categories []CategoryModel
	for rows.Next() {
		var category CategoryModel
		err := rows.Scan(&category.ID, &category.Title, &category.CreatedAt, &category.UpdatedAt)
		if err != nil {
			return nil, err
		}
		categories = append(categories, category)
	}
	defer rows.Close()
	return categories, nil
}
