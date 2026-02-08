package upload

import (
	"errors"
	"github.com/duke-git/lancet/v2/cryptor"
	"github.com/spf13/cast"
	image2 "image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"sync"
	"time"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/flipped-aurora/gin-vue-admin/server/utils"
	"go.uber.org/zap"
)

var mu sync.Mutex

type Local struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: UploadFile
//@description: 上传文件
//@param: file *multipart.FileHeader
//@return: string, string, error

func (*Local) UploadFile(file *multipart.FileHeader, req UploadFileExtReq) (resp OssUploadFileResp, err error) {
	// 读取文件后缀
	ext := filepath.Ext(file.Filename)
	// 读取文件名并加密
	name := strings.TrimSuffix(file.Filename, ext)
	name = utils.MD5V([]byte(name))
	// 拼接新文件名
	filename := name + "_" + time.Now().Format("20060102150405") + ext

	// 拼接路径和文件名
	if req.StoreDir != "" {
		if !strings.HasSuffix(req.StoreDir, "/") {
			req.StoreDir += "/"
		}
	}

	// 尝试创建此路径
	storePath := global.GVA_CONFIG.Local.StorePath + "/" + req.StoreDir
	mkdirErr := os.MkdirAll(storePath, os.ModePerm)
	if mkdirErr != nil {
		global.GVA_LOG.Error("function os.MkdirAll() failed", zap.Any("err", mkdirErr.Error()))
		err = errors.New("function os.MkdirAll() failed, err:" + mkdirErr.Error())
		return
	}

	p := storePath + filename
	tmpFilepath := global.GVA_CONFIG.Local.Path + "/" + req.StoreDir + filename

	f, openError := file.Open() // 读取文件
	if openError != nil {
		global.GVA_LOG.Error("function file.Open() failed", zap.Any("err", openError.Error()))
		err = errors.New("function file.Open() failed, err:" + openError.Error())
		return
	}
	defer f.Close() // 创建文件 defer 关闭

	out, createErr := os.Create(p)
	if createErr != nil {
		global.GVA_LOG.Error("function os.Create() failed", zap.Any("err", createErr.Error()))
		err = errors.New("function os.Create() failed, err:" + createErr.Error())
		return
	}
	defer out.Close() // 创建文件 defer 关闭

	_, copyErr := io.Copy(out, f) // 传输（拷贝）文件
	if copyErr != nil {
		global.GVA_LOG.Error("function io.Copy() failed", zap.Any("err", copyErr.Error()))
		err = errors.New("function io.Copy() failed, err:" + copyErr.Error())
		return
	}
	fileHash, fileHashErr := cryptor.Md5File(p)
	if fileHashErr != nil {
		global.GVA_LOG.Error("cal file hash err", zap.Error(fileHashErr))
		err = errors.New("cal file hash err, err:" + fileHashErr.Error())
		return
	}

	image := []string{".jpeg", ".jpg", ".png"}
	video := []string{".mp4", ".wmv"}

	if slices.Contains(image, ext) {
		myFile, fileErr := os.Open(p)
		if fileErr != nil {
			err = fileErr
			global.GVA_LOG.Error("打开文件失败!", zap.Error(fileErr))
			return
		}
		defer myFile.Close()
		fileConfig, _, fileConfigErr := image2.DecodeConfig(myFile)
		if fileConfigErr != nil {
			err = fileConfigErr
			global.GVA_LOG.Error("文件解码失败!", zap.Error(fileConfigErr))
			return
		}
		resp.Width = fileConfig.Width
		resp.Height = fileConfig.Height
	} else if slices.Contains(video, ext) {
		fmpegLib := utils.FfmpegUtil{Url: p}
		videoInfo, videoErr := fmpegLib.GetVideoMetaInfo()
		if videoErr != nil {
			err = videoErr
			global.GVA_LOG.Error("获取视频信息失败!", zap.Error(videoErr))
			return
		}
		resp.Width = videoInfo.Width
		resp.Height = videoInfo.Height
		resp.Duration = cast.ToInt(videoInfo.Duration * 1000)
		resp.Bitrate = videoInfo.Bitrate
		resp.Fps = videoInfo.Fps
	}

	resp.Filepath = tmpFilepath
	resp.Filename = filename
	resp.Hash = fileHash
	resp.Size = file.Size

	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@author: [ccfish86](https://github.com/ccfish86)
//@author: [SliverHorn](https://github.com/SliverHorn)
//@object: *Local
//@function: DeleteFile
//@description: 删除文件
//@param: key string
//@return: error

func (*Local) DeleteFile(key string) error {
	// 检查 key 是否为空
	if key == "" {
		return errors.New("key不能为空")
	}

	// 验证 key 是否包含非法字符或尝试访问存储路径之外的文件
	if strings.Contains(key, "..") || strings.ContainsAny(key, `\/:*?"<>|`) {
		return errors.New("非法的key")
	}

	p := filepath.Join(global.GVA_CONFIG.Local.StorePath, key)

	// 检查文件是否存在
	if _, err := os.Stat(p); os.IsNotExist(err) {
		return errors.New("文件不存在")
	}

	// 使用文件锁防止并发删除
	mu.Lock()
	defer mu.Unlock()

	err := os.Remove(p)
	if err != nil {
		return errors.New("文件删除失败: " + err.Error())
	}

	return nil
}
