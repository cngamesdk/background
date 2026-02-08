package utils

import (
	"fmt"
	"github.com/duke-git/lancet/v2/fileutil"
	"github.com/pkg/errors"
	"github.com/spf13/cast"
	"os/exec"
	"regexp"
	"strings"
)

type VideoMetaInfo struct {
	Duration float32
	Bitrate  int
	Width    int
	Height   int
	Fps      int
	Ext      string
}

type FfmpegUtil struct {
	Url string
}

// Scale 将资源压缩成指定的尺寸
func (receiver FfmpegUtil) Scale(width, height int, dst string) (err error) {
	url := receiver.Url
	if err = receiver.check(); err != nil {
		err = errors.Wrap(err, url)
		return
	}
	cmd := exec.Command("ffmpeg", "-i", url, "-vf", fmt.Sprintf("scale=%d:%d", width, height), dst)

	_, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		err = cmdErr
		return
	}
	return
}

func (receiver FfmpegUtil) check() (err error) {
	url := receiver.Url
	if exists := fileutil.IsExist(receiver.Url); !exists {
		err = errors.New("文件不存在" + url)
		return
	}
	return
}

func (receiver FfmpegUtil) ScreenshotByFps(fps float32, out string) (err error) {
	url := receiver.Url
	if err = receiver.check(); err != nil {
		err = errors.Wrap(err, url)
		return
	}

	cmd := exec.Command("ffmpeg", "-i", url, "-vf", fmt.Sprintf("fps=%f", fps), out)

	_, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		err = cmdErr
		return
	}
	return
}

func (receiver FfmpegUtil) GetVideoMetaInfo() (metaInfo VideoMetaInfo, err error) {
	url := receiver.Url
	if err = receiver.check(); err != nil {
		err = errors.Wrap(err, url)
		return
	}

	cmd := exec.Command("ffprobe", "-i", url)

	cmdOut, cmdErr := cmd.CombinedOutput()
	if cmdErr != nil {
		err = cmdErr
		return
	}
	cmdOutStr := string(cmdOut)
	regResult := regexp.MustCompile(`Duration: (\d+:\d+:\d+.\d+), start: (.+?), bitrate: (\d+) kb/s`) // Duration: 00:05:30.11, start: 0.000000, bitrate: 470 kb/s
	findResult := regResult.FindString(cmdOutStr)

	if findResult != "" {
		// 时长抽取
		durationReg := regexp.MustCompile(`Duration: (\d+:\d+:\d+.\d+)`)
		durationRegResult := durationReg.FindString(findResult)
		if durationRegResult != "" {
			durationStr := regexp.MustCompile(`\d+:\d+:\d+.\d+`).FindString(durationRegResult)
			if durationStr != "" {
				durationFormatArr := strings.Split(durationStr, ":")
				metaInfo.Duration += cast.ToFloat32(durationFormatArr[0]) * 3600
				metaInfo.Duration += cast.ToFloat32(durationFormatArr[1]) * 60
				metaInfo.Duration += cast.ToFloat32(durationFormatArr[2])
			}
		}

		//比特率抽取
		bitrateRegex := regexp.MustCompile(` bitrate: (\d+) kb/s`)
		bitrateRegexResult := bitrateRegex.FindString(findResult)
		bitrateStr := regexp.MustCompile(`\d+`).FindString(bitrateRegexResult)
		if bitrateStr != "" {
			metaInfo.Bitrate = cast.ToInt(bitrateStr)
		}
	}

	//宽高、fps抽取
	//480x272, 381 kb/s, 18 fps
	//1280x720 [SAR 1:1 DAR 16:9], 841 kb/s, 30 fps, 30 tbr, 90k tbn, 30 tbc (default)
	//0x31637661), yuv420p, 480x272, 381 kb/s, 18 fps
	reg2 := regexp.MustCompile(`\d+x\d+[,|\s](.*?) \d+ kb/s, \d+ fps`)
	reg2Result := reg2.FindString(cmdOutStr)
	if reg2Result != "" {
		widthHeightResult := regexp.MustCompile(`\d+x\d+`).FindString(reg2Result)
		if widthHeightResult != "" {
			widthHeightSplit := strings.Split(widthHeightResult, "x")
			metaInfo.Width = cast.ToInt(widthHeightSplit[0])
			metaInfo.Height = cast.ToInt(widthHeightSplit[1])
		}

		fpsResult := regexp.MustCompile(`\d+ fps`).FindString(reg2Result)
		if fpsResult != "" {
			metaInfo.Fps = cast.ToInt(strings.Trim(strings.ReplaceAll(fpsResult, "fps", ""), " "))
		}
	}
	return
}
