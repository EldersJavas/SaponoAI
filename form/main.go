// Created by EldersJavas(EldersJavas&gmail.com)

package form

import (
	"fmt"
	"github.com/EldersJavas/SaponoAI/app"
	"github.com/ying32/govcl/vcl"
	"log"
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
