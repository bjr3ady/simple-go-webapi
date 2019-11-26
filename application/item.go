package application

import (
	"errors"

	logger "github.com/bjr3ady/go-logger"

	models "git.r3ady.com/golang/school-board/models/orm"
)

//QueryItems query all items
func QueryItems(startIndex, count int, maps interface{}) ([]models.Item, error) {
	var items []models.Item
	item := &models.Item{}
	result, err := item.GetSome(startIndex, count, maps)
	if err != nil {
		return nil, err
	}
	items, ok := result.([]models.Item)
	if !ok {
		err := errors.New("failed to cast query result to item collection")
		logger.Info(err)
		return nil, err
	}
	return items, nil
}

//TotalItems query the count of items
func TotalItems(maps interface{}) (int, error) {
	item := &models.Item{}
	return item.GetTotal(maps)
}

//NewItem creates new item model
func NewItem(name, link, subCateID string, index int) error {
	if name == "" || link == "" || subCateID == "" {
		err := errors.New("parameters contain empty value while creating new item")
		logger.Info(err)
		return err
	}
	item := &models.Item{Name: name, Link: link, SubCateID: subCateID, Index: index}
	return item.Create()
}

//GetItemByID query the specific item by id
func GetItemByID(itemID string) (*models.Item, error) {
	if itemID == "" {
		err := errors.New("parameter contains empty value while delete item")
		logger.Info(err)
		return nil, err
	}
	item := &models.Item{ItemID: itemID}
	if err := item.GetSingle(); err != nil {
		return nil, err
	}
	return item, nil
}

//EditItem updates specific item data
func EditItem(itemID, name, link, subCateID string, index int) error {
	if itemID == "" || name == "" || link == "" || subCateID == "" {
		err := errors.New("parameters contain empty value while updating item")
		logger.Info(err)
		return err
	}
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if errSubCate := subCate.GetSingle(); errSubCate != nil {
		logger.Info("failed to find target sub-category by id while updating item", errSubCate)
		return errSubCate
	}
	item := &models.Item{ItemID: itemID}
	if err := item.GetSingle(); err != nil {
		logger.Info("failed to find target item by id while updating item", err)
		return err
	}
	item.Name = name
	item.Link = link
	item.SubCateID = subCateID
	item.Index = index
	return item.Edit()
}

//RemoveItem delete specific item data
func RemoveItem(itemID string) error {
	if itemID == "" {
		err := errors.New("parameter contains empty value while delete item")
		logger.Info(err)
		return err
	}
	item := &models.Item{ItemID: itemID}
	return item.Delete()
}