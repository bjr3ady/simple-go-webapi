package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//Home is the ORM school_board_home table object.
type Home struct {
	Model
	HomeID       string   `json:"home_id" gorm:"index"`
	CategoryID   string   `json:"category_id"`
	Category     Category `json:"category" gorm:"association_foreignkey:category_id"`
	IsDirectLink int      `json:"is_direct_link"`
	Link         string   `json:"link"`
	SizeMode     int      `json:"size_mode"`
	Index        int      `json:"index"`
}

//GetRelatedCategory query the specific category of home
func (home *Home) GetRelatedCategory() error {
	cate := &Category{
		CategoryID: home.CategoryID,
	}
	if err := cate.GetSingle(); err != nil {
		logger.Error("Failed to qurry home related category", err)
		return err
	}
	home.Category = *cate
	return nil
}

//Create creates new home data.
func (home *Home) Create() error {
	home.HomeID = util.GUID()
	if err := db.Create(&home).Error; err != nil {
		logger.Error("Failed to create new home", err)
		return err
	}
	return nil
}

//GetSingle query the specific home data
func (home *Home) GetSingle() error {
	if err := db.Where("home_id = ?", home.HomeID).First(&home).Error; err != nil {
		logger.Error("Failed to get specific home", err)
		return err
	}
	if errCate := home.GetRelatedCategory(); errCate != nil {
		return errCate
	}
	return nil
}

//GetSome query home by pagging
func (Home) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var homes []Home
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&homes).Error; err != nil {
		logger.Error("Failed to query homes by pagging", err)
		return nil, err
	}
	for index, home := range homes {
		if err := home.GetRelatedCategory(); err != nil {
			logger.Error(err)
			return nil, err
		}
		homes[index] = home
	}
	return homes, nil
}

//GetTotal query the count of homes
func (Home) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Home{}).Where(maps).Count(&count).Error; err != nil {
		logger.Error("Failed to query the count of homes", err)
		return -1, err
	}
	return count, nil
}

//Edit updates home data
func (home *Home) Edit() error {
	if err := db.Model(&Home{}).Where("home_id = ?", home.HomeID).Updates(home).Error; err != nil {
		logger.Error("Failed to update home data", err)
		return err
	}
	return nil
}

//Delete deletes specific home data
func (home *Home) Delete() error {
	if err := db.Where("home_id = ?", home.HomeID).Delete(home).Error; err != nil {
		logger.Error("Failed to delete specific home data", err)
		return err
	}
	return nil
}
