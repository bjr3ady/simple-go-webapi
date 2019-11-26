package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	"errors"

	"github.com/unknwon/com"
	logger "github.com/bjr3ady/go-logger"

	"git.r3ady.com/golang/school-board/application"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/e"
)

//GetAllHomes query all home lists
func GetAllHomes(w http.ResponseWriter, r *http.Request) {
	params := getURLParams(r)
	startIndexStr := params.Get("startindex")
	countStr := params.Get("count")
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if startIndexStr == "" || countStr == "" {
		code = e.INVALID_PARAMS
	} else {
		startIndex := com.StrTo(startIndexStr).MustInt()
		count := com.StrTo(countStr).MustInt()
		records, err := application.QueryHomes(startIndex, count, "")
		if err != nil {
			code = e.NO_HOME_RECORD_FOUND
		} else {
			count, err := application.TotalHomes("")
			if err != nil {
				code = e.GET_TOTAL_FAILED
			} else {
				res.Data = &CollectionResult{Collection: records, Count: count}
			}
		}
	}
	res.Code = code
	res.Response()
}

//CreateHome create new home
func CreateHome(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		home := &models.Home{}
		if err := json.Unmarshal(reqBytes, &home); err != nil {
			logger.Info(errors.New("failed to cast home parameters while creating new home item"), err)
			code = e.INVALID_PARAMS
			res.Data = false
		} else {
			if err := application.NewHome(home.CategoryID, home.Link, home.Index, home.SizeMode, home.IsDirectLink); err != nil {
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

//GetOneHome query specific home item by id
func GetOneHome(w http.ResponseWriter, r *http.Request) {
	homeID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if homeID == "" {
		code = e.INVALID_PARAMS
	} else {
		if home, err := application.GetHomeByID(homeID); err != nil {
			code = e.NO_HOME_RECORD_FOUND
		} else {
			res.Data = home
		}
	}
	res.Code = code
	res.Response()
}

//UpdateHome updates specific home item
func UpdateHome(w http.ResponseWriter, r *http.Request) {
	homeID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if homeID == "" || err != nil {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		home := &models.Home{}
		if err := json.Unmarshal(reqBytes, &home); err != nil {
			code = e.INVALID_PARAMS
			res.Data = false
		} else {
			if err := application.EditHome(homeID, home.CategoryID, home.Link, home.Index, home.SizeMode, home.IsDirectLink); err != nil {
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

//DeleteHome delete specific home item
func DeleteHome(w http.ResponseWriter, r *http.Request) {
	homeID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if homeID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		if err := application.RemoveHome(homeID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}