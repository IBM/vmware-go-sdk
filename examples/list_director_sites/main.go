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
	listDirectorSitesOptions := vmwareService.NewListDirectorSitesOptions()

	directorSiteCollection, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSiteCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
