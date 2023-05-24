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
	getOidcConfigurationOptions := vmwareService.NewGetOidcConfigurationOptions(
		"site_id",
	)

	oidc, response, err := vmwareService.GetOidcConfiguration(getOidcConfigurationOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(oidc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
