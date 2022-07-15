package response

import (
	"sixstar/shopstar-micro/client/app/util"
)

/**
 * @Com www.sixstaredu.com
 * @Author 六星教育-shineyork老师
 */

type AttributeList struct {
	//Attrs []goods.Attribute
}

type CateSku struct {
	//Sku [][]goods.AttributeDetail
}

type Tbdetail struct {
	Data map[int64]*util.GrapAtt
}