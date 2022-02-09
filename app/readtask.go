// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"errors"
	"github.com/gen2brain/dlgs"
	"github.com/ying32/govcl/vcl"
	"log"
	"path/filepath"
)

//ReadTaskFormDir 文件
func (AppM *App) ReadTaskFormDir(t int) {
	tfn, tname := GetTypeInfo(t)
	b, s := vcl.SelectDirectory2("选择"+tname+"文件夹", ".", false)
	if !b {
		return
	}
	if s == "" {
		log.Println(errors.New("文件夹读取错误").Error())
		_, err := dlgs.Error("错误", "无法读取文件夹."+errors.New("文件夹dlgs返回错误").Error())
		if err != nil {
			log.Println(errors.New("dlgs错误").Error(), err.Error())
			return
		}
		return
	}
	var Multi []string
	for _, t := range tfn {
		tm, err := GetFileFromDir(s, t)
		if err != nil {
			_, err := dlgs.Error("错误", "无法读取文件夹."+errors.New("文件夹内容读取错误").Error())
			if err != nil {
				log.Println(errors.New("dlgs错误").Error(), err.Error())
				return
			}
			return
		}
		Multi = append(Multi, tm...)
	}

	AppM.AddTasks(NewTasks(Multi, t)...)
}

// ReadTaskFormFiles ss
func (AppM *App) ReadTaskFormFiles(t int) {
	var Multi []string
	tfn, tname := GetTypeInfo(t)
	tmulti, b, err := dlgs.FileMulti("选择"+tname, "")
	if !b {
		return
	}
	if err != nil {
		log.Println(err)
		_, err := dlgs.Error("错误", "无法读取文件."+err.Error())
		if err != nil {
			log.Println(errors.New("dlgs错误").Error(), err.Error())
			return
		}
		return
	}
	for _, s := range tmulti {
		for _, s2 := range tfn {
			if filepath.Ext(s) == s2 {
				Multi = append(Multi, s)
			}
		}
	}
	AppM.AddTasks(NewTasks(Multi, t)...)
}

func GetTypeInfo(t int) (tfn []string, tname string) {
	switch t {
	case TASK_TYPE_PIC:
		tname = "图片"
		tfn = []string{".png", ".jpg"}
	case TASK_TYPE_VID:
		tname = "视频"
		tfn = []string{".mp4"}
	}
	return
}
