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

package routers

import (
	"easygoadmin/app/controllers"
	"easygoadmin/app/middleware"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/beego/beego/v2/server/web/filter/cors"
)

func init() {

	// 跨域解决方案
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		AllowMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	// 登录验证中间件
	middleware.CheckLogin()

	// 系统登录
	//beego.Router("/", &controllers.IndexController{}, "get:Index")
	beego.Router("/login", &controllers.LoginController{}, "post:Login")
	beego.Router("/captcha", &controllers.LoginController{}, "get:Captcha")
	beego.Router("/updateUserInfo", &controllers.IndexController{}, "put:UpdateUserInfo")
	beego.Router("/updatePwd", &controllers.IndexController{}, "put:UpdatePwd")
	beego.Router("/logout", &controllers.IndexController{}, "get:Logout")

	// 系统主页
	beego.Router("/index/menu", &controllers.IndexController{}, "get:Menu")
	beego.Router("/index/user", &controllers.IndexController{}, "get:User")

	// 普通图片上传
	beego.Router("/upload/uploadImage", &controllers.UploadController{}, "post:UploadImage")

	// 职级管理
	beego.Router("/level/list", &controllers.LevelController{}, "get:List")
	beego.Router("/level/detail/:id", &controllers.LevelController{}, "get:Detail")
	beego.Router("/level/add", &controllers.LevelController{}, "post:Add")
	beego.Router("/level/update", &controllers.LevelController{}, "put:Update")
	beego.Router("/level/delete/:id", &controllers.LevelController{}, "delete:Delete")
	beego.Router("/level/status", &controllers.LevelController{}, "put:Status")
	beego.Router("/level/getLevelList", &controllers.LevelController{}, "get:GetLevelList")

	// 岗位管理
	beego.Router("/position/list", &controllers.PositionController{}, "get:List")
	beego.Router("/position/detail/:id", &controllers.PositionController{}, "get:Detail")
	beego.Router("/position/add", &controllers.PositionController{}, "post:Add")
	beego.Router("/position/update", &controllers.PositionController{}, "put:Update")
	beego.Router("/position/delete/:id", &controllers.PositionController{}, "delete:Delete")
	beego.Router("/position/status", &controllers.PositionController{}, "put:Status")
	beego.Router("/position/getPositionList", &controllers.PositionController{}, "get:GetPositionList")

	// 角色管理
	beego.Router("/role/list", &controllers.RoleController{}, "get:List")
	beego.Router("/role/detail/:id", &controllers.RoleController{}, "get:Detail")
	beego.Router("/role/add", &controllers.RoleController{}, "post:Add")
	beego.Router("/role/update", &controllers.RoleController{}, "put:Update")
	beego.Router("/role/delete/:id", &controllers.RoleController{}, "delete:Delete")
	beego.Router("/role/status", &controllers.RoleController{}, "put:Status")
	beego.Router("/role/getRoleList", &controllers.RoleController{}, "get:GetRoleList")

	// 角色菜单
	beego.Router("/rolemenu/index/:roleId", &controllers.RoleMenuController{}, "get:Index")
	beego.Router("/rolemenu/save", &controllers.RoleMenuController{}, "post:Save")

	// 菜单管理
	beego.Router("/menu/list", &controllers.MenuController{}, "get:List")
	beego.Router("/menu/detail/:id", &controllers.MenuController{}, "get:Detail")
	beego.Router("/menu/add", &controllers.MenuController{}, "post:Add")
	beego.Router("/menu/update", &controllers.MenuController{}, "put:Update")
	beego.Router("/menu/delete/:id", &controllers.MenuController{}, "delete:Delete")

	// 部门管理
	beego.Router("/dept/list", &controllers.DeptController{}, "get:List")
	beego.Router("/dept/detail/:id", &controllers.DeptController{}, "get:Detail")
	beego.Router("/dept/add", &controllers.DeptController{}, "post:Add")
	beego.Router("/dept/update", &controllers.DeptController{}, "put:Update")
	beego.Router("/dept/delete/:id", &controllers.DeptController{}, "delete:Delete")
	beego.Router("/dept/getDeptList", &controllers.DeptController{}, "get:GetDeptList")

	// 用户管理
	beego.Router("/user/list", &controllers.UserController{}, "get:List")
	beego.Router("/user/detail/:id", &controllers.UserController{}, "get:Detail")
	beego.Router("/user/add", &controllers.UserController{}, "post:Add")
	beego.Router("/user/update", &controllers.UserController{}, "put:Update")
	beego.Router("/user/delete/:id", &controllers.UserController{}, "delete:Delete")
	beego.Router("/user/resetPwd", &controllers.UserController{}, "put:ResetPwd")

	// 城市管理
	beego.Router("/city/list", &controllers.CityController{}, "get:List")
	beego.Router("/city/detail/:id", &controllers.CityController{}, "get:Detail")
	beego.Router("/city/add", &controllers.CityController{}, "post:Add")
	beego.Router("/city/update", &controllers.CityController{}, "put:Update")
	beego.Router("/city/delete/:id", &controllers.CityController{}, "delete:Delete")
	beego.Router("/city/getChilds", &controllers.CityController{}, "post:GetChilds")

	// 会员等级管理
	beego.Router("/memberlevel/list", &controllers.MemberLevelController{}, "get:List")
	beego.Router("/memberlevel/detail/:id", &controllers.MemberLevelController{}, "get:Detail")
	beego.Router("/memberlevel/add", &controllers.MemberLevelController{}, "post:Add")
	beego.Router("/memberlevel/update", &controllers.MemberLevelController{}, "put:Update")
	beego.Router("/memberlevel/delete/:id", &controllers.MemberLevelController{}, "delete:Delete")
	beego.Router("/memberlevel/getMemberLevelList", &controllers.MemberLevelController{}, "get:GetMemberLevelList")

	// 会员管理
	beego.Router("/member/list", &controllers.MemberController{}, "get:List")
	beego.Router("/member/detail/:id", &controllers.MemberController{}, "get:Detail")
	beego.Router("/member/add", &controllers.MemberController{}, "post:Add")
	beego.Router("/member/update", &controllers.MemberController{}, "put:Update")
	beego.Router("/member/delete/:id", &controllers.MemberController{}, "delete:Delete")
	beego.Router("/member/status", &controllers.MemberController{}, "put:Status")

	// 友链管理
	beego.Router("/link/list", &controllers.LinkController{}, "get:List")
	beego.Router("/link/detail/:id", &controllers.LinkController{}, "get:Detail")
	beego.Router("/link/add", &controllers.LinkController{}, "post:Add")
	beego.Router("/link/update", &controllers.LinkController{}, "put:Update")
	beego.Router("/link/delete/:id", &controllers.LinkController{}, "delete:Delete")
	beego.Router("/link/status", &controllers.LinkController{}, "put:Status")

	// 字典管理
	beego.Router("/dict/list", &controllers.DictController{}, "get:List")
	beego.Router("/dict/detail/:id", &controllers.DictController{}, "get:Detail")
	beego.Router("/dict/add", &controllers.DictController{}, "post:Add")
	beego.Router("/dict/update", &controllers.DictController{}, "put:Update")
	beego.Router("/dict/delete/:id", &controllers.DictController{}, "delete:Delete")

	// 字典数据
	beego.Router("/dictdata/list", &controllers.DictDataController{}, "get:List")
	beego.Router("/dictdata/detail/:id", &controllers.DictDataController{}, "get:Detail")
	beego.Router("/dictdata/add", &controllers.DictDataController{}, "post:Add")
	beego.Router("/dictdata/update", &controllers.DictDataController{}, "put:Update")
	beego.Router("/dictdata/delete/:id", &controllers.DictDataController{}, "delete:Delete")

	// 配置管理
	beego.Router("/config/list", &controllers.ConfigController{}, "get:List")
	beego.Router("/config/detail/:id", &controllers.ConfigController{}, "get:Detail")
	beego.Router("/config/add", &controllers.ConfigController{}, "post:Add")
	beego.Router("/config/update", &controllers.ConfigController{}, "put:Update")
	beego.Router("/config/delete/:id", &controllers.ConfigController{}, "delete:Delete")

	// 配置数据
	beego.Router("/configdata/list", &controllers.ConfigDataController{}, "get:List")
	beego.Router("/configdata/add", &controllers.ConfigDataController{}, "post:Add")
	beego.Router("/configdata/detail/:id", &controllers.ConfigDataController{}, "get:Detail")
	beego.Router("/configdata/update", &controllers.ConfigDataController{}, "put:Update")
	beego.Router("/configdata/delete/:id", &controllers.ConfigDataController{}, "delete:Delete")
	beego.Router("/configdata/status", &controllers.ConfigDataController{}, "put:Status")

	// 网站配置
	beego.Router("/configweb/index", &controllers.ConfigWebController{}, "get:Index")
	beego.Router("/configweb/save", &controllers.ConfigWebController{}, "put:Save")

	// 通知公告管理
	beego.Router("/notice/list", &controllers.NoticeController{}, "get:List")
	beego.Router("/notice/detail/:id", &controllers.NoticeController{}, "get:Detail")
	beego.Router("/notice/add", &controllers.NoticeController{}, "post:Add")
	beego.Router("/notice/update", &controllers.NoticeController{}, "put:Update")
	beego.Router("/notice/delete/:id", &controllers.NoticeController{}, "delete:Delete")
	beego.Router("/notice/status", &controllers.NoticeController{}, "put:Status")

	// 站点管理
	beego.Router("/item/list", &controllers.ItemController{}, "get:List")
	beego.Router("/item/detail/:id", &controllers.ItemController{}, "get:Detail")
	beego.Router("/item/add", &controllers.ItemController{}, "post:Add")
	beego.Router("/item/update", &controllers.ItemController{}, "put:Update")
	beego.Router("/item/delete/:id", &controllers.ItemController{}, "delete:Delete")
	beego.Router("/item/status", &controllers.ItemController{}, "put:Status")
	beego.Router("/item/getItemList", &controllers.ItemController{}, "get:GetItemList")

	// 栏目管理
	beego.Router("/itemcate/list", &controllers.ItemCateController{}, "get:List")
	beego.Router("/itemcate/detail/:id", &controllers.ItemCateController{}, "get:Detail")
	beego.Router("/itemcate/add", &controllers.ItemCateController{}, "post:Add")
	beego.Router("/itemcate/update", &controllers.ItemCateController{}, "put:Update")
	beego.Router("/itemcate/delete/:id", &controllers.ItemCateController{}, "delete:Delete")
	beego.Router("/itemcate/getCateList", &controllers.ItemCateController{}, "get:GetCateList")

	// 广告位管理
	beego.Router("/adsort/list", &controllers.AdSortController{}, "get:List")
	beego.Router("/adsort/detail/:id", &controllers.AdSortController{}, "get:Detail")
	beego.Router("/adsort/add", &controllers.AdSortController{}, "post:Add")
	beego.Router("/adsort/update", &controllers.AdSortController{}, "put:Update")
	beego.Router("/adsort/delete/:id", &controllers.AdSortController{}, "delete:Delete")
	beego.Router("/adsort/getAdSortList", &controllers.AdSortController{}, "get:GetAdSortList")

	// 广告管理
	beego.Router("/ad/list", &controllers.AdController{}, "get:List")
	beego.Router("/ad/detail/:id", &controllers.AdController{}, "get:Detail")
	beego.Router("/ad/add", &controllers.AdController{}, "post:Add")
	beego.Router("/ad/update", &controllers.AdController{}, "put:Update")
	beego.Router("/ad/delete/:id", &controllers.AdController{}, "delete:Delete")
	beego.Router("/ad/status", &controllers.AdController{}, "put:Status")

	// 代码生成器
	beego.Router("/generate/list", &controllers.GenerateController{}, "get:List")
	beego.Router("/generate/generate", &controllers.GenerateController{}, "post:Generate")
	beego.Router("/generate/batchGenerate", &controllers.GenerateController{}, "post:BatchGenerate")
}
