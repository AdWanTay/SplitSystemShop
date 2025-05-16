package repositories

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"context"
	"gorm.io/gorm"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *models.Article) error
	GetByID(ctx context.Context, id uint) (*models.Article, error)
	GetAll(ctx context.Context) ([]models.Article, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req dto.NewArticleRequest) error
}

type articleRepository struct {
	db *gorm.DB
}

func NewArticleRepository(db *gorm.DB) ArticleRepository {
	return &articleRepository{db: db}
}

func (r *articleRepository) Create(ctx context.Context, article *models.Article) error {
	return r.db.WithContext(ctx).Create(article).Error
}

func (r *articleRepository) GetByID(ctx context.Context, id uint) (*models.Article, error) {
	var article models.Article
	err := r.db.WithContext(ctx).First(&article, id).Error
	return &article, err
}

func (r *articleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.WithContext(ctx).Find(&articles).Error
	return articles, err
}

func (r *articleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Article{}, id).Error
}

func (r *articleRepository) Update(ctx context.Context, id uint, req dto.NewArticleRequest) error {
	return r.db.WithContext(ctx).Model(&models.Article{}).Where("id = ?", id).Updates(models.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		ImageURL:    req.ImageURL,
	}).Error
}
