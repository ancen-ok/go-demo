package controller

import (
	"encoding/json"
	"gitee.com/molonglove/goboot/gin"
	"go-demo/app/models/entity"
	"go-demo/app/models/vo"
	"runtime"
	"time"
)

type BaseController struct {
}

func (b *BaseController) FunctionName() string {
	pc := make([]uintptr, 1)
	runtime.Callers(2, pc)
	f := runtime.FuncForPC(pc[0])
	return f.Name()
}

// Paese 解析前置数据
func (b *BaseController) Parse(c *gin.Context, title string, businessType vo.BusinessType, param any) (*vo.UserClaims, *entity.Operate) {
	var (
		value  any
		ip     string
		claims *vo.UserClaims
		bytes  []byte
		now    time.Time
	)
	value, _ = c.Get(vo.ClaimsInfo)
	claims = value.(*vo.UserClaims)
	ip = c.ClientIP()

	bytes, _ = json.Marshal(param)
	now = time.Now()
	return claims, &entity.Operate{
		Title:         title,
		BusinessType:  int(businessType),
		Method:        b.FunctionName(),
		RequestMethod: c.Request.Method,
		OperatorType:  1,
		OperName:      claims.Username,
		DeptName:      claims.DeptName,
		OperUrl:       c.Request.URL.Path,
		OperIp:        ip,
		OperParam:     string(bytes),
		OperTime:      &now,
	}

}
