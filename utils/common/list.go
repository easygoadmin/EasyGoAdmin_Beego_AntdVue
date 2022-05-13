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
 * 数组处理类
 * @author 半城风雨
 * @since 2021/11/16
 * @File : list
 */
package common

type ArrayList struct {
	elements []interface{}
	size     int
}

func New(values ...interface{}) *ArrayList {
	list := &ArrayList{}
	list.elements = make([]interface{}, 10)
	if len(values) > 0 {
		list.Add(values...)
	}
	return list
}

func (list *ArrayList) Add(values ...interface{}) {
	if list.size+len(values) >= len(list.elements)-1 {
		newElements := make([]interface{}, list.size+len(values)+1)
		copy(newElements, list.elements)
		list.elements = newElements
	}

	for _, value := range values {
		list.elements[list.size] = value
		list.size++
	}

}

func (list *ArrayList) Remove(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}

	curEle := list.elements[index]
	list.elements[index] = nil
	copy(list.elements[index:], list.elements[index+1:list.size])
	list.size--
	return curEle
}

func (list *ArrayList) Get(index int) interface{} {
	if index < 0 || index >= list.size {
		return nil
	}
	return list.elements[index]
}

func (list *ArrayList) IsEmpty() bool {
	return list.size == 0
}

func (list *ArrayList) Size() int {
	return list.size
}
func (list *ArrayList) Contains(value interface{}) bool {
	for _, curValue := range list.elements {
		if curValue == value {
			return true
		}
	}
	return false
}
