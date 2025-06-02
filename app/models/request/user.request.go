package request

type UserLoginParam struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
	Captcha  string `json:"captcha" binding:"required"` // 验证码
	Uuid     string `json:"uuid" binding:"required"`    // 验证码ID
}

type PageParam struct {
	CommonPage
	DeptId int64  `json:"deptId"`
	Status int    `json:"status"`
	Name   string `json:"name"`
}

// UserCreateRequest 用户创建
type UserCreateRequest struct {
	UserName   string  `json:"userName"`   // 用户名称
	NickName   string  `json:"nickName"`   // 用户昵称
	Password   string  `json:"password"`   // 密码
	DeptId     int64   `json:"deptId"`     // 部门ID
	Phone      string  `json:"phone"`      // 手机号
	Email      string  `json:"email"`      // 邮箱
	Sex        int     `json:"sex"`        // 性别
	Status     int     `json:"status"`     // 状态
	PostId     []int64 `json:"postId"`     // 岗位
	RoleId     []int64 `json:"roleId"`     // 角色
	Remark     string  `json:"remark"`     // 备注
	CreateName string  `json:"createName"` // 创建人名称
}
