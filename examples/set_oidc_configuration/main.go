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
	setOidcConfigurationOptions := vmwareService.NewSetOidcConfigurationOptions(
		"site_id",
	)

	oidc, response, err := vmwareService.SetOidcConfiguration(setOidcConfigurationOptions)
	if err != nil {
		panic(err)
	}
	b, _ := json.MarshalIndent(oidc, "", "  ")
	fmt.Println(string(b))
	fmt.Println(response)
}
