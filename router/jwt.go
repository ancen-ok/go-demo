package router

import (
	"encoding/json"
	"errors"
	"fmt"
	"gitee.com/molonglove/goboot/gin"
	"go-demo/core"
	"go-demo/vo"
)

// JwtMiddle 中间件配置
func JwtMiddle() gin.HandlerFunc {
	return func(context *gin.Context) {
		var (
			parseJwt = func(token string) (*vo.UserClaims, bool, float64, error) {
				var (
					key    string
					value  string
					expire float64
					user   vo.UserClaims
					err    error
				)
				key = fmt.Sprintf("%s:%s", vo.RedisToken, token)
				// 判断是否过期
				if expire = core.Cache.IsExpire(key); expire == 2 {
					return nil, false, 0, nil
				}
				//获取数据
				if value, err = core.Cache.GetKey(key); err != nil {
					return nil, false, 0, errors.New("获取用户数据失败")
				}
				//序列化
				if err = json.Unmarshal([]byte(value), &user); err != nil {
					return nil, false, 0, errors.New("解析用户数据失败")
				}
				return &user, false, expire, nil
			}
		)
	}
}
