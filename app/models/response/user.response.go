package response

import "time"

// CaptchaImageResponse 验证码返回
type CaptchaImageResponse struct {
	Uuid       string `json:"uuid"`       //uuid码
	Image      string `json:"image"`      //base64图片
	ExpireTime int64  `json:"expireTime"` //到期时间
}

// UserLoginResponse 用户登录返回
type UserLoginResponse struct {
	Token      string `json:"token"`      //token信息
	ExpireTime int64  `json:"expireTime"` //到期时间
}

// UserInfoResponse 用户信息
type UserInfoResponse struct {
	UserId   int64   `json:"userId"`   // 用户ID
	UserName string  `json:"userName"` // 用户名称
	NickName string  `json:"nickName"` // 用户昵称
	DeptId   int64   `json:"deptId"`   // 部门ID
	Phone    string  `json:"phone"`    // 手机号
	Email    string  `json:"email"`    // 邮箱
	Sex      int     `json:"sex"`      // 性别
	Status   int     `json:"status"`   // 状态
	PostId   []int64 `json:"postId"`   // 岗位
	RoleId   []int64 `json:"roleId"`   // 角色
	Remark   string  `json:"remark"`   // 备注
}

// UserPageResponse 用户分页
type UserPageResponse struct {
	UserId     int64     `json:"userId"`     // 用户ID
	UserName   string    `json:"userName"`   // 用户名称
	NickName   string    `json:"nickName"`   // 昵称
	DeptName   string    `json:"deptName"`   // 部门名称
	Phone      string    `json:"phone"`      // 手机号
	Status     int       `json:"status"`     // 状态
	CreateTime time.Time `json:"createTime"` // 创建时间
	IsSuper    bool      `json:"isSuper"`    // 是否为超级用户
}
