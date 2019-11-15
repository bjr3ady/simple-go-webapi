package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//SubCategory is the ORM school_board_sub_category table object.
type SubCategory struct {
	Model
	SubCategoryID string `json:"sub_category_id" gorm:"index"`
	Name string `json:"name"`
	Category Category `json:"category"`
}

//Create creates new sub category data
func (subCate *SubCategory) Create() error {
	subCate.SubCategoryID = util.GUID()
	if err := subCate.Create(); err != nil {
		logger.Error("Failed to create new sub category", err)
		return err
	}
	return nil
}

//GetSingle query the specific sub-cateogry data
func (subCate *SubCategory) GetSingle() error {
	if err := db.Where("sub_category_id = ?", subCate.SubCategoryID).First(&subCate).Error; err != nil {
		logger.Error("Failed to query specific sub-category", err)
		return err
	}
	return nil
}