package service

import (
	"context"
	"latian_clean_arch/model/domain"
	"latian_clean_arch/model/dto"
	repository "latian_clean_arch/repository"
	"net/http"
)

type BookServiceImpl struct {
	BookRepository repository.BookRepository
}

func NewBookServiceImpl(repository repository.BookRepository) BookService {
	return &BookServiceImpl{
		BookRepository: repository,
	}
}

func (service *BookServiceImpl) Create(ctx context.Context, request dto.BookRequestBody) (int, interface{}, error) {
	book, err := service.BookRepository.Save(ctx, ToBookDomain(request))
	var statusCode int
	if err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return statusCode, toBookResponse(book), err
}

func (service *BookServiceImpl) Update(ctx context.Context, request dto.BookRequestBody, bookId uint) (int, interface{}, error) {
	var statusCode int
	count, err := service.BookRepository.CountById(ctx, bookId)
	if err == nil {
		if count == 1 {
			book, err := service.BookRepository.Update(ctx, ToBookDomain(request), bookId)
			if err == nil {
				statusCode = http.StatusOK
				book.ID = bookId
				return statusCode, toBookResponse(book), err
			}
		} else {
			if count == 0 {
				statusCode = http.StatusNotFound
				return statusCode, nil, err
			}
		}
	}
	statusCode = http.StatusInternalServerError
	return statusCode, nil, err
}

func (service *BookServiceImpl) Delete(ctx context.Context, bookId uint) (int, error) {
	var statusCode int
	count, err := service.BookRepository.CountById(ctx, bookId)
	if err == nil {
		if count == 1 {
			err := service.BookRepository.Delete(ctx, bookId)
			if err == nil {
				statusCode = http.StatusOK
				return statusCode, nil
			}
		} else {
			if count == 0 {
				statusCode = http.StatusNotFound
				return statusCode, err
			}
		}
	}
	statusCode = http.StatusInternalServerError
	return statusCode, err
}

func (service *BookServiceImpl) FindById(ctx context.Context, bookId uint) (int, interface{}, error) {
	code, book, err := service.BookRepository.FindById(ctx, bookId)
	if code == http.StatusOK {
		return code, toBookResponse(book), err
	} else {
		return code, nil, err
	}
}

func (service *BookServiceImpl) FindAll(ctx context.Context) (int, []dto.BookResponse, error) {
	books, err := service.BookRepository.FindAll(ctx)
	var bookResponses []dto.BookResponse
	var statusCode int
	if err == nil {
		for _, book := range books {
			bookResponses = append(bookResponses, toBookResponse(book))
		}
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return statusCode, bookResponses, err
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
