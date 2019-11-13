package orm

import (
	// logger "github.com/bjr3ady/go-logger"
	// util "github.com/bjr3ady/go-util"
)

type Admin struct {
	Model
	AdminID	string `json:"admin_id" gorm:"index"`
	Name string `json:"name"`
	Token string `json:"token"`
	TokenExpire string `json:"token_expire"`
	Pwd string `json:"pwd"`
	Role []Role `json:"role" gorm:"many2many:admin_role"`
}