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
	listDirectorSitesPvdcsClustersOptions := vmwareService.NewListDirectorSitesPvdcsClustersOptions(
		"site_id",
		"pvdc_id",
	)

	clusterCollection, response, err := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(clusterCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
