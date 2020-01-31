package services

import (
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/dtos"
	"github.com/vtomar01/user-service/src/main/mappers"
	"github.com/vtomar01/user-service/src/main/models"
	"github.com/vtomar01/user-service/src/main/repo"
	"github.com/vtomar01/user-service/src/main/utils/uuid"
)

type UserService interface {
	CreateUser(ctx *context.Context, request *dtos.CreateUserRequest) (*dtos.User, error)
	GetUser(ctx *context.Context, id string) (*dtos.User, error)
	UpdateUser(ctx *context.Context, id string, request *dtos.UpdateUserRequest) (*dtos.User, error)
}

type UserServiceImpl struct {
	userRepo repo.UserRepo
}

func NewUserService(userRepo repo.UserRepo) UserService {
	return &UserServiceImpl{userRepo: userRepo}
}

func (u *UserServiceImpl) CreateUser(ctx *context.Context, request *dtos.CreateUserRequest) (*dtos.User, error) {
	user := &models.User{
		Id:    uuid.V4(),
		Name:  request.Name,
		Age:   request.Age,
		Phone: request.Phone,
	}
	err := u.userRepo.Insert(ctx, user)
	if err != nil {
		return nil, err
	}
	return mappers.MapUserModelToDto(user), nil
}

func (u *UserServiceImpl) GetUser(ctx *context.Context, id string) (*dtos.User, error) {
	user, err := u.userRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return mappers.MapUserModelToDto(user), nil
}

func (u UserServiceImpl) UpdateUser(ctx *context.Context, id string,
	request *dtos.UpdateUserRequest) (*dtos.User, error) {
	user := &models.User{
		Id:    id,
		Name:  request.Name,
		Age:   request.Age,
		Phone: request.Phone,
	}
	err := u.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}
	return mappers.MapUserModelToDto(user), nil
}
