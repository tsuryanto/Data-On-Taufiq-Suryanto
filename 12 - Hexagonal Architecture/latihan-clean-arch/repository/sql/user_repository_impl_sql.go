package sql

import (
	"context"
	"latian_clean_arch/app/helper"
	"latian_clean_arch/model/domain"
	. "latian_clean_arch/model/domain"
	"latian_clean_arch/repository"
	"net/http"

	"gorm.io/gorm"
)

type UserRepositoryImplSql struct {
	DB *gorm.DB
}

func NewUserRepositoryImplSql(db *gorm.DB) repository.UserRepository {
	return &UserRepositoryImplSql{
		DB: db,
	}
}

func (repository *UserRepositoryImplSql) Save(ctx context.Context, user domain.User) (domain.User, error) {
	err := repository.DB.WithContext(ctx).Save(&user).Error
	return user, err
}

func (repository *UserRepositoryImplSql) Update(ctx context.Context, user domain.User, userId uint) (domain.User, error) {
	err := repository.DB.WithContext(ctx).Model(&user).Where("id = ?", userId).Updates(&user).Error
	return user, err
}

func (repository *UserRepositoryImplSql) Delete(ctx context.Context, userId uint) error {
	err := repository.DB.WithContext(ctx).Delete(&User{}, userId).Error
	return err
}

func (repository *UserRepositoryImplSql) FindById(ctx context.Context, userId uint) (int, domain.User, error) {
	var user = User{}
	err := repository.DB.WithContext(ctx).First(&user, userId).Error
	code := helper.GetQueryErrorResponseCode(err)

	if code == http.StatusOK {
		return code, user, nil
	} else {
		return code, User{}, err
	}
}

func (repository *UserRepositoryImplSql) FindAll(ctx context.Context) ([]domain.User, error) {
	var users = []User{}
	query := repository.DB.WithContext(ctx).Find(&users)
	if err := query.Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepositoryImplSql) CountById(ctx context.Context, userId uint) (int, error) {
	var count int64
	query := repository.DB.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Count(&count)
	if err := query.Error; err != nil {
		return -1, err
	}
	return int(count), nil
}
