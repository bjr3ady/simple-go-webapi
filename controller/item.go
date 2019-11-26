package controller

import (
	"net/http"
	"io/ioutil"
	"encoding/json"
	
	"github.com/unknwon/com"
	logger "github.com/bjr3ady/go-logger"

	"git.r3ady.com/golang/school-board/pkg/e"
	"git.r3ady.com/golang/school-board/application"
	models "git.r3ady.com/golang/school-board/models/orm"
)

//GetAllItems query all items
func GetAllItems(w http.ResponseWriter, r *http.Request) {
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
		if records, err := application.QueryItems(startIndex, count, ""); err != nil {
			code = e.NO_SUB_CATEGORY_RECORD_FOUND
		} else {
			count, err := application.TotalItems("")
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

//CreateItem creates new item model
func CreateItem(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		item := &models.Item{}
		if err := json.Unmarshal(reqBytes, &item); err != nil {
			code = e.INVALID_PARAMS
			res.Data = false
		} else {
			if err := application.NewItem(item.Name, item.Link, item.SubCateID, item.Index); err != nil {
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

//GetOneItem query the specific item by id
func GetOneItem(w http.ResponseWriter, r *http.Request) {
	itemID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if itemID == "" {
		code = e.INVALID_PARAMS
	} else {
		item, err := application.GetItemByID(itemID)
		if err != nil {
			code = e.NO_ITEM_RECORD_FOUND
		} else {
			res.Data = item
		}
	}
	res.Code = code
	res.Response()
}

//UpdateItem updates specific item data
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	itemID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	res := &JSONResponse{Writer: w}
	code :=e.SUCCESS
	if err != nil {
		code = e.INVALID_PARAMS
		res.Data = false
	} else {
		item := &models.Item{}
		if err := json.Unmarshal(reqBytes, &item); err != nil {
			logger.Info("failed to cast item parameters while updating item", err)
			code = e.INVALID_PARAMS
		} else {
			if err := application.EditItem(itemID, item.Name, item.Link, item.SubCateID, item.Index); err != nil {
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

//DeleteItem remove specific item data
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	itemID := reqID(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if itemID == "" {
		code = e.INVALID_PARAMS
		res.Data = true
	} else {
		if err := application.RemoveItem(itemID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}