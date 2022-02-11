// Created by EldersJavas(EldersJavas&gmail.com)

package form

import (
	"fmt"
	"github.com/EldersJavas/SaponoAI/app"
	"github.com/ying32/govcl/vcl"
	"log"
	"math/rand"
	"time"
)

var AppM app.App

func (f *TMainform) Logout(s ...interface{}) {
	now := time.Now()
	t := now.Format("2006-01-02 15:04:05")
	log.Println(s...)
	vcl.ThreadSync(func() {
		f.LogList.AddItem(fmt.Sprintf("%v | %v", t, s), nil)
	})
}

func(f *TMainform) Testlistaddimg()  {
	f.TaskList.Items().BeginUpdate()
	for i := 1; i <= 100; i++ {
		item := f.TaskList.Items().Add()
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i+rand.Int()))
		item.SubItems().Add(fmt.Sprintf("值：%d", i+rand.Int()))
	}
	f.TaskList.Items().EndUpdate()
}