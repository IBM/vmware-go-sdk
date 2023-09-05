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
	listDirectorSiteHostProfilesOptions := vmwareService.NewListDirectorSiteHostProfilesOptions()

	directorSiteHostProfileCollection, response, err := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSiteHostProfileCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
