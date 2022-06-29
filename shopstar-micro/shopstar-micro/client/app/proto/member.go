package proto

//注册请求
type RegisterReq struct {
	Username string
	Password string
	Phone string
}
//登陆请求
type LoginReq struct {
	Username string
	Password string
}
type MemberReq struct {
	Id        string
	AddressId int
	Username  string
}
//返回用户信息
type MemberResp struct {
	Id        string
	Username  string
	Status    int
	Nickname  string
	Headerimg string
	Phone     string
}
// 会员地址
type MemberAddressReq struct {
	UserId    string	// 用户id
	Recipient string	// 收货用户名
	Phone     string	// 手机号码
	Province  int		// 省份
	City      int		// 城市
	District  int		// 地区
	Address   string	// 详细地址
	IsDefault byte
}

type MemberAddressResp struct {
	AddressId int
	UserId    string	// 用户id
	Recipient string	// 收货用户名
	Phone     string	// 手机号码
	Province  int		// 省份
	City      int		// 城市
	District  int		// 地区
	Address   string	// 详细地址
	IsDefault byte
}


type Member interface {

	Login(req LoginReq,resp *MemberResp) error
	Register(req RegisterReq,resp *MemberResp) error
	UserInfo(req MemberReq, resp *MemberResp) error
	// 新怎会员地址
	AddMemberAddress(req MemberAddressReq, resp *MemberAddressResp) error
	GetMemberAddressInfo(req MemberReq, resp *MemberAddressResp) error
}
