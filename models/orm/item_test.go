package orm

import (
	"errors"
	"testing"
	"git.r3ady.com/golang/school-board/pkg/setting"
)

var testItem NameSpecifier
const testItemName = "test"

func init() {
	ConnectDb(setting.Cfg)
	testItem = &Item{
		Name: testItemName,
	}
}

func getTestItem(t *testing.T) *Item {
	actual, err := testItem.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	items, ok := actual.([]Item)
	if !ok {
		t.Error(errors.New("Failed to cast query result to item collection"))
		return nil
	}
	return &items[0]
}

func prepareSubCategory(t *testing.T) *SubCategory {
	prepareCategory(t)
	subCate := &SubCategory{}
	total, _ := subCate.GetTotal("")
	if total == 0 {
		TestCreateSubCategory(t)
	}
	return getTestSubCategory(t)
}

func clearSubCategory(t *testing.T) {
	subCate := prepareSubCategory(t)
	subCate.Delete()
}

func TestCreateItem(t *testing.T) {
	subCate := prepareSubCategory(t)
	item, _ := testItem.(*Item)
	item.SubCateID = subCate.SubCategoryID
	if err := testItem.Create(); err != nil {
		t.Error(err)
	}
}

func TestGetSingleItem (t *testing.T) {
	testItem = getTestItem(t)
	if testItem != nil {
		if err := testItem.GetSingle(); err != nil {
			t.Error(err)
		}
	}
}

func TestGetItemByName (t *testing.T) {
	if err := testItem.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestHasItemName(t *testing.T) {
	exists, err := testItem.HasName()
	if err != nil {
		t.Error(err)
		return
	}
	if !exists {
		t.Error(errors.New("No item record found"))
	}
}

func TestGetSomeItems(t *testing.T) {
	getTestItem(t)
}

func TestGetTotalItems(t *testing.T) {
	count, err := testItem.GetTotal("")
	if err != nil {
		t.Error(err)
		return
	}
	if count == 0 {
		t.Error(errors.New("No item record found"))
	}
}

func TestEditItem (t *testing.T) {
	const updatedItemName = "new"
	testItem = getTestItem(t)
	item, _ := testItem.(*Item)
	item.Name = updatedItemName
	if err := item.Edit(); err != nil {
		t.Error(err)
		return
	}
	item = getTestItem(t)
	if item.Name != updatedItemName {
		t.Error(errors.New("Failed to update item name"))
	}
}

func TestDeleteItem(t *testing.T) {
	testItem := getTestItem(t)
	if err := testItem.Delete(); err != nil {
		t.Error(err)
	}
	clearSubCategory(t)
	clearCategory(t)
}