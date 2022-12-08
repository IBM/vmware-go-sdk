// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwareasaserviceapiv1"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

/**
 * This file contains an integration test for the vmwareasaserviceapiv1 package.
 *
 * Notes:
 *
 * The integration test will automatically skip tests if the required config file is not available.
 */

var _ = Describe(`VMwareAsAServiceApiV1 Integration Tests`, func() {
	const externalConfigFile = "../v_mware_as_a_service_api_v1.env"

	var (
		err          error
		vMwareAsAServiceApiService *vmwareasaserviceapiv1.VMwareAsAServiceApiV1
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
			config, err = core.GetServiceProperties(vmwareasaserviceapiv1.DefaultServiceName)
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
			vMwareAsAServiceApiServiceOptions := &vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{}

			vMwareAsAServiceApiService, err = vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(vMwareAsAServiceApiServiceOptions)
			Expect(err).To(BeNil())
			Expect(vMwareAsAServiceApiService).ToNot(BeNil())
			Expect(vMwareAsAServiceApiService.Service.Options.URL).To(Equal(serviceURL))

			core.SetLogger(core.NewLogger(core.LevelDebug, log.New(GinkgoWriter, "", log.LstdFlags), log.New(GinkgoWriter, "", log.LstdFlags)))
			vMwareAsAServiceApiService.EnableRetries(4, 30*time.Second)
		})
	})

	Describe(`CreateDirectorSites - Create a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
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

			createDirectorSitesOptions := &vmwareasaserviceapiv1.CreateDirectorSitesOptions{
				Name: core.StringPtr("testString"),
				ResourceGroup: core.StringPtr("testString"),
				Pvdcs: []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSites - List director site instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
			listDirectorSitesOptions := &vmwareasaserviceapiv1.ListDirectorSitesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			listDirectorSites, response, err := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDirectorSites).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSite - Get a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
			getDirectorSiteOptions := &vmwareasaserviceapiv1.GetDirectorSiteOptions{
				SiteID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSite).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcs - List the provider virtual data centers in a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions)`, func() {
			listDirectorSitesPvdcsOptions := &vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			listPvdCs, response, err := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listPvdCs).ToNot(BeNil())
		})
	})

	Describe(`CreateDirectorSitesPvdcs - Create a provider virtual data center instance in a specified director site`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions)`, func() {
			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
			}

			clusterOrderInfoModel := &vmwareasaserviceapiv1.ClusterOrderInfo{
				Name: core.StringPtr("testString"),
				StorageType: core.StringPtr("nfs"),
				HostCount: core.Int64Ptr(int64(2)),
				FileShares: fileSharesModel,
				HostProfile: core.StringPtr("testString"),
			}

			createDirectorSitesPvdcsOptions := &vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				Name: core.StringPtr("testString"),
				DataCenter: core.StringPtr("testString"),
				Clusters: []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdcResponse, response, err := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdcResponse).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorSitesPvdcs - Get the specified provider virtual data center in a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions)`, func() {
			getDirectorSitesPvdcsOptions := &vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions{
				SiteID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdcSummary, response, err := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(pvdcSummary).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSitesPvdcsClusters - List clusters`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions)`, func() {
			listDirectorSitesPvdcsClustersOptions := &vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions{
				SiteID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			listClusters, response, err := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listClusters).ToNot(BeNil())
		})
	})

	Describe(`GetDirectorInstancesPvdcsCluster - Get a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions)`, func() {
			getDirectorInstancesPvdcsClusterOptions := &vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			cluster, response, err := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
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
			jsonPatchOperationModel := &vmwareasaserviceapiv1.JSONPatchOperation{
				Op: core.StringPtr("add"),
				Path: core.StringPtr("testString"),
				From: core.StringPtr("testString"),
				Value: core.StringPtr("testString"),
			}

			updateDirectorSitesPvdcsClusterOptions := &vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				Body: []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			updateClusterResponse, response, err := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(updateClusterResponse).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteRegions - List regions`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions)`, func() {
			listDirectorSiteRegionsOptions := &vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSiteRegions, response, err := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSiteRegions).ToNot(BeNil())
		})
	})

	Describe(`ListDirectorSiteHostProfiles - List host profiles`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions)`, func() {
			listDirectorSiteHostProfilesOptions := &vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			listHostProfiles, response, err := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listHostProfiles).ToNot(BeNil())
		})
	})

	Describe(`ReplaceOrgAdminPassword - Replace the password of VMware Cloud Director tenant portal`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions)`, func() {
			replaceOrgAdminPasswordOptions := &vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions{
				SiteID: core.StringPtr("testString"),
			}

			newPassword, response, err := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(newPassword).ToNot(BeNil())
		})
	})

	Describe(`ListPrices - List billing metrics`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListPrices(listPricesOptions *ListPricesOptions)`, func() {
			listPricesOptions := &vmwareasaserviceapiv1.ListPricesOptions{
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSitePricingInfo, response, err := vMwareAsAServiceApiService.ListPrices(listPricesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSitePricingInfo).ToNot(BeNil())
		})
	})

	Describe(`GetVcddPrice - Quote price`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions)`, func() {
			fileSharesModel := &vmwareasaserviceapiv1.FileShares{
				STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
				STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
				STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
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

			getVcddPriceOptions := &vmwareasaserviceapiv1.GetVcddPriceOptions{
				Name: core.StringPtr("testString"),
				ResourceGroup: core.StringPtr("testString"),
				Pvdcs: []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel},
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSitePriceQuoteResponse, response, err := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(201))
			Expect(directorSitePriceQuoteResponse).ToNot(BeNil())
		})
	})

	Describe(`ListVdcs - List Virtual Data Centers`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
			listVdcsOptions := &vmwareasaserviceapiv1.ListVdcsOptions{
				AcceptLanguage: core.StringPtr("testString"),
			}

			listVdCs, response, err := vMwareAsAServiceApiService.ListVdcs(listVdcsOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listVdCs).ToNot(BeNil())
		})
	})

	Describe(`CreateVdc - Create a Virtual Data Center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
			vdcDirectorSiteClusterModel := &vmwareasaserviceapiv1.VDCDirectorSiteCluster{
				ID: core.StringPtr("testString"),
			}

			newVdcDirectorSiteModel := &vmwareasaserviceapiv1.NewVDCDirectorSite{
				ID: core.StringPtr("testString"),
				Cluster: vdcDirectorSiteClusterModel,
			}

			newVdcEdgeModel := &vmwareasaserviceapiv1.NewVDCEdge{
				Size: core.StringPtr("medium"),
				Type: core.StringPtr("dedicated"),
			}

			newVdcResourceGroupModel := &vmwareasaserviceapiv1.NewVDCResourceGroup{
				ID: core.StringPtr("testString"),
			}

			createVdcOptions := &vmwareasaserviceapiv1.CreateVdcOptions{
				Name: core.StringPtr("testString"),
				DirectorSite: newVdcDirectorSiteModel,
				Edge: newVdcEdgeModel,
				ResourceGroup: newVdcResourceGroupModel,
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vMwareAsAServiceApiService.CreateVdc(createVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`GetVdc - Get a Virtual Data Center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
			getVdcOptions := &vmwareasaserviceapiv1.GetVdcOptions{
				VdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vMwareAsAServiceApiService.GetVdc(getVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`DeleteVdc - Delete a Virtual Data Center`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
			deleteVdcOptions := &vmwareasaserviceapiv1.DeleteVdcOptions{
				VdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
			}

			vdc, response, err := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(vdc).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSitesPvdcsCluster - Delete a cluster`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions)`, func() {
			deleteDirectorSitesPvdcsClusterOptions := &vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions{
				SiteID: core.StringPtr("testString"),
				ClusterID: core.StringPtr("testString"),
				PvdcID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			pvdcResponse, response, err := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(pvdcResponse).ToNot(BeNil())
		})
	})

	Describe(`DeleteDirectorSite - Delete a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
			deleteDirectorSiteOptions := &vmwareasaserviceapiv1.DeleteDirectorSiteOptions{
				SiteID: core.StringPtr("testString"),
				AcceptLanguage: core.StringPtr("testString"),
				XGlobalTransactionID: core.StringPtr("testString"),
			}

			directorSite, response, err := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})
})

//
// Utility functions are declared in the unit test file
//
