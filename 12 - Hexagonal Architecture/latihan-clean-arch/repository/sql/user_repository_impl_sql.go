package sql

import (
	"context"
	"latian_clean_arch/model/domain"
	. "latian_clean_arch/model/domain"
	"latian_clean_arch/repository"

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
	updatedUser := User{}
	repository.DB.WithContext(ctx).First(&updatedUser, userId)
	return updatedUser, err
}

func (repository *UserRepositoryImplSql) Delete(ctx context.Context, userId uint) error {
	err := repository.DB.WithContext(ctx).Delete(&User{}, userId).Error
	return err
}

func (repository *UserRepositoryImplSql) FindById(ctx context.Context, userId uint) (domain.User, error) {
	var user = User{}
	err := repository.DB.WithContext(ctx).First(&user, userId).Error
	if err == nil {
		return user, nil
	} else {
		return User{}, err
	}
}

func (repository *UserRepositoryImplSql) FindAll(ctx context.Context) ([]domain.User, error) {
	var users = []User{}
	err := repository.DB.WithContext(ctx).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (repository *UserRepositoryImplSql) CountById(ctx context.Context, userId uint) (int, error) {
	var count int64
	err := repository.DB.WithContext(ctx).Model(&User{}).Where("id = ?", userId).Count(&count).Error
	if err != nil {
		return -1, err
	}
	return int(count), nil
}
