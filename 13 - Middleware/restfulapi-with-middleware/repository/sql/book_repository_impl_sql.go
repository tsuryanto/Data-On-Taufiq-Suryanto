package sql

import (
	"context"
	"latian_clean_arch/app/helper"
	. "latian_clean_arch/model/domain"
	"latian_clean_arch/repository"
	"net/http"

	"gorm.io/gorm"
)

type BookRepositoryImplSql struct {
	DB *gorm.DB
}

func NewBookRepositoryImplSql(db *gorm.DB) repository.BookRepository {
	return &BookRepositoryImplSql{
		DB: db,
	}
}

func (repository *BookRepositoryImplSql) Save(ctx context.Context, book Book) (Book, error) {
	err := repository.DB.WithContext(ctx).Save(&book).Error
	return book, err
}

func (repository *BookRepositoryImplSql) Update(ctx context.Context, book Book, bookId uint) (Book, error) {
	err := repository.DB.WithContext(ctx).Model(&book).Where("id = ?", bookId).Updates(&book).Error
	return book, err
}

func (repository *BookRepositoryImplSql) Delete(ctx context.Context, bookId uint) error {
	err := repository.DB.WithContext(ctx).Delete(&Book{}, bookId).Error
	return err
}

func (repository *BookRepositoryImplSql) FindById(ctx context.Context, bookId uint) (int, Book, error) {
	var book = Book{}
	err := repository.DB.WithContext(ctx).First(&book, bookId).Error
	code := helper.GetQueryErrorResponseCode(err)

	if code == http.StatusOK {
		return code, book, nil
	} else {
		return code, Book{}, err
	}
}

func (repository *BookRepositoryImplSql) FindAll(ctx context.Context) ([]Book, error) {
	var books = []Book{}
	query := repository.DB.WithContext(ctx).Find(&books)
	if err := query.Error; err != nil {
		return nil, err
	}
	return books, nil
}

func (repository *BookRepositoryImplSql) CountById(ctx context.Context, bookId uint) (int, error) {
	var count int64
	query := repository.DB.WithContext(ctx).Model(&Book{}).Where("id = ?", bookId).Count(&count)
	if err := query.Error; err != nil {
		return -1, err
	}
	return int(count), nil
}
