package controller

import (
	"encoding/json"
	application "git.r3ady.com/golang/school-board/application"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/e"
	logger "github.com/bjr3ady/go-logger"
	"github.com/unknwon/com"
	"io/ioutil"
	"net/http"
)

//GetDefaultSubCategory gets the default sub-category data
func GetDefaultSubCategory(w http.ResponseWriter, r *http.Request) {
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	res.Data = application.GetDefaultCategory()
	res.Code = code
	res.Response()
}

//GetOneSubCategory gets the id specific sub-category data.
func GetOneSubCategory(w http.ResponseWriter, r *http.Request) {
	id := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if id == "" {
		code = e.INVALID_PARAMS
	} else {
		subCate, err := application.GetSubCategoryByID(id)
		if err != nil {
			code = e.NO_SUB_CATEGORY_RECORD_FOUND
		} else {
			res.Data = subCate
		}
	}
	res.Code = code
	res.Response()
}

//GetSubCategoryByName gets the name specific sub-category data
func GetSubCategoryByName(w http.ResponseWriter, r *http.Request) {
	name := reqName(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if name == "" {
		code = e.INVALID_PARAMS
	} else {
		subCate, err := application.GetSubCategoryByName(name)
		if err != nil {
			code = e.NO_SUB_CATEGORY_RECORD_FOUND
		} else {
			res.Data = subCate
		}
	}
	res.Code = code
	res.Response()
}

//GetAllSubCategories query all sub-categories
func GetAllSubCategories(w http.ResponseWriter, r *http.Request) {
	params := getURLParams(r)
	startIndexStr := params.Get("startindex")
	countStr := params.Get("count")
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS

	if startIndexStr == "" || countStr == "" {
		code = e.INVALID_PARAMS
	} else {
		startIndex := com.StrTo(startIndexStr).MustInt()
		count := com.StrTo(countStr).MustInt()
		records, err := application.QuerySubCategories(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			res.Data = records
		}
	}
	res.Code = code
	res.Response()
}

//CreateSubCategory creates new sub-category
func CreateSubCategory(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		subCate := &models.SubCategory{}
		if err = json.Unmarshal(reqBytes, &subCate); err != nil {
			logger.Info("failed to unmarshal sub-category parameters", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.NewSubCategory(subCate.Name, subCate.CategoryID); err != nil {
				code = e.CREATE_FAILED
				res.Data = false
			} else {
				res.Data = true
			}
		}
	}
	res.Code = code
	res.Response()
}

//UpdateSubCategory updates sub-category
func UpdateSubCategory(w http.ResponseWriter, r *http.Request) {
	subCateID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil || subCateID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		subCate := &models.SubCategory{}
		if err = json.Unmarshal(reqBytes, &subCate); err != nil {
			logger.Info("failed to unmarshal sub-category parameters", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.EditSubCategory(subCateID, subCate.Name, subCate.CategoryID); err != nil {
				code = e.UPDATE_FAILED
				res.Data = false
			} else {
				res.Data = true
			}
		}
	}
	res.Code = code
	res.Response()
}

//DeleteSubCategory deletes sub-category
func DeleteSubCategory(w http.ResponseWriter, r *http.Request) {
	subCateID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if subCateID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		if err := application.RemoveSubCategory(subCateID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}
