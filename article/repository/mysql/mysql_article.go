package mysql

import (
	"context"
	"database/sql"
	"study/cp-go-clean-arch/domain"
)

type mysqlArticleRepository struct {
	Conn *sql.DB
}

func NewMysqlArticleRepository(conn *sql.DB) domain.ArticleRepository {
	return &mysqlArticleRepository{conn}
}

func (m *mysqlArticleRepository) Fetch(ctx context.Context, cursor string, num int64) (res []domain.Article, nextCursor string, err error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlArticleRepository) GetByID(ctx context.Context, id int64) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlArticleRepository) Update(ctx context.Context, ar *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlArticleRepository) GetByTitle(ctx context.Context, title string) (domain.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlArticleRepository) Store(ctx context.Context, article *domain.Article) error {
	//TODO implement me
	panic("implement me")
}

func (m *mysqlArticleRepository) Delete(ctx context.Context, id int64) error {
	//TODO implement me
	panic("implement me")
}
