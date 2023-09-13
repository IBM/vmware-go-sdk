//go:build examples
// +build examples

/**
 * (C) Copyright IBM Corp. 2023.
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
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

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
var _ = Describe(`VmwareV1 Examples Tests`, func() {

	const externalConfigFile = "../vmware_v1.env"

	var (
		vmwareService *vmwarev1.VmwareV1
		config        map[string]string
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

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name:        core.StringPtr("testString"),
				HostCount:   core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("testString"),
				FileShares:  fileSharesPrototypeModel,
			}

			pvdcPrototypeModel := &vmwarev1.PVDCPrototype{
				Name:           core.StringPtr("testString"),
				DataCenterName: core.StringPtr("testString"),
				Clusters:       []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
			}

			createDirectorSitesOptions := vmwareService.NewCreateDirectorSitesOptions(
				"testString",
				[]vmwarev1.PVDCPrototype{*pvdcPrototypeModel},
			)

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
				"testString",
			)

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
		It(`ListDirectorSitesPvdcs request example`, func() {
			fmt.Println("\nListDirectorSitesPvdcs() result:")
			// begin-list_director_sites_pvdcs

			listDirectorSitesPvdcsOptions := vmwareService.NewListDirectorSitesPvdcsOptions(
				"testString",
			)

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

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name:        core.StringPtr("testString"),
				HostCount:   core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("testString"),
				FileShares:  fileSharesPrototypeModel,
			}

			createDirectorSitesPvdcsOptions := vmwareService.NewCreateDirectorSitesPvdcsOptions(
				"testString",
				"testString",
				"testString",
				[]vmwarev1.ClusterPrototype{*clusterPrototypeModel},
			)

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
				"testString",
				"testString",
			)

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
				"testString",
				"testString",
			)

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

			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{}

			createDirectorSitesPvdcsClustersOptions := vmwareService.NewCreateDirectorSitesPvdcsClustersOptions(
				"testString",
				"testString",
				"testString",
				int64(2),
				"testString",
				fileSharesPrototypeModel,
			)

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
				"testString",
				"testString",
				"testString",
			)

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

			clusterPatchModel := &vmwarev1.ClusterPatch{}
			clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDirectorSitesPvdcsClusterOptions := vmwareService.NewUpdateDirectorSitesPvdcsClusterOptions(
				"testString",
				"testString",
				"testString",
				clusterPatchModelAsPatch,
			)

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
				ID: core.StringPtr("testString"),
			}

			vdcDirectorSitePrototypeModel := &vmwarev1.VDCDirectorSitePrototype{
				ID:   core.StringPtr("testString"),
				Pvdc: directorSitePvdcModel,
			}

			createVdcOptions := vmwareService.NewCreateVdcOptions(
				"testString",
				vdcDirectorSitePrototypeModel,
			)

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
				"testString",
			)

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

			vdcPatchModel := &vmwarev1.VDCPatch{}
			vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVdcOptions := vmwareService.NewUpdateVdcOptions(
				"testString",
				vdcPatchModelAsPatch,
			)

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
		It(`DeleteDirectorSite request example`, func() {
			fmt.Println("\nDeleteDirectorSite() result:")
			// begin-delete_director_site

			deleteDirectorSiteOptions := vmwareService.NewDeleteDirectorSiteOptions(
				"testString",
			)

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
		It(`DeleteDirectorSitesPvdcsCluster request example`, func() {
			fmt.Println("\nDeleteDirectorSitesPvdcsCluster() result:")
			// begin-delete_director_sites_pvdcs_cluster

			deleteDirectorSitesPvdcsClusterOptions := vmwareService.NewDeleteDirectorSitesPvdcsClusterOptions(
				"testString",
				"testString",
				"testString",
			)

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
				"testString",
			)

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
	})
})
