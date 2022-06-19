package service

import (
	"context"
	"latian_clean_arch/model/dto"
)

type BookService interface {
	Create(ctx context.Context, request dto.BookRequestBody) (interface{}, error)
	Update(ctx context.Context, request dto.BookRequestBody, bookId uint) (interface{}, error)
	Delete(ctx context.Context, bookId uint) error
	FindById(ctx context.Context, bookId uint) (interface{}, error)
	FindAll(ctx context.Context) ([]dto.BookResponse, error)
}
