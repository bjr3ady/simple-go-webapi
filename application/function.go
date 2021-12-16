package application

import (
	"errors"

	logger "github.com/bjr3ady/go-logger"
	constant "github.com/bjr3ady/simple-go-webapi/models/constant"
	models "github.com/bjr3ady/simple-go-webapi/models/orm"
)

var funcModel models.NameSpecifier

//NewFunc creates new system function model.
func NewFunc(name string) error {
	hasName, _ := FuncHasName(name)
	if hasName {
		err := errors.New("name of system function already exists")
		logger.Info(err)
		return err
	}
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.Name = name
	return funcm.Create()
}

//GetTheDefaultFunc query the default system function model.
func GetTheDefaultFunc() (models.Func, error) {
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.Name = constant.DefaultFunction
	err := funcm.GetByName()
	if err != nil {
		logger.Info("Failed to get default system function.")
		return models.Func{}, err
	}
	return *funcm, nil
}

//GetFuncByID gets the ID specific system function model.
func GetFuncByID(funcID string) (models.Func, error) {
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.FuncID = funcID
	if err := funcm.GetSingle(); err != nil {
		return models.Func{}, err
	}
	return *funcm, nil
}

//FuncHasName determines the specific name of system function model exists.
func FuncHasName(funcName string) (bool, error) {
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.Name = funcName
	return funcm.HasName()
}

//GetFuncByName gets the name specific system function model.
func GetFuncByName(funcName string) (models.Func, error) {
	hasName, err := FuncHasName(funcName)
	if !hasName {
		err = errors.New("the specific name of system function does not exist")
		return models.Func{}, err
	}
	if err != nil {
		return models.Func{}, err
	}
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.Name = funcName
	if err = funcm.GetByName(); err != nil {
		return models.Func{}, err
	}
	return *funcm, nil
}

//QueryFuncs get multiple system function models.
func QueryFuncs(pageIndex, pageSize int, where interface{}) ([]models.Func, error) {
	var funcs []models.Func
	funcModel = &models.Func{}
	result, err := funcModel.GetSome(pageIndex, pageSize, where)
	if err != nil {
		return funcs, err
	}
	funcs, ok := result.([]models.Func)
	if !ok {
		err = errors.New("failed to cast query result to system function models")
		return funcs, err
	}
	return funcs, nil
}

//TotalFuncs get total number of system function models.
func TotalFuncs(where interface{}) (int, error) {
	funcModel = &models.Func{}
	return funcModel.GetTotal(where)
}

//EditFunc modify name of system function model.
func EditFunc(funcID, name string) error {
	funcModel = &models.Func{}
	funcm := funcModel.(*models.Func)
	funcm.FuncID = funcID
	err := funcm.GetSingle()
	if err != nil {
		return err
	}
	funcm.Name = name
	return funcm.Edit()
}

//RemoveFunc deletes system function model.
func RemoveFunc(funcID string) error {
	funcm, err := GetFuncByID(funcID)
	if err != nil {
		return err
	}
	if err = funcm.Delete(); err != nil {
		return err
	}
	return nil
}
