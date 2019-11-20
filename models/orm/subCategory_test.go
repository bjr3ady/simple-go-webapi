package orm

import (
	"git.r3ady.com/golang/school-board/pkg/setting"
	"testing"
	"errors"
)

var testSubCate NameSpecifier
const testSubCateName = "test"

func init() {
	ConnectDb(setting.Cfg)
	testSubCate = &SubCategory{
		Name: testSubCateName,
	}
}

func clearCategory(t *testing.T) {
	cate := getTestCate(t)
	if cate != nil {
		cate.Delete()
	}
}

func getTestSubCategory(t *testing.T) *SubCategory {
	actual, err := testSubCate.GetSome(0, 1, "")
	if err != nil {
		t.Error(errors.New("No sub-category record found"))
		return nil
	}
	subCates, ok := actual.([]SubCategory)
	if !ok {
		t.Error("Failed to cast query result to sub-category collection")
		return nil
	}
	return &subCates[0]
}

func TestCreateSubCategory(t *testing.T) {
	subCate, _ := testSubCate.(*SubCategory)
	cate := prepareCategory(t)
	subCate.CategoryID = cate.CategoryID
	if err := subCate.Create(); err != nil {
		t.Error(err)
	}
}

func TestGetSingleSubCategory(t *testing.T) {
	testSubCate = getTestSubCategory(t)
	if testSubCate != nil {
		if err := testSubCate.GetSingle(); err != nil {
			t.Error(err)
		}
	}
}

func TestGetSomeSubCategories(t *testing.T) {
	getTestSubCategory(t)
}

func TestGetTotalSubCategories(t *testing.T) {
	count, err := testSubCate.GetTotal("")
	if err != nil {
		t.Error(err)
		return
	}
	if count == 0 {
		t.Error(errors.New("No sub-category record found"))
	}
}

func TestHasNameBySubCategory(t *testing.T) {
	exists, err := testSubCate.HasName()
	if err != nil {
		t.Error(err)
		return
	}
	if !exists {
		t.Error(errors.New("Determines the name of sub-category exists false"))
	}
}

func TestGetSubCategoryByName(t *testing.T) {
	expect := getTestSubCategory(t)
	if expect != nil {
		actual := &SubCategory{
			Name: testSubCateName,
		}
		if err := actual.GetByName(); err != nil {
			t.Error(err)
			return
		}
		if expect.SubCategoryID != actual.SubCategoryID {
			t.Error(errors.New("Get the test sub-category by name not match"))
		}
	}
}

func TestEditSubCategory(t *testing.T) {
	const updatedSubCategoryName = "changed"
	testSubCate = getTestSubCategory(t)
	if testSubCate != nil {
		subCate, _ := testSubCate.(*SubCategory)
		subCate.Name = updatedSubCategoryName
		if err := subCate.Edit(); err != nil {
			t.Error(err)
			return
		}
	}
	testSubCate = getTestSubCategory(t)
	subCate, _ := testSubCate.(*SubCategory)
	if subCate.Name != updatedSubCategoryName {
		t.Error(errors.New("SubCategory edit failed"))
	}
}

func TestDeleteSubCategory(t *testing.T) {
	testSubCate = getTestSubCategory(t)
	if testSubCate != nil {
		if err := testSubCate.Delete(); err != nil {
			t.Error(err)
		}
	}
}