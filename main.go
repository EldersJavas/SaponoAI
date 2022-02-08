package main

import (
	"github.com/EldersJavas/SaponoAI/form"
	"github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.SetTitle("SaponoAI")
    vcl.Application.SetScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&form.Mainform)
	form.Mainform.SetDoubleBuffered(true)
	vcl.Application.Run()

}
