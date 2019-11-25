package application

import (
	"errors"
	"time"

	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/application/auth"
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

//PasswordChange is the JSON struct for updating admin password
type PasswordChange struct {
	OriginalPwd string `json:"original_pwd"`
	NewPwd      string `json:"new_pwd"`
}

//AdminLoginResult is the JSON struct for admin login
type AdminLoginResult struct {
	Func    []models.Func `json:"funcs"`
	AdminID string        `json:"adminid"`
	Token   string        `json:"token"`
}

var adminModel models.NameSpecifier

func getAdminFunctions(roles []models.Role) ([]models.Func, error) {
	var funcs []models.Func
	for _, role := range roles {
		role, err := GetRoleByID(role.RoleID)
		if err != nil {
			return funcs, err
		}
		for _, fun := range role.Func {
			funcs = append(funcs, fun)
		}
	}
	return funcs, nil
}

func updateAdminToken(id string) (string, error) {
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.AdminID = id
	err := admin.GetSingle()
	if err != nil {
		return "", err
	}
	token := util.GUID()
	admin.Token = token
	admin.TokenExpire = time.Now().Add(time.Minute * 30).Unix()
	return token, admin.Edit()
}

//NewAdmin creates new admin model.
func NewAdmin(name, pwd string) error {
	hasName, _ := AdminHasName(name)
	if hasName {
		err := errors.New("name of admin already exists")
		logger.Error(err)
		return err
	}

	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.Name = name
	admin.Pwd = pwd
	admin.Token = util.GUID()
	admin.TokenExpire = time.Now().Add(time.Second * 30).Unix()
	admin.Role = []models.Role{}

	return admin.Create()
}

//LoginAdmin handle admin login process
func LoginAdmin(name, pwd string) (AdminLoginResult, error) {
	admin, err := GetAdminByName(name)
	if err != nil {
		return AdminLoginResult{}, err
	}
	if pwd != admin.Pwd {
		return AdminLoginResult{}, errors.New("password not match")
	}
	funcs, err := getAdminFunctions(admin.Role)
	if err != nil {
		return AdminLoginResult{}, err
	}
	token, err := updateAdminToken(admin.AdminID)
	if err != nil {
		return AdminLoginResult{}, err
	}
	return AdminLoginResult{Func: funcs, Token: token, AdminID: admin.AdminID}, nil
}

//GetAdminByID gets the ID specific admin model.
func GetAdminByID(adminID string) (models.Admin, error) {
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.AdminID = adminID
	if err := admin.GetSingle(); err != nil {
		return models.Admin{}, err
	}
	return *admin, nil
}

//AdminHasName determines the specific name of admin model exists.
func AdminHasName(adminName string) (bool, error) {
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.Name = adminName
	return admin.HasName()
}

//GetAdminByName gets the name specific admin model.
func GetAdminByName(adminName string) (models.Admin, error) {
	hasName, err := AdminHasName(adminName)
	if !hasName {
		err = errors.New("the specific name of admin does not exist")
		return models.Admin{}, err
	}
	if err != nil {
		return models.Admin{}, err
	}
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.Name = adminName
	if err = admin.GetByName(); err != nil {
		return models.Admin{}, err
	}
	return *admin, nil
}

//QueryAdmins get multiple admin models.
func QueryAdmins(pageIndex, pageSize int, where interface{}) ([]models.Admin, error) {
	var admins []models.Admin
	adminModel = &models.Admin{}
	result, err := adminModel.GetSome(pageIndex, pageSize, where)
	if err != nil {
		return admins, err
	}
	admins, ok := result.([]models.Admin)
	if !ok {
		err = errors.New("failed to cast query result to admin models")
		return admins, err
	}
	return admins, nil
}

//TotalAdmins get total number of admin models.
func TotalAdmins(where interface{}) (int, error) {
	adminModel = &models.Admin{}
	count, err := adminModel.GetTotal(where)
	if err != nil {
		return -1, err
	}
	return count, nil
}

//EditAdmin modify name, roles of admin model.
func EditAdmin(adminID, adminName string, roleIDs []string) error {
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.AdminID = adminID
	err := admin.GetSingle()
	if err != nil {
		return err
	}
	roles := []models.Role{}
	for _, roleID := range roleIDs {
		roles = append(roles, models.Role{RoleID: roleID})
	}
	admin.Role = roles
	admin.Name = adminName
	admin.Token = util.GUID()
	admin.TokenExpire = time.Now().Add(time.Minute * 30).Unix()
	return admin.Edit()
}

//UpdateAdminPassword modify password of admin model.
func UpdateAdminPassword(adminID, originalPwd, newPwd string) error {
	adminModel = &models.Admin{}
	admin := adminModel.(*models.Admin)
	admin.AdminID = adminID
	err := admin.GetSingle()
	if err != nil {
		return err
	}
	if originalPwd != admin.Pwd {
		err = errors.New("original password not match")
		return err
	}
	admin.Pwd = newPwd
	admin.Token = util.GUID()
	admin.TokenExpire = time.Now().Add(time.Minute * 30).Unix()
	return admin.Edit()
}

//RemoveAdmin deletes admin model.
func RemoveAdmin(adminID string) error {
	admin, err := GetAdminByID(adminID)
	if err != nil {
		return err
	}
	if err = admin.Delete(); err != nil {
		return err
	}
	return nil
}

func GenerateBearerToken(id, token, url string) string {
	return auth.GenerateBearerToken(token, id, url)
}