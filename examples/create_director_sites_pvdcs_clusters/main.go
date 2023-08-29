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
	fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

	createDirectorSitesPvdcsClustersOptions := vmwareService.NewCreateDirectorSitesPvdcsClustersOptions(
		"site_id",
		"pvdc_id",
		"cluster_1",
		int64(3),
		"BM_2S_20_CORES_192_GB",
		fileSharesPrototypeModel,
	)

	cluster, response, err := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(cluster, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
