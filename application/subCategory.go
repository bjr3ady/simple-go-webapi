package application

import (
	"errors"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/models/constant"
	logger "github.com/bjr3ady/go-logger"
)

//GetDefaultSubCategory query the default sub-category model
func GetDefaultSubCategory() (*models.SubCategory) {
	subCate := &models.SubCategory{}
	subCate.Name = constant.DEFAULT_SUB_CATEGORY
	if err := subCate.GetByName(); err != nil || subCate.CategoryID == "" {
		subCate.Create()
		subCate.GetByName()
	}
	return subCate
}

//GetSubCategoryByID query the specific sub-category model by id
func GetSubCategoryByID(subCateID string) (*models.SubCategory, error) {
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if err := subCate.GetSingle(); err != nil {
		return nil, err
	}
	return subCate, nil
}

//GetSubCategoryByName query the specific sub-category model by name
func GetSubCategoryByName(subCateName string) (*models.SubCategory, error) {
	subCate := &models.SubCategory{Name: subCateName}
	if err := subCate.GetByName(); err != nil {
		return nil, err
	}
	return subCate, nil
}

//QuerySubCategories gets all sub-category models
func QuerySubCategories(startIndex, count int, maps interface{}) ([]models.SubCategory, error) {
	var subCates []models.SubCategory
	subCate := &models.SubCategory{}
	result, err := subCate.GetSome(startIndex, count, maps)
	if err != nil {
		return nil, err
	}
	subCates, ok := result.([]models.SubCategory)
	if !ok {
		err = errors.New("failed to cast query result to sub-category collection")
		logger.Info(err)
		return nil, err
	}
	return subCates, nil
}

//TotalSubCategories query the count of sub-categories
func TotalSubCategories(maps interface{}) (int, error) {
	subCate := &models.SubCategory{}
	return subCate.GetTotal(maps)
}

//SubCategoryHasName determines the specific name of sub-category already exists.
func SubCategoryHasName(name string) (bool, error) {
	subCata := &models.SubCategory{Name: name}
	return subCata.HasName()
}

//NewSubCategory creates new sub-category model
func NewSubCategory(name, cateID string) error {
	hasName, _ := SubCategoryHasName(name)
	if hasName {
		err := errors.New("name of sub-category already exists")
		logger.Error(err)
		return err
	}
	subCate := &models.SubCategory{Name: name, CategoryID: cateID}
	return subCate.Create()
}

//EditSubCategory updates sub-category model
func EditSubCategory(subCateID, name, cateID string) error {
	hasName, _ := SubCategoryHasName(name)
	if hasName {
		err := errors.New("name of sub-category already exists")
		logger.Error(err)
		return err
	}
	cate := &models.Category{CategoryID: cateID}
	if errCate := cate.GetSingle(); errCate != nil {
		err := errors.New("no related category found while update sub-category")
		logger.Info(err)
		return err
	}
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if err := subCate.GetSingle(); err != nil {
		return err
	}
	subCate.Name = name
	subCate.CategoryID = cateID
	return subCate.Edit()
}

//RemoveSubCategory delete specific sub-category model
func RemoveSubCategory(subCateID string) error {
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if err := subCate.GetSingle(); err != nil {
		logger.Fatal(err)
		return err 
	}
	return subCate.Delete()
}