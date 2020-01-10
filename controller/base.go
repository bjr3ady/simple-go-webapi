package controller

import (
	"encoding/json"
	"net/http"
	"net/url"

	"git.r3ady.com/golang/simple-go-webapi/pkg/e"
	logger "github.com/bjr3ady/go-logger"
	"github.com/gorilla/mux"
)

//CollectionResult is the type of collection query result
type CollectionResult struct {
	Collection interface{} `json:"collection"`
	Count int `json:"count"`
}

//JSONResponse is the common type of RESTful api response data strcut.
type JSONResponse struct {
	Code   int         `json:"code"`
	Data   interface{} `json:"data"`
	Msg    string      `json:"msg"`
	Writer http.ResponseWriter
}

//Response write json response to http pipline.
func (jr *JSONResponse) Response() {
	jr.Msg = e.GetMsg(jr.Code)
	jr.Writer.Header().Set("Content-Type", "application/json")
	jr.Writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(jr.Writer).Encode(jr); err != nil {
		logger.Error("Failed to encode json response:", err)
	}
}

func reqID(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["id"]
}

func reqName(r *http.Request) string {
	vars := mux.Vars(r)
	return vars["name"]
}

func getURLParams(r *http.Request) url.Values {
	uq, _ := url.Parse(r.URL.String())
	params, _ := url.ParseQuery(uq.RawQuery)
	return params
}