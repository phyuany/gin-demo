package repository

import (
	"github.com/jinzhu/gorm"
	"jkdev.cn/api/common"
	"jkdev.cn/api/model"
)

type CategoryRepository struct {
	DB *gorm.DB
}

func NewCategoryRepository() CategoryRepository {
	return CategoryRepository{DB: common.GetDB()}
}

func (c CategoryRepository) Create(name string) (*model.Category, error) {
	category := model.Category{
		Name: name,
	}

	if err := c.DB.Create(&category).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) Update(category model.Category, name string) (*model.Category, error) {
	if err := c.DB.Model(&category).Update("name", name).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) SelectById(id int) (*model.Category, error) {
	var category model.Category
	if err := c.DB.First(&category, id).Error; err != nil {
		return nil, err
	}

	return &category, nil
}

func (c CategoryRepository) DeleteById(id int) error {
	if err := c.DB.Delete(model.Category{}, id).Error; err != nil {
		return err
	}

	return nil
}
