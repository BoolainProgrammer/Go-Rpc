package util

import (
	"encoding/json"
	"strconv"
	"time"
	"fmt"
	"github.com/dgrijalva/jwt-go"

	"toboefa/core/cache"
)

const VKey = "key_version:"

var (
	JwtExp     = time.Minute.Truncate(3)
	VersionJwt = "key_version_jwt"
)

var JWT = &Jwt{
	Iat: time.Now().Unix(),
	// pc设置为15分钟
	// 手机端设计2小时 - 30分钟
	Exp: time.Now().Add(JwtExp),
}

type AuthData map[string]interface{}

func NewAuthData(memberId string) AuthData {
	return AuthData{
		"member_id": memberId,
	}
}
func (auth AuthData) MemberId() string {
	return auth["member_id"].(string)
}

type Jwt struct {
	Iat int64 // 生效时间
	Exp time.Time
}

//func NewJwt() *Jwt{
//	return &Jwt{
//		Iat: 1,
//		Exp: time.Now(),
//	}
//}
// 生成jwt token
func (j *Jwt) GenerateToken(key string, Data AuthData) (string, error) {
	// 设置token过期
	Data[VersionJwt] = strconv.Itoa(int(time.Now().Unix()))
	// 更新本地版本
	err := j.invalid(Data)

	if err != nil {
		fmt.Printf("[GenerateToken] err = %v\n", err)
	}

	data, err := json.Marshal(Data)

	if err != nil {
		return "", err
	}
	// 创建jwttoken
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iat":  j.Iat,
		"exp":  j.Exp,
		"data": string(data),
	})

	token,err := claims.SignedString([]byte(key))

	return token,err
}

// jwt的token解析
func (j *Jwt) ParseToken(tokenStr, key string) (v AuthData, err error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(key), nil
	})
	if err != nil {
		return nil, err
	}
	// token 信息验证
	if err := token.Claims.Valid(); err != nil {
		return nil, err
	}

	data := token.Claims.(jwt.MapClaims)["data"].(string)
	json.Unmarshal([]byte(data), &v)

	return
}

// jwt的刷新
// 刷新-》先解析-》获取信息-》再生成token
func (j *Jwt) RefreshToken(tokenStr, key string) (string, error) {
	// 等到解析的正常token
	data, err := j.ParseToken(tokenStr, key)
	if err != nil {
		return "", err
	}

	// 去对token设置过期
	return j.GenerateToken(key, data)
}

// 对内设置过期
// 对jwttoken的版本更新
func (j *Jwt) invalid(data AuthData) error {

	return cache.CacheManager.Set(
		VKey+data.MemberId(),
		//"key",
		[]byte(data[VersionJwt].(string)),
		&cache.Options{Expiration: JwtExp},
	)

}

// 对外的设置过期
func (j *Jwt) Invalid(tokenStr, key string) error {
	data, err := j.ParseToken(tokenStr, key)
	if err != nil {
		return err
	}
	// 设置token过期
	data[VersionJwt] = strconv.Itoa(int(time.Now().Unix()))
	// 更新本地版本
	j.invalid(data)
	return nil
}

// 判断是否过期
// 1. 解析失败不通过
// 2. 缓存存在信息（无论是设置过期，刷新token 都会存在缓存）； 但是token中version不存在
// 3. 本地版本与token版本校验失败
// true: 没有 ； false 过期
func (j *Jwt) Expired(tokenStr, key string) bool {
	data, err := j.ParseToken(tokenStr, key)

	if err != nil {
		return false
	}
	// 因为第一次生成的token一直在用，所以缓存没有信息
	v, err := cache.CacheManager.Get(VKey+data.MemberId())

	if err != nil {
		return false
	}

	if v == nil {
		return false
	}
	// 用户在缓存中存在信息，说明正常的选择退出或者刷新令牌了；
	// 所以目前可能是第一次生成token，表示过期

	if data[VersionJwt] == nil {
		return false
	}

	// token的比对
	if string(v.([]byte)) == data[VersionJwt].(string) {
		return true
	}

	return false
}
