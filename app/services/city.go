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

var City = new(cityService)

type cityService struct{}

func (s *cityService) GetList(req dto.CityQueryReq) ([]vo.CityListVo, error) {
	// 创建查询实例
	query := orm.NewOrm().QueryTable(new(models.City)).Filter("mark", 1)
	// 上级城市ID
	query = query.Filter("pid", req.Pid)
	// 城市名称
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []models.City
	_, err := query.All(&list)

	// 数据解析
	var result = make([]vo.CityListVo, 0)
	// 遍历数据
	for _, v := range list {
		item := vo.CityListVo{}
		item.City = v
		if v.Level < 3 {
			item.HaveChild = true
		} else {
			item.HaveChild = false
		}
		result = append(result, item)
	}
	return result, err
}

func (s *cityService) Add(req dto.CityAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.City
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.Level = req.Level
	entity.Citycode = req.Citycode
	entity.PAdcode = req.PAdcode
	entity.Adcode = req.Adcode
	entity.Lng = req.Lng
	entity.Lat = req.Lat
	entity.Sort = req.Sort
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入记录
	return entity.Insert()
}

func (s *cityService) Update(req dto.CityUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.City{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, err
	}
	// 设置对象属性
	entity.Name = req.Name
	entity.Pid = req.Pid
	entity.Level = req.Level
	entity.Citycode = req.Citycode
	entity.PAdcode = req.PAdcode
	entity.Adcode = req.Adcode
	entity.Lng = req.Lng
	entity.Lat = req.Lat
	entity.Sort = req.Sort
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *cityService) GetChilds(cityCode string) ([]models.City, error) {
	var info models.City
	o := orm.NewOrm()
	err := o.QueryTable(new(models.City)).Filter("citycode", cityCode).Filter("mark", 1).One(&info)
	if err != nil {
		return nil, errors.New("城市不能存在")
	}
	list := make([]models.City, 0)
	o.QueryTable(new(models.City)).Filter("pid", info.Id).Filter("mark", 1).All(&list)
	if err != nil {
		return nil, err
	}
	return list, nil
}

func (s *cityService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.City{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.City{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

func (s *cityService) GetCityName(cityCode string, delimiter string) string {
	//info := &model.City{Citycode: cityCode}
	var info models.City
	err := orm.NewOrm().QueryTable(new(models.City)).Filter("citycode", cityCode).One(&info)
	if err != nil {
		return ""
	}
	// 城市ID
	cityId := info.Id
	// 声明数组
	list := make([]string, 0)
	for {
		if cityId <= 0 {
			// 退出
			break
		}
		// 业务处理
		info = models.City{}
		err2 := orm.NewOrm().QueryTable(new(models.City)).Filter("id", cityId).One(&info)
		if err2 != nil {
			break
		}
		// 上级栏目ID
		cityId = info.Pid
		// 加入数组
		list = append(list, info.Name)
	}
	// 结果数据处理
	if len(list) > 0 {
		// 数组反转
		utils.Reverse(&list)
		// 拼接字符串
		return strings.Join(list, delimiter)
	}
	return ""
}
