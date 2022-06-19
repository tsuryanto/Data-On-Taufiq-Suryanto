package sql

import (
	"context"
	"latian_clean_arch/model/domain"
	. "latian_clean_arch/model/domain"
	"latian_clean_arch/repository"

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
	updatedBook := Book{}
	repository.DB.WithContext(ctx).First(&updatedBook, bookId)
	return updatedBook, err
}

func (repository *BookRepositoryImplSql) Delete(ctx context.Context, bookId uint) error {
	err := repository.DB.WithContext(ctx).Delete(&Book{}, bookId).Error
	return err
}

func (repository *BookRepositoryImplSql) FindById(ctx context.Context, bookId uint) (domain.Book, error) {
	var book = Book{}
	err := repository.DB.WithContext(ctx).First(&book, bookId).Error
	if err == nil {
		return book, nil
	} else {
		return Book{}, err
	}
}

func (repository *BookRepositoryImplSql) FindAll(ctx context.Context) ([]domain.Book, error) {
	var books = []Book{}
	err := repository.DB.WithContext(ctx).Find(&books).Error
	if err != nil {
		return nil, err
	}
	return books, nil
}

func (repository *BookRepositoryImplSql) CountById(ctx context.Context, bookId uint) (int, error) {
	var count int64
	err := repository.DB.WithContext(ctx).Model(&Book{}).Where("id = ?", bookId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return int(count), nil
}
