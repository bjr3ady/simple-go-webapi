package orm

import (
	"errors"
	"testing"
	"github.com/bjr3ady/simple-go-webapi/pkg/setting"
)

var testFunc NameSpecifier

func init() {
	ConnectDb(setting.Cfg)
	testFunc = &Func{
		Name: "test",
	}
}

func TestCreateFunc(t *testing.T) {
	if err := testFunc.Create(); err != nil {
		t.Error(err)
	}
}

func getTestFunc(t *testing.T) *Func {
	result, err := testFunc.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	funcs, ok := result.([]Func)
	if !ok {
		t.Error(errors.New("Failed to cast query result to system functions"))
		return nil
	}
	if len(funcs) == 0 {
		return nil
	}
	return &funcs[0]
}

func TestHasFuncByName(t *testing.T) {
	testFunc = getTestFunc(t)
	if exists, err := testFunc.HasName(); !exists || err != nil {
		t.Error(err)
	}
}

func TestGetFuncByName(t *testing.T) {
	if err := testFunc.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestGetSomeFuncs(t *testing.T) {
	actual, err := testFunc.GetSome(0, 10, "")
	if err != nil {
		t.Error(err)
	}
	funcs, ok := actual.([]Func)
	if !ok || len(funcs) == 0 {
		t.Error(errors.New("Failed to cast query result to system function models"))
	}
}

func TestGetSingleFunc(t *testing.T) {
	testFunc = getTestFunc(t)
	if err := testFunc.GetSingle(); err != nil {
		t.Error(err)
	}
}

func TestGetTotalFuncs(t *testing.T) {
	if count, err := testFunc.GetTotal(""); count == 0 || err != nil {
		t.Error(err)
	}
}

func TestEditFunc(t *testing.T) {
	testFunc = getTestFunc(t)
	fun, ok := testFunc.(*Func)
	if !ok {
		t.Error(errors.New("Failed to cast query result to system function"))
	}
	fun.Name = "new"
	if err := fun.Edit(); err != nil {
		t.Error(err)
	}
	if err := fun.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestDeleteFunc(t *testing.T) {
	testFunc = getTestFunc(t)
	if err := testFunc.Delete(); err != nil {
		t.Error(err)
	}
}