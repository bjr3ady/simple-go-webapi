package auth

import (
	"errors"
	app "git.r3ady.com/golang/school-board/application"
	models "git.r3ady.com/golang/school-board/models/orm"
	"git.r3ady.com/golang/school-board/pkg/setting"
	"testing"
)

func init() {
	models.ConnectDb(setting.Cfg)
}

func prepareAdminData() *models.Admin {
	adminModel := &models.Admin{}
	result, _ := adminModel.GetSome(0, 1, "")
	admins, _ := result.([]models.Admin)
	if len(admins) == 0 {
		app.NewAdmin("test", "123")
	}
	admin, _ := app.GetAdminByName("test")
	return &admin
}

func clearAdminData() {
	admin, _ := app.GetAdminByName("test")
	admin.Delete()
}

func mockAdminLogin(admin *models.Admin) app.AdminLoginResult {
	result, _ := app.LoginAdmin(admin.Name, admin.Pwd)
	return result
}

func TestAuth(t *testing.T) {
	admin := prepareAdminData()
	loginResult := mockAdminLogin(admin)
	adminAuth := &AdminAuth{
		AdminID: admin.AdminID,
	Bearer: "",
	Token: loginResult.Token,
	}
	if success, _ := adminAuth.Auth(); success {
		t.Error(errors.New("admin authenticte failed"))
	}
	clearAdminData()
}   
