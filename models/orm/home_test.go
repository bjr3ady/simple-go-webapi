package orm

import (
	"testing"
	"errors"
	"git.r3ady.com/golang/school-board/pkg/setting"
)

var testHome NormalModeler

func init() {
	ConnectDb(setting.Cfg)
	testHome = &Home{
		Index: 0,
	}
}

func prepareCategory(t *testing.T) *Category {
	cate := getTestCate(t)
	if cate == nil {
		TestCreateCategory(t)
		return getTestCate(t)
	}
	return cate
}

func getTestHome(t *testing.T) *Home {
	actual, err := testHome.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	homes, ok := actual.([]Home)
	if !ok {
		t.Error(errors.New("Failed to cast query result to home collection"))
		return nil
	}
	return &homes[0]
}

func TestCreateHome(t *testing.T) {
	home, _ := testHome.(*Home)
	cate := prepareCategory(t)
	home.CategoryID = (*cate).CategoryID
	if err := home.Create(); err != nil {
		t.Error(err)	
	}
}

func TestGetSingleHome(t *testing.T) {
	testHome = getTestHome(t)
	if testHome != nil {
		if err := testHome.GetSingle(); err != nil {
			t.Error(err)
		}
	}
}

func TestGetSomeHomes(t *testing.T) {
	getTestHome(t)
}

func TestGetTotalHomes(t *testing.T) {
	count, err := testHome.GetTotal("")
	if err != nil {
		t.Error(err)
		return
	}
	if count <= 0 {
		t.Error(errors.New("No home record found"))
	}
}

func TestEditHome(t *testing.T) {
	testHome = getTestHome(t)
	if testHome != nil {
		home, _ := testHome.(*Home)
		home.IsDirectLink = 1
		if err := home.Edit(); err != nil {
			t.Error(err)
			return
		}
		testHome = getTestHome(t)
		newHome, _ := testHome.(*Home)
		if newHome.IsDirectLink != home.IsDirectLink {
			t.Error(errors.New("Update home's is_direct_link field failed"))
		}
	}
}

func TestDeleteHome(t *testing.T) {
	testHome = getTestHome(t)
	if err := testHome.Delete(); err != nil {
		t.Error(err)
	}
	cate := prepareCategory(t)
	cate.Delete()
}