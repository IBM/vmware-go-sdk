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
	getDirectorSiteOptions := vmwareService.NewGetDirectorSiteOptions(
		"site_id",
	)

	directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSite, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
