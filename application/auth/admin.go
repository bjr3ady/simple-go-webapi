package auth

import (
	"errors"
	"net/http"

	"git.r3ady.com/golang/simple-go-webapi/models/orm"
)

//AdminAuth is the bearer token object.
type AdminAuth struct {
	//AdminID is the uid headers item.
	AdminID string
	//Bearer is the bearer token item.
	Bearer string
	//Token is the admin token item.
	Token string
	//Req is the pointer to the http.Request
	Req *http.Request
}

//Auth authenticate admin request.
func (auth *AdminAuth) Auth() (bool, error) {
	if auth.Bearer != "" && auth.AdminID != "" {
		adminModel := &orm.Admin{
			AdminID: auth.AdminID,
		}
		if err := adminModel.GetSingle(); err != nil {
			return false, err
		}
		bearer := generateBearerToken(adminModel.Token, adminModel.AdminID, auth.Req)
		if bearer != auth.Bearer {
			return false, errors.New("bearer token not match")
		}
		return true, nil
	}
	return false, errors.New("authenticate parameters invalid")
}
