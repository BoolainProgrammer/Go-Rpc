package request
/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */
type Pay struct {
	OrderId int `form:"order_id"`
}

type AliUnifyPay struct {
	AppId          string `form:"app_id"`
	AuthAppId      string `form:"auth_app_id"`
	BuyerId        string `form:"buyer_id"`
	BuyerPayAmount string `form:"buyer_pay_amount"`
	FundBillList   string `form:"fund_bill_list"`
	InvoiceAmount  string `form:"invoice_amount"`
	NotifyId       string `form:"notify_id"`
	NotifyTime     string `form:"notify_time"`
	NotifyType     string `form:"notify_type"`
	OutTradeNo     string `form:"out_trade_no"`
	PointAmount    string `form:"point_amount"`
	ReceiptAmount  string `form:"receipt_amount"`
	SellerId       string `form:"seller_id"`
	Sign           string `form:"sign"`
	SignType       string `form:"sign_type"`
	Subject        string `form:"subject"`
	TradeNo        string `form:"trade_no"`
	TradeStatus    string `form:"trade_status"`
	Version        string `form:"version"`
}

type Ali struct {
}
