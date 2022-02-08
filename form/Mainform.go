// 由res2go IDE插件自动生成，不要编辑。
package form

import (
	_ "embed"
	"github.com/ying32/govcl/vcl"
)

type TMainform struct {
	*vcl.TForm
	LeftBox    *vcl.TGroupBox
	RightBox   *vcl.TGroupBox
	PBpres     *vcl.TProgressBar
	Background *vcl.TImage
	ListView1  *vcl.TListView
	Pres       *vcl.TStaticText
	MainMenu   *vcl.TMainMenu
	MenuItem1  *vcl.TMenuItem
	Maddpic    *vcl.TMenuItem
	Maddvd     *vcl.TMenuItem
	MenuItem4  *vcl.TMenuItem
	MenuItem5  *vcl.TMenuItem
	MPF        *vcl.TMenuItem
	MPD        *vcl.TMenuItem
	MPU        *vcl.TMenuItem
	MVF        *vcl.TMenuItem
	MVD        *vcl.TMenuItem
	MVU        *vcl.TMenuItem

	//::private::
	TMainformFields
}

var Mainform *TMainform

// vcl.Application.CreateForm(&Mainform)

func NewMainform(owner vcl.IComponent) (root *TMainform) {
	vcl.CreateResForm(owner, &root)
	return
}

//go:embed resources/Mainform.gfm
var mainformBytes []byte

// 注册Form资源
var _ = vcl.RegisterFormResource(Mainform, &mainformBytes)
