// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainform struct {
    *vcl.TForm
    GroupBox1  *vcl.TGroupBox
    GroupBox2  *vcl.TGroupBox
    PBpres     *vcl.TProgressBar
    Pres       *vcl.TEdit
    Background *vcl.TImage
    ListView1  *vcl.TListView

    //::private::
    TMainformFields
}

var Mainform *TMainform




// vcl.Application.CreateForm(&Mainform)

func NewMainform(owner vcl.IComponent) (root *TMainform)  {
    vcl.CreateResForm(owner, &root)
    return
}

//go:embed resources/Mainform.gfm
var mainformBytes []byte

// 注册Form资源  
var _ = vcl.RegisterFormResource(Mainform, &mainformBytes)
