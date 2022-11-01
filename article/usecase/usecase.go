package usecase

import (
	"context"
	"study/cp-go-clean-arch/domain"
	"time"
)

type articleUsecase struct {
	articleRepo    domain.ArticleRepository
	authorRepo     domain.AuthorRepository
	contextTimeout time.Duration
}

func NewArticleUsecase(a domain.ArticleRepository, ar domain.AuthorRepository, timeout time.Duration) domain.ArticleUsecase {
	return &articleUsecase{
		articleRepo:    a,
		authorRepo:     ar,
		contextTimeout: timeout,
	}
}

func (a *articleUsecase) Fetch(ctx context.Context, cursor string, num int64) ([]domain.Article, string, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleUsecase) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleUsecase) Update(ctx context.Context, ar *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleUsecase) GetByTitle(ctx context.Context, title string) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a *articleUsecase) Store(ctx context.Context, article *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (a *articleUsecase) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
