package service

import (
	"errors"
	"go-unit-test/entity"
	"go-unit-test/internal/repository"
)

type CategoryService struct {
	Repository repository.CategoryRepositry
}


func (service CategoryService) Get(id string) (*entity.Category, error) {
	category := service.Repository.FindById(id)
	if  category == nil {
		return nil, errors.New("category not found")	
	}else {
		return  category, nil
	}
}