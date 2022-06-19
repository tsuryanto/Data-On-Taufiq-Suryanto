package repository

import (
	"context"
	"latian_clean_arch/model/domain"
)

type BookRepository interface {
	Save(ctx context.Context, book domain.Book) (domain.Book, error)
	Update(ctx context.Context, book domain.Book, bookId uint) (domain.Book, error)
	Delete(ctx context.Context, bookId uint) error
	FindById(ctx context.Context, bookId uint) (domain.Book, error)
	FindAll(ctx context.Context) ([]domain.Book, error)
	CountById(ctx context.Context, bookId uint) (int, error)
}
