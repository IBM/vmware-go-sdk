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
	listMultitenantDirectorSitesOptions := vmwareService.NewListMultitenantDirectorSitesOptions()

	multitenantDirectorSiteCollection, response, err := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(multitenantDirectorSiteCollection, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
