package orm

import (
	"testing"
	"errors"
	"git.r3ady.com/golang/school-board/pkg/setting"
)

var testContent NormalModeler
const testContentTxt = "11"

func init() {
	ConnectDb(setting.Cfg)
	testContent = &Content{
		Content: testContentTxt,
	}
}

func getTestContent(t *testing.T) *Content {
	actual, err := testContent.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	contents, ok := actual.([]Content)
	if !ok {
		t.Error(errors.New("No content record found"))
		return nil
	}
	return &contents[0]
}

func TestCreateContent(t *testing.T) {
	subCate := prepareSubCategory(t)
	content, _ := testContent.(*Content)
	content.SubCategoryID = subCate.SubCategoryID
	if err := content.Create(); err != nil {
		t.Error(err)
	}
}

func TestGetSingleContent(t *testing.T) {
	testContent = getTestContent(t)
	if testContent != nil {
		if err := testContent.GetSingle(); err != nil {
			t.Error(err)
		}
	}
}

func TestGetSomeContents(t *testing.T) {
	getTestContent(t)
}

func TestGetTotalContents(t *testing.T) {
	count, err := testContent.GetTotal("")
	if err != nil {
		t.Error(err)
		return
	}
	if count == 0 {
		t.Error(errors.New("No content record found"))
	}
}

func TestEditContent(t *testing.T) {
	const updatedContentTxt = "22"
	testContent = getTestContent(t)
	if testContent != nil {
		content, _ := testContent.(*Content)
		content.Content = updatedContentTxt
		if err := content.Edit(); err != nil {
			t.Error(err)
		}
		content = getTestContent(t)
		if content.Content != updatedContentTxt {
			t.Error(errors.New("Failed to update content's content"))
		}
	}
}

func TestDeleteContent(t *testing.T) {
	content := getTestContent(t)
	if err := content.Delete(); err != nil {
		t.Error(err)
	}
	clearSubCategory(t)
	clearCategory(t)
}