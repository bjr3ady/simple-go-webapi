package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"git.r3ady.com/golang/school-board/application"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/e"
	logger "github.com/bjr3ady/go-logger"
	"github.com/unknwon/com"
)

//GetAllContents query all contents
func GetAllContents(w http.ResponseWriter, r *http.Request) {
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
		records, err := application.QueryContents(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			res.Data = records
		}
	}
	res.Code = code
	res.Response()
}

//CreateContent create new content
func CreateContent(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		content := &models.Content{}
		if err := json.Unmarshal(reqBytes, &content); err != nil {
			logger.Info("Failed to unmarshal content parameters while creating new one", err)
			code = e.INVALID_PARAMS
			res.Data = false
		} else {
			if err := application.NewContent(content.Content, content.SubCategoryID, content.VideoSrc); err != nil {
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

//GetOneContent query the specific content by id
func GetOneContent(w http.ResponseWriter, r *http.Request) {
	contentID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if contentID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		content := &models.Content{ContentID: contentID}
		if err := content.GetSingle(); err != nil {
			code = e.NO_SUB_CATEGORY_RECORD_FOUND
			res.Data = false
		} else {
			res.Data = content
		}
	}
	res.Code = code
	res.Response()
}

//UpdateContent updates specific content model
func UpdateContent(w http.ResponseWriter, r *http.Request) {
	contentID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil || contentID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		content := &models.Content{}
		if err := json.Unmarshal(reqBytes, &content); err != nil {
			logger.Info("failed to cast content parameters", err)
			code = e.INVALID_PARAMS
			res.Data = false
		} else {
			if err := application.EditContent(content.Content, content.SubCategoryID, content.VideoSrc); err != nil {
				logger.Info(err)
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

//DeleteContent remove specific content
func DeleteContent(w http.ResponseWriter, r *http.Request) {
	contentID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if contentID == "" {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		if err := application.RemoveContent(contentID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}
