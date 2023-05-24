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
	getDirectorSitesPvdcsOptions := vmwareService.NewGetDirectorSitesPvdcsOptions(
		"site_id",
		"pvdc-id",
	)

	pvdc, response, err := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(pvdc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
