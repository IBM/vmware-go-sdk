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
	deleteDirectorSiteOptions := vmwareService.NewDeleteDirectorSiteOptions(
		"id",
	)

	directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(directorSite, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
