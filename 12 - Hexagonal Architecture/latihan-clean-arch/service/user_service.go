package service

import (
	"context"
	"latian_clean_arch/model/dto"
)

type UserService interface {
	Create(ctx context.Context, request dto.UserRequestBody) (int, interface{}, error)
	Update(ctx context.Context, request dto.UserRequestBody, userId uint) (int, interface{}, error)
	Delete(ctx context.Context, userId uint) (int, error)
	FindById(ctx context.Context, userId uint) (int, interface{}, error)
	FindAll(ctx context.Context) (int, []dto.UserResponse, error)
}
