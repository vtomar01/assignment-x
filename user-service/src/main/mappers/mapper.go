package mappers

import (
	"github.com/vtomar01/user-service/src/main/dtos"
	"github.com/vtomar01/user-service/src/main/models"
)

func MapUserModelToDto(m *models.User) *dtos.User {
	return &dtos.User{
		Id:    m.Id,
		Name:  m.Name,
		Age:   m.Age,
		Phone: m.Phone,
	}
}
