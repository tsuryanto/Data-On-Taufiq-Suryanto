package service

import (
	"context"
	. "latian_clean_arch/app/constant"
	"latian_clean_arch/model/domain"
	"latian_clean_arch/model/dto"
	repository "latian_clean_arch/repository"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(repository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: repository,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request dto.BookRequestBody) (interface{}, error) {
	book, err := service.BookRepository.Save(ctx, ToBookDomain(request))
	if err != nil {
		return nil, err
	} else {
		return toBookResponse(book), nil
	}
}

func (service *BookServiceImpl) Update(ctx context.Context, request dto.UpdateBookRequestBody, bookId uint) (interface{}, error) {
	count, err := service.BookRepository.CountById(ctx, bookId)
	if err == nil {
		if count == 1 {
			book, err := service.BookRepository.Update(ctx, UpdateToBookDomain(request), bookId)
			if err == nil {
				book.ID = bookId
				return toBookResponse(book), nil
			}
		} else {
			if count == 0 {
				return nil, ErrNotFound
			}
		}
	}
	return nil, ErrServerError
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId uint) error {
	count, err := service.BookRepository.CountById(ctx, bookId)
	if err == nil {
		if count == 1 {
			err := service.BookRepository.Delete(ctx, bookId)
			if err == nil {
				return nil
			}
		} else {
			if count == 0 {
				return ErrNotFound
			}
		}
	}
	return ErrServerError
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId uint) (interface{}, error) {
	count, err := service.BookRepository.CountById(ctx, bookId)
	if err == nil {
		if count == 1 {
			book, err := service.BookRepository.FindById(ctx, bookId)
			if err == nil {
				return toBookResponse(book), nil
			}
		} else {
			if count == 0 {
				return nil, ErrNotFound
			}
		}
	}
	return nil, ErrServerError
}

func (service *BookServiceImpl) FindAll(ctx context.Context) ([]dto.BookResponse, error) {
	books, err := service.BookRepository.FindAll(ctx)
	var bookResponses []dto.BookResponse
	if err == nil {
		for _, book := range books {
			bookResponses = append(bookResponses, toBookResponse(book))
		}
		return bookResponses, nil
	} else {
		return bookResponses, ErrServerError
	}
}

func toBookResponse(book domain.Book) dto.BookResponse {
	return dto.BookResponse{
		ID:        book.ID,
		Name:      book.Name,
		Publisher: book.Publisher,
		Author:    book.Author,
		CreatedAt: book.CreatedAt,
		UpdatedAt: book.UpdatedAt,
	}
}

func ToBookDomain(req dto.BookRequestBody) domain.Book {
	book := domain.Book{
		Name:      req.Name,
		Publisher: req.Publisher,
		Author:    req.Author,
	}
	return book
}

func UpdateToBookDomain(req dto.UpdateBookRequestBody) domain.Book {
	book := domain.Book{}

	if req.Name != nil {
		book.Name = *req.Name
	}

	if req.Publisher != nil {
		book.Publisher = *req.Publisher
	}

	if req.Author != nil {
		book.Author = *req.Author
	}
	return book
}
