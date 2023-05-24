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
	directorSitePvdcModel := &vmwarev1.DirectorSitePVDC{
		ID: core.StringPtr("pvdc_uuid"),
	}

	vdcDirectorSitePrototypeModel := &vmwarev1.VDCDirectorSitePrototype{
		ID:   core.StringPtr("directorsiteuuid"),
		Pvdc: directorSitePvdcModel,
	}

	createVdcOptions := vmwareService.NewCreateVdcOptions(
		"sampleVDC",
		vdcDirectorSitePrototypeModel,
	)

	vdc, response, err := vmwareService.CreateVdc(createVdcOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(vdc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
