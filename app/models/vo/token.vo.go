package vo

// UserClaims 自定义jwt数据体
type UserClaims struct {
	UserId        int64  `json:"userId"`
	DeptId        int64  `json:"deptId"`
	DeptName      string `json:"deptName"`
	Username      string `json:"username"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	IsSuper       bool   `json:"isSuper"`       // 是否为超级用户
	Token         string `json:"token"`         // 会话 Token
	IpAdder       string `json:"ipAdder"`       // IP地址
	LoginLocation string `json:"loginLocation"` // 登录地址
	LoginTime     int64  `json:"loginTime"`     // 登录时间
	Os            string `json:"os"`            // 操作系统
	Browser       string `json:"browser"`       // 浏览器
}
