package sms

// 简单工厂模式

type Sms interface {
	Send(phone, content string) (string, error)
}

func NewSms(service string) (sms Sms) {
	switch service {
	case "smsbao":
		sms = &Smsbao{}
	}
	return
}