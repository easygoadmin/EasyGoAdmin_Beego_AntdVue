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
	"easygoadmin/app/constant"
	"easygoadmin/app/dto"
	"easygoadmin/app/models"
	"easygoadmin/app/vo"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"errors"
	"github.com/beego/beego/v2/client/orm"
	"strings"
	"time"
)

var Member = new(memberService)

type memberService struct{}

func (s *memberService) GetList(req dto.MemberPageReq) ([]vo.MemberListVo, int64, error) {
	// 初始化查询实例
	query := orm.NewOrm().QueryTable(new(models.Member)).Filter("mark", 1)
	// 用户名查询
	if req.Username != "" {
		query = query.Filter("username__contains", req.Username)
	}
	// 用户性别查询
	if req.Gender > 0 {
		query = query.Filter("gender", req.Gender)
	}
	// 排序
	query = query.OrderBy("-id")
	// 查询总数
	count, _ := query.Count()
	// 分页设置
	offset := (req.Page - 1) * req.Limit
	query = query.Limit(req.Limit, offset)
	// 查询列表
	lists := make([]models.Member, 0)
	query.All(&lists)
	// 数据处理
	var result = make([]vo.MemberListVo, 0)
	for _, v := range lists {
		item := vo.MemberListVo{}
		item.Member = v
		// 头像
		if v.Avatar != "" {
			item.Avatar = utils.GetImageUrl(v.Avatar)
		}
		// 性别
		if v.Gender > 0 {
			item.GenderName = constant.GENDER_LIST[v.Gender]
		}
		// 设备类型
		if v.Device > 0 {
			item.DeviceName = constant.MEMBER_DEVICE_LIST[v.Device]
		}
		// 会员来源
		if v.Source > 0 {
			item.SourceName = constant.MEMBER_SOURCE_LIST[v.Source]
		}
		// 所属城市
		if v.DistrictCode != "" {
			item.CityName = City.GetCityName(v.DistrictCode, ">>")
		}
		// 省市区
		cityList := make([]string, 0)
		// 省份编号
		cityList = append(cityList, item.ProvinceCode)
		// 城市编号
		cityList = append(cityList, item.CityCode)
		// 县区编号
		cityList = append(cityList, item.DistrictCode)
		item.City = cityList
		// 加入数组
		result = append(result, item)
	}
	return result, count, nil
}

func (s *memberService) Add(req dto.MemberAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Member
	entity.Username = req.Username
	entity.MemberLevel = req.MemberLevel
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	// 日期处理
	tm2, _ := time.Parse("2006-01-02", req.Birthday)
	entity.Birthday = tm2
	entity.Address = req.Address
	entity.Intro = req.Intro
	entity.Signature = req.Signature
	entity.Device = req.Device
	entity.Source = req.Source
	entity.Status = req.Status
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "member")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}

	// 省市区处理
	cityArr := strings.Split(req.City, ",")
	if len(cityArr) == 3 {
		entity.ProvinceCode = cityArr[0]
		entity.CityCode = cityArr[1]
		entity.DistrictCode = cityArr[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 插入数据
	return entity.Insert()
}

func (s *memberService) Update(req dto.MemberUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Member{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	entity.Username = req.Username
	entity.MemberLevel = req.MemberLevel
	entity.Realname = req.Realname
	entity.Nickname = req.Nickname
	entity.Gender = req.Gender
	// 日期处理
	tm2, _ := time.Parse("2006-01-02", req.Birthday)
	entity.Birthday = tm2
	entity.Address = req.Address
	entity.Intro = req.Intro
	entity.Signature = req.Signature
	entity.Device = req.Device
	entity.Source = req.Source
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()

	// 头像处理
	if req.Avatar != "" {
		avatar, err := utils.SaveImage(req.Avatar, "member")
		if err != nil {
			return 0, err
		}
		entity.Avatar = avatar
	}

	// 省市区处理
	cityArr := strings.Split(req.City, ",")
	if len(cityArr) == 3 {
		entity.ProvinceCode = cityArr[0]
		entity.CityCode = cityArr[1]
		entity.DistrictCode = cityArr[2]
	}

	// 密码
	if req.Password != "" {
		password, _ := utils.Md5(req.Password + req.Username)
		entity.Password = password
	}

	// 调用更新方法
	return entity.Update()
}

func (s *memberService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Member{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Member{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *memberService) Status(req dto.MemberStatusReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录是否存在
	entity := &models.Member{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	entity.Status = req.Status
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	return entity.Update()
}
