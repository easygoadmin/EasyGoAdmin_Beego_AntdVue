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

// 分页查询条件
type LinkPageReq struct {
	Name     string `form:"name"`     // 友链名称
	Type     int    `form:"type"`     // 友链类型
	Platform int    `form:"platform"` // 投放平台
	Page     int    `form:"page"`     // 页码
	Limit    int    `form:"limit"`    // 每页数
}

// 添加友链
type LinkAddReq struct {
	Name     string `form:"name" validate:"required"` // 友链名称
	Type     int    `form:"type" validate:"int"`      // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                      // 友链地址
	ItemId   int    `form:"itemId"`                   // 站点ID
	CateId   int    `form:"cateId"`                   // 栏目ID
	Platform int    `form:"platform" validate:"int"`  // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form" validate:"int"`      // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                    // 友链图片
	Status   int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Sort     int    `form:"sort" validate:"int"`      // 显示顺序
	Note     string `form:"note"`                     // 备注
}

// 添加友链表单验证
func (v LinkAddReq) Messages() map[string]string {
	return validate.MS{
		"Name.required": "友链名称不能为空.",
		"Type.int":      "请选择友链类型.",
		"Platform.int":  "请选择友链平台.",
		"Form.int":      "请选择友链形式.",
		"Status.int":    "请选择友链状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 修改友链
type LinkUpdateReq struct {
	Id       int    `form:"id" validate:"int"`
	Name     string `form:"name" validate:"required"` // 友链名称
	Type     int    `form:"type" validate:"int"`      // 类型：1友情链接 2合作伙伴
	Url      string `form:"url"`                      // 友链地址
	ItemId   int    `form:"itemId"`                   // 站点ID
	CateId   int    `form:"cateId"`                   // 栏目ID
	Platform int    `form:"platform" validate:"int"`  // 平台：1PC站 2WAP站 3微信小程序 4APP应用
	Form     int    `form:"form" validate:"int"`      // 友链形式：1文字链接 2图片链接
	Image    string `form:"image"`                    // 友链图片
	Status   int    `form:"status" validate:"int"`    // 状态：1在用 2停用
	Sort     int    `form:"sort" validate:"int"`      // 显示顺序
	Note     string `form:"note"`                     // 备注
}

// 更新友链表单验证
func (v LinkUpdateReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":        "友链ID不能为空.",
		"Name.required": "友链名称不能为空.",
		"Type.int":      "请选择友链类型.",
		"Platform.int":  "请选择友链平台.",
		"Form.int":      "请选择友链形式.",
		"Status.int":    "请选择友链状态.",
		"Sort.int":      "排序不能为空.",
	}
}

// 设置状态
type LinkStatusReq struct {
	Id     int `form:"id" validate:"int"`
	Status int `form:"status" validate:"int"`
}

func (v LinkStatusReq) Messages() map[string]string {
	return validate.MS{
		"Id.int":     "友链ID不能为空.",
		"Status.int": "请选择友链状态.",
	}
}
