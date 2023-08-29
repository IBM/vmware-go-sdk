package main

import (
	"encoding/json"
	"fmt"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

var (
	vmwareService *vmwarev1.VmwareV1
)

func main() {
	getDirectorInstancesPvdcsClusterOptions := vmwareService.NewGetDirectorInstancesPvdcsClusterOptions(
		"site_id",
		"id",
		"pvdc_id",
	)

	cluster, response, err := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(cluster, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
