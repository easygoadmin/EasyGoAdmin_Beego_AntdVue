// +----------------------------------------------------------------------
// | EasyGoAdmin敏捷开发框架 [ 赋能开发者，助力企业发展 ]
// +----------------------------------------------------------------------
// | 版权所有 2019~2022 深圳EasyGoAdmin研发中心
// +----------------------------------------------------------------------
// | Licensed LGPL-3.0 EasyGoAdmin并不是自由软件，未经许可禁止去掉相关版权
// +----------------------------------------------------------------------
// | 官方网站: http://www.easygoadmin.vip
// +----------------------------------------------------------------------
// | Author: @半城风雨 团队荣誉出品
// +----------------------------------------------------------------------
// | 版权和免责声明:
// | 本团队对该软件框架产品拥有知识产权（包括但不限于商标权、专利权、著作权、商业秘密等）
// | 均受到相关法律法规的保护，任何个人、组织和单位不得在未经本团队书面授权的情况下对所授权
// | 软件框架产品本身申请相关的知识产权，禁止用于任何违法、侵害他人合法权益等恶意的行为，禁
// | 止用于任何违反我国法律法规的一切项目研发，任何个人、组织和单位用于项目研发而产生的任何
// | 意外、疏忽、合约毁坏、诽谤、版权或知识产权侵犯及其造成的损失 (包括但不限于直接、间接、
// | 附带或衍生的损失等)，本团队不承担任何法律责任，本软件框架禁止任何单位和个人、组织用于
// | 任何违法、侵害他人合法利益等恶意的行为，如有发现违规、违法的犯罪行为，本团队将无条件配
// | 合公安机关调查取证同时保留一切以法律手段起诉的权利，本软件框架只能用于公司和个人内部的
// | 法律所允许的合法合规的软件产品研发，详细声明内容请阅读《框架免责声明》附件；
// +----------------------------------------------------------------------

package dto

import "github.com/gookit/validate"

// 列表查询条件
type MenuQueryReq struct {
	Title string `form:"title"` // 菜单标题
}

// 添加菜单
type MenuAddReq struct {
	Id          int    `form:"id"`
	ParentId    int    `form:"parentId" validate:"int"`   // 上级ID
	Title       string `form:"title" validate:"required"` // 菜单标题
	Icon        string `form:"icon" validate:"required"`  // 图标
	Path        string `form:"path" validate:"required"`  // URL地址
	Component   string `form:"component"`                 // 菜单组件
	Target      int    `form:"target" validate:"int"`     // 打开方式：0组件 1内链 2外链
	Permission  string `form:"permission"`                // 权限标识
	Type        int    `form:"type" validate:"int"`       // 类型：0菜单 1节点
	Status      int    `form:"status" validate:"int"`     // 是否显示：1显示 2不显示
	Hide        int    `form:"hide" validate:"int"`       // 是否可见：1是 2否
	Note        string `form:"note"`                      // 菜单备注
	Sort        int    `form:"sort" validate:"int"`       // 显示顺序
	CheckedList string `form:"checkedList"`               // 权限节点
}

// 添加菜单表单验证
func (v MenuAddReq) Messages() map[string]string {
	return validate.MS{
		"ParentId.int":   "请选择上级菜单.",
		"Title.required": "菜单标题不能为空.",
		"Icon.required":  "请选择菜单图标.",
		"Path.required":  "菜单URL地址不能为空.",
		"Target.int":     "请选择菜单打开模式.",
		"Type.int":       "请选择菜单类型.",
		"Status.int":     "请选择菜单状态.",
		"Hide.int":       "请选择是否隐藏.",
		"Sort.int":       "排序不能为空.",
	}
}

// 更新菜单
type MenuUpdateReq struct {
	Id          int    `form:"id" validate:"int"`
	ParentId    int    `form:"parentId" validate:"int"`   // 上级ID
	Title       string `form:"title" validate:"required"` // 菜单标题
	Icon        string `form:"icon" validate:"required"`  // 图标
	Path        string `form:"path" validate:"required"`  // URL地址
	Component   string `form:"component"`                 // 菜单组件
	Target      int    `form:"target" validate:"int"`     // 打开方式：0组件 1内链 2外链
	Permission  string `form:"permission"`                // 权限标识
	Type        int    `form:"type" validate:"int"`       // 类型：0菜单 1节点
	Status      int    `form:"status" validate:"int"`     // 是否显示：1显示 2不显示
	Hide        int    `form:"hide" validate:"int"`       // 是否可见：1是 2否
	Note        string `form:"note"`                      // 菜单备注
	Sort        int    `form:"sort" validate:"int"`       // 显示顺序
	CheckedList string `form:"checkedList"`               // 权限节点
}

// 更新菜单表单验证
func (v MenuUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":         "菜单ID不能为空.",
		"ParentId.int":   "请选择上级菜单.",
		"Title.required": "菜单标题不能为空.",
		"Icon.required":  "请选择菜单图标.",
		"Path.required":  "菜单URL地址不能为空.",
		"Target.int":     "请选择菜单打开模式.",
		"Type.int":       "请选择菜单类型.",
		"Status.int":     "请选择菜单状态.",
		"Hide.int":       "请选择是否隐藏.",
		"Sort.int":       "排序不能为空.",
	}
}
