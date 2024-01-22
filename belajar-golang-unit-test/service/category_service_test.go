package service

import (
	"belajar-golang-unit-test/entity"
	"belajar-golang-unit-test/repository"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

var categoryRepsository = &repository.CategoryRepositoryMock{
	Mock: mock.Mock{},
}

var categoryService = CategoryService{
	Repository: categoryRepsository,
}

func TestCategoryService_GetNotFound(t *testing.T) {
	categoryRepsository.Mock.On("FindById", "1").Return(nil)

	category, err := categoryService.Get("1")
	assert.Nil(t, category)
	assert.NotNil(t, err)
}

func TestCategoryService_GetSuccess(t *testing.T) {
	category := entity.Category{
		Id:   "2",
		Name: "Laptop",
	}

	categoryRepsository.Mock.On("FindById", "2").Return(category)

	result, err := categoryService.Get("2")
	assert.Nil(t, err)
	assert.NotNil(t, result)
}
