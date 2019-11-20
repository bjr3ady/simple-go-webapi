package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//Item is the ORM school_board_item table object.
type Item struct {
	Model
	ItemID      string      `json:"item_id" gorm:"index"`
	Name        string      `json:"name"`
	Link        string      `json:"link"`
	Index       int         `json:"index"`
	SubCateID   string      `json:"sub_cate_id"`
	SubCategory SubCategory `json:"sub_category" gorm:"association_foreignkey:sub_cate_id"`
}

//Create creates new item data.
func (item *Item) Create() error {
	item.ItemID = util.GUID()
	if err := db.Create(&item).Error; err != nil {
		logger.Error("Failed to create new item", err)
		return err
	}
	return nil
}

//GetRelatedSubCategory query the specific sub-category of item
func (item *Item) GetRelatedSubCategory() error {
	subCate := &SubCategory{
		SubCategoryID: item.SubCateID,
	}
	if err := subCate.GetSingle(); err != nil {
		logger.Error("Failed to query item related sub-category", err)
		return err
	}
	item.SubCategory = *subCate
	return nil
}

//GetSingle query the specific item data
func (item *Item) GetSingle() error {
	if err := db.Where("item_id = ?", item.ItemID).First(&item).Error; err != nil {
		logger.Error("Failed to query the specific item", err)
		return err
	}
	if errSubCate := item.GetRelatedSubCategory(); errSubCate != nil {
		return errSubCate
	}
	return nil
}

//GetByName query the name of specific item.
func (item *Item) GetByName() error {
	if err := db.Where("name = ?", item.Name).First(&item).Error; err != nil {
		logger.Error("Failed to query the name of specific item", err)
		return err
	}
	return nil
}

//HasName determines the specific name of item already exists.
func (item *Item) HasName() (bool, error) {
	if err := db.Select("item_id").Where("name = ?", item.Name).First(&item).Error; err != nil {
		logger.Error("Failed to find the specific name of item", err)
		return false, err
	}
	return item.ItemID != "", nil
}

//GetSome query items by pagging
func (Item) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var items []Item
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&items).Error; err != nil {
		logger.Error("Failed to get some items.", err)
		return nil, err
	}
	for index, item := range items {
		if errSubCate := item.GetRelatedSubCategory(); errSubCate != nil {
			return nil, errSubCate
		}
		items[index] = item
	}
	return items, nil
}

//GetTotal query the count of items
func (Item) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Item{}).Where(maps).Count(&count).Error; err != nil {
		logger.Error("Failed to query the count of items.", err)
		return -1, err
	}
	return count, nil
}

//Edit updates the specific item data
func (item *Item) Edit() error {
	if err := db.Model(&Item{}).Where("item_id = ?", item.ItemID).Updates(item).Error; err != nil {
		logger.Error("Failed to update specific item.", err)
		return err
	}
	return nil
}

//Delete deletes the specific item data
func (item *Item) Delete() error {
	if err := db.Where("item_id = ?", item.ItemID).Delete(item).Error; err != nil {
		logger.Error("Failed to delete specific item.", err)
		return err
	}
	return nil
}
