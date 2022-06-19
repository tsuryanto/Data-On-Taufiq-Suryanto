package service

import (
	"context"
	"latian_clean_arch/model/dto"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserRequestBody) (interface{}, error)
	Update(ctx context.Context, request dto.UserRequestBody, userId uint) (interface{}, error)
	Delete(ctx context.Context, userId uint) error
	FindById(ctx context.Context, userId uint) (interface{}, error)
	FindAll(ctx context.Context) ([]dto.UserResponse, error)
}
