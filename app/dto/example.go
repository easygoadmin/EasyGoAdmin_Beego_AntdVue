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

/**
 * 演示一Dto
 * @author 半城风雨
 * @since 2022-05-13
 * @File : example
 */
package dto

import "github.com/gookit/validate"

// 分页查询
type ExamplePageReq struct {
	Page  int    `form:"page"`  // 页码
	Limit int    `form:"limit"` // 每页数

	
	Name   string `form:"name"`   // 测试名称
	

	
	Status   int    `form:"status"`   // 状态：1正常 2停用
	

	
	Type   int    `form:"type"`   // 类型：1京东 2淘宝 3拼多多 4唯品会
	

	
	IsVip   int    `form:"isVip"`   // 是否VIP：1是 2否
	

}

// 添加演示一
type ExampleAddReq struct {

	
	Name  string `form:"name" validate:"required"`   // 测试名称
	

	
	Avatar  string `form:"avatar" validate:"required"`    // 头像
	

	
	Content  string `form:"content" validate:"required"`   // 内容
	

	
	Status  int    `form:"status" validate:"int"`    // 状态：1正常 2停用
	

	
	Type  int    `form:"type" validate:"int"`    // 类型：1京东 2淘宝 3拼多多 4唯品会
	

	
	IsVip  int    `form:"isVip" validate:"int"`    // 是否VIP：1是 2否
	

	
	Sort  int    `form:"sort" validate:"int"`    // 排序号
	

}

// 添加表单验证
func (v ExampleAddReq) Messages() map[string]string {
	return validate.MS{
	
		
		"Name.required": "测试名称不能为空.", // 测试名称
		
	
		
		"Avatar.required": "头像不能为空.", // 头像
		
	
		
		"Content.required": "内容不能为空.", // 内容
		
	
		
		"Status.int":    "请选择状态.", // 状态：1正常 2停用
		
	
		
		"Type.int":    "请选择类型.", // 类型：1京东 2淘宝 3拼多多 4唯品会
		
	
		
		"IsVip.int":    "请选择是否VIP.", // 是否VIP：1是 2否
		
	
		
		"Sort.int":      "排序号不能为空.", // 排序号
		
	
	}
}

// 编辑演示一
type ExampleUpdateReq struct {
	Id     int    `form:"id" validate:"int"`

	
	Name  string `form:"name" validate:"required"`   // 测试名称
	

	
	Avatar  string `form:"avatar" validate:"required"`    // 头像
	

	
	Content  string `form:"content" validate:"required"`   // 内容
	

	
	Status  int    `form:"status" validate:"int"`    // 状态：1正常 2停用
	

	
	Type  int    `form:"type" validate:"int"`    // 类型：1京东 2淘宝 3拼多多 4唯品会
	

	
	IsVip  int    `form:"isVip" validate:"int"`    // 是否VIP：1是 2否
	

	
	Sort  int    `form:"sort" validate:"int"`    // 排序号
	

}

// 更新表单验证
func (v ExampleUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "记录ID不能为空.",
	
		
		"Name.required": "测试名称不能为空.", // 测试名称
		
	
		
		"Avatar.required": "头像不能为空.", // 头像
		
	
		
		"Content.required": "内容不能为空.", // 内容
		
	
		
		"Status.int":    "请选择状态.", // 状态：1正常 2停用
		
	
		
		"Type.int":    "请选择类型.", // 类型：1京东 2淘宝 3拼多多 4唯品会
		
	
		
		"IsVip.int":    "请选择是否VIP.", // 是否VIP：1是 2否
		
	
		
		"Sort.int":      "排序号不能为空.", // 排序号
		
	
	}
}









// 设置状态
type ExampleStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

// 设置状态参数验证
func (v ExampleStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "记录ID不能为空.",
		"Status.int": "请选择状态：1正常 2停用.",
	}
}





// 设置是否VIP
type ExampleIsVipReq struct {
	Id     int `form:"id" validate:"int"`
	IsVip int `form:"isVip" validate:"int"`
}

// 设置状态参数验证
func (v ExampleIsVipReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "记录ID不能为空.",
		"IsVip.int": "请选择是否VIP：1是 2否.",
	}
}



