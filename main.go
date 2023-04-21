package main

import (
	_ "embed"

	"github.com/vishnushankarsg/metachain/command/root"
	"github.com/vishnushankarsg/metachain/licenses"
)

var (
	//go:embed LICENSE
	license string
)

func main() {
	licenses.SetLicense(license)

	root.NewRootCommand().Execute()
}
