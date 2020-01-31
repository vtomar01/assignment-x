package repo

import (
	"errors"
	"github.com/jinzhu/gorm"
	"github.com/vtomar01/user-service/src/main/context"
	"github.com/vtomar01/user-service/src/main/models"
)

type UserRepo interface {
	Insert(ctx *context.Context, user *models.User) error
	Update(ctx *context.Context, user *models.User) error
	Get(ctx *context.Context, id string) (*models.User, error)
}

type UserRepoImpl struct {
	dbCon *gorm.DB
}

func NewUserRepo(dbCon *gorm.DB) UserRepo {
	return &UserRepoImpl{dbCon: dbCon}
}

func (u *UserRepoImpl) Insert(ctx *context.Context, user *models.User) error {
	db := u.dbCon.Create(user)
	if db.RowsAffected != 1 {
		return errors.New("error in insertion")
	}
	return nil
}

func (u *UserRepoImpl) Update(ctx *context.Context, user *models.User) error {
	db := u.dbCon.Table("users").Where("id = ?", user.Id).Update(user)
	if db.RowsAffected != 1 {
		return errors.New("error in update")
	}
	return nil
}

func (u *UserRepoImpl) Get(ctx *context.Context, id string) (*models.User, error) {
	var users []*models.User
	db := u.dbCon.Where(&models.User{Id: id}).First(&users)
	if db.Error != nil {
		return nil, db.Error
	}
	if len(users) == 0 {
		return nil, errors.New("no user found")
	}
	return users[0], nil
}
