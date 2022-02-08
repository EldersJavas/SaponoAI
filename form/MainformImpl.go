
package form

import (
    _ "embed"
    "github.com/ying32/govcl/vcl"
    "github.com/ying32/govcl/vcl/types"
)


//::private::
type TMainformFields struct {

}

func (f *TMainform) OnFormCreate(sender vcl.IObject) {

}


func (f *TMainform) OnFormShow(sender vcl.IObject) {

}


func (f *TMainform) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
    *canClose = vcl.MessageDlg("是否退出?", types.MtInformation, types.MbYes, types.MbNo) == types.MrYes
    if *canClose {vcl.Application.Terminate()}
}
func (f *TMainform) OnBackgroundClick(sender vcl.IObject) {

}

