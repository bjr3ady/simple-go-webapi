package orm

import (
	"fmt"
	"time"

	"github.com/go-ini/ini"

	"github.com/bjr3ady/gorm"
	_ "github.com/bjr3ady/gorm/dialects/mysql"
	
)

var db *gorm.DB

// Model is the base struct of all models.
type Model struct {
	ID         int `gorm:"primary_key;auto_increment" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

// NormalModeler is the normal model interface
// which contains simple CRUD methods.
type NormalModeler interface {
	Create() error
	GetSingle() error
	GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error)
	GetTotal(maps interface{}) (int, error)
	Edit() error
	Delete() error
}

// NameSpecifier is ther interfaces for the model
// which name's unique.
type NameSpecifier interface {
	NormalModeler
	HasName() (bool, error)
	GetByName() error
}

//SerialNoSpecifier is the interfaces for the serial number model
type SerialNoSpecifier interface {
	NormalModeler
	HasSerialNo() (bool, error)
	GetBySerialNo() error
}

//ConnectDb connect mysql database.
func ConnectDb(conf *ini.File) error {
	var (
		err                                          error
		dbType, dbName, user, pwd, host, tablePrefix string
	)

	sec, err := conf.GetSection("database")
	if err != nil {
		return err
	}

	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	pwd = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		pwd,
		host,
		dbName))
		// db.LogMode(true)
	if err != nil {
		return err
	}

	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName
	}

	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
	return nil
}

//BeforeCreate append created on value before create new model row.
func (model *Model) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())
	return nil
}

//BeforeUpdate append modified on value before update exists model row.
func (model *Model) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}

//CloseDB close the current db connection.
func CloseDB() {
	defer db.Close()
}
