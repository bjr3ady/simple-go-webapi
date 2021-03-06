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

//GetDefaultRole gets the default role
func GetDefaultRole(w http.ResponseWriter, r *http.Request) {
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	role, err := application.GetTheDefaultRole()
	if err != nil {
		code = e.NO_ROLE_RECORD_FOUND
	} else {
		res.Data = role
	}
	res.Code = code
	res.Response()
}

//GetOneRole gets the id specific role data
func GetOneRole(w http.ResponseWriter, r *http.Request) {
	roleID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if roleID == "" {
		code = e.INVALID_PARAMS
	} else {
		role, err := application.GetRoleByID(roleID)
		if err != nil {
			code = e.NO_ROLE_RECORD_FOUND
		} else {
			res.Data = role
		}
	}
	res.Code = code
	res.Response()
}

//GetRoleByName gets the name specific role data
func GetRoleByName(w http.ResponseWriter, r *http.Request) {
	name := reqName(r)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if name == "" {
		code = e.INVALID_PARAMS
	} else {
		role, err := application.GetRoleByName(name)
		if err != nil {
			code = e.NO_FUNC_RECORD_FOUND
		} else {
			res.Data = role
		}
	}
	res.Code = code
	res.Response()
}

//GetAllRoles get all roles by pagging
func GetAllRoles(w http.ResponseWriter, r *http.Request) {
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
		records, err := application.QueryRoles(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			count, err := application.TotalRoles("")
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

//CreateRole creates new role
func CreateRole(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		role := models.Role{}
		if err = json.Unmarshal(reqBytes, &role); err != nil {
			logger.Info("Failed to unmarshal role parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			funcIDs := []string{}
			for _, funcm := range role.Func {
				funcIDs = append(funcIDs, funcm.FuncID)
			}
			if err = application.NewRole(role.Name, funcIDs); err != nil {
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

//UpdateRole updates specific role.
func UpdateRole(w http.ResponseWriter, r *http.Request) {
	roleID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil || roleID == "" {
		code = e.INVALID_PARAMS
	} else {
		role := models.Role{}
		if err = json.Unmarshal(reqBytes, &role); err != nil {
			logger.Info("Faile to unmarshal role parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			funcNames := []string{}
			for _, funcm := range role.Func {
				funcNames = append(funcNames, funcm.Name)
			}
			if err = application.EditRole(roleID, role.Name, funcNames); err != nil {
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

//DeleteRole deletes specific role
func DeleteRole(w http.ResponseWriter, r *http.Request) {
	roleID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if roleID == "" {
		code = e.INVALID_PARAMS
	} else {
		if err := application.RemoveRole(roleID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}
