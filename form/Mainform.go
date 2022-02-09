// 由res2go IDE插件自动生成，不要编辑。
package form

import (
    "github.com/ying32/govcl/vcl"
    _ "embed"
)

type TMainform struct {
    *vcl.TForm
    Background *vcl.TImage
    LeftBox    *vcl.TGroupBox
    RightBox   *vcl.TGroupBox
    PBpres     *vcl.TProgressBar
    ListView1  *vcl.TListView
    Pres       *vcl.TStaticText
    MainMenu   *vcl.TMainMenu
    MenuItem1  *vcl.TMenuItem
    Maddpic    *vcl.TMenuItem
    MPF        *vcl.TMenuItem
    MPD        *vcl.TMenuItem
    MPU        *vcl.TMenuItem
    Maddvd     *vcl.TMenuItem
    MVF        *vcl.TMenuItem
    MVD        *vcl.TMenuItem
    MVU        *vcl.TMenuItem
    MenuItem4  *vcl.TMenuItem
    MenuItem5  *vcl.TMenuItem
    Button1    *vcl.TButton
    Button2    *vcl.TButton
    PopupList  *vcl.TPopupMenu
    MenuItem2  *vcl.TMenuItem
    PPF        *vcl.TMenuItem
    PPD        *vcl.TMenuItem
    PPU        *vcl.TMenuItem
    MenuItem8  *vcl.TMenuItem
    MenuItem9  *vcl.TMenuItem
    MenuItem10 *vcl.TMenuItem
    MenuItem11 *vcl.TMenuItem
    PVF        *vcl.TMenuItem
    PVD        *vcl.TMenuItem
    PVU        *vcl.TMenuItem

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
