package vo

type UserClaims struct {
	UserId   int64
	DeptId   int64
	DeptName string
	UserName string
	Email    string
	Phone    string
	IsSuper  bool //是否超级用户
}
