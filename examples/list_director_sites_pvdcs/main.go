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
	listDirectorSitesPvdcsOptions := vmwareService.NewListDirectorSitesPvdcsOptions(
		"site_id",
	)

	pvdcCollection, response, err := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(pvdcCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
