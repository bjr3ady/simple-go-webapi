package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/unknwon/com"

	models "github.com/bjr3ady/simple-go-webapi/models/orm"
	"github.com/bjr3ady/simple-go-webapi/application"
	"github.com/bjr3ady/simple-go-webapi/pkg/e"
	"github.com/bjr3ady/go-logger"
)

//GetDefaultFunc get default system function.
func GetDefaultFunc(w http.ResponseWriter, r *http.Request) {
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	funcm, err := application.GetTheDefaultFunc()
	if err != nil {
		code = e.NO_FUNC_RECORD_FOUND
	} else {
		res.Data = funcm
	}
	res.Code = code
	res.Response()
}

//GetOneFunc get id specific system function
func GetOneFunc(w http.ResponseWriter, r *http.Request) {
	funcID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if funcID == "" {
		code = e.INVALID_PARAMS
	} else {
		funcm, err := application.GetFuncByID(funcID)
		if err != nil {
			code = e.NO_FUNC_RECORD_FOUND
		} else {
			res.Data = funcm
		}
	}
	res.Code = code
	res.Response()
}

//GetFuncByName get name specific system function data
func GetFuncByName(w http.ResponseWriter, r *http.Request) {
	name := reqName(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if name == "" {
		code = e.INVALID_PARAMS
	} else {
		fun, err := application.GetFuncByName(name)
		if err != nil {
			code = e.NO_FUNC_RECORD_FOUND
		} else {
			res.Data = fun
		}
	}
	res.Code = code
	res.Response()
}

//GetAllFuncs get all system functions by pagging
func GetAllFuncs(w http.ResponseWriter, r *http.Request) {
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
		records, err := application.QueryFuncs(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			count, err := application.TotalFuncs("")
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

//CreateFunc creates new system function
func CreateFunc(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		funcm := models.Func{}
		if err = json.Unmarshal(reqBytes, &funcm); err != nil {
			logger.Info("Failed to unmarshal system function parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.NewFunc(funcm.Name); err != nil {
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

//UpdateFunc updates specific system function.
func UpdateFunc(w http.ResponseWriter, r *http.Request) {
	funcID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}

	if funcID == "" {
		code = e.INVALID_PARAMS
	} else {
		funcm := models.Func{}
		if err = json.Unmarshal(reqBytes, &funcm); err != nil {
			logger.Info("Faile to unmarshal system function parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.EditFunc(funcID, funcm.Name); err != nil {
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

//DeleteFunc deletes specific system function.
func DeleteFunc(w http.ResponseWriter, r *http.Request) {
	funcID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if funcID == "" {
		code = e.INVALID_PARAMS
	} else {
		if err := application.RemoveFunc(funcID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}
