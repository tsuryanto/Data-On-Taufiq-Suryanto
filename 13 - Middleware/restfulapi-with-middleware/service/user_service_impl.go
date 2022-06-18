package service

import (
	"context"
	"latian_clean_arch/model/domain"
	"latian_clean_arch/model/dto"
	repository "latian_clean_arch/repository"
	"net/http"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
}

func NewUserServiceImpl(repository repository.UserRepository) UserService {
	return &UserServiceImpl{
		UserRepository: repository,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, request dto.UserRequestBody) (int, interface{}, error) {
	user, err := service.UserRepository.Save(ctx, ToUserDomain(request))
	var statusCode int
	if err == nil {
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return statusCode, toUserResponse(user), err
}

func (service *UserServiceImpl) Update(ctx context.Context, request dto.UserRequestBody, userId uint) (int, interface{}, error) {
	var statusCode int
	count, err := service.UserRepository.CountById(ctx, userId)
	if err == nil {
		if count == 1 {
			user, err := service.UserRepository.Update(ctx, ToUserDomain(request), userId)
			if err == nil {
				statusCode = http.StatusOK
				user.ID = userId
				return statusCode, toUserResponse(user), err
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

func (service *UserServiceImpl) Delete(ctx context.Context, userId uint) (int, error) {
	var statusCode int
	count, err := service.UserRepository.CountById(ctx, userId)
	if err == nil {
		if count == 1 {
			err := service.UserRepository.Delete(ctx, userId)
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

func (service *UserServiceImpl) FindById(ctx context.Context, userId uint) (int, interface{}, error) {
	code, user, err := service.UserRepository.FindById(ctx, userId)
	if code == http.StatusOK {
		return code, toUserResponse(user), err
	} else {
		return code, nil, err
	}
}

func (service *UserServiceImpl) FindAll(ctx context.Context) (int, []dto.UserResponse, error) {
	users, err := service.UserRepository.FindAll(ctx)
	var userResponses []dto.UserResponse
	var statusCode int
	if err == nil {
		for _, user := range users {
			userResponses = append(userResponses, toUserResponse(user))
		}
		statusCode = http.StatusOK
	} else {
		statusCode = http.StatusInternalServerError
	}
	return statusCode, userResponses, err
}

func (service *UserServiceImpl) Auth(ctx context.Context, request dto.UserRequestAuth) (int, interface{}, error) {
	isValid, err := service.UserRepository.MatchingEmailAndPassword(ctx, request.Identifier, request.Password)
	if err != nil {
		return http.StatusInternalServerError, nil, err
	}

	if isValid {
		code, user, err := service.UserRepository.FindByIdentifier(ctx, request.Identifier)
		if code == http.StatusOK {
			return code, toUserResponse(user), err
		} else {
			return code, nil, err
		}
	} else {
		return http.StatusUnauthorized, nil, err
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
	// now := time.Now()
	// if tag == "create" {
	// 	user.CreatedAt = now
	// 	user.CreatedAt = now
	// } else if tag == "update" {
	// 	user.CreatedAt = now
	// }
	return user
}
