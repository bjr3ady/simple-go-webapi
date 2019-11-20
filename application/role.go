package application

import (
	"errors"

	models "git.r3ady.com/golang/school-board/models/orm"
	constant "git.r3ady.com/golang/school-board/models/constant"
	
	logger "github.com/bjr3ady/go-logger"
)


var roleModel models.NameSpecifier

//GetTheDefaultRole query the default system role model.
func GetTheDefaultRole() (models.Role, error) {
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.Name = constant.DEFAULT_ROLE
	err := role.GetByName()
	if err != nil {
		logger.Error("Failed to get the default role.")
		return models.Role{}, err
	}
	return *role, nil
}

//NewRole creates new role model.
func NewRole(name string, funcIDs []string) error {
	hasName, _ := RoleHasName(name)
	if hasName {
		err := errors.New("name of role already exists")
		logger.Error(err)
		return err
	}
	funcs := []models.Func{}
	for _, funcID := range funcIDs {
		funcs = append(funcs, models.Func{FuncID: funcID})
	}
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.Name = name
	role.Func = funcs
	return role.Create()
}

//GetRoleByID gets the ID specific role model.
func GetRoleByID(roleID string) (models.Role, error) {
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.RoleID = roleID
	if err := role.GetSingle(); err != nil {
		return models.Role{}, err
	}
	return *role, nil
}

//RoleHasName determines the specific name of role model exists.
func RoleHasName(name string) (bool, error) {
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.Name = name
	return role.HasName()
}

//GetRoleByName gets the name specific role model.
func GetRoleByName(name string) (models.Role, error) {
	hasName, err := RoleHasName(name)
	if !hasName {
		err = errors.New("the specific name of role does not exist")
		return models.Role{}, err
	}
	if err != nil {
		return models.Role{}, err
	}
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.Name = name
	if err = role.GetByName(); err != nil {
		return models.Role{}, err
	}
	return *role, nil
}

//QueryRoles get multiple role models.
func QueryRoles(pageIndex, pageNum int, where interface{}) ([]models.Role, error) {
	var roles []models.Role
	roleModel = &models.Role{}
	result, err := roleModel.GetSome(pageIndex, pageNum, where)
	if err != nil {
		return roles, err
	}
	roles, ok := result.([]models.Role)
	if !ok {
		err = errors.New("failed to cast query result to role models")
		return roles, err
	}
	return roles, nil
}

//TotalRoles get total number of role models.
func TotalRoles(where interface{}) (int, error) {
	roleModel = &models.Role{}
	count, err := roleModel.GetTotal(where)
	if err != nil {
		return -1, err
	}
	return count, nil
}

//EditRole modify name, functions of role model.
func EditRole(roleID, name string, funcNames []string) error {
	roleModel = &models.Role{}
	role := roleModel.(*models.Role)
	role.RoleID = roleID
	err := role.GetSingle()
	if err != nil {
		return err
	}
	funcs := []models.Func{}
	for _, funcName := range funcNames {
		funcm, err := GetFuncByName(funcName)
		if err != nil {
			return err
		}
		funcs = append(funcs, funcm)
	}
	role.Func = funcs
	role.Name = name
	return role.Edit()
}

//RemoveRole deletes role model.
func RemoveRole(roleID string) error {
	role, err := GetRoleByID(roleID)
	if err != nil {
		return err
	}
	if err = role.Delete(); err != nil {
		return err
	}
	return nil
}