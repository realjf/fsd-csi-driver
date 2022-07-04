package main

import (
	"flag"
	"fsd-csi-driver/pkg/fsd"

	"k8s.io/klog"
)

var (
	endpoint string
	nodeID   string
)

func main() {
	flag.StringVar(&endpoint, "endpoint", "", "CSI Endpoint")
	flag.StringVar(&nodeID, "nodeid", "", "node id")

	klog.InitFlags(nil)
	flag.Parse()

	d := fsd.NewDriver(nodeID, endpoint)
	d.Run()
}
