package models

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

/*
CREATE TABLE `member` (
  `id` varchar(255) NOT NULL,
  `username` varchar(255) NOT NULL,
  `password` varchar(255) NOT NULL,
  `status` varchar(255) NOT NULL DEFAULT '1',
  `nickname` varchar(255) DEFAULT NULL,
  `headerimg` varchar(255) DEFAULT NULL,
  `phone` varchar(255) NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=MyISAM DEFAULT CHARSET=utf8;\
*/

type Address struct {
	AddressId int
	UserId    string
	Recipient string
	Phone     string
	Province  int
	City      int
	District  int
	Address   string
	IsDefault byte
}

func (model *Address) TableName() string {
	return "member_address"
}
