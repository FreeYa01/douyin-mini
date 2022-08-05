package util

import (
	"bytes"
	"douyin-mini/global"
	"fmt"
	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"os"
)

func GetSnapshot(videoPath,coverPath string,frameNum int) (err error){
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select",ffmpeg.Args{fmt.Sprintf("gte(n,%d)",frameNum)}).
		// vframes 设置输出内容的帧数
		Output("pipe:",ffmpeg.KwArgs{"vframes":frameNum,"format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf,os.Stdout).
		Run()
	if err != nil {
		global.Lg.Info("生成缩略图失败")
		return
	}
	img,err := imaging.Decode(buf)
	if err != nil {
		global.Lg.Info("生成缩略图失败")
		return
	}
	err = imaging.Save(img,coverPath)
	if err != nil{
		global.Lg.Info("生成缩略图失败")
		return
	}
	return err
}

