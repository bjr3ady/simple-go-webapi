package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//Category is the ORM school_board_category table object.
type Category struct {
	Model
	CategoryID    string `json:"category_id" gorm:"index"`
	Name          string `json:"name"`
	Icon          string `json:"icon"`
	BannerBgColor string `json:"banner_bg_color"`
	Thumb         string `json:"thumb"`
}

//Create creates new category data
func (cate *Category) Create() error {
	cate.CategoryID = util.GUID()
	if err := db.Create(&cate).Error; err != nil {
		logger.Info("Failed to create new category", err)
		return err
	}
	return nil
}

//GetSingle query the specific category data
func (cate *Category) GetSingle() error {
	if err := db.Where("category_id = ?", cate.CategoryID).First(&cate).Error; err != nil {
		logger.Info("Failed to query specific category", err)
		return err
	}
	return nil
}

//GetSome query categories with pagging
func (Category) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var categories []Category
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&categories).Error; err != nil {
		logger.Info("Failed to get some categories", err)
		return nil, err
	}
	return categories, nil
}

//GetTotal query count of categories
func (Category) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Category{}).Where(maps).Count(&count).Error; err != nil {
		logger.Info("Failed to query count of categories", err)
		return -1, err
	}
	return count, nil
}

//HasName determines if specific name of category exists
func (cate *Category) HasName() (bool, error) {
	if err := db.Where("name = ?", cate.Name).First(&cate).Error; err != nil {
		logger.Info("Failed to find specific name of category", err)
		return false, err
	}
	return cate.CategoryID != "", nil
}

//GetByName query the specific name of category
func (cate *Category) GetByName() error {
	if err := db.Where("name = ?", cate.Name).First(&cate).Error; err != nil {
		logger.Info("Failed to find specific name of category", err)
		return err
	}
	return nil
}

//Edit updates category data
func (cate *Category) Edit() error {
	if err := db.Model(&Category{}).Where("category_id = ?", cate.CategoryID).Updates(cate).Error; err != nil {
		logger.Info("Failed to update category", err)
		return err
	}
	return nil
}

//Delete deletes category data
func (cate *Category) Delete() error {
	if err := db.Where("category_id = ?", cate.CategoryID).Delete(cate).Error; err != nil {
		logger.Info("Failed to delete category", err)
		return err
	}
	return nil
}
