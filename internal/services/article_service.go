package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"SplitSystemShop/internal/utils"
	"context"
	"strings"
)

type ArticleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(ctx context.Context, req dto.NewArticleRequest) (*models.Article, error) {
	imagePath := ""

	// Обработка превью-картинки
	if strings.HasPrefix(req.ImageBase64, "data:image") {
		path, err := utils.SaveBase64Image(req.ImageBase64)
		if err != nil {
			return nil, err
		}
		imagePath = path
	} else {
		imagePath = "/web/static/uploads/article_images/placeholder.jpg"
	}

	// Обработка base64-картинок в HTML
	content, err := utils.ReplaceBase64ImagesInHTML(req.Content)
	if err != nil {
		return nil, err
	}

	article := &models.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     content,
		ImageURL:    imagePath,
	}
	if err := s.repo.Create(ctx, article); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) GetByID(ctx context.Context, id uint) (*models.Article, error) {
	return s.repo.GetByID(ctx, id)
}

// Для "похожие статьи"
func (s *ArticleService) GetRandomExcept(ctx context.Context, excludeID uint, limit int) ([]models.Article, error) {
	return s.repo.GetRandomExcept(ctx, excludeID, limit)
}

func (s *ArticleService) GetAll(ctx context.Context) ([]models.Article, error) {
	return s.repo.GetAll(ctx)
}

func (s *ArticleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ArticleService) Update(ctx context.Context, id uint, req dto.NewArticleRequest) (*models.Article, error) {
	return s.repo.Update(ctx, id, req)
}
