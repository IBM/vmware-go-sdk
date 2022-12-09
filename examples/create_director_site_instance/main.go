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
	fileSharesModel := &vmwarev1.FileShares{
		STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
		STORAGETWOIOPSGB:          core.Int64Ptr(int64(24000)),
		STORAGEFOURIOPSGB:         core.Int64Ptr(int64(24000)),
		STORAGETENIOPSGB:          core.Int64Ptr(int64(8000)),
	}

	clusterOrderInfoModel := &vmwarev1.ClusterOrderInfo{
		Name:        core.StringPtr("cluster_1"),
		StorageType: core.StringPtr("nfs"),
		HostCount:   core.Int64Ptr(int64(3)),
		FileShares:  fileSharesModel,
		HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
	}

	pvdcOrderInfoModel := &vmwarev1.PVDCOrderInfo{
		Name:       core.StringPtr("pvdc-1"),
		DataCenter: core.StringPtr("dal10"),
		Clusters:   []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel},
	}

	createDirectorSitesOptions := vmwareService.NewCreateDirectorSitesOptions(
		"my_director_site",
		"...",
		[]vmwarev1.PVDCOrderInfo{*pvdcOrderInfoModel},
	)

	directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSite, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
