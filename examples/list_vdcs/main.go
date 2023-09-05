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
	listVdcsOptions := vmwareService.NewListVdcsOptions()

	vdcCollection, response, err := vmwareService.ListVdcs(listVdcsOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(vdcCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
