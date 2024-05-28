
package main

import (
	"flag"

	"github.com/kubevirt/device-plugin-manager/pkg/dpm"
	"github.com/extend-k8s.io/k8s-device-plugin-example/pkg"
)

func main() {
	flag.Parse()
	manager := dpm.NewManager(pkg.Lister{})
	manager.Run()
}