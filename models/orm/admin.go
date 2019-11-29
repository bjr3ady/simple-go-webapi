package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

// Admin is the ORM school_board_admin table object.
type Admin struct {
	Model
	AdminID     string `json:"admin_id" gorm:"index"`
	Name        string `json:"name"`
	Token       string `json:"token"`
	TokenExpire int64  `json:"token_expire"`
	Pwd         string `json:"pwd"`
	Role        []Role `json:"role" gorm:"many2many:admin_role"`
}

//Create creates new admin data
func (admin *Admin) Create() error {
	var roleIDs []string
	for _, role := range admin.Role {
		roleIDs = append(roleIDs, role.RoleID)
	}
	roles := []Role{}
	if errRole := db.Where("role_id in (?)", roleIDs).Find(&roles).Error; errRole != nil {
		logger.Info("Failed to get roles while create new admin.", errRole)
		return errRole
	}
	admin.AdminID = util.GUID()
	admin.Role = roles
	if err := db.Create(&admin).Association("Role").Append(roles).Error; err != nil {
		logger.Info("Failed to create new admin", err)
		return err
	}
	return nil
}

//GetSingle query specific admin by id.
func (admin *Admin) GetSingle() error {
	if errAdmin := db.Where("admin_id = ?", admin.AdminID).First(&admin).Error; errAdmin != nil {
		logger.Info("Failed to get specific admin.", errAdmin)
		return errAdmin
	}
	if errRole := db.Model(&admin).Related(&admin.Role, "Role").Error; errRole != nil {
		logger.Info("Failed to get admin relaed roles.", errRole)
		return errRole
	}
	return nil
}

//GetSome query admins by pagging
func (Admin) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var admins []Admin
	if err := db.Preload("Role").Where(maps).Offset(pageNum).Limit(pageSize).Find(&admins).Error; err != nil {
		logger.Info("Faile to query admins by pagging", err)
		return nil, err
	}
	return admins, nil
}

//GetTotal query the count of admins
func (Admin) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Preload("Role").Model(&Admin{}).Where(maps).Count(&count).Error; err != nil {
		logger.Info("Failed to query the count of admins.", err)
		return -1, err
	}
	return count, nil
}

//HasName determines the specific name of admin already exists.
func (admin *Admin) HasName() (bool, error) {
	if err := db.Where("name = ?", admin.Name).First(&admin).Error; err != nil {
		logger.Info("Failed to query count of admins", err)
		return false, err
	}
	return admin.AdminID != "", nil
}

//GetByName query the specific name of admin.
func (admin *Admin) GetByName() error {
	if err := db.Where("name = ?", admin.Name).First(&admin).Error; err != nil {
		logger.Info("Failed to query the specific name of admin", err)
		return err
	}
	if errRole := db.Model(&admin).Related(&admin.Role, "Role").Error; errRole != nil {
		logger.Info("Failed to get admin related roles.", errRole)
		return errRole
	}
	return nil
}

//Edit updates specific admin data.
func (admin *Admin) Edit() error {
	var roleIDs []string
	for _, role := range admin.Role {
		roleIDs = append(roleIDs, role.RoleID)
	}
	roles := []Role{}
	if errRole := db.Where("role_id in (?)", roleIDs).Find(&roles).Error; errRole != nil {
		logger.Info("Failed to find admin related roles.", errRole)
		return errRole
	}
	db.Model(&admin).Association("Role").Replace(roles)
	if err := db.Model(&Admin{}).Updates(admin).Error; err != nil {
		logger.Info("Failed to update admin.", err)
		return err
	}
	return nil
}

//Delete deletes specific admin data.
func (admin *Admin) Delete() error {
	if errAdmin := db.Where("admin_id = ?", admin.AdminID).First(&admin).Error; errAdmin != nil {
		logger.Info("Failed to query specific admin", errAdmin)
		return errAdmin
	}
	db.Model(&admin).Association("Role").Clear()
	if err := db.Where("admin_id = ?", admin.AdminID).Delete(admin).Error; err != nil {
		logger.Info("Failed to delete specific admin", err)
		return err
	}
	return nil
}
