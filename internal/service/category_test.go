package service

import (
	"go-unit-test/entity"
	"go-unit-test/internal/repository"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var categoryRepository = repository.CategoryRepositryMock{Mock: mock.Mock{}}
var categoryService = CategoryService{Repository: &categoryRepository}

func TestCategoryMock_GetNotFound(t *testing.T) {

	//program
	//mengsetting skenarion jika mengambi data dengan id=1 data di kosong/tidak ada
	categoryRepository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)

}

func TestCategoryMock_GetSucsess(t *testing.T) {
	category := entity.Category{
		Id: "2",
		Name: "Laptop",
	}

	//program mock
	//mengsetting skenario ketia mengambil id=2 dan datanya ada  maka tampilkan data nya
	categoryRepository.Mock.On("FindById", "2").Return(category)
	
	result, err := categoryService.Get("2")
	assert.Nil(t, err) //mensetting error karena pada skenario mengambil data id=2 dan datanya ada
	assert.NotNil(t, result)//menseting result sebagai notnil/tidak kosong karena datanya ada
	assert.Equal(t, category.Id, result.Id) //(testing, data_di_database, hasil_yang_diharapkan)
	assert.Equal(t, category.Name, result.Name)

}