package server

import (
	"sixstar/shopstar-micro/member-server/app/logic"
	"sixstar/shopstar-micro/member-server/app/proto"
)

type MemberServer struct {
}

func NewMemberServer() *MemberServer {
	return &MemberServer{}
}

func (*MemberServer) Login(req proto.LoginReq, resp *proto.MemberResp) error {
	return logic.MemberLogic.Login(req, resp)
}
func (s *MemberServer) Register(req proto.RegisterReq, resp *proto.MemberResp) error {
	return logic.MemberLogic.Register(req, resp)
}

func (s *MemberServer) UserInfo(req proto.MemberReq, resp *proto.MemberResp) error {
	return logic.MemberLogic.UserInfo(req, resp)
}

func (s *MemberServer) AddMemberAddress(req proto.MemberAddressReq, resp *proto.MemberAddressResp) error {
	return logic.MemberAddressLogic.AddMemberAddress(req, resp)
}

func (s *MemberServer) GetMemberAddressInfo(req proto.MemberReq, resp *proto.MemberAddressResp) error {
	return logic.MemberAddressLogic.GetMemberAddressInfo(req, resp)
}