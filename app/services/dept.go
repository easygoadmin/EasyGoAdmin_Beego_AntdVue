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
	"reflect"
	"strings"
	"time"
)

var Dept = new(deptService)

type deptService struct{}

func (s *deptService) GetList(req dto.DeptPageReq) ([]models.Dept, error) {
	// 创建查询实例
	query := orm.NewOrm().QueryTable(new(models.Dept)).Filter("mark", 1)
	// 部门名称
	if req.Name != "" {
		query = query.Filter("name__contains", req.Name)
	}
	// 排序
	query = query.OrderBy("sort")
	// 查询数据
	var list []models.Dept
	_, err := query.All(&list)
	return list, err
}

func (s *deptService) Add(req dto.DeptAddReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 实例化对象
	var entity models.Dept
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.CreateUser = userId
	entity.CreateTime = time.Now()
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	entity.Mark = 1
	// 插入记录
	return entity.Insert()
}

func (s *deptService) Update(req dto.DeptUpdateReq, userId int) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	// 查询记录
	entity := &models.Dept{Id: req.Id}
	err := entity.Get()
	if err != nil {
		return 0, errors.New("记录不存在")
	}
	// 设置参数
	entity.Name = req.Name
	entity.Code = req.Code
	entity.Fullname = req.Fullname
	entity.Type = req.Type
	entity.Pid = req.Pid
	entity.Sort = req.Sort
	entity.Note = req.Note
	entity.UpdateUser = userId
	entity.UpdateTime = time.Now()
	// 更新记录
	return entity.Update()
}

func (s *deptService) Delete(ids string) (int64, error) {
	if utils.AppDebug() {
		return 0, errors.New("演示环境，暂无权限操作")
	}
	idsArr := strings.Split(ids, ",")
	if len(idsArr) == 1 {
		// 单个删除
		entity := &models.Dept{Id: gconv.Int(ids)}
		rows, err := entity.Delete()
		return rows, err
	} else {
		// 批量删除
		count := 0
		for _, v := range idsArr {
			entity := &models.Dept{Id: gconv.Int(v)}
			rows, err := entity.Delete()
			if err != nil || rows == 0 {
				continue
			}
			count++
		}
		return int64(count), nil
	}
}

// 获取子级菜单
func (s *deptService) GetDeptTreeList() ([]*vo.DeptTreeNode, error) {
	var deptNode vo.DeptTreeNode
	// 查询列表
	list := make([]models.Dept, 0)
	orm.NewOrm().QueryTable(new(models.Dept)).Filter("mark", 1).OrderBy("sort").All(&list)
	makeDeptTree(list, &deptNode)
	return deptNode.Children, nil
}

//递归生成分类列表
func makeDeptTree(cate []models.Dept, tn *vo.DeptTreeNode) {
	for _, c := range cate {
		if c.Pid == tn.Id {
			child := &vo.DeptTreeNode{}
			child.Dept = c
			tn.Children = append(tn.Children, child)
			makeDeptTree(cate, child)
		}
	}
}

// 数据源转换
func (s *deptService) MakeList(data []*vo.DeptTreeNode) map[int]string {
	deptList := make(map[int]string, 0)
	if reflect.ValueOf(data).Kind() == reflect.Slice {
		// 一级栏目
		for _, val := range data {
			deptList[val.Id] = val.Name

			// 二级栏目
			for _, v := range val.Children {
				deptList[v.Id] = "|--" + v.Name

				// 三级栏目
				for _, vt := range v.Children {
					deptList[vt.Id] = "|--|--" + vt.Name
				}
			}
		}
	}
	return deptList
}
