package upload

import (
	"mime/multipart"

	"github.com/flipped-aurora/gin-vue-admin/server/global"
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"github.com/pkg/errors"
)

var HuaWeiObs = new(Obs)

type Obs struct{}

func NewHuaWeiObsClient() (client *obs.ObsClient, err error) {
	return obs.New(global.GVA_CONFIG.HuaWeiObs.AccessKey, global.GVA_CONFIG.HuaWeiObs.SecretKey, global.GVA_CONFIG.HuaWeiObs.Endpoint)
}

func (o *Obs) UploadFile(file *multipart.FileHeader, req UploadFileExtReq) (resp OssUploadFileResp, err error) {
	// var open multipart.File
	open, openErr := file.Open()
	if openErr != nil {
		err = openErr
		return
	}
	defer open.Close()
	filename := file.Filename
	input := &obs.PutObjectInput{
		PutObjectBasicInput: obs.PutObjectBasicInput{
			ObjectOperationInput: obs.ObjectOperationInput{
				Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
				Key:    filename,
			},
			HttpHeader: obs.HttpHeader{
				ContentType: file.Header.Get("content-type"),
			},
		},
		Body: open,
	}

	client, clientErr := NewHuaWeiObsClient()
	if clientErr != nil {
		err = errors.Wrap(err, "获取华为对象存储对象失败!")
		return
	}

	_, putErr := client.PutObject(input)
	if putErr != nil {
		err = errors.Wrap(err, "文件上传失败!")
		return
	}
	resp.Filepath = global.GVA_CONFIG.HuaWeiObs.Path + "/" + filename
	resp.Filename = filename
	return
}

func (o *Obs) DeleteFile(key string) error {
	client, err := NewHuaWeiObsClient()
	if err != nil {
		return errors.Wrap(err, "获取华为对象存储对象失败!")
	}
	input := &obs.DeleteObjectInput{
		Bucket: global.GVA_CONFIG.HuaWeiObs.Bucket,
		Key:    key,
	}
	var output *obs.DeleteObjectOutput
	output, err = client.DeleteObject(input)
	if err != nil {
		return errors.Wrapf(err, "删除对象(%s)失败!, output: %v", key, output)
	}
	return nil
}
