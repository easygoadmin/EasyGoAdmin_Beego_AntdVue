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

import (
	"github.com/gookit/validate"
)

// 用户分页查询条件
type UserPageReq struct {
	Realname string `form:"realname"` // 用户名
	Gender   int    `form:"gender"`   // 性别
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加用户
type UserAddReq struct {
	Realname   string `form:"realname" validate:"required"`
	Nickname   string `form:"nickname" validate:"required"`
	Gender     int    `form:"gender" validate:"int"`
	Avatar     string `form:"avatar" validate:"required"`
	Mobile     string `form:"mobile" validate:"required"`
	Email      string `form:"email" validate:"required"`
	Birthday   string `form:"birthday" validate:"required"`
	DeptId     int    `form:"deptId" validate:"int"`
	LevelId    int    `form:"levelId" validate:"int"`
	PositionId int    `form:"positionId" validate:"int"`
	City       string `form:"city" validate:"required"` // 省市区
	Address    string `form:"address"`
	Username   string `form:"username" validate:"required"`
	Password   string `form:"password"`
	Intro      string `form:"intro"`
	Status     int    `form:"status" validate:"int"`
	Note       string `form:"note"`
	Sort       int    `form:"sort" validate:"int"`
	RoleIds    string `form:"roleIds"` // 用户角色
}

// 添加用户表单验证
func (v UserAddReq) Messages() map[string]string {
	return validate.MS{
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.int":        "请选择用户性别.",
		"Avatar.required":   "请上传头像.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Birthday.required": "请选择出生日期.",
		"City.required":     "请选择省市区.",
		"DeptId.int":        "请选择所属部门.",
		"LevelId.int":       "请选择职级.",
		"PositionId.int":    "请选择用户.",
		"Username.required": "用户名不能为空.",
		"Status.int":        "请选择用户状态.",
		"Sort.int":          "排序不能为空.",
	}
}

// 添加用户
type UserUpdateReq struct {
	Id         int    `form:"id" validate:"int"`
	Realname   string `form:"realname" validate:"required"`
	Nickname   string `form:"nickname" validate:"required"`
	Gender     int    `form:"gender" validate:"int"`
	Avatar     string `form:"avatar" validate:"required"`
	Mobile     string `form:"mobile" validate:"required"`
	Email      string `form:"email" validate:"required"`
	Birthday   string `form:"birthday" validate:"required"`
	DeptId     int    `form:"deptId" validate:"int"`
	LevelId    int    `form:"levelId" validate:"int"`
	PositionId int    `form:"positionId" validate:"int"`
	City       string `form:"city" validate:"required"` // 省市区
	Address    string `form:"address"`
	Username   string `form:"username" validate:"required"`
	Password   string `form:"password"`
	Intro      string `form:"intro"`
	Status     int    `form:"status" validate:"int"`
	Note       string `form:"note"`
	Sort       int    `form:"sort" validate:"int"`
	RoleIds    string `form:"roleIds"` // 用户角色
}

// 更新用户表单验证
func (v UserUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":            "用户ID不能为空.",
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.int":        "请选择用户性别.",
		"Avatar.required":   "请上传头像.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Birthday.required": "请选择出生日期.",
		"City.required":     "请选择省市区.",
		"DeptId.int":        "请选择所属部门.",
		"LevelId.int":       "请选择职级.",
		"PositionId.int":    "请选择用户.",
		"Username.required": "用户名不能为空.",
		"Status.int":        "请选择用户状态.",
		"Sort.int":          "排序不能为空.",
	}
}

// 设置状态
type UserStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status"    validate:"int"`
}

func (v UserStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "用户ID不能为空.",
		"Status.int": "请选择用户状态.",
	}
}

// 用户中心
type UserInfoReq struct {
	Avatar   string `form:"avatar"`                       // 头像
	Realname string `form:"realname" validate:"required"` // 真实姓名
	Nickname string `form:"nickname" validate:"required"` // 昵称
	Gender   int    `form:"gender" validate:"int"`        // 性别:1男 2女 3保密
	Mobile   string `form:"mobile" validate:"required"`   // 手机号码
	Email    string `form:"email" validate:"required"`    // 邮箱地址
	Address  string `form:"address"`                      // 详细地址
	Intro    string `form:"intro"`                        // 个人简介
}

// 更新个人中心表单验证
func (v UserInfoReq) Messages() map[string]string {
	return validate.MS{
		"Realname.required": "用户名称不能为空.",
		"Nickname.required": "用户昵称不能为空.",
		"Gender.int":        "请选择用户性别.",
		"Mobile.required":   "手机号码不能为空.",
		"Email.required":    "电子邮件不能为空.",
		"Address.required":  "用户名不能为空.",
	}
}

// 更新密码
type UpdatePwd struct {
	OldPassword string `form:"oldPassword" validate:"required"` // 旧密码
	NewPassword string `form:"newPassword" validate:"required"` // 新密码
	RePassword  string `form:"rePassword" validate:"required"`  // 确认密码
}

// 更新密码表单验证
func (v UpdatePwd) Messages() map[string]string {
	return validate.MS{
		"OldPassword.required": "旧密码不能为空.",
		"NewPassword.required": "新密码不能为空.",
		"RePassword.required":  "确认密码不能为空.",
	}
}

// 重置密码
type UserResetPwdReq struct {
	Id int `form:"id" validate:"int"`
}

// 更新密码表单验证
func (v UserResetPwdReq) Messages() map[string]string {
	return validate.MS{
		"Id.int": "用户ID不能为空.",
	}
}
