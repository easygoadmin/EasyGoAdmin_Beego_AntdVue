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
 *
 * @author 半城风雨
 * @since 2021/11/19
 * @File : gregex
 */
package gregex

import (
	"regexp"
	"sync"
)

var (
	regexMu = sync.RWMutex{}
	// Cache for regex object.
	// Note that:
	// 1. It uses sync.RWMutex ensuring the concurrent safety.
	// 2. There's no expiring logic for this map.
	regexMap = make(map[string]*regexp.Regexp)
)

func MatchString(pattern string, src string) ([]string, error) {
	if r, err := getRegexp(pattern); err == nil {
		return r.FindStringSubmatch(src), nil
	} else {
		return nil, err
	}
}

func getRegexp(pattern string) (regex *regexp.Regexp, err error) {
	// Retrieve the regular expression object using reading lock.
	regexMu.RLock()
	regex = regexMap[pattern]
	regexMu.RUnlock()
	if regex != nil {
		return
	}
	// If it does not exist in the cache,
	// it compiles the pattern and creates one.
	regex, err = regexp.Compile(pattern)
	if err != nil {
		return
	}
	// Cache the result object using writing lock.
	regexMu.Lock()
	regexMap[pattern] = regex
	regexMu.Unlock()
	return
}
