package form

import (
	_ "embed"
	"github.com/EldersJavas/SaponoAI/app"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"log"
)

//::private::
type TMainformFields struct {
}

func (f *TMainform) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	*canClose = vcl.MessageDlg("是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes
	if *canClose {
		vcl.Application.Terminate()
	}
}

func (f *TMainform) OnMaddpicClick(sender vcl.IObject) {

}

func (f *TMainform) OnMaddvdClick(sender vcl.IObject) {

}

func (f *TMainform) OnBackgroundClick(sender vcl.IObject) {

}

func (f *TMainform) OnMenuItem1Click(sender vcl.IObject) {

}

func (f *TMainform) OnMPFClick(sender vcl.IObject) {
	AppM.ReadTaskFormFiles(app.TASK_TYPE_PIC)
	log.Println(AppM.Tasklist)
}

func (f *TMainform) OnMPDClick(sender vcl.IObject) {
	AppM.ReadTaskFormDir(app.TASK_TYPE_PIC)
	log.Println(AppM.Tasklist)
}

func (f *TMainform) OnMPUClick(sender vcl.IObject) {
	AppM.ReadTaskFormFiles(app.TASK_TYPE_VID)
	log.Println(AppM.Tasklist)
}

func (f *TMainform) OnMVFClick(sender vcl.IObject) {
	AppM.ReadTaskFormDir(app.TASK_TYPE_VID)
	log.Println(AppM.Tasklist)
}

func (f *TMainform) OnFormShow(sender vcl.IObject) {

}


func (f *TMainform) OnMVDClick(sender vcl.IObject) {

}


func (f *TMainform) OnMVUClick(sender vcl.IObject) {

}


func (f *TMainform) OnMenuItem5Click(sender vcl.IObject) {

}


func (f *TMainform) OnPPFClick(sender vcl.IObject) {

}



func (f *TMainform) OnStartBtnMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x int32, y int32) {

}


func (f *TMainform) OnStartBtnMouseLeave(sender vcl.IObject) {

}


func (f *TMainform) OnStartBtnMouseMove(sender vcl.IObject, shift types.TShiftState, x int32, y int32) {

	f.StartBtn.Picture().LoadFromBytes([]byte{})
}


func (f *TMainform) OnStartBtnMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x int32, y int32) {

}

