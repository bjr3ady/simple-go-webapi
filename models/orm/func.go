package orm

import (
	logger "github.com/bjr3ady/go-logger"
	util "github.com/bjr3ady/go-util"
)

// Func is the ORM ordermachine_func table object.
type Func struct {
	Model
	FuncID string `json:"func_id" gorm:"index"`
	Name string `json:"name"`
}

// Create creates new system function.
func (fun *Func) Create() error {
	fun.FuncID = util.GUID()
	if err := db.Create(&fun).Error; err != nil {
		logger.Error("Failed to create system function.", err)
		return err
	}
	return nil
}

// GetSome query system functions by pagging.
func (fun *Func) GetSome(pageNum, pageSize int, maps interface{}) (interface{}, error) {
	var funcs []Func
	if err := db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&funcs).Error; err != nil {
		logger.Error("Failed to get some system functions.", err)
		return nil, err
	}
	return funcs, nil
}

//GetTotal query count of system functions.
func (fun *Func) GetTotal(maps interface{}) (int, error) {
	var count int
	if err := db.Model(&Func{}).Where(maps).Count(&count).Error; err != nil {
		logger.Error("Failed to query count of system functions.", err)
		return -1, err
	}
	return count, nil
}

//GetSingle query specific system function.
func (fun *Func) GetSingle() error {
	if err := db.Where("func_id=?", fun.FuncID).First(&fun).Error; err != nil {
		logger.Error("Failed to query specific system function by id.", err)
		return err
	}
	return nil
}

//Edit updates specific system function
func (fun *Func) Edit() error {
	if err := db.Model(&Func{}).Where("func_id=?", fun.FuncID).Updates(fun).Error; err != nil {
		logger.Error("Failed to update system function.", err)
		return err
	}
	return nil
}

//Delete deletes specific system function
func (fun *Func) Delete() error {
	if err := db.Where("func_id=?", fun.FuncID).Delete(&Func{}).Error; err != nil {
		logger.Error("Failed to delete specific system function.", err)
		return err
	}
	return nil
}

//HasName determines the specific name of system functino exists.
func (fun *Func) HasName() (bool, error) {
	if err := db.Select("func_id").Where("name=?", fun.Name).First(&fun).Error; err != nil {
		logger.Error("Failed to find specific name of system function.", err)
		return false, err
	}
	return fun.FuncID != "", nil
}

//GetByName query specific system function by name.
func (fun *Func) GetByName() error {
	if err := db.Where("name=?", fun.Name).First(&fun).Error; err != nil {
		logger.Error("Failed to find specific system function by name.", err)
		return err
	}
	return nil
}