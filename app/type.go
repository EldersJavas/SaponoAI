// Created by EldersJavas(EldersJavas&gmail.com)

package app

import "github.com/EldersJavas/SaponoAI/app/tool"

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

func NewTask(path string, tType int) *Task {
	n := tool.GetFileBaseName(path)
	return &Task{Name: n, Path: path, TaskType: tType}
}

func (a *App) AddTask(task Task) {
	a.Tasklist = append(a.Tasklist, task)
}
