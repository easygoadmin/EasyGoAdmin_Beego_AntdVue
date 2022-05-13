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
	"easygoadmin/conf"
	"easygoadmin/utils"
	"easygoadmin/utils/gconv"
	"easygoadmin/utils/gregex"
	"easygoadmin/utils/gstr"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/server/web/context"
	"github.com/google/uuid"
	"io"
	"os"
	"path"
	"strings"
	"time"
)

var Upload = new(uploadService)

type uploadService struct{}

// 上传得文件信息
type FileInfo struct {
	FileName string `json:"fileName"`
	FileSize int64  `json:"fileSize"`
	FileUrl  string `json:"fileUrl"`
	FileType string `json:"fileType"`
}

func (s *uploadService) UploadImage(ctx *context.Context) (FileInfo, error) {
	if utils.AppDebug() {
		return FileInfo{}, errors.New("演示环境，暂无权限操作")
	}
	// 获取文件(注意这个地方的file要和html模板中的name一致)
	file, h, err := ctx.Request.FormFile("file")
	if err != nil {
		return FileInfo{}, err
	}
	// 关闭
	defer file.Close()
	//获取文件名称
	fmt.Println(h.Filename)
	//文件大小
	fmt.Println(h.Size)
	//获取文件的后缀名
	fileExt := path.Ext(h.Filename)
	fmt.Println(fileExt)
	// 允许上传文件后缀
	allowExt := conf.CONFIG.Attachment.FileExt
	// 检查上传文件后缀
	if !checkFileExt(fileExt, allowExt) {
		return FileInfo{}, errors.New("上传文件格式不正确，文件后缀只允许为：" + allowExt + "的文件")
	}
	// 允许文件上传最大值（如：1M）
	allowSize := conf.CONFIG.Attachment.FileSize + "M"
	// 检查上传文件大小
	isvalid, err := checkFileSize(h.Size, allowSize)
	if err != nil {
		return FileInfo{}, err
	}
	if !isvalid {
		return FileInfo{}, errors.New("上传文件大小不得超过：" + allowSize)
	}

	// 创建附件目录
	uploadDir := conf.CONFIG.Attachment.FilePath
	_, err = os.Stat(uploadDir)
	if err != nil {
		err = os.MkdirAll(uploadDir, os.ModePerm)
	}

	// 文件存放相对路径
	savePath := uploadDir + "/temp/" + time.Now().Format("20060102")

	// 创建临时文件目录
	ok := utils.CreateDir(savePath)
	if !ok {
		return FileInfo{}, errors.New("存储路径创建失败")
	}
	// 使用UUID作为新的文件名
	fileName := uuid.New().String() + fileExt
	// 保存上传文件
	filePath := savePath + "/" + fileName //filepath.Join(savePath, "/", fileName)
	// 写入文件
	f, err := os.OpenFile(filePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0666)
	if err != nil {
		return FileInfo{}, err
	}
	defer f.Close()
	// 复制文件
	io.Copy(f, file)

	// 文件URL地址
	fileUrl := strings.ReplaceAll(filePath, uploadDir, "")

	// 返回结果
	result := FileInfo{
		FileName: h.Filename,
		FileSize: h.Size,
		FileUrl:  fileUrl,
	}
	return result, nil
}

// 检查文件格式是否合法
func checkFileExt(fileExt string, typeString string) bool {
	// 允许上传文件后缀
	exts := gstr.Split(typeString, ",")
	// 是否验证通过
	isValid := false
	for _, v := range exts {
		// 对比文件后缀
		if gstr.Equal(fileExt, "."+v) {
			isValid = true
			break
		}
	}
	return isValid
}

// 检查上传文件大小
func checkFileSize(fileSize int64, maxSize string) (bool, error) {
	// 匹配上传文件最大值
	match, err := gregex.MatchString(`^([0-9]+)(?i:([a-z]*))$`, maxSize)
	if err != nil {
		return false, err
	}
	if len(match) == 0 {
		err = errors.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB）")
		return false, err
	}
	var cfSize int64
	switch gstr.ToUpper(match[2]) {
	case "MB", "M":
		cfSize = gconv.Int64(match[1]) * 1024 * 1024
	case "KB", "K":
		cfSize = gconv.Int64(match[1]) * 1024
	case "":
		cfSize = gconv.Int64(match[1])
	}
	if cfSize == 0 {
		err = errors.New("上传文件大小未设置，请在后台配置，格式为（30M,30k,30MB），最大单位为MB")
		return false, err
	}
	return cfSize >= fileSize, nil
}
