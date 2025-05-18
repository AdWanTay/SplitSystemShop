package repositories

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/utils"
	"context"
	"gorm.io/gorm"
	"strings"
)

type ArticleRepository interface {
	Create(ctx context.Context, article *models.Article) error
	GetByID(ctx context.Context, id uint) (*models.Article, error)
	GetAll(ctx context.Context) ([]models.Article, error)
	Delete(ctx context.Context, id uint) error
	Update(ctx context.Context, id uint, req dto.NewArticleRequest) (*models.Article, error)
	GetRandomExcept(ctx context.Context, id uint, limit int) ([]models.Article, error)
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

func (r *articleRepository) GetRandomExcept(ctx context.Context, excludeID uint, limit int) ([]models.Article, error) {
	var articles []models.Article
	if err := r.db.WithContext(ctx).
		Where("id != ?", excludeID).
		Order("RANDOM()").
		Limit(limit).
		Find(&articles).Error; err != nil {
		return nil, err
	}
	return articles, nil
}

func (r *articleRepository) GetAll(ctx context.Context) ([]models.Article, error) {
	var articles []models.Article
	err := r.db.WithContext(ctx).Find(&articles).Error
	return articles, err
}

func (r *articleRepository) Delete(ctx context.Context, id uint) error {
	return r.db.WithContext(ctx).Delete(&models.Article{}, id).Error
}

func (r *articleRepository) Update(ctx context.Context, id uint, req dto.NewArticleRequest) (*models.Article, error) {
	imagePath := ""

	// Обработка base64 превью, если есть
	if strings.HasPrefix(req.ImageBase64, "data:image") {
		path, err := utils.SaveBase64Image(req.ImageBase64)
		if err != nil {
			return nil, err
		}
		imagePath = path
	} else {
		// можно оставить старую картинку без изменений, если base64 нет
		// для этого нужно сначала получить текущую статью
		var existing models.Article
		if err := r.db.WithContext(ctx).First(&existing, id).Error; err != nil {
			return nil, err
		}
		imagePath = existing.ImageURL
	}

	// Обработка base64 в HTML (если нужно)
	content, err := utils.ReplaceBase64ImagesInHTML(req.Content)
	if err != nil {
		return nil, err
	}

	// Обновление записи
	newArticle := models.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     content,
		ImageURL:    imagePath,
	}
	err = r.db.WithContext(ctx).Model(&models.Article{}).Where("id = ?", id).Updates(newArticle).Error
	if err != nil {
		return nil, err
	}
	return &newArticle, nil
}
