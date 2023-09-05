package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
)

var (
	vmwareService *vmwarev1.VmwareV1
)

func main() {
	fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

	clusterPrototypeModel := &vmwarev1.ClusterPrototype{
		Name:        core.StringPtr("cluster_1"),
		HostCount:   core.Int64Ptr(int64(2)),
		HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
		FileShares:  fileSharesPrototypeModel,
	}

	pvdcPrototypeModel := &vmwarev1.PVDCPrototype{
		Name:           core.StringPtr("pvdc-1"),
		DataCenterName: core.StringPtr("dal10"),
		Clusters:       []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
	}

	createDirectorSitesOptions := vmwareService.NewCreateDirectorSitesOptions(
		"my_director_site",
		[]vmwarev1.PVDCPrototype{*pvdcPrototypeModel},
	)

	directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSite, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
