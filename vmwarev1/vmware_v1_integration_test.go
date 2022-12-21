//go:build integration
// +build integration

package vmwarev1_test

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

var _ = Describe("vmware integration test", Ordered, func() {
	const externalConfigFile = "../vmware_v1.env"
	createdInstanceID := "3a54fa9c-3f79-495f-922b-cbb8bab1b31e"

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

	Describe(`CreateDirectorSites - Create a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
			fileSharesModel := &vmwarev1.FileShares{
				STORAGETWOIOPSGB: core.Int64Ptr(int64(24000)),
			}

			clusterOrderInfoModel := &vmwarev1.ClusterOrderInfo{
				Name:        core.StringPtr("cluster_1"),
				StorageType: core.StringPtr("nfs"),
				HostCount:   core.Int64Ptr(int64(2)),
				FileShares:  fileSharesModel,
				HostProfile: core.StringPtr("BM_2S_32_CORES_192_GB"),
			}

			pvdcOrderInfoModel := &vmwarev1.PVDCOrderInfo{
				Name:       core.StringPtr("pvdc_1"),
				DataCenter: core.StringPtr("tok02"),
				Clusters:   []vmwarev1.ClusterOrderInfo{*clusterOrderInfoModel},
			}

			createDirectorSitesOptions := &vmwarev1.CreateDirectorSitesOptions{
				Name:                core.StringPtr("sdk_test_4"),
				ResourceGroup:       core.StringPtr("Default"),
				Pvdcs:               []vmwarev1.PVDCOrderInfo{*pvdcOrderInfoModel},
				IBMAuthRefreshToken: core.StringPtr(config["AUTH_REFRESH_TOKEN"]),
			}

			res2B, _ := json.Marshal(createDirectorSitesOptions)
			fmt.Println(string(res2B))

			directorSite, response, err := vmwareService.CreateDirectorSites(createDirectorSitesOptions)
			result, _ := json.Marshal(directorSite)
			fmt.Println(string(result))

			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
			Expect(directorSite.ID).ToNot(BeNil())
			createdInstanceID = *directorSite.ID
		})
	})

	Describe(`ListDirectorSites - List director site instances`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
			listDirectorSitesOptions := &vmwarev1.ListDirectorSitesOptions{}

			listDirectorSites, response, err := vmwareService.ListDirectorSites(listDirectorSitesOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(listDirectorSites).ToNot(BeNil())

			created := vmwarev1.DirectorSite{}
			for _, ds := range listDirectorSites.DirectorSites {
				if *(ds.ID) == createdInstanceID {
					created = ds
				}
			}
			Expect(*(created.ID)).To(Equal(createdInstanceID))
		})
	})

	Describe(`GetDirectorSite - Get a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
			getDirectorSiteOptions := &vmwarev1.GetDirectorSiteOptions{
				SiteID: core.StringPtr(createdInstanceID),
			}

			directorSite, response, err := vmwareService.GetDirectorSite(getDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(200))
			Expect(directorSite).ToNot(BeNil())
			Expect(*(directorSite.Status)).To(Equal(vmwarev1.DirectorSite_Status_Creating))
		})
	})

	Describe(`DeleteDirectorSite - Delete a director site instance`, func() {
		BeforeEach(func() {
			shouldSkipTest()
		})
		It(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
			deleteDirectorSiteOptions := &vmwarev1.DeleteDirectorSiteOptions{
				SiteID:              core.StringPtr("e42668d0-7f1f-434b-a710-e96a15de15e6"),
				IBMAuthRefreshToken: core.StringPtr(config["AUTH_REFRESH_TOKEN"]),
			}

			directorSite, response, err := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptions)
			Expect(err).To(BeNil())
			Expect(response.StatusCode).To(Equal(202))
			Expect(directorSite).ToNot(BeNil())
		})
	})

})
