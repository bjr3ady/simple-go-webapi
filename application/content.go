package application

import (
	"errors"
	models "git.r3ady.com/golang/school-board/models/orm"
	logger "github.com/bjr3ady/go-logger"
)

//GetContentByID query the specific content by id
func GetContentByID(contentID string) (*models.Content, error) {
	if contentID == "" {
		err := errors.New("parameter [contentID] is an empty string while find target content model")
		return nil, err
	}
	content := &models.Content{ContentID: contentID}
	if err := content.GetSingle(); err != nil {
		return nil, err
	}
	return content, nil
}

//QueryContents gets all contents
func QueryContents(startIndex, count int, maps string) ([]models.Content, error) {
	var contents []models.Content
	content := &models.Content{}
	result, err := content.GetSome(startIndex, count, maps)
	if err != nil {
		return nil, err
	}
	contents, ok := result.([]models.Content)
	if !ok {
		err = errors.New("failed to cast query result to content collection")
		logger.Info(err)
		return nil, err
	}
	return contents, nil
}

//TotalContents query the count of contents
func TotalContents(maps interface{}) (int, error) {
	content := &models.Content{}
	return content.GetTotal(maps)
}

//NewContent creates new content model
func NewContent(content, subCateID, videoSrc string) error {
	if content == "" {
		err := errors.New("parameter [content] is an empty string while creating new content model")
		return err
	}
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if errCate := subCate.GetSingle(); errCate != nil {
		err := errors.New("failed to find target sub-category while creating new content")
		logger.Info(err, "=>", errCate)
		return err
	}
	contentModel := &models.Content{
		Content: content,
		SubCategoryID: subCateID,
		VideoSrc: videoSrc,
	}
	return contentModel.Create()
}

//EditContent updates specific content model
func EditContent(content, subCateID, videoSrc string) error {
	if content == "" {
		err := errors.New("parameter [content] is an empty string while updating content model")
		return err
	}
	subCate := &models.SubCategory{SubCategoryID: subCateID}
	if errSubCate := subCate.GetSingle(); errSubCate != nil {
		err := errors.New("failed to find target sub-category while updating content")
		logger.Info(err, "=>", errSubCate)
		return err
	}
	contentModel := &models.Content{
		Content: content,
		SubCategoryID: subCateID,
		VideoSrc: videoSrc,
	}
	return contentModel.Edit()
}

//RemoveContent delete specific content	
func RemoveContent(contentID string) error {
	content := &models.Content{ContentID: contentID}
	if contentID != "" {
		if err := content.GetSingle(); err != nil {
			logger.Info(errors.New("failed to find target content by id to delete"), err)
			return err
		}
		return content.Delete()
	}
	return errors.New("parameter is an empty string as ContentID")
}