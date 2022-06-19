package service

import (
	"context"
	. "latian_clean_arch/app/constant"
	"latian_clean_arch/model/domain"
	"latian_clean_arch/model/dto"
	repository "latian_clean_arch/repository"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.UserRequestBody) (interface{}, error) {
	user, err := service.UserRepository.Save(ctx, ToUserDomain(request))
	if err != nil {
		return nil, err
	} else {
		return toUserResponse(user), nil
	}
}

func (service *UserServiceImpl) Update(ctx context.Context, request dto.UserRequestBody, userId uint) (interface{}, error) {
	count, err := service.UserRepository.CountById(ctx, userId)
	if err == nil {
		if count == 1 {
			user, err := service.UserRepository.Update(ctx, ToUserDomain(request), userId)
			if err == nil {
				user.ID = userId
				return toUserResponse(user), nil
			}
		} else {
			if count == 0 {
				return nil, ErrNotFound
			}
		}
	}
	return nil, ErrServerError
}

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) error {
	count, err := service.UserRepository.CountById(ctx, userId)
	if err == nil {
		if count == 1 {
			err := service.UserRepository.Delete(ctx, userId)
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

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) (interface{}, error) {
	count, err := service.UserRepository.CountById(ctx, userId)
	if err == nil {
		if count == 1 {
			user, err := service.UserRepository.FindById(ctx, userId)
			if err == nil {
				return toUserResponse(user), nil
			}
		} else {
			if count == 0 {
				return nil, ErrNotFound
			}
		}
	}
	return nil, ErrServerError
}

func (service *UserServiceImpl) FindAll(ctx context.Context) ([]dto.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx)
	var userResponses []dto.UserResponse
	if err == nil {
		for _, user := range users {
			userResponses = append(userResponses, toUserResponse(user))
		}
		return userResponses, nil
	} else {
		return userResponses, ErrServerError
	}
}

func (service *UserServiceImpl) Auth(ctx context.Context, request dto.UserRequestAuth) (interface{}, error) {
	isValid, err := service.UserRepository.MatchingEmailAndPassword(ctx, request.Identifier, request.Password)
	if err != nil {
		return nil, ErrServerError
	}

	if isValid {
		user, err := service.UserRepository.FindByIdentifier(ctx, request.Identifier)
		if err == nil {
			return toUserResponse(user), nil
		} else {
			return nil, ErrServerError
		}
	} else {
		return nil, ErrUnauthorized
	}
}

func toUserResponse(user domain.User) dto.UserResponse {
	return dto.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserDomain(req dto.UserRequestBody) domain.User {
	user := domain.User{
		Name:     req.Name,
		Email:    req.Email,
		Password: req.Password,
	}
	return user
}
