package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/unknwon/com"

	logger "github.com/bjr3ady/go-logger"
	
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/application"
	"git.r3ady.com/golang/school-board/pkg/e"
)

//GetDefaultCategory gets the default category
func GetDefaultCategory(w http.ResponseWriter, r *http.Request) {
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	res.Data = application.GetDefaultCategory()
	res.Code = code
	res.Response()
}

//GetOneCategory get specific category data
func GetOneCategory(w http.ResponseWriter, r *http.Request) {
	cateID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if cateID == "" {
		code = e.INVALID_PARAMS
	} else {
		cate , err := application.GetCategoryByID(cateID)
		if err != nil {
			code = e.NO_CATEGORY_RECORD_FOUND
		} else {
			res.Data = cate
		}
	}
	res.Code = code
	res.Response()
}

//GetAllCategories get all categories by pagging
func GetAllCategories(w http.ResponseWriter, r *http.Request) {
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
		records, err := application.QueryCategories(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			res.Data = records
		}
	}
	res.Code = code
	res.Response()
}

//CreateCategory creates new category
func CreateCategory(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		category := models.Category{}
		if err = json.Unmarshal(reqBytes, &category); err != nil {
			logger.Error("Failed to unmarshal admin parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.NewCategory(category.Name, category.Icon, category.BannerBgColor, category.Thumb); err != nil {
				code = e.CREATE_FAILED
			} else {
				res.Data = true
			}
		}
	}
	res.Code = code
	res.Response()
}

//UpdateCategory updates specifc category
func UpdateCategory(w http.ResponseWriter, r *http.Request) {
	cateID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}

	if cateID == "" || err != nil {
		code = e.INVALID_PARAMS
	} else {
		cate := &models.Category{}
		if err = json.Unmarshal(reqBytes, &cate); err != nil {
			logger.Error("Failed to unmarshal category parameters", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.EditCategory(cateID, cate.Name, cate.Icon, cate.BannerBgColor, cate.Thumb); err != nil {
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

//DeleteCategory delete specific category model
func DeleteCategory(w http.ResponseWriter, r *http.Request) {
	cateID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if cateID == "" {
		code = e.INVALID_PARAMS
	} else {
		if err := application.RemoveCategory(cateID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}