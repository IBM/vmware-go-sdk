package main

import (
	"encoding/json"
	"fmt"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
)

var (
	vmwareService *vmwarev1.VmwareV1
)

func main() {
	vdcPatchModel := &vmwarev1.VDCPatch{}
	vdcPatchModelAsPatch, _ := vdcPatchModel.AsPatch()

	updateVdcOptions := vmwareService.NewUpdateVdcOptions(
		"vdc_id",
		vdcPatchModelAsPatch,
	)

	vdc, response, err := vmwareService.UpdateVdc(updateVdcOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(vdc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
