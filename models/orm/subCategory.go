package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//SubCategory is the ORM school_board_sub_category table object.
type SubCategory struct {
	Model
	SubCategoryID string   `json:"sub_category_id" gorm:"index"`
	Name          string   `json:"name"`
	CategoryID    string   `json:"category_id"`
	Category      Category `json:"category" gorm:"association_foreignkey:category_id"`
}

//Create creates new sub category data
func (subCate *SubCategory) Create() error {
	subCate.SubCategoryID = util.GUID()
	if err := db.Create(&subCate).Error; err != nil {
		logger.Error("Failed to create new sub category", err)
		return err
	}
	return nil
}

//GetRelatedCategory query the specific category of sub-category
func (subCate *SubCategory) GetRelatedCategory() error {
	cate := &Category{
		CategoryID: subCate.CategoryID,
	}
	if errCate := cate.GetSingle(); errCate != nil {
		logger.Error("Failed to query sub-category related category", errCate)
		return errCate
	}
	subCate.Category = *cate
	return nil
}

//GetSingle query the specific sub-cateogry data
func (subCate *SubCategory) GetSingle() error {
	if err := db.Where("sub_category_id = ?", subCate.SubCategoryID).First(&subCate).Error; err != nil {
		logger.Error("Failed to query specific sub-category", err)
		return err
	}
	if errCate := subCate.GetRelatedCategory(); errCate != nil {
		return errCate
	}
	return nil
}

//HasName determines the specific name of sub-category already exists.
func (subCate *SubCategory) HasName() (bool, error) {
	if err := db.Where("name = ?", subCate.Name).First(&subCate).Error; err != nil {
		logger.Error("Failed to query the specific name of sub-category", err)
		return false, err
	}
	return subCate.SubCategoryID != "", nil
}

//GetByName query the specific name of sub-category data
func (subCate *SubCategory) GetByName() error {
	if err := db.Where("name = ?", subCate.Name).First(&subCate).Error; err != nil {
		logger.Error("Failed to find the specific name of sub-category", err)
		return err
	}
	return nil
}

//GetSome query some sub-categories
func (SubCategory) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var subCates []SubCategory
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&subCates).Error; err != nil {
		logger.Error("Failed to query some sub-categories", err)
		return nil, err
	}
	for index, subCate := range subCates {
		if err := subCate.GetRelatedCategory(); err != nil {
			return nil, err
		}
		subCates[index] = subCate
	}
	return subCates, nil
}

//GetTotal query the count of sub-categories
func (SubCategory) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&SubCategory{}).Where(maps).Count(&count).Error; err != nil {
		logger.Error("Failed to query the count of sub-categories", err)
		return -1, err
	}
	return count, nil
}

//Edit updates the specific sub-category data
func (subCate *SubCategory) Edit() error {
	if err := db.Model(&SubCategory{}).Where("sub_category_id = ?", subCate.SubCategoryID).Updates(subCate).Error; err != nil {
		logger.Error("Failed to update specific sub-category", err)
		return err
	}
	return nil
}

//Delete deletes the specific sub-category data
func (subCate *SubCategory) Delete() error {
	if err := db.Where("sub_category_id = ?", subCate.SubCategoryID).Delete(subCate).Error; err != nil {
		logger.Error("Failed to delete the specific sub-category data", err)
		return err
	}
	return nil
}