package orm

import (
	// logger "github.com/bjr3ady/go-logger"
	// util "github.com/bjr3ady/go-util"
)

type Role struct {
	Model
	RoleID string `json:"role_id" gorm:"index"`
	Name string `json:"name"`
	Func []Func `json:"func" gorm:"many2many:role_func"`
}

