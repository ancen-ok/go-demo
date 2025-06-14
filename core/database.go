package core

import (
	"embed"
	"gitee.com/molonglove/goboot/gorm"
	"gitee.com/molonglove/goboot/gorm/driver/mysql"
	"gitee.com/molonglove/goboot/gorm/schema"
	entity2 "go-demo/app/models/entity"
	"go-demo/utils"
)

// GetTable 获取模式下的表名
const GetTable = "SELECT table_name FROM information_schema.tables WHERE table_schema = ?"

var DB *gorm.DB

//go:embed mapper
var tpl embed.FS

func InitDb() {

	var (
		db          *gorm.DB
		tablePrefix = "sys_"
		tables      []string
		err         error
		tableMap    = map[string]any{
			"sys_dept":      entity2.Dept{},     // 部门表
			"sys_user":      entity2.User{},     // 用户表
			"sys_post":      entity2.Post{},     // 岗位表
			"sys_role":      entity2.Role{},     // 角色表
			"sys_menu":      entity2.Menu{},     // 菜单表
			"sys_user_role": entity2.UserRole{}, // 用户角色表
			"sys_role_menu": entity2.RoleMenu{}, // 角色菜单表
			"sys_role_dept": entity2.RoleDept{}, // 角色部门表
			"sys_user_post": entity2.UserPost{}, // 用户岗位表
			"sys_operate":   entity2.Operate{},  // 操作日志记录
			"sys_dict_type": entity2.DictType{}, // 字典类型表
			"sys_dict_data": entity2.DictData{}, // 字典数据表
			"sys_setting":   entity2.Setting{},  // 参数配置表
			"sys_visit":     entity2.Visit{},    // 系统访问记录
		}
	)
	dns := mysql.Open(Config.Db.Link())
	if db, err = gorm.Open(dns, &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   tablePrefix,
			SingularTable: true,
		},
		//配置日志输出级别
		Logger2:  Log,
		MapperFs: tpl,
	}); err != nil {
		panic("数据库连接失败=>" + err.Error())
	}
	Log.Info("数据库连接成功")
	if err = db.Raw(GetTable, Config.Db.DbName).Scan(&tables).Error; err == nil {
		for key, value := range tableMap {
			if exist, _ := utils.In[string](key, tables); !exist {
				if err = db.AutoMigrate(&value); err != nil {
					Log.Error("生成数据表失败 => %s:%s", key, err.Error())
					panic("初始化数据表失败")
				}
				Log.Info("生成表[%s]完成", key)
			}
		}
	}
	DB = db
}
