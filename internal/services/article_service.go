package services

import (
	"SplitSystemShop/internal/dto"
	"SplitSystemShop/internal/models"
	"SplitSystemShop/internal/repositories"
	"context"
)

type ArticleService struct {
	repo repositories.ArticleRepository
}

func NewArticleService(repo repositories.ArticleRepository) *ArticleService {
	return &ArticleService{repo: repo}
}

func (s *ArticleService) Create(ctx context.Context, req dto.NewArticleRequest) (*models.Article, error) {
	article := &models.Article{
		Title:       req.Title,
		Description: req.Description,
		Content:     req.Content,
		ImageURL:    req.ImageURL,
	}
	if err := s.repo.Create(ctx, article); err != nil {
		return nil, err
	}
	return article, nil
}

func (s *ArticleService) GetByID(ctx context.Context, id uint) (*models.Article, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *ArticleService) GetAll(ctx context.Context) ([]models.Article, error) {
	return s.repo.GetAll(ctx)
}

func (s *ArticleService) Delete(ctx context.Context, id uint) error {
	return s.repo.Delete(ctx, id)
}

func (s *ArticleService) Update(ctx context.Context, id uint, req dto.NewArticleRequest) error {
	return s.repo.Update(ctx, id, req)
}
