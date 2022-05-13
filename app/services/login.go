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

package services

import (
	"easygoadmin/app/models"
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/gstr"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	"github.com/beego/beego/v2/server/web/context"
	"time"
)

var Login = new(loginService)

type loginService struct{}

// 系统登录
func (s *loginService) UserLogin(username, password string, ctx *context.Context) (string, error) {
	// 查询用户
	var user models.User
	err := orm.NewOrm().QueryTable(new(models.User)).Filter("username", username).Filter("mark", 1).One(&user)
	if err != nil {
		return "", errors.New("用户名或者密码不正确")
	}
	// 密码校验
	pwd, _ := utils.Md5(password + user.Username)
	if user.Password != pwd {
		return "", errors.New("密码不正确")
	}
	// 判断当前用户状态
	if user.Status != 1 {
		return "", errors.New("您的账号已被禁用,请联系管理员")
	}

	// 更新登录时间、登录IP
	o := orm.NewOrm()
	entity := models.User{Id: user.Id}
	entity.LoginTime = time.Now()
	entity.LoginIp = ctx.Input.IP()
	entity.UpdateTime = time.Now()
	o.Update(&entity, "LoginTime", "LoginIp", "UpdateTime")

	// 生成Token
	token, _ := utils.GenerateToken(user.Id, user.Username, user.Password)
	fmt.Println("生成的token:", token)

	// 返回token
	return token, nil
}

// 获取个人信息
func (s *loginService) GetProfile(userId int) *models.User {
	user := &models.User{Id: userId}
	err := user.Get()
	if err != nil {
		return nil
	}
	// 头像
	if user.Avatar != "" && !gstr.Contains(user.Avatar, conf.CONFIG.EGAdmin.Image) {
		user.Avatar = utils.GetImageUrl(user.Avatar)
	}
	return user
}
