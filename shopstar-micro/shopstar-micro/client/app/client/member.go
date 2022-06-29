package client

import (
	"sixstar/shopstar-micro/client/app/proto"
)

type MemberClient struct {
}


func (c *MemberClient) Login(req proto.LoginReq, resp *proto.MemberResp) error {
	return c.call( "MemberServer.Login", req, resp)
}

func (c *MemberClient) Register(req proto.RegisterReq, resp *proto.MemberResp) error {
	return c.call( "MemberServer.Register", req, resp)
}

func (c *MemberClient) UserInfo(req proto.MemberReq, resp *proto.MemberResp) error {
	return c.call( "MemberServer.UserInfo", req, resp)
}

func (c *MemberClient) AddMemberAddress(req proto.MemberAddressReq, resp *proto.MemberAddressResp) error {
	return c.call( "MemberServer.AddMemberAddress", req, resp)
}

func (c *MemberClient) GetMemberAddressInfo(req proto.MemberReq, resp *proto.MemberAddressResp) error {
	return c.call( "MemberServer.GetMemberAddressInfo", req, resp)
}

func (c *MemberClient) call(serverMethod string, req interface{}, resp interface{}) error {
	return call("member", serverMethod, req, resp)
}