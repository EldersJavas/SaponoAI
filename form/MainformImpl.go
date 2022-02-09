package form

import (
	_ "embed"
	"github.com/EldersJavas/SaponoAI/app"
	"github.com/EldersJavas/SaponoAI/app/tool"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
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
	f.Logout(AppM.Tasklist)
}

func (f *TMainform) OnMPDClick(sender vcl.IObject) {
	AppM.ReadTaskFormDir(app.TASK_TYPE_PIC)
	f.Logout(AppM.Tasklist)
}

func (f *TMainform) OnMPUClick(sender vcl.IObject) {
	AppM.ReadTaskFormFiles(app.TASK_TYPE_VID)
	f.Logout(AppM.Tasklist)
}

func (f *TMainform) OnMVFClick(sender vcl.IObject) {
	AppM.ReadTaskFormDir(app.TASK_TYPE_VID)
	f.Logout(AppM.Tasklist)
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
	vcl.ThreadSync(func() {
		f.StartBtn.Picture().LoadFromFile(tool.FileFormatPath(".\\res\\startbtn2.png"))
	})
}

func (f *TMainform) OnStartBtnMouseLeave(sender vcl.IObject) {
	vcl.ThreadSync(func() {
		f.StartBtn.Picture().LoadFromFile(tool.FileFormatPath(".\\res\\startbtn1.png"))
	})
}

func (f *TMainform) OnStartBtnMouseMove(sender vcl.IObject, shift types.TShiftState, x int32, y int32) {
	vcl.ThreadSync(func() {
		f.StartBtn.Picture().LoadFromFile(tool.FileFormatPath(".\\res\\startbtn2.png"))
	})

}

func (f *TMainform) OnStartBtnMouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x int32, y int32) {
	vcl.ThreadSync(func() {
		f.StartBtn.Picture().LoadFromFile(tool.FileFormatPath(".\\res\\startbtn1.png"))
	})
}

func (f *TMainform) OnFormCreate(sender vcl.IObject) {

}

func (f *TMainform) OnLogListClick(sender vcl.IObject) {
	vcl.ThreadSync(func() {
		//f.LogList.AddItem("test",nil)
	})
}
