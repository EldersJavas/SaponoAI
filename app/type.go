// Created by EldersJavas(EldersJavas&gmail.com)

package app

import (
	"github.com/EldersJavas/SaponoAI/app/tool"
)

type App struct {
	Tasklist []Task
}

func NewApp() *App {
	return &App{}
}

type Task struct {
	Name     string
	Path     string
	TaskType int
}

func NewTask(p string, tType int) *Task {
	n := tool.GetFileBaseName(tool.LinuxDir(p))
	return &Task{Name: n, Path: tool.FileFormatPath(p), TaskType: tType}
}

func NewTasks(path []string, tType int) (tl []Task) {
	for _, p := range path {
		n := tool.GetFileBaseName(tool.LinuxDir(p))
		tl = append(tl, Task{Name: n, Path: tool.FileFormatPath(p), TaskType: tType})
	}
	return
}

func (AppM *App) AddTasks(task ...Task) {
	AppM.Tasklist = append(AppM.Tasklist, task...)
}
