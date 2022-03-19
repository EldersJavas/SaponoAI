// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/SaponoAI/app/tool"
	uuid "github.com/satori/go.uuid"
)

type App struct {
	Tasklist []Task
}

func NewApp() *App {
	return &App{}
}

type Task struct {
	Inpath   string `json:"Inpath"`
	Outpath  string `json:"Outpath"`
	Scale    int    `json:"Scale"`
	Model    int    `json:"Model"`
	Uid      string `json:"Uid"`
	TaskType int    `json:"TaskType"`
	Name     string `json:"Name"`
}

func NewTask(p string, tType int) *Task {
	m := 0
	if tType == TASK_TYPE_PIC {
		m = Realesrgan_x4plus
	} else {
		m = RealESRGANv2_animevideo_xsx4
	}
	n := tool.GetFileBaseName(tool.LinuxDir(p))
	return &Task{
		Inpath:   tool.FileFormatPath(p),
		Outpath:  tool.FileFormatPath(tool.GetAppRootDir() + "\\output\\"),
		Scale:    4,
		Model:    m,
		Uid:      uuid.Must(uuid.NewV4(), nil).String(),
		TaskType: tType,
		Name:     n,
	}
}

func NewTasks(path []string, tType int) (tl []Task) {
	for _, p := range path {
		tl = append(tl, *NewTask(p, tType))
	}
	return
}

func (AppM *App) AddTasks(task ...Task) {
	AppM.Tasklist = append(AppM.Tasklist, task...)
}
