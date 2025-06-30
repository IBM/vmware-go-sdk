//go:build examples

/**
 * (C) Copyright IBM Corp. 2025.
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

package vmwarev1_test

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

//
// This file provides an example of how to use the vmware service.
//
// The following configuration properties are assumed to be defined:
// VMWARE_URL=<service base url>
// VMWARE_AUTH_TYPE=iam
// VMWARE_APIKEY=<IAM apikey>
// VMWARE_AUTH_URL=<IAM token service base URL - omit this if using the production environment>
//
// These configuration properties can be exported as environment variables, or stored
// in a configuration file and then:
// export IBM_CREDENTIALS_FILE=<name of configuration file>
//
var _ = Describe(`VmwareV1 Examples Tests`, func() {

	const externalConfigFile = "../vmware_v1.env"

	var (
		vmwareService *vmwarev1.VmwareV1
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
			config, err = core.GetServiceProperties(vmwarev1.DefaultServiceName)
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

			vmwareServiceOptions := &vmwarev1.VmwareV1Options{}

			vmwareService, err = vmwarev1.NewVmwareV1UsingExternalConfig(vmwareServiceOptions)

			if err != nil {
				panic(err)
			}

			// end-common

			Expect(vmwareService).ToNot(BeNil())
		})
	})

	Describe(`VmwareV1 request examples`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites request example`, func() {
			fmt.Println("\nCreateDirectorSites() result:")
			// begin-create_director_sites

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
			}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name: core.StringPtr("cluster_1"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
				FileShares: fileSharesPrototypeModel,
			}

			pvdcPrototypeModel := &vmwarev1.PVDCPrototype{
				Name: core.StringPtr("pvdc-1"),
				DataCenterName: core.StringPtr("dal10"),
				Clusters: []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
			}

			createDirectorSitesOptions := vmwareService.NewCreateDirectorSitesOptions(
				"my_director_site",
				[]vmwarev1.PVDCPrototype{*pvdcPrototypeModel},
			)
			createDirectorSitesOptions.SetAcceptLanguage("en-us")
			createDirectorSitesOptions.SetXGlobalTransactionID("transaction1")

			directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
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

			listDirectorSitesOptions := vmwareService.NewListDirectorSitesOptions()
			listDirectorSitesOptions.SetAcceptLanguage("en-us")
			listDirectorSitesOptions.SetXGlobalTransactionID("transaction1")

			directorSiteCollection, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSiteCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteCollection).ToNot(BeNil())
		})
		It(`GetDirectorSite request example`, func() {
			fmt.Println("\nGetDirectorSite() result:")
			// begin-get_director_site

			getDirectorSiteOptions := vmwareService.NewGetDirectorSiteOptions(
				"site_id",
			)
			getDirectorSiteOptions.SetAcceptLanguage("en-us")
			getDirectorSiteOptions.SetXGlobalTransactionID("transaction1")

			directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
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
		It(`EnableVeeamOnPvdcsList request example`, func() {
			fmt.Println("\nEnableVeeamOnPvdcsList() result:")
			// begin-enable_veeam_on_pvdcs_list

			enableVeeamOnPvdcsListOptions := vmwareService.NewEnableVeeamOnPvdcsListOptions(
				"site_id",
				true,
			)
			enableVeeamOnPvdcsListOptions.SetAcceptLanguage("en-us")
			enableVeeamOnPvdcsListOptions.SetXGlobalTransactionID("transaction1")

			serviceEnabled, response, err := vmwareService.EnableVeeamOnPvdcsList(enableVeeamOnPvdcsListOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceEnabled, "", "  ")
			fmt.Println(string(b))

			// end-enable_veeam_on_pvdcs_list

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceEnabled).ToNot(BeNil())
		})
		It(`EnableVcdaOnDataCenter request example`, func() {
			fmt.Println("\nEnableVcdaOnDataCenter() result:")
			// begin-enable_vcda_on_data_center

			enableVcdaOnDataCenterOptions := vmwareService.NewEnableVcdaOnDataCenterOptions(
				"site_id",
				true,
			)
			enableVcdaOnDataCenterOptions.SetAcceptLanguage("en-us")
			enableVcdaOnDataCenterOptions.SetXGlobalTransactionID("transaction1")

			serviceEnabled, response, err := vmwareService.EnableVcdaOnDataCenter(enableVcdaOnDataCenterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(serviceEnabled, "", "  ")
			fmt.Println(string(b))

			// end-enable_vcda_on_data_center

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceEnabled).ToNot(BeNil())
		})
		It(`CreateDirectorSitesVcdaConnectionEndpoints request example`, func() {
			fmt.Println("\nCreateDirectorSitesVcdaConnectionEndpoints() result:")
			// begin-create_director_sites_vcda_connection_endpoints

			createDirectorSitesVcdaConnectionEndpointsOptions := vmwareService.NewCreateDirectorSitesVcdaConnectionEndpointsOptions(
				"site_id",
				"private",
				"dal10",
			)
			createDirectorSitesVcdaConnectionEndpointsOptions.SetAcceptLanguage("en-us")
			createDirectorSitesVcdaConnectionEndpointsOptions.SetXGlobalTransactionID("transaction1")

			vcdaConnection, response, err := vmwareService.CreateDirectorSitesVcdaConnectionEndpoints(createDirectorSitesVcdaConnectionEndpointsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vcdaConnection, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites_vcda_connection_endpoints

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaConnection).ToNot(BeNil())
		})
		It(`UpdateDirectorSitesVcdaConnectionEndpoints request example`, func() {
			fmt.Println("\nUpdateDirectorSitesVcdaConnectionEndpoints() result:")
			// begin-update_director_sites_vcda_connection_endpoints

			updateDirectorSitesVcdaConnectionEndpointsOptions := vmwareService.NewUpdateDirectorSitesVcdaConnectionEndpointsOptions(
				"site_id",
				"vcda_connections_id",
			)
			updateDirectorSitesVcdaConnectionEndpointsOptions.SetAcceptLanguage("en-us")
			updateDirectorSitesVcdaConnectionEndpointsOptions.SetXGlobalTransactionID("transaction1")

			updatedVcdaConnection, response, err := vmwareService.UpdateDirectorSitesVcdaConnectionEndpoints(updateDirectorSitesVcdaConnectionEndpointsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updatedVcdaConnection, "", "  ")
			fmt.Println(string(b))

			// end-update_director_sites_vcda_connection_endpoints

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updatedVcdaConnection).ToNot(BeNil())
		})
		It(`CreateDirectorSitesVcdaC2cConnection request example`, func() {
			fmt.Println("\nCreateDirectorSitesVcdaC2cConnection() result:")
			// begin-create_director_sites_vcda_c2c_connection

			createDirectorSitesVcdaC2cConnectionOptions := vmwareService.NewCreateDirectorSitesVcdaC2cConnectionOptions(
				"site_id",
				"dal10",
				"ddirw002-gr80d10vcda",
				"dirw274t02vcda",
				"jp-tok",
			)
			createDirectorSitesVcdaC2cConnectionOptions.SetAcceptLanguage("en-us")
			createDirectorSitesVcdaC2cConnectionOptions.SetXGlobalTransactionID("transaction1")

			vcdaC2c, response, err := vmwareService.CreateDirectorSitesVcdaC2cConnection(createDirectorSitesVcdaC2cConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vcdaC2c, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites_vcda_c2c_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaC2c).ToNot(BeNil())
		})
		It(`UpdateDirectorSitesVcdaC2cConnection request example`, func() {
			fmt.Println("\nUpdateDirectorSitesVcdaC2cConnection() result:")
			// begin-update_director_sites_vcda_c2c_connection

			updateDirectorSitesVcdaC2cConnectionOptions := vmwareService.NewUpdateDirectorSitesVcdaC2cConnectionOptions(
				"site_id",
				"connection_id",
				"Text of the note...",
			)
			updateDirectorSitesVcdaC2cConnectionOptions.SetAcceptLanguage("en-us")
			updateDirectorSitesVcdaC2cConnectionOptions.SetXGlobalTransactionID("transaction1")

			updatedVcdaC2c, response, err := vmwareService.UpdateDirectorSitesVcdaC2cConnection(updateDirectorSitesVcdaC2cConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updatedVcdaC2c, "", "  ")
			fmt.Println(string(b))

			// end-update_director_sites_vcda_c2c_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updatedVcdaC2c).ToNot(BeNil())
		})
		It(`GetOidcConfiguration request example`, func() {
			fmt.Println("\nGetOidcConfiguration() result:")
			// begin-get_oidc_configuration

			getOidcConfigurationOptions := vmwareService.NewGetOidcConfigurationOptions(
				"site_id",
			)
			getOidcConfigurationOptions.SetAcceptLanguage("en-us")

			oidc, response, err := vmwareService.GetOidcConfiguration(getOidcConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(oidc, "", "  ")
			fmt.Println(string(b))

			// end-get_oidc_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(oidc).ToNot(BeNil())
		})
		It(`SetOidcConfiguration request example`, func() {
			fmt.Println("\nSetOidcConfiguration() result:")
			// begin-set_oidc_configuration

			setOidcConfigurationOptions := vmwareService.NewSetOidcConfigurationOptions(
				"site_id",
			)
			setOidcConfigurationOptions.SetAcceptLanguage("en-us")

			oidc, response, err := vmwareService.SetOidcConfiguration(setOidcConfigurationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(oidc, "", "  ")
			fmt.Println(string(b))

			// end-set_oidc_configuration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(oidc).ToNot(BeNil())
		})
		It(`ListDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nListDirectorSitesPvdcs() result:")
			// begin-list_director_sites_pvdcs

			listDirectorSitesPvdcsOptions := vmwareService.NewListDirectorSitesPvdcsOptions(
				"site_id",
			)
			listDirectorSitesPvdcsOptions.SetAcceptLanguage("en-us")
			listDirectorSitesPvdcsOptions.SetXGlobalTransactionID("transaction1")

			pvdcCollection, response, err := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdcCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdcCollection).ToNot(BeNil())
		})
		It(`CreateDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nCreateDirectorSitesPvdcs() result:")
			// begin-create_director_sites_pvdcs

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
			}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name: core.StringPtr("cluster_1"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
				FileShares: fileSharesPrototypeModel,
			}

			createDirectorSitesPvdcsOptions := vmwareService.NewCreateDirectorSitesPvdcsOptions(
				"site_id",
				"pvdc-1",
				"dal10",
				[]vmwarev1.ClusterPrototype{*clusterPrototypeModel},
			)
			createDirectorSitesPvdcsOptions.SetAcceptLanguage("en-us")
			createDirectorSitesPvdcsOptions.SetXGlobalTransactionID("transaction1")

			pvdc, response, err := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdc, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdc).ToNot(BeNil())
		})
		It(`GetDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nGetDirectorSitesPvdcs() result:")
			// begin-get_director_sites_pvdcs

			getDirectorSitesPvdcsOptions := vmwareService.NewGetDirectorSitesPvdcsOptions(
				"site_id",
				"pvdc_id",
			)
			getDirectorSitesPvdcsOptions.SetAcceptLanguage("en-us")
			getDirectorSitesPvdcsOptions.SetXGlobalTransactionID("transaction1")

			pvdc, response, err := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(pvdc, "", "  ")
			fmt.Println(string(b))

			// end-get_director_sites_pvdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdc).ToNot(BeNil())
		})
		It(`ListDirectorSitesPvdcsClusters request example`, func() {
			fmt.Println("\nListDirectorSitesPvdcsClusters() result:")
			// begin-list_director_sites_pvdcs_clusters

			listDirectorSitesPvdcsClustersOptions := vmwareService.NewListDirectorSitesPvdcsClustersOptions(
				"site_id",
				"pvdc_id",
			)
			listDirectorSitesPvdcsClustersOptions.SetAcceptLanguage("en-us")
			listDirectorSitesPvdcsClustersOptions.SetXGlobalTransactionID("transaction1")

			clusterCollection, response, err := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_director_sites_pvdcs_clusters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterCollection).ToNot(BeNil())
		})
		It(`CreateDirectorSitesPvdcsClusters request example`, func() {
			fmt.Println("\nCreateDirectorSitesPvdcsClusters() result:")
			// begin-create_director_sites_pvdcs_clusters

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
			}

			createDirectorSitesPvdcsClustersOptions := vmwareService.NewCreateDirectorSitesPvdcsClustersOptions(
				"site_id",
				"pvdc_id",
				"cluster_1",
				int64(2),
				"BM_2S_20_CORES_192_GB",
				fileSharesPrototypeModel,
			)
			createDirectorSitesPvdcsClustersOptions.SetAcceptLanguage("en-us")
			createDirectorSitesPvdcsClustersOptions.SetXGlobalTransactionID("transaction1")

			cluster, response, err := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(cluster, "", "  ")
			fmt.Println(string(b))

			// end-create_director_sites_pvdcs_clusters

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(cluster).ToNot(BeNil())
		})
		It(`GetDirectorInstancesPvdcsCluster request example`, func() {
			fmt.Println("\nGetDirectorInstancesPvdcsCluster() result:")
			// begin-get_director_instances_pvdcs_cluster

			getDirectorInstancesPvdcsClusterOptions := vmwareService.NewGetDirectorInstancesPvdcsClusterOptions(
				"site_id",
				"cluster_id",
				"pvdc_id",
			)
			getDirectorInstancesPvdcsClusterOptions.SetAcceptLanguage("en-us")
			getDirectorInstancesPvdcsClusterOptions.SetXGlobalTransactionID("transaction1")

			cluster, response, err := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
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

			clusterPatchModel := &vmwarev1.ClusterPatch{
			}
			clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDirectorSitesPvdcsClusterOptions := vmwareService.NewUpdateDirectorSitesPvdcsClusterOptions(
				"site_id",
				"cluster_id",
				"pvdc_id",
				clusterPatchModelAsPatch,
			)
			updateDirectorSitesPvdcsClusterOptions.SetAcceptLanguage("en-us")
			updateDirectorSitesPvdcsClusterOptions.SetXGlobalTransactionID("transaction1")

			updateCluster, response, err := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(updateCluster, "", "  ")
			fmt.Println(string(b))

			// end-update_director_sites_pvdcs_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateCluster).ToNot(BeNil())
		})
		It(`ListDirectorSiteRegions request example`, func() {
			fmt.Println("\nListDirectorSiteRegions() result:")
			// begin-list_director_site_regions

			listDirectorSiteRegionsOptions := vmwareService.NewListDirectorSiteRegionsOptions()
			listDirectorSiteRegionsOptions.SetAcceptLanguage("en-us")
			listDirectorSiteRegionsOptions.SetXGlobalTransactionID("transaction1")

			directorSiteRegionCollection, response, err := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSiteRegionCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_director_site_regions

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteRegionCollection).ToNot(BeNil())
		})
		It(`ListMultitenantDirectorSites request example`, func() {
			fmt.Println("\nListMultitenantDirectorSites() result:")
			// begin-list_multitenant_director_sites

			listMultitenantDirectorSitesOptions := vmwareService.NewListMultitenantDirectorSitesOptions()
			listMultitenantDirectorSitesOptions.SetAcceptLanguage("en-us")
			listMultitenantDirectorSitesOptions.SetXGlobalTransactionID("transaction1")

			multitenantDirectorSiteCollection, response, err := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(multitenantDirectorSiteCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_multitenant_director_sites

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(multitenantDirectorSiteCollection).ToNot(BeNil())
		})
		It(`ListDirectorSiteHostProfiles request example`, func() {
			fmt.Println("\nListDirectorSiteHostProfiles() result:")
			// begin-list_director_site_host_profiles

			listDirectorSiteHostProfilesOptions := vmwareService.NewListDirectorSiteHostProfilesOptions()
			listDirectorSiteHostProfilesOptions.SetAcceptLanguage("en-us")
			listDirectorSiteHostProfilesOptions.SetXGlobalTransactionID("transaction1")

			directorSiteHostProfileCollection, response, err := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(directorSiteHostProfileCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_director_site_host_profiles

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteHostProfileCollection).ToNot(BeNil())
		})
		It(`ListVdcs request example`, func() {
			fmt.Println("\nListVdcs() result:")
			// begin-list_vdcs

			listVdcsOptions := vmwareService.NewListVdcsOptions()
			listVdcsOptions.SetAcceptLanguage("en-us")

			vdcCollection, response, err := vmwareService.ListVdcs(listVdcsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vdcCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_vdcs

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdcCollection).ToNot(BeNil())
		})
		It(`CreateVdc request example`, func() {
			fmt.Println("\nCreateVdc() result:")
			// begin-create_vdc

			directorSitePvdcModel := &vmwarev1.DirectorSitePVDC{
				ID: core.StringPtr("ds_pvdc_id"),
			}

			vdcDirectorSitePrototypeModel := &vmwarev1.VDCDirectorSitePrototype{
				ID: core.StringPtr("ds_id"),
				Pvdc: directorSitePvdcModel,
			}

			createVdcOptions := vmwareService.NewCreateVdcOptions(
				"sampleVDC",
				vdcDirectorSitePrototypeModel,
			)
			createVdcOptions.SetAcceptLanguage("en-us")

			vdc, response, err := vmwareService.CreateVdc(createVdcOptions)
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

			getVdcOptions := vmwareService.NewGetVdcOptions(
				"vdc_id",
			)
			getVdcOptions.SetAcceptLanguage("en-us")

			vdc, response, err := vmwareService.GetVdc(getVdcOptions)
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
		It(`UpdateVdc request example`, func() {
			fmt.Println("\nUpdateVdc() result:")
			// begin-update_vdc

			vdcPatchModel := &vmwarev1.VDCPatch{
			}
			vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVdcOptions := vmwareService.NewUpdateVdcOptions(
				"vdc_id",
				vdcPatchModelAsPatch,
			)
			updateVdcOptions.SetAcceptLanguage("en-us")

			vdc, response, err := vmwareService.UpdateVdc(updateVdcOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vdc, "", "  ")
			fmt.Println(string(b))

			// end-update_vdc

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
		It(`AddTransitGatewayConnections request example`, func() {
			fmt.Println("\nAddTransitGatewayConnections() result:")
			// begin-add_transit_gateway_connections

			addTransitGatewayConnectionsOptions := vmwareService.NewAddTransitGatewayConnectionsOptions(
				"vdc_id",
				"edge_id",
				"transit_gateway_id",
			)
			addTransitGatewayConnectionsOptions.SetAcceptLanguage("en-us")

			transitGateway, response, err := vmwareService.AddTransitGatewayConnections(addTransitGatewayConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(transitGateway, "", "  ")
			fmt.Println(string(b))

			// end-add_transit_gateway_connections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(transitGateway).ToNot(BeNil())
		})
		It(`SwapHaEdgeSites request example`, func() {
			fmt.Println("\nSwapHaEdgeSites() result:")
			// begin-swap_ha_edge_sites

			swapHaEdgeSitesOptions := vmwareService.NewSwapHaEdgeSitesOptions(
				"vdc_id",
				"edge_id",
			)
			swapHaEdgeSitesOptions.SetAcceptLanguage("en-us")

			swapHaEdgeSitesResponse, response, err := vmwareService.SwapHaEdgeSites(swapHaEdgeSitesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(swapHaEdgeSitesResponse, "", "  ")
			fmt.Println(string(b))

			// end-swap_ha_edge_sites

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(swapHaEdgeSitesResponse).ToNot(BeNil())
		})
		It(`ListLicenses request example`, func() {
			fmt.Println("\nListLicenses() result:")
			// begin-list_licenses

			listLicensesOptions := vmwareService.NewListLicensesOptions()

			licenseCollection, response, err := vmwareService.ListLicenses(listLicensesOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(licenseCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_licenses

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(licenseCollection).ToNot(BeNil())
		})
		It(`ListUsageMeterRegistrations request example`, func() {
			fmt.Println("\nListUsageMeterRegistrations() result:")
			// begin-list_usage_meter_registrations

			listUsageMeterRegistrationsOptions := vmwareService.NewListUsageMeterRegistrationsOptions()

			usageMeterRegistrationCollection, response, err := vmwareService.ListUsageMeterRegistrations(listUsageMeterRegistrationsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(usageMeterRegistrationCollection, "", "  ")
			fmt.Println(string(b))

			// end-list_usage_meter_registrations

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(usageMeterRegistrationCollection).ToNot(BeNil())
		})
		It(`CreateUsageMeterRegistration request example`, func() {
			fmt.Println("\nCreateUsageMeterRegistration() result:")
			// begin-create_usage_meter_registration

			usageMeterIdentityModel := &vmwarev1.UsageMeterIdentity{
				ID: core.StringPtr("4242b01d-2db2-4d7b-ad5d-0792c61295a8"),
			}

			createUsageMeterRegistrationOptions := vmwareService.NewCreateUsageMeterRegistrationOptions(
				"string",
				usageMeterIdentityModel,
			)

			usageMeterRegistration, response, err := vmwareService.CreateUsageMeterRegistration(createUsageMeterRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(usageMeterRegistration, "", "  ")
			fmt.Println(string(b))

			// end-create_usage_meter_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(usageMeterRegistration).ToNot(BeNil())
		})
		It(`GetUsageMeterRegistration request example`, func() {
			fmt.Println("\nGetUsageMeterRegistration() result:")
			// begin-get_usage_meter_registration

			getUsageMeterRegistrationOptions := vmwareService.NewGetUsageMeterRegistrationOptions(
				"testString",
			)

			usageMeterRegistration, response, err := vmwareService.GetUsageMeterRegistration(getUsageMeterRegistrationOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(usageMeterRegistration, "", "  ")
			fmt.Println(string(b))

			// end-get_usage_meter_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(usageMeterRegistration).ToNot(BeNil())
		})
		It(`DeleteDirectorSite request example`, func() {
			fmt.Println("\nDeleteDirectorSite() result:")
			// begin-delete_director_site

			deleteDirectorSiteOptions := vmwareService.NewDeleteDirectorSiteOptions(
				"site_id",
			)
			deleteDirectorSiteOptions.SetAcceptLanguage("en-us")
			deleteDirectorSiteOptions.SetXGlobalTransactionID("transaction1")

			directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
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
		It(`DeleteDirectorSitesVcdaConnectionEndpoints request example`, func() {
			fmt.Println("\nDeleteDirectorSitesVcdaConnectionEndpoints() result:")
			// begin-delete_director_sites_vcda_connection_endpoints

			deleteDirectorSitesVcdaConnectionEndpointsOptions := vmwareService.NewDeleteDirectorSitesVcdaConnectionEndpointsOptions(
				"site_id",
				"vcda_connections_id",
			)
			deleteDirectorSitesVcdaConnectionEndpointsOptions.SetAcceptLanguage("en-us")
			deleteDirectorSitesVcdaConnectionEndpointsOptions.SetXGlobalTransactionID("transaction1")

			vcdaConnection, response, err := vmwareService.DeleteDirectorSitesVcdaConnectionEndpoints(deleteDirectorSitesVcdaConnectionEndpointsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vcdaConnection, "", "  ")
			fmt.Println(string(b))

			// end-delete_director_sites_vcda_connection_endpoints

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaConnection).ToNot(BeNil())
		})
		It(`DeleteDirectorSitesVcdaC2cConnection request example`, func() {
			fmt.Println("\nDeleteDirectorSitesVcdaC2cConnection() result:")
			// begin-delete_director_sites_vcda_c2c_connection

			deleteDirectorSitesVcdaC2cConnectionOptions := vmwareService.NewDeleteDirectorSitesVcdaC2cConnectionOptions(
				"site_id",
				"connection_id",
			)
			deleteDirectorSitesVcdaC2cConnectionOptions.SetAcceptLanguage("en-us")
			deleteDirectorSitesVcdaC2cConnectionOptions.SetXGlobalTransactionID("transaction1")

			vcdaC2c, response, err := vmwareService.DeleteDirectorSitesVcdaC2cConnection(deleteDirectorSitesVcdaC2cConnectionOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(vcdaC2c, "", "  ")
			fmt.Println(string(b))

			// end-delete_director_sites_vcda_c2c_connection

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaC2c).ToNot(BeNil())
		})
		It(`DeleteDirectorSitesPvdcsCluster request example`, func() {
			fmt.Println("\nDeleteDirectorSitesPvdcsCluster() result:")
			// begin-delete_director_sites_pvdcs_cluster

			deleteDirectorSitesPvdcsClusterOptions := vmwareService.NewDeleteDirectorSitesPvdcsClusterOptions(
				"site_id",
				"cluster_id",
				"pvdc_id",
			)
			deleteDirectorSitesPvdcsClusterOptions.SetAcceptLanguage("en-us")
			deleteDirectorSitesPvdcsClusterOptions.SetXGlobalTransactionID("transaction1")

			clusterSummary, response, err := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(clusterSummary, "", "  ")
			fmt.Println(string(b))

			// end-delete_director_sites_pvdcs_cluster

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(clusterSummary).ToNot(BeNil())
		})
		It(`DeleteVdc request example`, func() {
			fmt.Println("\nDeleteVdc() result:")
			// begin-delete_vdc

			deleteVdcOptions := vmwareService.NewDeleteVdcOptions(
				"vdc_id",
			)
			deleteVdcOptions.SetAcceptLanguage("en-us")

			vdc, response, err := vmwareService.DeleteVdc(deleteVdcOptions)
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
		It(`RemoveTransitGatewayConnections request example`, func() {
			fmt.Println("\nRemoveTransitGatewayConnections() result:")
			// begin-remove_transit_gateway_connections

			removeTransitGatewayConnectionsOptions := vmwareService.NewRemoveTransitGatewayConnectionsOptions(
				"vdc_id",
				"edge_id",
				"transit_gateway_id",
			)
			removeTransitGatewayConnectionsOptions.SetAcceptLanguage("en-us")

			transitGateway, response, err := vmwareService.RemoveTransitGatewayConnections(removeTransitGatewayConnectionsOptions)
			if err != nil {
				panic(err)
			}
			b, _ := json.MarshalIndent(transitGateway, "", "  ")
			fmt.Println(string(b))

			// end-remove_transit_gateway_connections

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(transitGateway).ToNot(BeNil())
		})
		It(`DeleteUsageMeterRegistration request example`, func() {
			// begin-delete_usage_meter_registration

			deleteUsageMeterRegistrationOptions := vmwareService.NewDeleteUsageMeterRegistrationOptions(
				"testString",
			)

			response, err := vmwareService.DeleteUsageMeterRegistration(deleteUsageMeterRegistrationOptions)
			if err != nil {
				panic(err)
			}
			if response.StatusCode != 204 {
				fmt.Printf("\nUnexpected response status code received from DeleteUsageMeterRegistration(): %d\n", response.StatusCode)
			}

			// end-delete_usage_meter_registration

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(204))
		})
	})
})
