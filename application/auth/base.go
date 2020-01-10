package auth

import (
	"net/http"
	"crypto/md5"
	"fmt"

	"github.com/bjr3ady/simple-go-webapi/pkg/setting"
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

//GenerateBearerToken generate new string token.
func GenerateBearerToken(token, ID, URL string) string {
	targetString := fmt.Sprintf("%s,%s,%s", URL, ID, token)
	data := []byte(targetString)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func generateBearerToken(token, ID string, req *http.Request) string {
	URL := fmt.Sprintf("%s://%s%s", setting.HTTPProto, req.Host, req.URL.String())
	return GenerateBearerToken(token, ID, URL)
}