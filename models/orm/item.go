package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//Item is the ORM school_board_item table object.
type Item struct {
	Model
	ItemID string `json:"item_id" gorm:"index"`
	Name string `json:"name"`
	Link string `json:"link"`
	Index int `json:"index"`
	SubCategory SubCategory `json:"sub_category"`
}

//Create creates new item data.
func (item *Item) Create() error {
	item.ItemID = util.GUID()
	if err := item.Create(); err != nil {
		logger.Error("Failed to create new item", err)
		return err
	}
	return nil
}

