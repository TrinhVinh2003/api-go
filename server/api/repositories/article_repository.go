package repository

import (
	"project/vnexpress/api/models"

	"gorm.io/gorm"
)

type ArticleRepository interface {
	FindAll() ([]models.Article, error)
	FindByID(id uint) (models.Article, error)
	Create(article *models.Article) error
	Update(article *models.Article) error
	Delete(id uint) error
}

type articleRepository struct {
	DB *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{DB: db}
}

func (r *articleRepository) FindAll() ([]models.Article, error) {
	var articles []models.Article
	if err := r.DB.Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) FindByID(id uint) (models.Article, error) {
	var article models.Article
	if err := r.DB.First(&article, id).Error; err != nil {
		return article, err
	}
	return article, nil
}

func (r *articleRepository) Create(article *models.Article) error {
	return r.DB.Create(article).Error
}

func (r *articleRepository) Update(article *models.Article) error {
	return r.DB.Save(article).Error
}

func (r *articleRepository) Delete(id uint) error {
	return r.DB.Delete(&models.Article{}, id).Error
}
