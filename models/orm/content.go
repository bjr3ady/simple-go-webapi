package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//Content is the school_board_content table object.
type Content struct {
	Model
	ContentID string `json:"content_id" gorm:"index"`
	Content string `json:"content"`
	SubCategory SubCategory `json:"sub_category"`
	VideoSrc string `json:"video_src"`
}

//GetRelatedSubCategory query the specific sub-category of content
func (content *Content) GetRelatedSubCategory() error {
	subCate := &SubCategory {
		SubCategoryID: content.SubCategory.SubCategoryID,
	}
	if errSubCate := subCate.GetSingle(); errSubCate != nil {
		logger.Error("Failed to find content related sub-category", errSubCate)
		return errSubCate
	}
	content.SubCategory = *subCate
	return nil
}

//Create creates new content data.
func (content *Content) Create() error {
	if errCate := content.GetRelatedSubCategory(); errCate != nil {
		return errCate
	}
	content.ContentID = util.GUID()
	if err := content.Create(); err != nil {
		logger.Error(err)
	}
	return nil
}

//GetSingle query the specific content data.
func (content *Content) GetSingle() error {
	if errCate := content.GetRelatedSubCategory(); errCate != nil {
		return errCate
	}
	if err := db.Where("content_id = ?", content.ContentID).First(&content).Error; err != nil {
		logger.Error("Failed to get specific content.", err)
		return err
	}
	return nil
}