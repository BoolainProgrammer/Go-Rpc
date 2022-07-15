package logic

import (
	"go.uber.org/zap"
	"gorm.io/gorm"
	"strconv"
	"strings"
	menu "toboefa/app/models/menu"
	menuApeRelation "toboefa/app/models/menuApiRelation"
	"toboefa/app/transform/request"
	"toboefa/app/util"
	"toboefa/global"
)

type Menu struct {
}

var MenuService = new(Menu)

func (*Menu) AddMenu(req *request.Menu) error {
	_, err := MenuService.FindMenu(&request.Menu{Name: req.Name})
	if err != global.ErrGormRecordNotFound {
		return global.ErrMenuNameExist
	}
	_, err = MenuService.FindMenu(&request.Menu{Code: req.Code})
	if err != global.ErrGormRecordNotFound {
		return global.ErrMenuCodeExist
	}
	menuData := MenuService.unmarshalMenu(req)
	menuData.CreatedUser = "admin"
	menuData.CreatedTime = util.TimeNowStr()
	menuData.UpdatedUser = "admin"
	menuData.UpdatedTime = util.TimeNowStr()
	req.ApiServiceList = strings.Trim(req.ApiServiceList, ",")
	apiIdList := strings.Split(req.ApiServiceList, ",")
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err = global.DB.Create(menuData).Error; err != nil {
			zap.Error(err)
			return err
		}
		var id []int
		global.DB.Raw("select LAST_INSERT_ID() as id").Pluck("id", &id)
		for _, s := range apiIdList {
			menuApiRelation := new(menuApeRelation.MenuApeRelation)
			menuApiRelation.MenuId = id[0]
			menuApiRelation.ApiServiceId, _ = strconv.Atoi(s)
			menuApiRelation.CreatedTime = util.TimeNowStr()
			menuApiRelation.UpdatedTime = util.TimeNowStr()
			menuApiRelation.CreatedUser = req.UpdatedUser
			menuApiRelation.UpdatedUser = req.UpdatedUser
			if err = global.DB.Omit("created_time").Omit("created_user").Save(&menuApiRelation).Error; err != nil {
				zap.Error(err)
				return err
			}
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (*Menu) SaveMenu(req *request.Menu) error {
	menuInfo, err := MenuService.FindMenu(&request.Menu{Name: req.Name})
	if err != global.ErrGormRecordNotFound && err != nil {
		return global.ErrUpdateMenu
	}
	if menuInfo.Id != req.Id {
		return global.ErrMenuNameExist
	}
	menuInfo, err = MenuService.FindMenu(&request.Menu{Code: req.Code})
	if err != global.ErrGormRecordNotFound && err != nil {
		return global.ErrUpdateMenu
	}
	if menuInfo.Id != req.Id {
		return global.ErrMenuCodeExist
	}
	menuData := MenuService.unmarshalMenu(req)
	menuData.Id = req.Id
	menuData.UpdatedUser = req.UpdatedUser
	menuData.UpdatedTime = util.TimeNowStr()
	req.ApiServiceList = strings.Trim(req.ApiServiceList, ",")
	apiIdList := strings.Split(req.ApiServiceList, ",")
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		if err = global.DB.Omit("created_time").Omit("created_user").Save(&menuData).Error; err != nil {
			zap.Error(err)
			return err
		}
		for _, s := range apiIdList {
			menuApiRelation := new(menuApeRelation.MenuApeRelation)
			menuApiRelation.MenuId = req.Id
			menuApiRelation.ApiServiceId, _ = strconv.Atoi(s)
			menuApiRelation.CreatedTime = util.TimeNowStr()
			menuApiRelation.UpdatedTime = util.TimeNowStr()
			menuApiRelation.CreatedUser = req.UpdatedUser
			menuApiRelation.UpdatedUser = req.UpdatedUser
			if err = global.DB.Omit("created_time").Omit("created_user").Save(&menuApiRelation).Error; err != nil {
				zap.Error(err)
				return err
			}
		}
		if err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}

	return nil
}

func (*Menu) FindMenu(req *request.Menu) (*menu.Menu, error) {
	menuInfo := new(menu.Menu)
	tx := global.DB

	if len(req.Name) > 0 {
		tx.Where("name = ? ", req.Name)
	}
	if req.Id > 0 {
		tx.Where("id = ? ", req.Id)
	}
	if len(req.Code) > 0 {
		tx.Where("code = ?", req.Code)
	}
	err := tx.First(&menuInfo).Error

	if err == gorm.ErrRecordNotFound {
		return nil, global.ErrGormRecordNotFound
	}
	if err != nil {
		zap.Error(err)
		return menuInfo, global.ErrGormRecord
	}
	return menuInfo, nil
}

func (*Menu) FindMenuList(name string, isDisplay string) ([]*menu.Menu, error) {
	menuList := ([]*menu.Menu)(nil)
	tx := global.DB
	if len(name) > 0 {
		tx.Where("name like ? ", "%"+name+"%")
	}
	if len(isDisplay) > 0 {
		tx.Where("is_display = ? ", isDisplay)
	}
	tx.Where("parent_id = 0")
	err := tx.Find(&menuList).Error
	if err != nil {
		zap.Error(err)
		return menuList, err
	}
	for i, menuSon := range menuList {
		menuList[i].Child, err = MenuService.FindMenuChildListByParentId(menuSon.Id)
		if err != nil {
			return menuList, nil
		}
	}
	return menuList, nil
}

func (*Menu) FindMenuChildListByParentId(parentId int) ([]menu.Menu, error) {
	var menuList []menu.Menu
	err := global.DB.Where("parent_id = ? ", parentId).Find(&menuList).Error
	if err == gorm.ErrRecordNotFound {
		return menuList, nil
	}
	if err != nil {
		zap.Error(err)
		return menuList, err
	}
	for i, m := range menuList {
		menuList[i].Child, err = MenuService.FindMenuChildListByParentId(m.Id)
		if err != nil {
			zap.Error(err)
			return nil, err
		}
	}
	return menuList, nil
}

func (*Menu) unmarshalMenu(req *request.Menu) (menuData *menu.Menu) {
	menuData.ParentId = req.ParentId
	menuData.Name = req.Name
	menuData.Code = req.Code
	menuData.Type = req.Type
	menuData.Content = req.Content
	menuData.DisplayOrder = req.DisplayOrder
	menuData.IsDisplay = req.IsDisplay
	menuData.Status = req.Status
	return menuData
}
