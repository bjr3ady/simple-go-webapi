package orm

import (
	"errors"
	"testing"
	"git.r3ady.com/golang/school-board/pkg/setting"
)

var testCategory NameSpecifier

func init() {
	ConnectDb(setting.Cfg)
	testCategory = &Category{
		Name: "test",
	}
}

func TestCreateCategory(t *testing.T) {
	if err := testCategory.Create(); err != nil {
		t.Error(err)
	}
}

func getTestCate(t *testing.T) *Category {
	result, err := testCategory.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	cates, ok := result.([]Category)
	if !ok {
		t.Error(errors.New("Failed to cast query result to categories"))
		return nil
	}
	if len(cates) == 0 {
		return nil
	}
	return &cates[0]
}

func TestHasCategoryByName(t *testing.T) {
	testCategory = getTestCate(t)
	if exists, err := testCategory.HasName(); !exists || err != nil {
		t.Error(err)
	}
}

func TestGetCategoryByName(t *testing.T) {
	if err := testCategory.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestGetSomeCategories(t *testing.T) {
	actual, err := testCategory.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
	}
	cates, ok := actual.([]Category)
	if !ok {
		t.Error(errors.New("Failed to cast query result to category collection"))
		return
	}
	if len(cates) == 0 {
		t.Error(errors.New("No category data found"))
	}
}

func TestGetTotalCategories(t *testing.T) {
	count, err := testCategory.GetTotal("")
	if err != nil {
		t.Error(err)
		return
	}
	if count == 0 {
		t.Error(errors.New("No category data found"))
	}
}

func TestEditCategory(t *testing.T) {
	testCategory = getTestCate(t)
	if testCategory != nil {
		const newName = "new"
		cate, _ := testCategory.(*Category)
		cate.Name = newName
		if err := cate.Edit(); err != nil {
			t.Error(err)
			return
		}
		testCategory = getTestCate(t)
		newCate, _ := testCategory.(*Category)
		if newCate.Name != newName {
			t.Error("Faile to update new name to category")
		}
	}
}

func TestDeleteCategory(t *testing.T) {
	testCategory = getTestCate(t)
	if err := testCategory.Delete(); err != nil {
		t.Error(err)
	}
}