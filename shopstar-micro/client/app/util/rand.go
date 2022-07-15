package util

import (
	"github.com/shopspring/decimal"
	"github.com/valyala/fastrand"
	"math/rand"
	"time"
)

// 随机器

var (
	rngFast fastrand.RNG
	rngMath = rand.New(rand.NewSource(time.Now().Unix()))
)

func RInt(n int) int {
	return int(rngFast.Uint32n(uint32(n)))
}

//
//  RDecimal
//  @Author：六星教育-shineyork老师
//  @Description: 随机金枪工具
//  @param value 随机价格的范围
//  @param exp 价格的精确度
//  @return decimal.Decimal
//
func RDecimal(value, exp int) decimal.Decimal {
	in := rngMath.Float64() + float64(RInt(value))
	return decimal.NewFromFloatWithExponent(in, int32((exp)))
}