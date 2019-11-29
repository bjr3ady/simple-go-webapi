package application

import (
	"errors"

	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/models/constant"
	logger "github.com/bjr3ady/go-logger"
)

//GetDefaultCategory query the default category model
func GetDefaultCategory() (*models.Category) {
	cate := &models.Category{}
	cate.Name = constant.DEFAULT_CATEGORY
	if err := cate.GetByName(); err != nil || cate.CategoryID == "" {
		cate.Create()
		cate.GetByName()
	}
	return cate
}

//GetCategoryByID gets the ID specific category model.
func GetCategoryByID(cateID string) (*models.Category, error) {
	cate := &models.Category{CategoryID: cateID}
	if err := cate.GetSingle(); err != nil {
		return &models.Category{}, err
	}
	return cate, nil
}

//GetCategoryByName gets the name specific category model
func GetCategoryByName(cateName string) (*models.Category, error) {
	cate := &models.Category{Name: cateName}
	if err := cate.GetByName(); err != nil {
		return &models.Category{}, err
	}
	return cate, nil
}

//QueryCategories gets all categorie models.
func QueryCategories(startIndex, count int, maps interface{}) ([]models.Category, error) {
	var cates []models.Category
	cate := &models.Category{}
	result, err := cate.GetSome(startIndex, count, maps)
	if err != nil {
		return nil, err
	}
	cates, ok := result.([]models.Category)
	if !ok {
		err = errors.New("failed to cast query result to category models")
		return nil, err
	}
	return cates, nil
}

//TotalCategories get the count of all categories
func TotalCategories(maps interface{}) (int, error) {
	cate := &models.Category{}
	return cate.GetTotal(maps)
}

//CategoryHasName determines the specific category name already exists.
func CategoryHasName(name string) (bool, error) {
	cate := &models.Category{Name: name}
	return cate.HasName()
}

//NewCategory creates new category model
func NewCategory(name, icon, bannerBgColor, thumb string) error {
	hasName, _ := CategoryHasName(name)
	if hasName {
		err := errors.New("name of category already exists")
		logger.Info(err)
		return err
	}
	cate := &models.Category{
		Name: name,
		Icon: icon,
		BannerBgColor: bannerBgColor,
		Thumb: thumb,
	}
	return cate.Create()
}

//EditCategory updates specific category model
func EditCategory(cateID, name, icon, bannerBgColor, thumb string) error {
	cate := &models.Category{CategoryID: cateID}
	if err := cate.GetSingle(); err != nil {
		return err
	}
	cate.Name = name
	cate.Icon = icon
	cate.BannerBgColor = bannerBgColor
	cate.Thumb = thumb
	return cate.Edit()
}

//RemoveCategory delete category model
func RemoveCategory(cateID string) error {
	cate, err := GetCategoryByID(cateID)
	if err != nil {
		return err
	}
	if err = cate.Delete(); err != nil {
		return err
	}
	return nil
}