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
	deleteDirectorSitesPvdcsClusterOptions := vmwareService.NewDeleteDirectorSitesPvdcsClusterOptions(
		"site_id",
		"id",
		"pvdc_id",
	)

	clusterSummary, response, err := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(clusterSummary, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
