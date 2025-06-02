package entity

import "time"

// Menu 菜单表
type Menu struct {
	MenuId     int64     `gorm:"column:menu_id;primaryKey;not null;autoIncrement;comment:菜单ID" json:"menuId"` // 主键
	MenuName   string    `gorm:"column:menu_name;not null;size:50;comment:菜单名称" json:"menuName"`              // 菜单名称
	MenuCode   string    `gorm:"column:menu_code;not null;size:100;comment:菜单权限/菜单编码" json:"menuCode"`        // 菜单编码/权限值
	ParentId   int64     `gorm:"column:parent_id;default:0;comment:父菜单ID" json:"parentId"`                    // 上级菜单
	OrderNum   int       `gorm:"column:order_num;default:1;comment:显示顺序" json:"orderNum"`                     // 显示顺序
	MenuPath   string    `gorm:"column:menu_path;size:200;comment:路由地址" json:"menuPath"`                      // 菜单路径
	MenuType   int       `gorm:"column:menu_type;not null;comment:菜单类型(1目录 2菜单 3按钮)" json:"menuType"`         // 菜单类型(1目录 2菜单 3按钮)
	IsPublic   int       `gorm:"column:is_public;default:1;comment:菜单状态(1公开 2私有)" json:"isPublic"`            // 是否是公共的
	IsShow     int       `gorm:"column:is_show;default:1;comment:菜单状态(1显示 2隐藏)" json:"isShow"`                // 是否显示
	Remark     string    `gorm:"column:remark;comment:菜单备注" json:"remark"`                                    // 备注
	Icon       string    `gorm:"column:icon;comment:图标" json:"icon"`                                          // 图标
	CreateId   int64     `gorm:"column:create_id;comment:创建人ID" json:"createId"`                              // 创建人
	CreateTime time.Time `gorm:"column:create_time;comment:创建时间" json:"createTime"`                           // 创建时间
	UpdateId   int64     `gorm:"column:update_id;comment:修改人ID" json:"updateId"`                              // 更新人
	UpdateTime time.Time `gorm:"column:update_time;comment:修改时间" json:"updateTime"`                           // 更新时间
	CanView    int       `gorm:"column:status;default:1;comment:菜单状态(1正常 2停用)" json:"canView"`                // 删除
}
