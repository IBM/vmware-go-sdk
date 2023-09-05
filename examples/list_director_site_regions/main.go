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
	listDirectorSiteRegionsOptions := vmwareService.NewListDirectorSiteRegionsOptions()

	directorSiteRegionCollection, response, err := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSiteRegionCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
