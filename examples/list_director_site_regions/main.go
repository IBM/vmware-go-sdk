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
	clusterPatchModel := &vmwarev1.ClusterPatch{}
	clusterPatchModelAsPatch, _ := clusterPatchModel.AsPatch()

	updateDirectorSitesPvdcsClusterOptions := vmwareService.NewUpdateDirectorSitesPvdcsClusterOptions(
		"site_id",
		"cluster_id",
		"pvdc_id",
		clusterPatchModelAsPatch,
	)

	updateCluster, response, err := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(updateCluster, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
