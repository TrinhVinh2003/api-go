package models

type Article struct {
	ID          uint   `json:"id" gorm:"primaryKey"`
	Title       string `json:"title" gorm:"not null;column:title;size:255;charset=utf8mb4;collate:utf8mb4_unicode_ci"`
	Link        string `json:"link" gorm:"not null;column:link;size:255;unique;charset=utf8mb4;collate:utf8mb4_unicode_ci"`
	Image       string `json:"image" gorm:"not null;column:image;size:255"`
	Description string `json:"description" gorm:"null;column:description;size:255;charset=utf8mb4;collate:utf8mb4_unicode_ci"`
}
