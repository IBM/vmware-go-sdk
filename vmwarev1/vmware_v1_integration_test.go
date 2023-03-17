//go:build integration
// +build integration

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
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
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
		err           error
		vmwareService *vmwarev1.VmwareV1
		serviceURL    string
		config        map[string]string
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

	//Describe(`CreateDirectorSites - Create a director site instance`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
	//		fileSharesModel := &vmwarev1.FileShares{
	//			STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
	//		}
	//
	//		clusterOrderInfoModel := &vmwarev1.ClusterOrderInfo{
	//			Name: core.StringPtr("testString"),
	//			HostCount: core.Int64Ptr(int64(2)),
	//			FileShares: fileSharesModel,
	//			HostProfile: core.StringPtr("testString"),
	//		}
	//
	//		pvdcOrderInfoModel := &vmwarev1.PVDCOrderInfo{
	//			Name: core.StringPtr("testString"),
	//			DataCenter: core.StringPtr("testString"),
	//			Clusters: []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel},
	//		}
	//
	//		createDirectorSitesOptions := &vmwarev1.CreateDirectorSitesOptions{
	//			XAuthRefreshToken: core.StringPtr("testString"),
	//			Name: core.StringPtr("testString"),
	//			ResourceGroup: core.StringPtr("testString"),
	//			Pvdcs: []vmwarev1.PVDCOrderInfo{*pvdcOrderInfoModel},
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(directorSite).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListDirectorSites - List director site instances`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
	//		listDirectorSitesOptions := &vmwarev1.ListDirectorSitesOptions{
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		listDirectorSites, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(listDirectorSites).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`GetDirectorSite - Get a director site instance`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
	//		getDirectorSiteOptions := &vmwarev1.GetDirectorSiteOptions{
	//			ID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(directorSite).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListDirectorSitesPvdcs - List the provider virtual data centers in a director site instance`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions)`, func() {
	//		listDirectorSitesPvdcsOptions := &vmwarev1.ListDirectorSitesPvdcsOptions{
	//			SiteID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		listPvdCs, response, err := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(listPvdCs).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`CreateDirectorSitesPvdcs - Create a provider virtual data center instance in a specified director site`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions)`, func() {
	//		fileSharesModel := &vmwarev1.FileShares{
	//			STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
	//		}
	//
	//		clusterOrderInfoModel := &vmwarev1.ClusterOrderInfo{
	//			Name: core.StringPtr("testString"),
	//			HostCount: core.Int64Ptr(int64(2)),
	//			FileShares: fileSharesModel,
	//			HostProfile: core.StringPtr("testString"),
	//		}
	//
	//		createDirectorSitesPvdcsOptions := &vmwarev1.CreateDirectorSitesPvdcsOptions{
	//			SiteID: core.StringPtr("testString"),
	//			XAuthRefreshToken: core.StringPtr("testString"),
	//			Name: core.StringPtr("testString"),
	//			DataCenter: core.StringPtr("testString"),
	//			Clusters: []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel},
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		pvdcResponse, response, err := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(pvdcResponse).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`GetDirectorSitesPvdcs - Get the specified provider virtual data center in a director site instance`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions)`, func() {
	//		getDirectorSitesPvdcsOptions := &vmwarev1.GetDirectorSitesPvdcsOptions{
	//			SiteID: core.StringPtr("testString"),
	//			ID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		pvdcResponse, response, err := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(pvdcResponse).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListDirectorSitesPvdcsClusters - List clusters`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions)`, func() {
	//		listDirectorSitesPvdcsClustersOptions := &vmwarev1.ListDirectorSitesPvdcsClustersOptions{
	//			SiteID: core.StringPtr("testString"),
	//			PvdcID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		listClusters, response, err := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(listClusters).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`GetDirectorInstancesPvdcsCluster - Get a cluster`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions)`, func() {
	//		getDirectorInstancesPvdcsClusterOptions := &vmwarev1.GetDirectorInstancesPvdcsClusterOptions{
	//			SiteID: core.StringPtr("testString"),
	//			ID: core.StringPtr("testString"),
	//			PvdcID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		cluster, response, err := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(cluster).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`UpdateDirectorSitesPvdcsCluster - Update a cluster`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions)`, func() {
	//		jsonPatchOperationModel := &vmwarev1.JSONPatchOperation{
	//			Op: core.StringPtr("add"),
	//			Path: core.StringPtr("testString"),
	//			From: core.StringPtr("testString"),
	//			Value: core.StringPtr("testString"),
	//		}
	//
	//		updateDirectorSitesPvdcsClusterOptions := &vmwarev1.UpdateDirectorSitesPvdcsClusterOptions{
	//			SiteID: core.StringPtr("testString"),
	//			ID: core.StringPtr("testString"),
	//			PvdcID: core.StringPtr("testString"),
	//			XAuthRefreshToken: core.StringPtr("testString"),
	//			Body: []vmwarev1.JSONPatchOperation{*jsonPatchOperationModel},
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		updateClusterResponse, response, err := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(updateClusterResponse).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListDirectorSiteRegions - List regions`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions)`, func() {
	//		listDirectorSiteRegionsOptions := &vmwarev1.ListDirectorSiteRegionsOptions{
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSiteRegions, response, err := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(directorSiteRegions).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListDirectorSiteHostProfiles - List host profiles`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions)`, func() {
	//		listDirectorSiteHostProfilesOptions := &vmwarev1.ListDirectorSiteHostProfilesOptions{
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		listHostProfiles, response, err := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(listHostProfiles).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ReplaceOrgAdminPassword - Replace the password of VMware Cloud Director tenant portal`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions)`, func() {
	//		replaceOrgAdminPasswordOptions := &vmwarev1.ReplaceOrgAdminPasswordOptions{
	//			SiteID: core.StringPtr("testString"),
	//		}
	//
	//		newPassword, response, err := vmwareService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(newPassword).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListPrices - List billing metrics`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListPrices(listPricesOptions *ListPricesOptions)`, func() {
	//		listPricesOptions := &vmwarev1.ListPricesOptions{
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSitePricingInfo, response, err := vmwareService.ListPrices(listPricesOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(directorSitePricingInfo).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`GetVcddPrice - Quote price`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions)`, func() {
	//		fileSharesModel := &vmwarev1.FileShares{
	//			STORAGEPOINTTWOFIVEIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETWOIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGEFOURIOPSGB: core.Int64Ptr(int64(0)),
	//			STORAGETENIOPSGB: core.Int64Ptr(int64(0)),
	//		}
	//
	//		clusterOrderInfoModel := &vmwarev1.ClusterOrderInfo{
	//			Name: core.StringPtr("testString"),
	//			HostCount: core.Int64Ptr(int64(2)),
	//			FileShares: fileSharesModel,
	//			HostProfile: core.StringPtr("testString"),
	//		}
	//
	//		pvdcOrderInfoModel := &vmwarev1.PVDCOrderInfo{
	//			Name: core.StringPtr("testString"),
	//			DataCenter: core.StringPtr("testString"),
	//			Clusters: []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel},
	//		}
	//
	//		getVcddPriceOptions := &vmwarev1.GetVcddPriceOptions{
	//			Name: core.StringPtr("testString"),
	//			ResourceGroup: core.StringPtr("testString"),
	//			Pvdcs: []vmwarev1.PVDCOrderInfo{*pvdcOrderInfoModel},
	//			Country: core.StringPtr("USA"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSitePriceQuoteResponse, response, err := vmwareService.GetVcddPrice(getVcddPriceOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(201))
	//		Expect(directorSitePriceQuoteResponse).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`ListVdcs - List Virtual Data Centers`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
	//		listVdcsOptions := &vmwarev1.ListVdcsOptions{
	//			AcceptLanguage: core.StringPtr("testString"),
	//		}
	//
	//		listVdCs, response, err := vmwareService.ListVdcs(listVdcsOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(listVdCs).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`CreateVdc - Create a Virtual Data Center`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
	//		directorSitePvdcModel := &vmwarev1.DirectorSitePVDC{
	//			ID: core.StringPtr("testString"),
	//		}
	//
	//		newVdcDirectorSiteModel := &vmwarev1.NewVDCDirectorSite{
	//			ID: core.StringPtr("testString"),
	//			Pvdc: directorSitePvdcModel,
	//		}
	//
	//		newVdcEdgeModel := &vmwarev1.NewVDCEdge{
	//			Size: core.StringPtr("medium"),
	//			Type: core.StringPtr("dedicated"),
	//		}
	//
	//		newVdcResourceGroupModel := &vmwarev1.NewVDCResourceGroup{
	//			ID: core.StringPtr("testString"),
	//		}
	//
	//		createVdcOptions := &vmwarev1.CreateVdcOptions{
	//			Name: core.StringPtr("testString"),
	//			DirectorSite: newVdcDirectorSiteModel,
	//			Edge: newVdcEdgeModel,
	//			ResourceGroup: newVdcResourceGroupModel,
	//			AcceptLanguage: core.StringPtr("testString"),
	//		}
	//
	//		vdc, response, err := vmwareService.CreateVdc(createVdcOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(vdc).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`GetVdc - Get a Virtual Data Center`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
	//		getVdcOptions := &vmwarev1.GetVdcOptions{
	//			ID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//		}
	//
	//		vdc, response, err := vmwareService.GetVdc(getVdcOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(200))
	//		Expect(vdc).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`DeleteVdc - Delete a Virtual Data Center`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
	//		deleteVdcOptions := &vmwarev1.DeleteVdcOptions{
	//			ID: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//		}
	//
	//		vdc, response, err := vmwareService.DeleteVdc(deleteVdcOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(vdc).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`DeleteDirectorSitesPvdcsCluster - Delete a cluster`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions)`, func() {
	//		deleteDirectorSitesPvdcsClusterOptions := &vmwarev1.DeleteDirectorSitesPvdcsClusterOptions{
	//			SiteID: core.StringPtr("testString"),
	//			ID: core.StringPtr("testString"),
	//			PvdcID: core.StringPtr("testString"),
	//			XAuthRefreshToken: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		clusterSummary, response, err := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(clusterSummary).ToNot(BeNil())
	//	})
	//})
	//
	//Describe(`DeleteDirectorSite - Delete a director site instance`, func() {
	//	BeforeEach(func() {
	//		shouldSkipTest()
	//	})
	//	It(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
	//		deleteDirectorSiteOptions := &vmwarev1.DeleteDirectorSiteOptions{
	//			ID: core.StringPtr("testString"),
	//			XAuthRefreshToken: core.StringPtr("testString"),
	//			AcceptLanguage: core.StringPtr("testString"),
	//			XGlobalTransactionID: core.StringPtr("testString"),
	//		}
	//
	//		directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
	//		Expect(err).To(BeNil())
	//		Expect(response.StatusCode).To(Equal(202))
	//		Expect(directorSite).ToNot(BeNil())
	//	})
	//})
})

//
// Utility functions are declared in the unit test file
//
