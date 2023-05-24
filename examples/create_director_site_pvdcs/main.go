package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

var (
	vmwareService *vmwarev1.VmwareV1
)

func main() {
	fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

	clusterPrototypeModel := &vmwarev1.ClusterPrototype{
		Name:        core.StringPtr("cluster_1"),
		HostCount:   core.Int64Ptr(int64(3)),
		HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
		FileShares:  fileSharesPrototypeModel,
	}

	createDirectorSitesPvdcsOptions := vmwareService.NewCreateDirectorSitesPvdcsOptions(
		"site_id",
		"pvdc-1",
		"dal10",
		[]vmwarev1.ClusterPrototype{*clusterPrototypeModel},
	)

	pvdc, response, err := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(pvdc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
