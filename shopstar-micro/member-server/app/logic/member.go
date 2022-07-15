package logic

import (
	"gorm.io/gorm"
	"sixstar/shopstar-micro/member-server/app/models"
	"sixstar/shopstar-micro/member-server/app/proto"
	"sixstar/shopstar-micro/member-server/app/util"
	coreModel "sixstar/shopstar-micro/member-server/core/model"
	"sixstar/shopstar-micro/member-server/global"
)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
type Member struct {
}
var MemberLogic = new(Member)
// 注意用户名不重复，
// 注意id生成规则 =》 防止别人根据用户id知道相关的数据量
func (logic *Member) Register(req proto.RegisterReq, resp *proto.MemberResp)  error {
	model := new(models.Member)
	// 判断用户名是否重复
	if _,err := logic.FindMemberByUsername(req.Username); err != global.ErrGormRecordNotFound {
		return  global.ErrUsernameExist
	}

	// 密码加密 =》 哈希加密
	pw, err := util.GenPasswordHash(req.Password)
	if err != nil {
		return global.ErrPasswordHash
	}
	// 完善model
	model.Id = coreModel.WUID()
	model.Password = string(pw)
	model.Username = req.Username
	model.Status   = 1
	model.Phone    = req.Password
	// 添加
	if err := global.DB.Create(model).Error; err != nil {
		// 系统需要记录异常，非对外异常
		return global.ErrRegister
	}

	//*resp = proto.MemberResp{
	//	Id:        model.Id,
	//	Username:  model.Username,
	//	Status:    model.Status,
	//}
	unmarshalMember(model, resp)
	return nil
}
// 登入
func (logic *Member) Login(req proto.LoginReq, resp *proto.MemberResp) error {
	// 查询用户
	model,err := logic.FindMemberByUsername(req.Username)
	if  err != nil {
		return err
	}
	// 密码校验
	if util.ValidatePasswordHash(req.Password, model.Password) {

		unmarshalMember(model, resp)

		return nil
	}

	return global.ErrPassword
}
// 用户信息
func (logic *Member) UserInfo(req proto.MemberReq, resp *proto.MemberResp) error {
	model := new(models.Member)
	err := global.DB.Debug().Where("id = ?", req.Id).First(model).Error

	if err != nil {
		return err
	}

	//*resp = proto.MemberResp{
	//	Id:        model.Id,
	//	Username:  model.Username,
	//	Status:    model.Status,
	//	Nickname:  model.Nickname,
	//	Headerimg: model.Headerimg,
	//	Phone:     model.Phone,
	//}
	unmarshalMember(model, resp)

	return nil
}

// 根据用户名查询
func (logic *Member) FindMemberByUsername(username string) (*models.Member, error) {
	model := new(models.Member)

	err := global.DB.Where("username = ?", username).First(model).Error

	if err == gorm.ErrRecordNotFound {
		return model, global.ErrGormRecordNotFound
	}

	if err != nil {
		return model, global.ErrGormRecord

	}

	return model, nil
}
func (logic *Member) FindMemberByPhone(phone string) (*models.Member, error) {
	model := new(models.Member)

	err := global.DB.Where("phone = ?", phone).First(model).Error

	if err == gorm.ErrRecordNotFound {
		return model, global.ErrGormRecordNotFound
	}

	if err != nil {
		return model, global.ErrGormRecord

	}

	return model, nil
}
// 获取购买者的用户信息
func (logic *Member) GetMember(id string) (*models.Member, error) {
	model := new(models.Member)
	err := global.DB.Debug().Where("id = ?", id).First(model).Error
	return model, err
}

func unmarshalMember(model *models.Member, resp *proto.MemberResp)  {
	*resp = proto.MemberResp{
		Id:        model.Id,
		Username:  model.Username,
		Status:    model.Status,
		Nickname:  model.Nickname,
		Headerimg: model.Headerimg,
		Phone:     model.Phone,
	}
}
