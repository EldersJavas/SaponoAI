// Created by EldersJavas(EldersJavas&gmail.com)

package form

import (
	"github.com/EldersJavas/SaponoAI/app"
	"github.com/ying32/govcl/vcl"
)

func (f *TMainform) UpdateTasklist(a app.App) {
	list := a.Tasklist
	for _, task := range list {
		var tls *vcl.TListItems
		var ic vcl.IComponent
		ic.
		f.TaskList.AddItem()
		f.TaskList.LargeImages().AddImages()
	}
}
