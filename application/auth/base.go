package auth

import (
	"net/http"
	"crypto/md5"
	"fmt"

	"git.r3ady.com/golang/school-board/pkg/setting"
)

//UserAuth is the bearer token object.
type UserAuth struct {
	//UserID is the uid headers item.
	UserID string
	//Token is the token headers item.
	Token string
	//UType is the type of item.
	UType int
	//Req is the pointer to the http.Request
	Req *http.Request
}

//Authler is the interface for authenticatable object check authentication.
type Authler interface {
	Auth() (bool, error)
}

func generateBearerToken(token, ID string, req *http.Request) string {
	URL := fmt.Sprintf("%s://%s%s", setting.HTTPProto, req.Host, req.URL.String())
	targetString := fmt.Sprintf("%s,%s,%s", URL, ID, token)
	data := []byte(targetString)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}