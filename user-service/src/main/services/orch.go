package services

import (
	"github.com/vtomar01/user-service/src/main/clients"
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/dtos"
)

type UserOrch interface {
	CreateUser(ctx *context.Context, request *dtos.CreateUserRequest) (*dtos.User, error)
	GetUser(ctx *context.Context, id string) (*dtos.User, error)
	UpdateUser(ctx *context.Context, id string, request *dtos.UpdateUserRequest) (*dtos.User, error)
}

type UserOrchImpl struct {
	userSvc           UserService
	phoneStandardizer *clients.PhoneStandardizerClient
}

func NewUserOrch(userSvc UserService, phoneStandardizer *clients.PhoneStandardizerClient) UserOrch {
	return &UserOrchImpl{userSvc: userSvc, phoneStandardizer: phoneStandardizer}
}

func (u *UserOrchImpl) CreateUser(ctx *context.Context,
	request *dtos.CreateUserRequest) (*dtos.User, error) {
	standardizedPhone, err := u.phoneStandardizer.Standardize(ctx,
		&dtos.PhoneStandardizerRequest{Phone: request.Phone})
	if err == nil {
		request.Phone = standardizedPhone.Phone
	}
	return u.userSvc.CreateUser(ctx, request)
}

func (u *UserOrchImpl) GetUser(ctx *context.Context, id string) (*dtos.User, error) {
	return u.userSvc.GetUser(ctx, id)
}

func (u *UserOrchImpl) UpdateUser(ctx *context.Context, id string,
	request *dtos.UpdateUserRequest) (*dtos.User, error) {
	standardizedPhone, err := u.phoneStandardizer.Standardize(ctx,
		&dtos.PhoneStandardizerRequest{Phone: request.Phone})
	if err == nil {
		request.Phone = standardizedPhone.Phone
	}
	return u.userSvc.UpdateUser(ctx, id, request)
}
