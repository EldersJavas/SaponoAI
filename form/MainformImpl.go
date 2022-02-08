package form

import (
	_ "embed"
	"errors"
	"github.com/gen2brain/dlgs"
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
	multi, b, err := dlgs.FileMulti("选择图片", "")
	if err != nil {
		log.Println(err)
		dlgs.Error("错误", "无法读取文件."+err.Error())
		return
	}

}

func (f *TMainform) OnMPDClick(sender vcl.IObject) {
	vcl.SelectDirectory2("选择图片文件夹", ".", false)
}

func (f *TMainform) OnMPUClick(sender vcl.IObject) {

}

func (f *TMainform) OnMVFClick(sender vcl.IObject) {

}

func (f *TMainform) OnMVDClick(sender vcl.IObject) {

}

func (f *TMainform) OnMVUClick(sender vcl.IObject) {

}
