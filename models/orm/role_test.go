package orm

import (
	"git.r3ady.com/golang/simple-go-webapi/pkg/setting"
	"errors"
	"testing"
)

var testRole NameSpecifier

func init() {
	ConnectDb(setting.Cfg)
	testRole = &Role{
		Name: "test",
	}
}

func prepareFunc(t *testing.T) *Func {
	fun := getTestFunc(t)
	if fun == nil {
		TestCreateFunc(t)
		return getTestFunc(t)
	}
	return fun
}

func TestCreateRole(t *testing.T) {
	fun := prepareFunc(t)
	role, _ := testRole.(*Role)
	role.Func = append(role.Func, *fun)
	if err := role.Create(); err != nil {
		t.Error(err)
	}
}

func getRole(t *testing.T) *Role{
	actual, err := testRole.GetSome(0, 1, "")
	if err != nil {
		t.Error(err)
		return nil
	}
	roles, ok := actual.([]Role)
	if !ok {
		t.Error(errors.New("Failed to cast result to roles"))
		return nil
	}
	if len(roles) == 0 {
		return nil
	}
	return &roles[0]
}

func TestGetSingleRole(t *testing.T) {
	testRole = getRole(t)
	if err := testRole.GetSingle(); err != nil {
		t.Error(err)
	}
}

func TestHasRoleName(t *testing.T) {
	testRole = getRole(t)
	if exists, err := testRole.HasName(); err != nil || !exists {
		t.Error(err)
	}
}

func TestGetRoleByName(t *testing.T) {
	if err := testRole.GetByName(); err != nil {
		t.Error(err)
	}
}

func TestGetSomeRoles(t *testing.T) {
	if actual, err := testRole.GetSome(0, 1, ""); err == nil {
		if roles, ok := actual.([]Role); !ok {
			t.Error(errors.New("Failed to cast result to roles"))
		} else if len(roles) == 0 {
			t.Log(errors.New("No role exists"))
		}
	} else {
		t.Error(err)
	}
}

func TestGetTotalRoles(t *testing.T) {
	count, err := testRole.GetTotal("")
	if err != nil {
		t.Error("Failed to get total roles")
	}
	if count == 0 {
		t.Error("No role record")
	}
}

func TestEditRole(t *testing.T) {
	testRole = getRole(t)
	role, ok := testRole.(*Role)
	if !ok {
		t.Error("Failed to cast role model")
	}
	role.Name = "new"
	if err := role.Edit(); err != nil {
		t.Error("Failed to edit role")
	}
	testRole = getRole(t)
	if newRole, _ := testRole.(*Role); newRole.Name != "new" {
		t.Error("Role name not changed")
	}
}

func TestDeleteRole(t *testing.T) {
	testRole = getRole(t)
	if err := testRole.Delete(); err != nil {
		t.Error(err)
		return
	}
	fun := getTestFunc(t)
	fun.Delete()
}