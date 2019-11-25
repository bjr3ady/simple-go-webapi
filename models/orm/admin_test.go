package orm

import (
	"errors"
	"git.r3ady.com/golang/school-board/pkg/setting"
	util "github.com/bjr3ady/go-util"
	"testing"
	"time"
)

var testAdmin NameSpecifier

func prepareRole(t *testing.T) *Role {
	role := getRole(t)
	if role == nil {
		prepareFunc(t)
		TestCreateRole(t)
		return getRole(t)
	}
	return role
}

func clearPreparedRoleFunc(t *testing.T) {
	fun := prepareFunc(t)
	fun.Delete()
	role := prepareRole(t)
	role.Delete()
}

func init() {
	ConnectDb(setting.Cfg)
	testAdmin = &Admin{
		Name:        "admin",
		Token:       util.GUID(),
		TokenExpire: time.Now().Unix(),
		Pwd:         "123",
		Role:        []Role{},
	}
}

func TestCreateAdmin(t *testing.T) {
	role := prepareRole(t)
	admin, _ := testAdmin.(*Admin)
	admin.Role = append(admin.Role, *role)
	if err := admin.Create(); err != nil {
		t.Error(err)
	}
}

func getTestAdmin(t *testing.T) *Admin {
	actual, err := testAdmin.GetSome(0, 1, "")
	if err != nil {
		t.Error(errors.New("Failed to query all admins"))
		return nil
	}
	admins, ok := actual.([]Admin)
	if !ok {
		t.Error(errors.New("Failed to cast query result to admins collection"))
		return nil
	}
	return &admins[0]
}

func TestGetSingleAdmin(t *testing.T) {
	admin := getTestAdmin(t)
	if admin != nil {
		if err := admin.GetSingle(); err != nil {
			t.Error(err)
		}
	}
}

func TestGetSomeAdmins(t *testing.T) {
	actual, err := testAdmin.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
	}
	admins, ok := actual.([]Admin)
	if !ok {
		t.Error(errors.New("Failed to cast query result to admins collection"))
	}
	if len(admins) == 0 {
		t.Error(errors.New("No admin record found"))
	}
}

func TestGetTotal(t *testing.T) {
	count, err := testAdmin.GetTotal("")
	if err != nil {
		t.Error(err)
	}
	if count == 0 {
		t.Error(errors.New("No admin record found"))
	}
}

func TestHasAdminName(t *testing.T) {
	exist, err := testAdmin.HasName()
	if err != nil {
		t.Error(err)
	}
	if !exist {
		t.Error(errors.New("Has name function error"))
	}
}

func TestGetAdminByName(t *testing.T) {
	if err := testAdmin.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestChangeAdminName(t *testing.T) {
	admin := getTestAdmin(t)
	if admin != nil {
		const newName = "new"
		admin.Name = newName
		if err := admin.Edit(); err != nil {
			t.Error(err)
			return
		}
		newAdmin := getTestAdmin(t)
		if newAdmin.Name != newName {
			t.Error(errors.New("Update admin name failed"))
		}
		return
	}
	t.Error(errors.New("Failed to get admin by name"))
}

func TestChangeAdminRole(t *testing.T) {
	//Create new Role
	newRole := &Role{
		Name: "newrole",
	}
	fun := getTestFunc(t)
	newRole.Func = append(newRole.Func, *fun)
	if errRole := newRole.Create(); errRole != nil {
		t.Error(errRole)
	}
	newRole.GetByName()

	admin := getTestAdmin(t)
	admin.Role = append(admin.Role, *newRole)
	if err := admin.Edit(); err != nil {
		t.Error(err)
	}
	updatedAdmin := getTestAdmin(t)
	updatedAdminRoles := updatedAdmin.Role
	if len(updatedAdminRoles) != 2 {
		t.Error(errors.New("Failed to add new role to admin"))
	}
	var existRole bool
	for _, role := range updatedAdminRoles {
		if role.RoleID == newRole.RoleID {
			existRole = true
			break
		}
	}
	if !existRole {
		t.Error(errors.New("Failed to determine new role in updated admin"))
	}
	
	newRole.Delete()
}

func TestDeleteAdmin (t *testing.T) {
	testAdmin = getTestAdmin(t)
	if testAdmin == nil {
		t.Error(errors.New("Failed to get test admin"))
		return
	}
	if err := testAdmin.Delete(); err != nil {
		t.Error(err)
	}
	clearPreparedRoleFunc(t)
}