package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

// Role is the ORM school_board_role table object.
type Role struct {
	Model
	RoleID string `json:"role_id" gorm:"index"`
	Name   string `json:"name"`
	Func   []Func `json:"func" gorm:"many2many:role_func"`
}

//Create creates new role data.
func (role *Role) Create() error {
	var funcIDs []string
	for _, fun := range role.Func {
		funcIDs = append(funcIDs, fun.FuncID)
	}
	funcs := []Func{}
	if errFunc := db.Where("func_id in (?)", funcIDs).Find(&funcs).Error; errFunc != nil {
		logger.Error("Failed to find role related system functions", errFunc)
		return errFunc
	}
	role.RoleID = util.GUID()
	role.Func = funcs
	if err := db.Create(&role).Association("Func").Append(funcs).Error; err != nil {
		logger.Error("Failed to create new role", err)
		return err
	}
	return nil
}

//GetSingle query specific role data.
func (role *Role) GetSingle() error {
	if err := db.Where("role_id = ?", role.RoleID).First(&role).Error; err != nil {
		logger.Error("Failed to get single role", err)
		return err
	}
	if errFunc := db.Model(&role).Related(&role.Func, "Func").Error; errFunc != nil {
		logger.Error("Failed to get role related system functions.", errFunc)
		return errFunc
	}
	return nil
}

//GetSome query roles by mapping
func (role *Role) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var roles []Role
	if err := db.Preload("Func").Where(maps).Offset(pageNum).Limit(pageSize).Find(&roles).Error; err != nil {
		logger.Error("Failed to query roles by pagging.", err)
		return nil, err
	}
	return roles, nil
}

//GetTotal query the count of roles.
func (Role)GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Role{}).Where(maps).Count(&count).Error; err != nil {
		logger.Error("Failed to query the count of roles", err)
		return -1, err
	}
	return count, nil
}

//Edit updates the specific role data.
func (role *Role) Edit() error {
	var funcIDs []string
	for _, fun := range role.Func {
		funcIDs = append(funcIDs, fun.FuncID)
	}
	funcs := []Func{}
	if errFunc := db.Where("func_id in (?)", funcIDs).Find(&funcs).Error; errFunc != nil {
		logger.Error("Failed to find role related system functions", errFunc)
		return errFunc
	}
	if errRole := db.Where("role_id = ?", role.RoleID).Find(&Role{}).Error; errRole != nil {
		logger.Error("Failed to find target role", errRole)
		return errRole
	}
	db.Model(&role).Association("Func").Replace(funcs)
	if err := db.Model(&role).Updates(role).Error; err != nil {
		logger.Error("Failed to update specific role", err)
		return err
	}
	return nil
}

//Delete deletes the specific role.
func (role *Role) Delete() error {
	if errFunc := db.Model(&role).Association("Func").Clear().Error; errFunc != nil {
		logger.Error("Failed to remove relationships between role and system functions", errFunc)
		return errFunc
	}
	if err := db.Where("role_id = ?", role.RoleID).Delete(role).Error; err != nil {
		logger.Error("Failed to delete specific role", err)
		return err
	}
	return nil
}

//HasName determines the specific role name exists.
func (role *Role) HasName() (bool, error) {
	if err := db.Where("name = ?", role.Name).First(&role).Error; err != nil {
		logger.Error("Failed to find role by name.", err)
		return false, err
	}
	return role.RoleID != "", nil
}

//GetByName query the specific role by name
func (role *Role) GetByName() error {
	if err := db.Where("name = ?", role.Name).First(&role).Error; err != nil {
		logger.Error("Failed to get role by name", err)
		return err
	}
	if errFunc := db.Model(&role).Related(&role.Func, "Func").Error; errFunc != nil {
		logger.Error("Failed to get role related system functions", errFunc)
		return errFunc
	}
	return nil
}