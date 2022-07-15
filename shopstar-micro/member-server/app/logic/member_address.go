package logic

import (
	"errors"
	"sixstar/shopstar-micro/member-server/app/proto"
	"sixstar/shopstar-micro/member-server/global"
	"sixstar/shopstar-micro/member-server/app/models"
)


type MemberAddress struct {
}

var MemberAddressLogic = new(MemberAddress)
// 添加会员地址
func (logic *MemberAddress) AddMemberAddress(req proto.MemberAddressReq, resp *proto.MemberAddressResp) error {
	adds, err := logic.GetMemberAddressAll(req.UserId)
	if err != nil {
		return err
	}

	if len(adds) >= 10{
		return errors.New("个人收件地址不能超过10个")
	}

	model := &models.Address {
		UserId:    req.UserId,
		Recipient: req.Recipient,
		Phone:     req.Phone,
		Province:  req.Province,
		City:      req.City,
		Address:   req.Address,
		District:  req.District,
	}

	err = global.DB.Create(model).Error
	if err != nil {
		return err
	}

	unmarshalMemberAddress(model, resp)

	return nil
}


func (logic *MemberAddress) GetMemberAddressInfo(req proto.MemberReq, resp *proto.MemberAddressResp) (err error) {
	var model models.Address
	err = global.DB.Where("user_id = ? and is_default = 1 ", req.Id).First(&model).Error
	if err != nil {
		return err
	}

	unmarshalMemberAddress(&model, resp)

	return
}


func (logic *MemberAddress) GetMemberAddress(addressId int) (address models.Address, err error){
	err = global.DB.Where("address_id = ? ", addressId).First(&address).Error
	return
}
func (logic *MemberAddress) GetMemberAddressAll(userId string) (address []*models.Address, err error){
	err = global.DB.Where("user_id = ? ", userId).First(&address).Error
	return
}
func unmarshalMemberAddress(model *models.Address, resp *proto.MemberAddressResp)  {
	*resp = proto.MemberAddressResp{
		AddressId: model.AddressId,
		UserId:    model.UserId,
		Recipient: model.Recipient,
		Phone:     model.Phone,
		Province:  model.Province,
		City:      model.City,
		District:  model.District,
		Address:   model.Address,
		IsDefault: model.IsDefault,
	}
}
