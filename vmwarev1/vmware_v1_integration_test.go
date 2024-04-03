//go:build integration

/**
 * (C) Copyright IBM Corp. 2024.
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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwarev1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the vmwarev1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`VmwareV1 Integration Tests`, func() {
	const externalConfigFile = "../vmware_v1.env"

	var (
		err          error
		vmwareService *vmwarev1.VmwareV1
		serviceURL   string
		config       map[string]string
	)

	var shouldSkipTest = func() {
		Skip("External configuration is not available, skipping tests...")
	}

	Describe(`External configuration`, func() {
		It("Successfully load the configuration", func() {
			_, err = os.Stat(externalConfigFile)
			if err != nil {
				Skip("External configuration file not found, skipping tests: " + err.Error())
			}

			os.Setenv("IBM_CREDENTIALS_FILE", externalConfigFile)
			config, err = core.GetServiceProperties(vmwarev1.DefaultServiceName)
			if err != nil {
				Skip("Error loading service properties, skipping tests: " + err.Error())
			}
			serviceURL = config["URL"]
			if serviceURL == "" {
				Skip("Unable to load service URL configuration property, skipping tests")
			}

			fmt.Fprintf(GinkgoWriter, "Service URL: %v\n", serviceURL)
			shouldSkipTest = func() {}
		})
	})

	Describe(`Client initialization`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It("Successfully construct the service client instance", func() {
			vmwareServiceOptions := &vmwarev1.VmwareV1Options{}

			vmwareService, err = vmwarev1.NewVmwareV1UsingExternalConfig(vmwareServiceOptions)
			Expect(err).To(BeNil())
			Expect(vmwareService).ToNot(BeNil())
			Expect(vmwareService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			vmwareService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateDirectorSites - Create a Cloud Director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
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

			resourceGroupIdentityModel := &vmwarev1.ResourceGroupIdentity{
				ID: core.StringPtr("some_resourcegroupid"),
			}

			serviceIdentityModel := &vmwarev1.ServiceIdentity{
				Name: core.StringPtr("veeam"),
			}

			createDirectorSitesOptions := &vmwarev1.CreateDirectorSitesOptions{
				Name: core.StringPtr("my_director_site"),
				Pvdcs: []vmwarev1.PVDCPrototype{*pvdcPrototypeModel},
				ResourceGroup: resourceGroupIdentityModel,
				Services: []vmwarev1.ServiceIdentity{*serviceIdentityModel},
				PrivateOnly: core.BoolPtr(true),
				ConsoleConnectionType: core.StringPtr("private"),
				IpAllowList: []string{"1.1.1.1/24", "2.2.2.2/24"},
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSites - List Cloud Director site instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
			listDirectorSitesOptions := &vmwarev1.ListDirectorSitesOptions{
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSiteCollection, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteCollection).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSite - Get a Cloud Director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
			getDirectorSiteOptions := &vmwarev1.GetDirectorSiteOptions{
				ID: core.StringPtr("site_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`EnableVeeamOnPvdcsList - Enable or disable Veeam on a Cloud Director site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`EnableVeeamOnPvdcsList(enableVeeamOnPvdcsListOptions *EnableVeeamOnPvdcsListOptions)`, func() {
			enableVeeamOnPvdcsListOptions := &vmwarev1.EnableVeeamOnPvdcsListOptions{
				SiteID: core.StringPtr("site_id"),
				Enable: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			serviceEnabled, response, err := vmwareService.EnableVeeamOnPvdcsList(enableVeeamOnPvdcsListOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceEnabled).ToNot(BeNil())
		})
	})

	Describe(`EnableVcdaOnDataCenter - Enable or disable VCDA on a Cloud Director site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`EnableVcdaOnDataCenter(enableVcdaOnDataCenterOptions *EnableVcdaOnDataCenterOptions)`, func() {
			enableVcdaOnDataCenterOptions := &vmwarev1.EnableVcdaOnDataCenterOptions{
				SiteID: core.StringPtr("site_id"),
				Enable: core.BoolPtr(true),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			serviceEnabled, response, err := vmwareService.EnableVcdaOnDataCenter(enableVcdaOnDataCenterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(serviceEnabled).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesVcdaConnectionEndpoints - Create a VCDA connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesVcdaConnectionEndpoints(createDirectorSitesVcdaConnectionEndpointsOptions *CreateDirectorSitesVcdaConnectionEndpointsOptions)`, func() {
			createDirectorSitesVcdaConnectionEndpointsOptions := &vmwarev1.CreateDirectorSitesVcdaConnectionEndpointsOptions{
				SiteID: core.StringPtr("site_id"),
				Type: core.StringPtr("private"),
				DataCenterName: core.StringPtr("dal10"),
				AllowList: []string{"1.1.1.1"},
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			vcdaConnection, response, err := vmwareService.CreateDirectorSitesVcdaConnectionEndpoints(createDirectorSitesVcdaConnectionEndpointsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaConnection).ToNot(BeNil())
		})
	})

	Describe(`UpdateDirectorSitesVcdaConnectionEndpoints - Update VCDA connection allowlist`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDirectorSitesVcdaConnectionEndpoints(updateDirectorSitesVcdaConnectionEndpointsOptions *UpdateDirectorSitesVcdaConnectionEndpointsOptions)`, func() {
			updateDirectorSitesVcdaConnectionEndpointsOptions := &vmwarev1.UpdateDirectorSitesVcdaConnectionEndpointsOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("vcda_connections_id"),
				AllowList: []string{"1.1.1.1/24", "2.2.2.2/24"},
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			updatedVcdaConnection, response, err := vmwareService.UpdateDirectorSitesVcdaConnectionEndpoints(updateDirectorSitesVcdaConnectionEndpointsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(updatedVcdaConnection).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesVcdaC2cConnection - Create a VCDA cloud-to-cloud connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesVcdaC2cConnection(createDirectorSitesVcdaC2cConnectionOptions *CreateDirectorSitesVcdaC2cConnectionOptions)`, func() {
			createDirectorSitesVcdaC2cConnectionOptions := &vmwarev1.CreateDirectorSitesVcdaC2cConnectionOptions{
				SiteID: core.StringPtr("site_id"),
				LocalDataCenterName: core.StringPtr("dal10"),
				LocalSiteName: core.StringPtr("ddirw002-gr80d10vcda"),
				PeerSiteName: core.StringPtr("dirw274t02vcda"),
				PeerRegion: core.StringPtr("jp-tok"),
				Note: core.StringPtr("Text of the note..."),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			vcdaC2c, response, err := vmwareService.CreateDirectorSitesVcdaC2cConnection(createDirectorSitesVcdaC2cConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaC2c).ToNot(BeNil())
		})
	})

	Describe(`UpdateDirectorSitesVcdaC2cConnection - Update note in the cloud-to-cloud connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDirectorSitesVcdaC2cConnection(updateDirectorSitesVcdaC2cConnectionOptions *UpdateDirectorSitesVcdaC2cConnectionOptions)`, func() {
			updateDirectorSitesVcdaC2cConnectionOptions := &vmwarev1.UpdateDirectorSitesVcdaC2cConnectionOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("connection_id"),
				Note: core.StringPtr("Text of the note..."),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			updatedVcdaC2c, response, err := vmwareService.UpdateDirectorSitesVcdaC2cConnection(updateDirectorSitesVcdaC2cConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updatedVcdaC2c).ToNot(BeNil())
		})
	})

	Describe(`GetOidcConfiguration - Get an OIDC configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetOidcConfiguration(getOidcConfigurationOptions *GetOidcConfigurationOptions)`, func() {
			getOidcConfigurationOptions := &vmwarev1.GetOidcConfigurationOptions{
				SiteID: core.StringPtr("site_id"),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			oidc, response, err := vmwareService.GetOidcConfiguration(getOidcConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(oidc).ToNot(BeNil())
		})
	})

	Describe(`SetOidcConfiguration - Set an OIDC configuration`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`SetOidcConfiguration(setOidcConfigurationOptions *SetOidcConfigurationOptions)`, func() {
			setOidcConfigurationOptions := &vmwarev1.SetOidcConfigurationOptions{
				SiteID: core.StringPtr("site_id"),
				ContentLength: core.Int64Ptr(int64(0)),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			oidc, response, err := vmwareService.SetOidcConfiguration(setOidcConfigurationOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(oidc).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcs - List the resource pools in a Cloud Director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions)`, func() {
			listDirectorSitesPvdcsOptions := &vmwarev1.ListDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("site_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			pvdcCollection, response, err := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdcCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesPvdcs - Create a resource pool instance in a specified Cloud Director site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterPrototypeModel := &vmwarev1.ClusterPrototype{
				Name: core.StringPtr("cluster_1"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
				FileShares: fileSharesPrototypeModel,
			}

			createDirectorSitesPvdcsOptions := &vmwarev1.CreateDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("site_id"),
				Name: core.StringPtr("pvdc-1"),
				DataCenterName: core.StringPtr("dal10"),
				Clusters: []vmwarev1.ClusterPrototype{*clusterPrototypeModel},
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			pvdc, response, err := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdc).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSitesPvdcs - Get the specified resource pool in a Cloud Director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions)`, func() {
			getDirectorSitesPvdcsOptions := &vmwarev1.GetDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("pvdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			pvdc, response, err := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdc).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcsClusters - List clusters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions)`, func() {
			listDirectorSitesPvdcsClustersOptions := &vmwarev1.ListDirectorSitesPvdcsClustersOptions{
				SiteID: core.StringPtr("site_id"),
				PvdcID: core.StringPtr("pvdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			clusterCollection, response, err := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(clusterCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesPvdcsClusters - Create a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions *CreateDirectorSitesPvdcsClustersOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			createDirectorSitesPvdcsClustersOptions := &vmwarev1.CreateDirectorSitesPvdcsClustersOptions{
				SiteID: core.StringPtr("site_id"),
				PvdcID: core.StringPtr("pvdc_id"),
				Name: core.StringPtr("cluster_1"),
				HostCount: core.Int64Ptr(int64(2)),
				HostProfile: core.StringPtr("BM_2S_20_CORES_192_GB"),
				FileShares: fileSharesPrototypeModel,
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			cluster, response, err := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(cluster).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorInstancesPvdcsCluster - Get a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions)`, func() {
			getDirectorInstancesPvdcsClusterOptions := &vmwarev1.GetDirectorInstancesPvdcsClusterOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("cluster_id"),
				PvdcID: core.StringPtr("pvdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			cluster, response, err := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(cluster).ToNot(BeNil())
		})
	})

	Describe(`UpdateDirectorSitesPvdcsCluster - Update a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions)`, func() {
			fileSharesPrototypeModel := &vmwarev1.FileSharesPrototype{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterPatchModel := &vmwarev1.ClusterPatch{
				FileShares: fileSharesPrototypeModel,
				HostCount: core.Int64Ptr(int64(2)),
			}
			clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateDirectorSitesPvdcsClusterOptions := &vmwarev1.UpdateDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("cluster_id"),
				PvdcID: core.StringPtr("pvdc_id"),
				Body: clusterPatchModelAsPatch,
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			updateCluster, response, err := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateCluster).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteRegions - List regions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions)`, func() {
			listDirectorSiteRegionsOptions := &vmwarev1.ListDirectorSiteRegionsOptions{
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSiteRegionCollection, response, err := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteRegionCollection).ToNot(BeNil())
		})
	})

	Describe(`ListMultitenantDirectorSites - Get all multitenant Cloud Director sites`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions *ListMultitenantDirectorSitesOptions)`, func() {
			listMultitenantDirectorSitesOptions := &vmwarev1.ListMultitenantDirectorSitesOptions{
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			multitenantDirectorSiteCollection, response, err := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(multitenantDirectorSiteCollection).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteHostProfiles - List host profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions)`, func() {
			listDirectorSiteHostProfilesOptions := &vmwarev1.ListDirectorSiteHostProfilesOptions{
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSiteHostProfileCollection, response, err := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteHostProfileCollection).ToNot(BeNil())
		})
	})

	Describe(`ListVdcs - List virtual data centers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
			listVdcsOptions := &vmwarev1.ListVdcsOptions{
				AcceptLanguage: core.StringPtr("en-us"),
			}

			vdcCollection, response, err := vmwareService.ListVdcs(listVdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdcCollection).ToNot(BeNil())
		})
	})

	Describe(`CreateVdc - Create a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
			vdcProviderTypeModel := &vmwarev1.VDCProviderType{
				Name: core.StringPtr("paygo"),
			}

			directorSitePvdcModel := &vmwarev1.DirectorSitePVDC{
				ID: core.StringPtr("pvdc_id"),
				ProviderType: vdcProviderTypeModel,
			}

			vdcDirectorSitePrototypeModel := &vmwarev1.VDCDirectorSitePrototype{
				ID: core.StringPtr("site_id"),
				Pvdc: directorSitePvdcModel,
			}

			vdcEdgePrototypeModel := &vmwarev1.VDCEdgePrototype{
				Size: core.StringPtr("medium"),
				Type: core.StringPtr("performance"),
			}

			resourceGroupIdentityModel := &vmwarev1.ResourceGroupIdentity{
				ID: core.StringPtr("some_resourcegroupid"),
			}

			createVdcOptions := &vmwarev1.CreateVdcOptions{
				Name: core.StringPtr("sampleVDC"),
				DirectorSite: vdcDirectorSitePrototypeModel,
				Edge: vdcEdgePrototypeModel,
				FastProvisioningEnabled: core.BoolPtr(true),
				ResourceGroup: resourceGroupIdentityModel,
				Cpu: core.Int64Ptr(int64(0)),
				Ram: core.Int64Ptr(int64(0)),
				RhelByol: core.BoolPtr(false),
				WindowsByol: core.BoolPtr(false),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			vdc, response, err := vmwareService.CreateVdc(createVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`GetVdc - Get a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
			getVdcOptions := &vmwarev1.GetVdcOptions{
				ID: core.StringPtr("vdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			vdc, response, err := vmwareService.GetVdc(getVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`UpdateVdc - Update a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`UpdateVdc(updateVdcOptions *UpdateVdcOptions)`, func() {
			vdcPatchModel := &vmwarev1.VDCPatch{
				Cpu: core.Int64Ptr(int64(0)),
				FastProvisioningEnabled: core.BoolPtr(true),
				Ram: core.Int64Ptr(int64(0)),
			}
			vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
			Expect(asPatchErr).To(BeNil())

			updateVdcOptions := &vmwarev1.UpdateVdcOptions{
				ID: core.StringPtr("vdc_id"),
				VDCPatch: vdcPatchModelAsPatch,
				AcceptLanguage: core.StringPtr("en-us"),
			}

			vdc, response, err := vmwareService.UpdateVdc(updateVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`AddTransitGatewayConnections - Add IBM Transit Gateway connections to edge`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`AddTransitGatewayConnections(addTransitGatewayConnectionsOptions *AddTransitGatewayConnectionsOptions)`, func() {
			addTransitGatewayConnectionsOptions := &vmwarev1.AddTransitGatewayConnectionsOptions{
				VdcID: core.StringPtr("vdc_id"),
				EdgeID: core.StringPtr("edge_id"),
				ID: core.StringPtr("transit_gateway_id"),
				ContentLength: core.Int64Ptr(int64(0)),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			transitGateway, response, err := vmwareService.AddTransitGatewayConnections(addTransitGatewayConnectionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(transitGateway).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSite - Delete a Cloud Director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
			deleteDirectorSiteOptions := &vmwarev1.DeleteDirectorSiteOptions{
				ID: core.StringPtr("site_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSitesVcdaConnectionEndpoints - Delete a VCDA connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSitesVcdaConnectionEndpoints(deleteDirectorSitesVcdaConnectionEndpointsOptions *DeleteDirectorSitesVcdaConnectionEndpointsOptions)`, func() {
			deleteDirectorSitesVcdaConnectionEndpointsOptions := &vmwarev1.DeleteDirectorSitesVcdaConnectionEndpointsOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("vcda_connections_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			vcdaConnection, response, err := vmwareService.DeleteDirectorSitesVcdaConnectionEndpoints(deleteDirectorSitesVcdaConnectionEndpointsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaConnection).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSitesVcdaC2cConnection - Delete a VCDA cloud-to-cloud connection`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSitesVcdaC2cConnection(deleteDirectorSitesVcdaC2cConnectionOptions *DeleteDirectorSitesVcdaC2cConnectionOptions)`, func() {
			deleteDirectorSitesVcdaC2cConnectionOptions := &vmwarev1.DeleteDirectorSitesVcdaC2cConnectionOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("connection_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			vcdaC2c, response, err := vmwareService.DeleteDirectorSitesVcdaC2cConnection(deleteDirectorSitesVcdaC2cConnectionOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vcdaC2c).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSitesPvdcsCluster - Delete a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions)`, func() {
			deleteDirectorSitesPvdcsClusterOptions := &vmwarev1.DeleteDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("site_id"),
				ID: core.StringPtr("cluster_id"),
				PvdcID: core.StringPtr("pvdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
				XGlobalTransactionID: core.StringPtr("transaction1"),
			}

			clusterSummary, response, err := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(clusterSummary).ToNot(BeNil())
		})
	})

	Describe(`DeleteVdc - Delete a virtual data center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
			deleteVdcOptions := &vmwarev1.DeleteVdcOptions{
				ID: core.StringPtr("vdc_id"),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			vdc, response, err := vmwareService.DeleteVdc(deleteVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`RemoveTransitGatewayConnections - Remove IBM Transit Gateway connections from edge`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`RemoveTransitGatewayConnections(removeTransitGatewayConnectionsOptions *RemoveTransitGatewayConnectionsOptions)`, func() {
			removeTransitGatewayConnectionsOptions := &vmwarev1.RemoveTransitGatewayConnectionsOptions{
				VdcID: core.StringPtr("vdc_id"),
				EdgeID: core.StringPtr("edge_id"),
				ID: core.StringPtr("transit_gateway_id"),
				AcceptLanguage: core.StringPtr("en-us"),
			}

			transitGateway, response, err := vmwareService.RemoveTransitGatewayConnections(removeTransitGatewayConnectionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(transitGateway).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
