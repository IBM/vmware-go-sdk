// +build examples

/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package vmwareasaserviceapiv1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwareasaserviceapiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the VMware as a Service API service.
//
// The following configuration properties are assumed to be defined:
// V_MWARE_AS_A_SERVICE_API_URL=<service base url>
// V_MWARE_AS_A_SERVICE_API_AUTH_TYPE=iam
// V_MWARE_AS_A_SERVICE_API_APIKEY=<IAM apikey>
// V_MWARE_AS_A_SERVICE_API_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`VMwareAsAServiceApiV1 Examples Tests`, func() {

	const externalConfigFile = "../v_mware_as_a_service_api_v1.env"

	var (
		vMwareAsAServiceApiService *vmwareasaserviceapiv1.VMwareAsAServiceApiV1
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping examples...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			var err error
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping examples: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(vmwareasaserviceapiv1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping examples: " + err.Error())
			} else if len(config) == 0 {
				Skip("Unable to load service properties, skipping examples")
			}

			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			var err error

			// begin-common

			vMwareAsAServiceApiServiceOptions := &vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{}

			vMwareAsAServiceApiService, err = vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(vMwareAsAServiceApiServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(vMwareAsAServiceApiService).ToNot(BeNil())
		})
	})

	Describe(`VMwareAsAServiceApiV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites request example`, func() {
			fmt.Println("\nCreateDirectorSites() result:")
			// begin-create_director_sites

			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
			}

			clusterOrderInfoModel := &vmwareasaserviceapiv1.ClusterOrderInfo{
				Name: core.StringPtr("testString"),
				StorageType: core.StringPtr("nfs"),
				HostCount: core.Int64Ptr(int64(2)),
				FileShares: fileSharesModel,
				HostProfile: core.StringPtr("testString"),
			}

			pvdcOrderInfoModel := &vmwareasaserviceapiv1.PVDCOrderInfo{
				Name: core.StringPtr("testString"),
				DataCenter: core.StringPtr("testString"),
				Clusters: []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel},
			}

			createDirectorSitesOptions := vMwareAsAServiceApiService.NewCreateDirectorSitesOptions(
				"testString",
				"testString",
				[]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel},
			)

			directorSite, response, err := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSite, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
		It(`ListDirectorSites request example`, func() {
			fmt.Println("\nListDirectorSites() result:")
			// begin-list_director_sites

			listDirectorSitesOptions := vMwareAsAServiceApiService.NewListDirectorSitesOptions()

			listDirectorSites, response, err := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listDirectorSites, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDirectorSites).ToNot(BeNil())
		})
		It(`GetDirectorSite request example`, func() {
			fmt.Println("\nGetDirectorSite() result:")
			// begin-get_director_site

			getDirectorSiteOptions := vMwareAsAServiceApiService.NewGetDirectorSiteOptions(
				"testString",
			)

			directorSite, response, err := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSite, "", "  ")
			fmt.Println(string(b))

			// end-get_director_site

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSite).ToNot(BeNil())
		})
		It(`ListDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nListDirectorSitesPvdcs() result:")
			// begin-list_director_sites_pvdcs

			listDirectorSitesPvdcsOptions := vMwareAsAServiceApiService.NewListDirectorSitesPvdcsOptions(
				"testString",
			)

			listPvdCs, response, err := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listPvdCs, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listPvdCs).ToNot(BeNil())
		})
		It(`CreateDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nCreateDirectorSitesPvdcs() result:")
			// begin-create_director_sites_pvdcs

			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
			}

			clusterOrderInfoModel := &vmwareasaserviceapiv1.ClusterOrderInfo{
				Name: core.StringPtr("testString"),
				StorageType: core.StringPtr("nfs"),
				HostCount: core.Int64Ptr(int64(2)),
				FileShares: fileSharesModel,
				HostProfile: core.StringPtr("testString"),
			}

			createDirectorSitesPvdcsOptions := vMwareAsAServiceApiService.NewCreateDirectorSitesPvdcsOptions(
				"testString",
				"testString",
				"testString",
				[]vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel},
			)

			pvdcResponse, response, err := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdcResponse, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdcResponse).ToNot(BeNil())
		})
		It(`GetDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nGetDirectorSitesPvdcs() result:")
			// begin-get_director_sites_pvdcs

			getDirectorSitesPvdcsOptions := vMwareAsAServiceApiService.NewGetDirectorSitesPvdcsOptions(
				"testString",
				"testString",
			)

			pvdcSummary, response, err := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdcSummary, "", "  ")
			fmt.Println(string(b))

			// end-get_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdcSummary).ToNot(BeNil())
		})
		It(`ListDirectorSitesPvdcsClusters request example`, func() {
			fmt.Println("\nListDirectorSitesPvdcsClusters() result:")
			// begin-list_director_sites_pvdcs_clusters

			listDirectorSitesPvdcsClustersOptions := vMwareAsAServiceApiService.NewListDirectorSitesPvdcsClustersOptions(
				"testString",
				"testString",
			)

			listClusters, response, err := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listClusters, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites_pvdcs_clusters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listClusters).ToNot(BeNil())
		})
		It(`GetDirectorInstancesPvdcsCluster request example`, func() {
			fmt.Println("\nGetDirectorInstancesPvdcsCluster() result:")
			// begin-get_director_instances_pvdcs_cluster

			getDirectorInstancesPvdcsClusterOptions := vMwareAsAServiceApiService.NewGetDirectorInstancesPvdcsClusterOptions(
				"testString",
				"testString",
				"testString",
			)

			cluster, response, err := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cluster, "", "  ")
			fmt.Println(string(b))

			// end-get_director_instances_pvdcs_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())
		})
		It(`UpdateDirectorSitesPvdcsCluster request example`, func() {
			fmt.Println("\nUpdateDirectorSitesPvdcsCluster() result:")
			// begin-update_director_sites_pvdcs_cluster

			jsonPatchOperationModel := &vmwareasaserviceapiv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
			}

			updateDirectorSitesPvdcsClusterOptions := vMwareAsAServiceApiService.NewUpdateDirectorSitesPvdcsClusterOptions(
				"testString",
				"testString",
				"testString",
				[]vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel},
			)

			updateClusterResponse, response, err := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updateClusterResponse, "", "  ")
			fmt.Println(string(b))

			// end-update_director_sites_pvdcs_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateClusterResponse).ToNot(BeNil())
		})
		It(`ListDirectorSiteRegions request example`, func() {
			fmt.Println("\nListDirectorSiteRegions() result:")
			// begin-list_director_site_regions

			listDirectorSiteRegionsOptions := vMwareAsAServiceApiService.NewListDirectorSiteRegionsOptions()

			directorSiteRegions, response, err := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSiteRegions, "", "  ")
			fmt.Println(string(b))

			// end-list_director_site_regions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteRegions).ToNot(BeNil())
		})
		It(`ListDirectorSiteHostProfiles request example`, func() {
			fmt.Println("\nListDirectorSiteHostProfiles() result:")
			// begin-list_director_site_host_profiles

			listDirectorSiteHostProfilesOptions := vMwareAsAServiceApiService.NewListDirectorSiteHostProfilesOptions()

			listHostProfiles, response, err := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listHostProfiles, "", "  ")
			fmt.Println(string(b))

			// end-list_director_site_host_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listHostProfiles).ToNot(BeNil())
		})
		It(`ReplaceOrgAdminPassword request example`, func() {
			fmt.Println("\nReplaceOrgAdminPassword() result:")
			// begin-replace_org_admin_password

			replaceOrgAdminPasswordOptions := vMwareAsAServiceApiService.NewReplaceOrgAdminPasswordOptions(
				"testString",
			)

			newPassword, response, err := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(newPassword, "", "  ")
			fmt.Println(string(b))

			// end-replace_org_admin_password

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(newPassword).ToNot(BeNil())
		})
		It(`ListPrices request example`, func() {
			fmt.Println("\nListPrices() result:")
			// begin-list_prices

			listPricesOptions := vMwareAsAServiceApiService.NewListPricesOptions()

			directorSitePricingInfo, response, err := vMwareAsAServiceApiService.ListPrices(listPricesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSitePricingInfo, "", "  ")
			fmt.Println(string(b))

			// end-list_prices

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSitePricingInfo).ToNot(BeNil())
		})
		It(`GetVcddPrice request example`, func() {
			fmt.Println("\nGetVcddPrice() result:")
			// begin-get_vcdd_price

			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
			}

			clusterOrderInfoModel := &vmwareasaserviceapiv1.ClusterOrderInfo{
				Name: core.StringPtr("testString"),
				StorageType: core.StringPtr("nfs"),
				HostCount: core.Int64Ptr(int64(2)),
				FileShares: fileSharesModel,
				HostProfile: core.StringPtr("testString"),
			}

			pvdcOrderInfoModel := &vmwareasaserviceapiv1.PVDCOrderInfo{
				Name: core.StringPtr("testString"),
				DataCenter: core.StringPtr("testString"),
				Clusters: []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel},
			}

			getVcddPriceOptions := vMwareAsAServiceApiService.NewGetVcddPriceOptions(
				"testString",
				"testString",
				[]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel},
			)

			directorSitePriceQuoteResponse, response, err := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSitePriceQuoteResponse, "", "  ")
			fmt.Println(string(b))

			// end-get_vcdd_price

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(directorSitePriceQuoteResponse).ToNot(BeNil())
		})
		It(`ListVdcs request example`, func() {
			fmt.Println("\nListVdcs() result:")
			// begin-list_vdcs

			listVdcsOptions := vMwareAsAServiceApiService.NewListVdcsOptions()

			listVdCs, response, err := vMwareAsAServiceApiService.ListVdcs(listVdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(listVdCs, "", "  ")
			fmt.Println(string(b))

			// end-list_vdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listVdCs).ToNot(BeNil())
		})
		It(`CreateVdc request example`, func() {
			fmt.Println("\nCreateVdc() result:")
			// begin-create_vdc

			vdcDirectorSiteClusterModel := &vmwareasaserviceapiv1.VDCDirectorSiteCluster{
				ID: core.StringPtr("testString"),
			}

			newVdcDirectorSiteModel := &vmwareasaserviceapiv1.NewVDCDirectorSite{
				ID: core.StringPtr("testString"),
				Cluster: vdcDirectorSiteClusterModel,
			}

			createVdcOptions := vMwareAsAServiceApiService.NewCreateVdcOptions(
				"testString",
				newVdcDirectorSiteModel,
			)

			vdc, response, err := vMwareAsAServiceApiService.CreateVdc(createVdcOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vdc, "", "  ")
			fmt.Println(string(b))

			// end-create_vdc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
		It(`GetVdc request example`, func() {
			fmt.Println("\nGetVdc() result:")
			// begin-get_vdc

			getVdcOptions := vMwareAsAServiceApiService.NewGetVdcOptions(
				"testString",
			)

			vdc, response, err := vMwareAsAServiceApiService.GetVdc(getVdcOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vdc, "", "  ")
			fmt.Println(string(b))

			// end-get_vdc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdc).ToNot(BeNil())
		})
		It(`DeleteVdc request example`, func() {
			fmt.Println("\nDeleteVdc() result:")
			// begin-delete_vdc

			deleteVdcOptions := vMwareAsAServiceApiService.NewDeleteVdcOptions(
				"testString",
			)

			vdc, response, err := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vdc, "", "  ")
			fmt.Println(string(b))

			// end-delete_vdc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
		It(`DeleteDirectorSitesPvdcsCluster request example`, func() {
			fmt.Println("\nDeleteDirectorSitesPvdcsCluster() result:")
			// begin-delete_director_sites_pvdcs_cluster

			deleteDirectorSitesPvdcsClusterOptions := vMwareAsAServiceApiService.NewDeleteDirectorSitesPvdcsClusterOptions(
				"testString",
				"testString",
				"testString",
			)

			pvdcResponse, response, err := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdcResponse, "", "  ")
			fmt.Println(string(b))

			// end-delete_director_sites_pvdcs_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdcResponse).ToNot(BeNil())
		})
		It(`DeleteDirectorSite request example`, func() {
			fmt.Println("\nDeleteDirectorSite() result:")
			// begin-delete_director_site

			deleteDirectorSiteOptions := vMwareAsAServiceApiService.NewDeleteDirectorSiteOptions(
				"testString",
			)

			directorSite, response, err := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSite, "", "  ")
			fmt.Println(string(b))

			// end-delete_director_site

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})
})
