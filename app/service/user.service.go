package service

import (
	"errors"
	"fmt"
	"gitee.com/molonglove/goboot/gin"
	"gitee.com/molonglove/goboot/gorm"
	"github.com/jinzhu/copier"
	"github.com/mojocn/base64Captcha"
	"go-demo/app/dao"
	"go-demo/app/models/entity"
	"go-demo/app/models/request"
	"go-demo/app/models/response"
	"go-demo/app/models/vo"
	"go-demo/core"
	"sync"
	"time"
)

var User = NewUserService()

type UserService struct {
	captcha *base64Captcha.Captcha
}

func NewUserService() *UserService {
	driver := base64Captcha.NewDriverDigit(40, 135, 5, 0.4, 72)
	store := &redisStore{expiration: time.Minute * 5}
	return &UserService{
		captcha: base64Captcha.NewCaptcha(driver, store),
	}
}

type redisStore struct {
	sync.Mutex
	expiration time.Duration
}

func (r *redisStore) Set(id string, value string) error {
	key := vo.RedisCaptcha + ":" + id
	_, err := core.Cache.SetKeyValue(key, value, time.Minute*5)
	return err
}

func (r *redisStore) Get(id string, clear bool) string {
	key := fmt.Sprintf("%s:%s", vo.RedisCaptcha, id)

	result, err := core.Cache.GetKey(key)
	if err == nil {
		return result
	}

	// 取失败时再判断是否清除
	if clear {
		_ = core.Cache.Delete(key) // Delete 只传 key，除非你接口真要3个参数
	}

	return ""
}

func (r *redisStore) Verify(id, answer string, clear bool) bool {
	return true
}

func (u *UserService) CaptchaImage() (*response.CaptchaImageResponse, *response.BusinessError) {
	var (
		id     string
		base64 string
		err    error
	)
	if id, base64, _, err = u.captcha.Generate(); err != nil {
		return nil, response.NewBusinessError(response.CaptchaImageError)
	}

	return &response.CaptchaImageResponse{
		Uuid:       id,
		Image:      base64,
		ExpireTime: time.Now().Add(time.Minute * 5).Unix(),
	}, nil
}

func (u *UserService) UserLogin(param *request.UserLoginParam, ctx *gin.Context) (*response.UserLoginResponse, *response.BusinessError) {
	var (
		captchaVerify = func(id, code string) error {
			if !core.Cache.Exist(vo.RedisCaptcha + ":" + id) {
				return errors.New("验证码已过期！")
			}
			if !u.captcha.Verify(id, code, true) {
				return errors.New("验证码不对")
			}
			return nil
		}
		//验证用户
		getUserInfoWithVerity = func(username, password string) (*vo.UserClaims, error) {
			var (
				user entity.User
				dept entity.Dept
				err  error
			)
			if user, err = dao.User.GetUserByUserName(username); err != nil {
				return nil, errors.New("登录用户" + username + "不存在！")
			}
			return &vo.UserClaims{
				UserId:   user.UserId,
				Username: user.UserName,
				DeptId:   user.DeptId,
				DeptName: dept.DeptName,
				Phone:    user.Phone,
				Email:    user.Email,
				IsSuper:  user.UserId == vo.SUPER_USER_ID,
			}, nil
		}
		//登录日志
		loginLogger = func(c *gin.Context, user *vo.UserClaims, msg string, status int) {

		}
		user            *vo.UserClaims
		err             error
		loginFailLogger = func(c *gin.Context, username string, msg string, status int) {

		}
	)

	// 验证码验证
	if err = captchaVerify(param.Uuid, param.Captcha); err != nil {
		loginFailLogger(ctx, param.UserName, err.Error(), 0)
		return nil, response.LoginBusinessError(err.Error())
	}
	//获取用户信息
	if user, err = getUserInfoWithVerity(param.UserName, param.Password); err != nil {
		loginFailLogger(ctx, param.UserName, err.Error(), 0)
		return nil, response.LoginBusinessError(err.Error())
	}
	loginLogger(ctx, user, "登录成功", 1)
	return &response.UserLoginResponse{
		Token:      fmt.Sprintf("%s_%s", param.UserName, param.Password),
		ExpireTime: 30,
	}, nil
}

// GetUserInfo 获取用户信息
func (u *UserService) GetUserInfo(userId int64) (*response.UserInfoResponse, *response.BusinessError) {
	var (
		user   entity.User
		result response.UserInfoResponse
		err    error
	)
	if err = core.DB.Builder().
		Select().
		From("sys_user").
		Where(gorm.Eq("user_id", userId)).
		QExecute(&user).
		Error; err != nil {
		core.Log.Error("获取用户信息失败：%s", err.Error())
		return nil, response.NewBusinessError(response.DataNotExist)
	}
	if err = copier.Copy(&result, user); err != nil {
		core.Log.Error("解析用户信息失败：%s", err.Error())
		return nil, response.NewBusinessError(response.DataNotExist)
	}

	return &result, nil

}

func (u *UserService) Page(param *request.PageParam) (*response.PageData, *response.BusinessError) {
	var (
		list  []response.UserPageResponse
		total int64
		err   error
	)
	fmt.Printf("Query parameters: Status=%d, DeptId=%d, Ancestors=%v\n", param.Status, param.DeptId, param.Name)

	if err = core.DB.Namespace("user").
		DQuery("selectUserPage", param).
		DOffset(param.Page, param.Size).
		TPage(&list, &total).
		Error; err != nil {
		core.Log.Error("查询用户信息失败，异常信息如下：%s", err.Error())
		return nil, response.CustomBusinessError(response.Failed, "获取用户数据失败")
	}
	for i := 0; i < len(list); i++ {
		if list[i].UserId == vo.SUPER_USER_ID {
			list[i].IsSuper = true
		}
	}
	return &response.PageData{
		Total: total,
		Page:  param.Page,
		Size:  param.Size,
		Data:  list,
	}, nil

}
