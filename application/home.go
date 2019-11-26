package application

import (
	"errors"

	logger "github.com/bjr3ady/go-logger"

	models "git.r3ady.com/golang/school-board/models/orm"
)

//QueryHomes query all home data
func QueryHomes(startIndex, count int, maps interface{}) ([]models.Home, error) {
	var homes []models.Home
	home := &models.Home{}
	result, err := home.GetSome(startIndex, count, "")
	if err != nil {
		return nil, err
	}
	homes, ok := result.([]models.Home)
	if !ok {
		err := errors.New("failed to cast query result to home collection")
		logger.Info(err)
		return nil, err
	}
	return homes, nil
}

//TotalHomes query the count of home items
func TotalHomes(maps interface{}) (int, error) {
	home := &models.Home{}
	return home.GetTotal(maps)
}

//NewHome create new home item
func NewHome(cateID, link string, index, sizeMode, isDirectLink int) error {
	if cateID == "" || link == "" {
		err := errors.New("parameters contain empty string while creating new home item")
		logger.Info(err)
		return err
	}
	cate := &models.Category{CategoryID: cateID}
	if err := cate.GetSingle(); err != nil {
		logger.Info(errors.New("failed to find target category by id while creating new home item"), err)
		return err
	}
	home := &models.Home{
		CategoryID: cateID,
		Link: link,
		Index: index,
		SizeMode: sizeMode,
		IsDirectLink: isDirectLink,
	}
	return home.Create()
}

//GetHomeByID query the specific home item by id
func GetHomeByID(homeID string) (*models.Home, error) {
	if homeID == "" {
		err := errors.New("parameter is empty string while creating new home item")
		return nil ,err
	}
	home := &models.Home{HomeID: homeID}
	if err := home.GetSingle(); err != nil {
		return nil, err
	}
	return home, nil
}

//EditHome updates specific home item
func EditHome(homeID, cateID, link string, index, sizeMode, isDirectLink int) error {
	if homeID == "" || cateID == "" || link == "" {
		err := errors.New("parameters contain empty string while updating home item")
		logger.Info(err)
		return err
	}
	cate := &models.Category{CategoryID: cateID}
	if err := cate.GetSingle(); err != nil {
		logger.Info(errors.New("no related category found while updating home item"), err)
		return err
	}
	home := &models.Home{HomeID: homeID}
	if err := home.GetSingle(); err != nil {
		logger.Info(errors.New("no home item found while updating home item"), err)
		return err
	}
	home.CategoryID = cateID
	home.Link = link
	home.Index = index
	home.SizeMode = sizeMode
	home.IsDirectLink = isDirectLink
	return home.Edit()
}

//RemoveHome deletes specific home item
func RemoveHome(homeID string) error {
	if homeID == "" {
		err := errors.New("parameter is empty string")
		logger.Info(err)
		return err
	}
	home := &models.Home{HomeID: homeID}
	if err := home.GetSingle(); err != nil {
		return err
	}
	return home.Delete()
}