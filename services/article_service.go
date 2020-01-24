package services

import (
	"github.com/darwin1224/go-test/models"
	"github.com/jinzhu/gorm"
)

// ArticleService -- service instance
type ArticleService struct {
	Repo *gorm.DB
}

// GetAllArticle -- Get all data
func (s *ArticleService) GetAllArticle() []models.Article {
	var articles []models.Article
	err := s.Repo.Find(&articles).Error
	if err != nil {
		panic(err)
	}
	return articles
}

// GetArticleByID -- Get data by id
func (s *ArticleService) GetArticleByID(id *int) models.Article {
	var article models.Article
	err := s.Repo.Where("ID = ?", id).First(&article).Error
	if err != nil {
		panic(err)
	}
	return article
}

// InsertArticle -- Insert data
func (s *ArticleService) InsertArticle(article *models.Article) {
	err := s.Repo.Save(article).Error
	if err != nil {
		panic(err)
	}
}

// UpdateArticle -- Update data by id
func (s *ArticleService) UpdateArticle(article *models.Article) {
	err := s.Repo.Save(article).Error
	if err != nil {
		panic(err)
	}
}

// DeleteArticle -- Delete data by id
func (s *ArticleService) DeleteArticle(article *models.Article) {
	err := s.Repo.Delete(article).Error
	if err != nil {
		panic(err)
	}
}
