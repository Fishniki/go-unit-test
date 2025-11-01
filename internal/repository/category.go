package repository

import "go-unit-test/entity"

type CategoryRepositry interface {
	FindById(id string) *entity.Category
}

