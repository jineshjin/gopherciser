package main

import (
	"github.com/qlik-oss/gopherciser/generatedocs/pkg/doccompiler"
	"github.com/qlik-oss/gopherciser/generatedocs/pkg/flags"
)

func main() {
	data := doccompiler.NewData()
	data.PopulateFromDataDir(flags.DataRoot())
	data.CompileToFile(flags.OutputFile())
}
