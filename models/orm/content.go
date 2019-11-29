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
	SubCategoryID string `json:"sub_category_id"`
	SubCategory SubCategory `json:"sub_category" gorm:"association_foreignkey:sub_category_id"`
	VideoSrc string `json:"video_src"`
}

//GetRelatedSubCategory query the specific sub-category of content
func (content *Content) GetRelatedSubCategory() error {
	subCate := &SubCategory {
		SubCategoryID: content.SubCategoryID,
	}
	if errSubCate := subCate.GetSingle(); errSubCate != nil {
		logger.Info("Failed to find content related sub-category", errSubCate)
		return errSubCate
	}
	content.SubCategory = *subCate
	return nil
}

//Create creates new content data.
func (content *Content) Create() error {
	content.ContentID = util.GUID()
	if err := db.Create(&content).Error; err != nil {
		logger.Info(err)
	}
	return nil
}

//GetSingle query the specific content data.
func (content *Content) GetSingle() error {
	if err := db.Where("content_id = ?", content.ContentID).First(&content).Error; err != nil {
		logger.Info("Failed to get specific content.", err)
		return err
	}
	if errSubCate := content.GetRelatedSubCategory(); errSubCate != nil {
		return errSubCate
	}
	return nil
}

//GetSome query some contents with pagging
func (Content) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var contents []Content
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&contents).Error; err != nil {
		logger.Info("Failed to query some contents with pagging", err)
		return nil, err
	}
	for index, content := range contents {
		if errSubCate := content.GetRelatedSubCategory(); errSubCate != nil {
			return nil, errSubCate
		}
		contents[index] = content
	}
	return contents, nil
}

//GetTotal query the count of contents
func (Content) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Content{}).Where(maps).Count(&count).Error; err != nil {
		logger.Info("Failed to query the count of contents", err)
		return -1, err
	}
	return count, nil
}

//Edit updates the specific content data
func (content *Content) Edit() error {
	if err := db.Model(&Content{}).Where("content_id = ?", content.ContentID).Updates(content).Error; err != nil {
		logger.Info("Failed to update the specific content", err)
		return err
	}
	return nil
}

//Delete deletes the specific content data
func (content *Content) Delete() error {
	if err := db.Where("content_id = ?", content.ContentID).Delete(content).Error; err != nil {
		logger.Info("Failed to delete the specific content.", err)
		return err
	}
	return nil
}