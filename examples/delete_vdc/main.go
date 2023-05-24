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
	deleteVdcOptions := vmwareService.NewDeleteVdcOptions(
		"vdc_id",
	)

	vdc, response, err := vmwareService.DeleteVdc(deleteVdcOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(vdc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
