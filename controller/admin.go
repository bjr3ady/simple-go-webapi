package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/unknwon/com"

	logger "github.com/bjr3ady/go-logger"

	"git.r3ady.com/golang/simple-go-webapi/application"
	models "git.r3ady.com/golang/simple-go-webapi/models/orm"
	"git.r3ady.com/golang/simple-go-webapi/pkg/e"
)

//GetOneAdmin get id specific admin
func GetOneAdmin(w http.ResponseWriter, r *http.Request) {
	adminID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if adminID == "" {
		code = e.INVALID_PARAMS
	} else {
		admin, err := application.GetAdminByID(adminID)
		if err != nil {
			code = e.NO_ADMIN_RECORD_FOUND
		} else {
			res.Data = admin
		}
	}
	res.Code = code
	res.Response()
}

//GetAdminByName get name specific admin data
func GetAdminByName(w http.ResponseWriter, r *http.Request) {
	name := reqName(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if name == "" {
		code = e.INVALID_PARAMS
	} else {
		admin, err := application.GetAdminByName(name)
		if err != nil {
			code = e.NO_ADMIN_RECORD_FOUND
		} else {
			res.Data = admin
		}
	}
	res.Code = code
	res.Response()
}

//GetAllAdmins get all admins by pagging
func GetAllAdmins(w http.ResponseWriter, r *http.Request) {
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
		records, err := application.QueryAdmins(startIndex, count, "")
		if err != nil {
			code = e.ERROR
		} else {
			count, err := application.TotalAdmins("")
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

//CreateAdmin creates new admin
func CreateAdmin(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}
	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		admin := models.Admin{}
		if err = json.Unmarshal(reqBytes, &admin); err != nil {
			logger.Info("Failed to unmarshal admin parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			if err = application.NewAdmin(admin.Name, admin.Pwd); err != nil {
				code = e.CREATE_FAILED
			} else {
				res.Data = true
			}
		}
	}
	res.Code = code
	res.Response()
}

//LoginAdmin handle admin login proecess.
func LoginAdmin(w http.ResponseWriter, r *http.Request) {
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}

	if err != nil {
		code = e.INVALID_PARAMS
	} else {
		admin := models.Admin{}
		if err = json.Unmarshal(reqBytes, &admin); err != nil {
			logger.Info("Failed to unmarshal admin parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			token, err := application.LoginAdmin(admin.Name, admin.Pwd)
			if err != nil {
				code = e.ADMIN_LOGIN_FAILED
			} else {
				res.Data = token
			}
		}
	}
	res.Code = code
	res.Response()
}

//UpdateAdmin updates specific admin
func UpdateAdmin(w http.ResponseWriter, r *http.Request) {
	adminID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}

	if adminID == "" || err != nil {
		code = e.INVALID_PARAMS
	} else {
		admin := models.Admin{}
		if err = json.Unmarshal(reqBytes, &admin); err != nil {
			logger.Info("Failed to unmarshal admin parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			roleIDs := []string{}
			for _, role := range admin.Role {
				roleIDs = append(roleIDs, role.RoleID)
			}
			if err = application.EditAdmin(adminID, admin.Name, roleIDs); err != nil {
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

//UpdateAdminPassword updates specific admin's password
func UpdateAdminPassword(w http.ResponseWriter, r *http.Request) {
	adminID := reqID(r)
	reqBytes, err := ioutil.ReadAll(r.Body)
	code := e.SUCCESS
	res := &JSONResponse{Writer: w}

	if adminID == "" || err != nil {
		code = e.INVALID_PARAMS
	} else {
		adminPwd := application.PasswordChange{}
		if err := json.Unmarshal(reqBytes, &adminPwd); err != nil {
			logger.Info("Failed to unmarshal admin password change parameters:", err)
			code = e.INVALID_PARAMS
		} else {
			if err := application.UpdateAdminPassword(adminID, adminPwd.OriginalPwd, adminPwd.NewPwd); err != nil {
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

//DeleteAdmin deletes specific admin
func DeleteAdmin(w http.ResponseWriter, r *http.Request) {
	adminID := reqID(r)
	res := &JSONResponse{Writer: w}
	code := e.SUCCESS
	if adminID == "" {
		code = e.INVALID_PARAMS
	} else {
		if err := application.RemoveAdmin(adminID); err != nil {
			code = e.REMOVE_FAILED
			res.Data = false
		} else {
			res.Data = true
		}
	}
	res.Code = code
	res.Response()
}

//TestT is the structure for test token json
type TestT struct {
	ID    string `json:"id"`
	Token string `json:"token"`
	URL   string `json:"url"`
}

//TestToken generate bearer token
func TestToken(w http.ResponseWriter, r *http.Request) {
	res := &JSONResponse{Writer: w}
	reqBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		res.Code = e.INVALID_PARAMS
	} else {
		auth := &TestT{}
		if err := json.Unmarshal(reqBytes, &auth); err != nil {
			res.Code = e.INVALID_PARAMS
		} else {
			res.Code = e.SUCCESS
			res.Data = application.GenerateBearerToken(auth.ID, auth.Token, auth.URL)
		}
	}
	res.Response()
}
