package util

import (
	"strconv"
	"strings"
)

// 采用装饰器对属性进行封装

type SplicComponent interface {
	Id() string
	Name() string
}

type AttrSplice struct {
	c SplicComponent
	attrId int64
	attrValId int64
	attrValName string
}
// attrId:attrValId; =>
func (d *AttrSplice) Id() string {
	if d.c == nil {
		return ""
	}
	var ret strings.Builder

	str := d.c.Id()
	if len(str) != 0 { // 注意注意注意
		ret.WriteString(str)
		ret.WriteString(";")
	}

	ret.WriteString(strconv.Itoa(int(d.attrId)))
	ret.WriteString(":")
	ret.WriteString(strconv.Itoa(int(d.attrValId)))

	return ret.String()
}

func (d *AttrSplice) Name() string {
	if d.c == nil {
		return ""
	}

	return d.c.Name() + d.attrValName + " "
}

func NewAttrSplice(c SplicComponent, attrId,attrValId int64, attrValName string) SplicComponent {
	return &AttrSplice{
		c: c,
		attrId: attrId,
		attrValId: attrValId,
		attrValName: attrValName,
	}
}