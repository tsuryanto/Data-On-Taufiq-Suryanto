package service

import (
	"context"
	"latian_clean_arch/model/dto"
)

type BookService interface {
	Create(ctx context.Context, request dto.BookRequestBody) (int, interface{}, error)
	Update(ctx context.Context, request dto.BookRequestBody, bookId uint) (int, interface{}, error)
	Delete(ctx context.Context, bookId uint) (int, error)
	FindById(ctx context.Context, bookId uint) (int, interface{}, error)
	FindAll(ctx context.Context) (int, []dto.BookResponse, error)
}
