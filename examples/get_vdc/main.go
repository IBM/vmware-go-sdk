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
	getVdcOptions := vmwareService.NewGetVdcOptions(
		"testString",
	)

	vdc, response, err := vmwareService.GetVdc(getVdcOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(vdc, "", "  ")
	fmt.Println(string(b))Z
	fmt.Println(response)
}
