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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/IBM/vmware-go-sdk/vmwareasaserviceapiv1"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe(`VMwareAsAServiceApiV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(vMwareAsAServiceApiService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(vMwareAsAServiceApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				URL: "https://vmwareasaserviceapiv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(vMwareAsAServiceApiService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"V_MWARE_AS_A_SERVICE_API_URL": "https://vmwareasaserviceapiv1/api",
				"V_MWARE_AS_A_SERVICE_API_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				})
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := vMwareAsAServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vMwareAsAServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vMwareAsAServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vMwareAsAServiceApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL: "https://testService/api",
				})
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vMwareAsAServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vMwareAsAServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vMwareAsAServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vMwareAsAServiceApiService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				})
				err := vMwareAsAServiceApiService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vMwareAsAServiceApiService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vMwareAsAServiceApiService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vMwareAsAServiceApiService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vMwareAsAServiceApiService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"V_MWARE_AS_A_SERVICE_API_URL": "https://vmwareasaserviceapiv1/api",
				"V_MWARE_AS_A_SERVICE_API_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
			})

			It(`Instantiate service client with error`, func() {
				Expect(vMwareAsAServiceApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"V_MWARE_AS_A_SERVICE_API_AUTH_TYPE":   "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1UsingExternalConfig(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(vMwareAsAServiceApiService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = vmwareasaserviceapiv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions) - Operation response error`, func() {
		createDirectorSitesPath := "/director_sites"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDirectorSites with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.ResourceGroup = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions)`, func() {
		createDirectorSitesPath := "/director_sites"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke CreateDirectorSites successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.ResourceGroup = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesWithContext(ctx, createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.CreateDirectorSitesWithContext(ctx, createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke CreateDirectorSites successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.ResourceGroup = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDirectorSites with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.ResourceGroup = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDirectorSitesOptions model with no property values
				createDirectorSitesOptionsModelNew := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateDirectorSites successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.ResourceGroup = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions) - Operation response error`, func() {
		listDirectorSitesPath := "/director_sites"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDirectorSites with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions)`, func() {
		listDirectorSitesPath := "/director_sites"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_sites": [{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSites successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListDirectorSitesWithContext(ctx, listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListDirectorSitesWithContext(ctx, listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_sites": [{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSites successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSites with error: Operation request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDirectorSites successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions) - Operation response error`, func() {
		getDirectorSitePath := "/director_sites/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDirectorSite with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions)`, func() {
		getDirectorSitePath := "/director_sites/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke GetDirectorSite successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.GetDirectorSiteWithContext(ctx, getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.GetDirectorSiteWithContext(ctx, getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke GetDirectorSite successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorSite with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorSiteOptions model with no property values
				getDirectorSiteOptionsModelNew := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDirectorSite successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions) - Operation response error`, func() {
		deleteDirectorSitePath := "/director_sites/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitePath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDirectorSite with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions)`, func() {
		deleteDirectorSitePath := "/director_sites/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke DeleteDirectorSite successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.DeleteDirectorSiteWithContext(ctx, deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.DeleteDirectorSiteWithContext(ctx, deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitePath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "Creating", "resource_group": "ResourceGroup", "creator": "Creator", "resource_group_id": "ResourceGroupID", "resource_group_crn": "ResourceGroupCrn", "pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke DeleteDirectorSite successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDirectorSite with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDirectorSiteOptions model with no property values
				deleteDirectorSiteOptionsModelNew := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteDirectorSite successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions) - Operation response error`, func() {
		listDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions)`, func() {
		listDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsWithContext(ctx, listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcsWithContext(ctx, listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pvdcs": [{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDirectorSitesPvdcsOptions model with no property values
				listDirectorSitesPvdcsOptionsModelNew := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions) - Operation response error`, func() {
		createDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenter = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions)`, func() {
		createDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}], "status": "Creating"}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenter = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcsWithContext(ctx, createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.CreateDirectorSitesPvdcsWithContext(ctx, createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}], "status": "Creating"}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenter = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenter = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDirectorSitesPvdcsOptions model with no property values
				createDirectorSitesPvdcsOptionsModelNew := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenter = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions) - Operation response error`, func() {
		getDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions)`, func() {
		getDirectorSitesPvdcsPath := "/director_sites/testString/pvdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}`)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcsWithContext(ctx, getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.GetDirectorSitesPvdcsWithContext(ctx, getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorSitesPvdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}]}`)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorSitesPvdcsOptions model with no property values
				getDirectorSitesPvdcsOptionsModelNew := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwareasaserviceapiv1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions) - Operation response error`, func() {
		listDirectorSitesPvdcsClustersPath := "/director_sites/testString/pvdcs/testString/clusters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsClustersPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions)`, func() {
		listDirectorSitesPvdcsClustersPath := "/director_sites/testString/pvdcs/testString/clusters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsClustersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "href": "Href", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "pvdc_id": "PvdcID", "director_site": "DirectorSite", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClustersWithContext(ctx, listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcsClustersWithContext(ctx, listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSitesPvdcsClustersPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "href": "Href", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "pvdc_id": "PvdcID", "director_site": "DirectorSite", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSitesPvdcsClusters with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDirectorSitesPvdcsClustersOptions model with no property values
				listDirectorSitesPvdcsClustersOptionsModelNew := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions) - Operation response error`, func() {
		getDirectorInstancesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorInstancesPvdcsClusterPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions)`, func() {
		getDirectorInstancesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorInstancesPvdcsClusterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "pvdc_id": "PvdcID", "director_site": "DirectorSite", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsClusterWithContext(ctx, getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.GetDirectorInstancesPvdcsClusterWithContext(ctx, getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getDirectorInstancesPvdcsClusterPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "instance_ordered": "2019-01-01T12:00:00.000Z", "instance_created": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "pvdc_id": "PvdcID", "director_site": "DirectorSite", "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"anyKey": "anyValue"}}`)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorInstancesPvdcsCluster with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorInstancesPvdcsClusterOptions model with no property values
				getDirectorInstancesPvdcsClusterOptionsModelNew := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions) - Operation response error`, func() {
		deleteDirectorSitesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions)`, func() {
		deleteDirectorSitesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}], "status": "Creating"}`)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsClusterWithContext(ctx, deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsClusterWithContext(ctx, deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center": "DataCenter", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "storage_type": "nfs", "host_count": 2, "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "host_profile": "HostProfile", "id": "ID", "href": "Href"}], "status": "Creating"}`)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDirectorSitesPvdcsClusterOptions model with no property values
				deleteDirectorSitesPvdcsClusterOptionsModelNew := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions) - Operation response error`, func() {
		updateDirectorSitesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions)`, func() {
		updateDirectorSitesPvdcsClusterPath := "/director_sites/testString/pvdcs/testString/clusters/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The request has been accepted."}`)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsClusterWithContext(ctx, updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsClusterWithContext(ctx, updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateDirectorSitesPvdcsClusterPath))
					Expect(req.Method).To(Equal("PATCH"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"message": "The request has been accepted."}`)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDirectorSitesPvdcsClusterOptions model with no property values
				updateDirectorSitesPvdcsClusterOptionsModelNew := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwareasaserviceapiv1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ClusterID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = []vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions) - Operation response error`, func() {
		listDirectorSiteRegionsPath := "/director_site_regions"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteRegionsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDirectorSiteRegions with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions)`, func() {
		listDirectorSiteRegionsPath := "/director_site_regions"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_regions": {"mapKey": {"data_centers": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}}}`)
				}))
			})
			It(`Invoke ListDirectorSiteRegions successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegionsWithContext(ctx, listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListDirectorSiteRegionsWithContext(ctx, listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteRegionsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_regions": {"mapKey": {"data_centers": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}}}`)
				}))
			})
			It(`Invoke ListDirectorSiteRegions successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSiteRegions with error: Operation request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDirectorSiteRegions successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions) - Operation response error`, func() {
		listDirectorSiteHostProfilesPath := "/director_site_host_profiles"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteHostProfilesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListDirectorSiteHostProfiles with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions)`, func() {
		listDirectorSiteHostProfilesPath := "/director_site_host_profiles"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteHostProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_host_profiles": [{"id": "ID", "cpu": 3, "family": "Family", "processor": "Processor", "ram": 3, "socket": 6, "speed": "Speed", "manufacturer": "Manufacturer", "features": ["Features"]}]}`)
				}))
			})
			It(`Invoke ListDirectorSiteHostProfiles successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfilesWithContext(ctx, listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListDirectorSiteHostProfilesWithContext(ctx, listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listDirectorSiteHostProfilesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_host_profiles": [{"id": "ID", "cpu": 3, "family": "Family", "processor": "Processor", "ram": 3, "socket": 6, "speed": "Speed", "manufacturer": "Manufacturer", "features": ["Features"]}]}`)
				}))
			})
			It(`Invoke ListDirectorSiteHostProfiles successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSiteHostProfiles with error: Operation request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListDirectorSiteHostProfiles successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwareasaserviceapiv1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) - Operation response error`, func() {
		replaceOrgAdminPasswordPath := "/director_site_password"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))
					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions)`, func() {
		replaceOrgAdminPasswordPath := "/director_site_password"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"password": "Password"}`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPasswordWithContext(ctx, replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ReplaceOrgAdminPasswordWithContext(ctx, replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(replaceOrgAdminPasswordPath))
					Expect(req.Method).To(Equal("PUT"))

					Expect(req.URL.Query()["site_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"password": "Password"}`)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ReplaceOrgAdminPassword with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ReplaceOrgAdminPasswordOptions model with no property values
				replaceOrgAdminPasswordOptionsModelNew := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ReplaceOrgAdminPassword successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				replaceOrgAdminPasswordOptionsModel := new(vmwareasaserviceapiv1.ReplaceOrgAdminPasswordOptions)
				replaceOrgAdminPasswordOptionsModel.SiteID = core.StringPtr("testString")
				replaceOrgAdminPasswordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrices(listPricesOptions *ListPricesOptions) - Operation response error`, func() {
		listPricesPath := "/director_site_pricing"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListPrices with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwareasaserviceapiv1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListPrices(listPricesOptions *ListPricesOptions)`, func() {
		listPricesPath := "/director_site_pricing"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_pricing": [{"metric": "Metric", "description": "Description", "price_list": [{"country": "Country", "currency": "Currency", "prices": [{"price": 5, "quantity_tier": 12}]}]}]}`)
				}))
			})
			It(`Invoke ListPrices successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwareasaserviceapiv1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListPricesWithContext(ctx, listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListPricesWithContext(ctx, listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listPricesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"director_site_pricing": [{"metric": "Metric", "description": "Description", "price_list": [{"country": "Country", "currency": "Currency", "prices": [{"price": 5, "quantity_tier": 12}]}]}]}`)
				}))
			})
			It(`Invoke ListPrices successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListPrices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwareasaserviceapiv1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListPrices with error: Operation request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwareasaserviceapiv1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListPrices successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := new(vmwareasaserviceapiv1.ListPricesOptions)
				listPricesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listPricesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listPricesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListPrices(listPricesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions) - Operation response error`, func() {
		getVcddPricePath := "/director_site_price_quote"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVcddPrice with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Name = core.StringPtr("testString")
				getVcddPriceOptionsModel.ResourceGroup = core.StringPtr("testString")
				getVcddPriceOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions)`, func() {
		getVcddPricePath := "/director_site_price_quote"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"base_charge": {"name": "Name", "currency": "Currency", "price": 5}, "clusters": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "count": 5, "currency": "Currency", "price": 5}]}]}], "currency": "Currency", "total": 5}`)
				}))
			})
			It(`Invoke GetVcddPrice successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Name = core.StringPtr("testString")
				getVcddPriceOptionsModel.ResourceGroup = core.StringPtr("testString")
				getVcddPriceOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.GetVcddPriceWithContext(ctx, getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.GetVcddPriceWithContext(ctx, getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVcddPricePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"base_charge": {"name": "Name", "currency": "Currency", "price": 5}, "clusters": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "currency": "Currency", "price": 5, "items": [{"name": "Name", "count": 5, "currency": "Currency", "price": 5}]}]}], "currency": "Currency", "total": 5}`)
				}))
			})
			It(`Invoke GetVcddPrice successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.GetVcddPrice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Name = core.StringPtr("testString")
				getVcddPriceOptionsModel.ResourceGroup = core.StringPtr("testString")
				getVcddPriceOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVcddPrice with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Name = core.StringPtr("testString")
				getVcddPriceOptionsModel.ResourceGroup = core.StringPtr("testString")
				getVcddPriceOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVcddPriceOptions model with no property values
				getVcddPriceOptionsModelNew := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke GetVcddPrice successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsModel := new(vmwareasaserviceapiv1.GetVcddPriceOptions)
				getVcddPriceOptionsModel.Name = core.StringPtr("testString")
				getVcddPriceOptionsModel.ResourceGroup = core.StringPtr("testString")
				getVcddPriceOptionsModel.Pvdcs = []vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}
				getVcddPriceOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVcddPriceOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getVcddPriceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.GetVcddPrice(getVcddPriceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVdcs(listVdcsOptions *ListVdcsOptions) - Operation response error`, func() {
		listVdcsPath := "/vdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVdcs with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwareasaserviceapiv1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVdcs(listVdcsOptions *ListVdcsOptions)`, func() {
		listVdcsPath := "/vdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vdcs": [{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwareasaserviceapiv1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVdcsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"vdcs": [{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.ListVdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwareasaserviceapiv1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVdcs with error: Operation request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwareasaserviceapiv1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVdcs successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwareasaserviceapiv1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVdc(createVdcOptions *CreateVdcOptions) - Operation response error`, func() {
		createVdcPath := "/vdcs"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateVdc with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwareasaserviceapiv1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateVdc(createVdcOptions *CreateVdcOptions)`, func() {
		createVdcPath := "/vdcs"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke CreateVdc successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwareasaserviceapiv1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.CreateVdcWithContext(ctx, createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.CreateVdcWithContext(ctx, createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createVdcPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke CreateVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.CreateVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwareasaserviceapiv1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVdc with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwareasaserviceapiv1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVdcOptions model with no property values
				createVdcOptionsModelNew := new(vmwareasaserviceapiv1.CreateVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke CreateVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				newVdcResourceGroupModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwareasaserviceapiv1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = newVdcDirectorSiteModel
				createVdcOptionsModel.Edge = newVdcEdgeModel
				createVdcOptionsModel.ResourceGroup = newVdcResourceGroupModel
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVdc(getVdcOptions *GetVdcOptions) - Operation response error`, func() {
		getVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVdc with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwareasaserviceapiv1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVdc(getVdcOptions *GetVdcOptions)`, func() {
		getVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke GetVdc successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwareasaserviceapiv1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.GetVdcWithContext(ctx, getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.GetVdcWithContext(ctx, getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVdcPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke GetVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.GetVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwareasaserviceapiv1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVdc with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwareasaserviceapiv1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVdcOptions model with no property values
				getVdcOptionsModelNew := new(vmwareasaserviceapiv1.GetVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.GetVdc(getVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwareasaserviceapiv1.GetVdcOptions)
				getVdcOptionsModel.VdcID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions) - Operation response error`, func() {
		deleteVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke DeleteVdc with error: Operation response processing error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vMwareAsAServiceApiService.EnableRetries(0, 0)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteVdc(deleteVdcOptions *DeleteVdcOptions)`, func() {
		deleteVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke DeleteVdc successfully with retries`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())
				vMwareAsAServiceApiService.EnableRetries(0, 0)

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vMwareAsAServiceApiService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vMwareAsAServiceApiService.DisableRetries()
				result, response, operationErr := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vMwareAsAServiceApiService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteVdcPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"id": "ID", "allocation_model": "paygo", "created_time": "2019-01-01T12:00:00.000Z", "crn": "Crn", "deleted_time": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "cluster": {"id": "ID"}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "type": "dedicated"}], "errors": [{"code": "Code", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_time": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "status": "Creating", "type": "dedicated"}`)
				}))
			})
			It(`Invoke DeleteVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vMwareAsAServiceApiService.DeleteVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteVdc with error: Operation validation and request error`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vMwareAsAServiceApiService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteVdcOptions model with no property values
				deleteVdcOptionsModelNew := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(202)
				}))
			})
			It(`Invoke DeleteVdc successfully`, func() {
				vMwareAsAServiceApiService, serviceErr := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vMwareAsAServiceApiService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwareasaserviceapiv1.DeleteVdcOptions)
				deleteVdcOptionsModel.VdcID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vMwareAsAServiceApiService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			vMwareAsAServiceApiService, _ := vmwareasaserviceapiv1.NewVMwareAsAServiceApiV1(&vmwareasaserviceapiv1.VMwareAsAServiceApiV1Options{
				URL:           "http://vmwareasaserviceapiv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewClusterOrderInfo successfully`, func() {
				name := "testString"
				storageType := "nfs"
				hostCount := int64(2)
				var fileShares *vmwareasaserviceapiv1.FileShares = nil
				hostProfile := "testString"
				_, err := vMwareAsAServiceApiService.NewClusterOrderInfo(name, storageType, hostCount, fileShares, hostProfile)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateDirectorSitesOptions successfully`, func() {
				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				Expect(fileSharesModel).ToNot(BeNil())
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				Expect(clusterOrderInfoModel).ToNot(BeNil())
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")
				Expect(clusterOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterOrderInfoModel.StorageType).To(Equal(core.StringPtr("nfs")))
				Expect(clusterOrderInfoModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterOrderInfoModel.FileShares).To(Equal(fileSharesModel))
				Expect(clusterOrderInfoModel.HostProfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				Expect(pvdcOrderInfoModel).ToNot(BeNil())
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				Expect(pvdcOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(pvdcOrderInfoModel.DataCenter).To(Equal(core.StringPtr("testString")))
				Expect(pvdcOrderInfoModel.Clusters).To(Equal([]vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}))

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsName := "testString"
				createDirectorSitesOptionsResourceGroup := "testString"
				createDirectorSitesOptionsPvdcs := []vmwareasaserviceapiv1.PVDCOrderInfo{}
				createDirectorSitesOptionsModel := vMwareAsAServiceApiService.NewCreateDirectorSitesOptions(createDirectorSitesOptionsName, createDirectorSitesOptionsResourceGroup, createDirectorSitesOptionsPvdcs)
				createDirectorSitesOptionsModel.SetName("testString")
				createDirectorSitesOptionsModel.SetResourceGroup("testString")
				createDirectorSitesOptionsModel.SetPvdcs([]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel})
				createDirectorSitesOptionsModel.SetAcceptLanguage("testString")
				createDirectorSitesOptionsModel.SetXGlobalTransactionID("testString")
				createDirectorSitesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDirectorSitesOptionsModel).ToNot(BeNil())
				Expect(createDirectorSitesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.Pvdcs).To(Equal([]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}))
				Expect(createDirectorSitesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDirectorSitesPvdcsOptions successfully`, func() {
				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				Expect(fileSharesModel).ToNot(BeNil())
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				Expect(clusterOrderInfoModel).ToNot(BeNil())
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")
				Expect(clusterOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterOrderInfoModel.StorageType).To(Equal(core.StringPtr("nfs")))
				Expect(clusterOrderInfoModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterOrderInfoModel.FileShares).To(Equal(fileSharesModel))
				Expect(clusterOrderInfoModel.HostProfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				siteID := "testString"
				createDirectorSitesPvdcsOptionsName := "testString"
				createDirectorSitesPvdcsOptionsDataCenter := "testString"
				createDirectorSitesPvdcsOptionsClusters := []vmwareasaserviceapiv1.ClusterOrderInfo{}
				createDirectorSitesPvdcsOptionsModel := vMwareAsAServiceApiService.NewCreateDirectorSitesPvdcsOptions(siteID, createDirectorSitesPvdcsOptionsName, createDirectorSitesPvdcsOptionsDataCenter, createDirectorSitesPvdcsOptionsClusters)
				createDirectorSitesPvdcsOptionsModel.SetSiteID("testString")
				createDirectorSitesPvdcsOptionsModel.SetName("testString")
				createDirectorSitesPvdcsOptionsModel.SetDataCenter("testString")
				createDirectorSitesPvdcsOptionsModel.SetClusters([]vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel})
				createDirectorSitesPvdcsOptionsModel.SetAcceptLanguage("testString")
				createDirectorSitesPvdcsOptionsModel.SetXGlobalTransactionID("testString")
				createDirectorSitesPvdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDirectorSitesPvdcsOptionsModel).ToNot(BeNil())
				Expect(createDirectorSitesPvdcsOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.DataCenter).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Clusters).To(Equal([]vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}))
				Expect(createDirectorSitesPvdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVdcOptions successfully`, func() {
				// Construct an instance of the VDCDirectorSiteCluster model
				vdcDirectorSiteClusterModel := new(vmwareasaserviceapiv1.VDCDirectorSiteCluster)
				Expect(vdcDirectorSiteClusterModel).ToNot(BeNil())
				vdcDirectorSiteClusterModel.ID = core.StringPtr("testString")
				Expect(vdcDirectorSiteClusterModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the NewVDCDirectorSite model
				newVdcDirectorSiteModel := new(vmwareasaserviceapiv1.NewVDCDirectorSite)
				Expect(newVdcDirectorSiteModel).ToNot(BeNil())
				newVdcDirectorSiteModel.ID = core.StringPtr("testString")
				newVdcDirectorSiteModel.Cluster = vdcDirectorSiteClusterModel
				Expect(newVdcDirectorSiteModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(newVdcDirectorSiteModel.Cluster).To(Equal(vdcDirectorSiteClusterModel))

				// Construct an instance of the NewVDCEdge model
				newVdcEdgeModel := new(vmwareasaserviceapiv1.NewVDCEdge)
				Expect(newVdcEdgeModel).ToNot(BeNil())
				newVdcEdgeModel.Size = core.StringPtr("medium")
				newVdcEdgeModel.Type = core.StringPtr("dedicated")
				Expect(newVdcEdgeModel.Size).To(Equal(core.StringPtr("medium")))
				Expect(newVdcEdgeModel.Type).To(Equal(core.StringPtr("dedicated")))

				// Construct an instance of the NewVDCResourceGroup model
				newVdcResourceGroupModel := new(vmwareasaserviceapiv1.NewVDCResourceGroup)
				Expect(newVdcResourceGroupModel).ToNot(BeNil())
				newVdcResourceGroupModel.ID = core.StringPtr("testString")
				Expect(newVdcResourceGroupModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsName := "testString"
				var createVdcOptionsDirectorSite *vmwareasaserviceapiv1.NewVDCDirectorSite = nil
				createVdcOptionsModel := vMwareAsAServiceApiService.NewCreateVdcOptions(createVdcOptionsName, createVdcOptionsDirectorSite)
				createVdcOptionsModel.SetName("testString")
				createVdcOptionsModel.SetDirectorSite(newVdcDirectorSiteModel)
				createVdcOptionsModel.SetEdge(newVdcEdgeModel)
				createVdcOptionsModel.SetResourceGroup(newVdcResourceGroupModel)
				createVdcOptionsModel.SetAcceptLanguage("testString")
				createVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVdcOptionsModel).ToNot(BeNil())
				Expect(createVdcOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.DirectorSite).To(Equal(newVdcDirectorSiteModel))
				Expect(createVdcOptionsModel.Edge).To(Equal(newVdcEdgeModel))
				Expect(createVdcOptionsModel.ResourceGroup).To(Equal(newVdcResourceGroupModel))
				Expect(createVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDirectorSiteOptions successfully`, func() {
				// Construct an instance of the DeleteDirectorSiteOptions model
				siteID := "testString"
				deleteDirectorSiteOptionsModel := vMwareAsAServiceApiService.NewDeleteDirectorSiteOptions(siteID)
				deleteDirectorSiteOptionsModel.SetSiteID("testString")
				deleteDirectorSiteOptionsModel.SetAcceptLanguage("testString")
				deleteDirectorSiteOptionsModel.SetXGlobalTransactionID("testString")
				deleteDirectorSiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDirectorSiteOptionsModel).ToNot(BeNil())
				Expect(deleteDirectorSiteOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDirectorSitesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				siteID := "testString"
				clusterID := "testString"
				pvdcID := "testString"
				deleteDirectorSitesPvdcsClusterOptionsModel := vMwareAsAServiceApiService.NewDeleteDirectorSitesPvdcsClusterOptions(siteID, clusterID, pvdcID)
				deleteDirectorSitesPvdcsClusterOptionsModel.SetSiteID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetClusterID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetPvdcID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVdcOptions successfully`, func() {
				// Construct an instance of the DeleteVdcOptions model
				vdcID := "testString"
				deleteVdcOptionsModel := vMwareAsAServiceApiService.NewDeleteVdcOptions(vdcID)
				deleteVdcOptionsModel.SetVdcID("testString")
				deleteVdcOptionsModel.SetAcceptLanguage("testString")
				deleteVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVdcOptionsModel).ToNot(BeNil())
				Expect(deleteVdcOptionsModel.VdcID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDirectorInstancesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				siteID := "testString"
				clusterID := "testString"
				pvdcID := "testString"
				getDirectorInstancesPvdcsClusterOptionsModel := vMwareAsAServiceApiService.NewGetDirectorInstancesPvdcsClusterOptions(siteID, clusterID, pvdcID)
				getDirectorInstancesPvdcsClusterOptionsModel.SetSiteID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetClusterID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetPvdcID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorInstancesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDirectorSiteOptions successfully`, func() {
				// Construct an instance of the GetDirectorSiteOptions model
				siteID := "testString"
				getDirectorSiteOptionsModel := vMwareAsAServiceApiService.NewGetDirectorSiteOptions(siteID)
				getDirectorSiteOptionsModel.SetSiteID("testString")
				getDirectorSiteOptionsModel.SetAcceptLanguage("testString")
				getDirectorSiteOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorSiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorSiteOptionsModel).ToNot(BeNil())
				Expect(getDirectorSiteOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDirectorSitesPvdcsOptions successfully`, func() {
				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				siteID := "testString"
				pvdcID := "testString"
				getDirectorSitesPvdcsOptionsModel := vMwareAsAServiceApiService.NewGetDirectorSitesPvdcsOptions(siteID, pvdcID)
				getDirectorSitesPvdcsOptionsModel.SetSiteID("testString")
				getDirectorSitesPvdcsOptionsModel.SetPvdcID("testString")
				getDirectorSitesPvdcsOptionsModel.SetAcceptLanguage("testString")
				getDirectorSitesPvdcsOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorSitesPvdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorSitesPvdcsOptionsModel).ToNot(BeNil())
				Expect(getDirectorSitesPvdcsOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVcddPriceOptions successfully`, func() {
				// Construct an instance of the FileShares model
				fileSharesModel := new(vmwareasaserviceapiv1.FileShares)
				Expect(fileSharesModel).ToNot(BeNil())
				fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterOrderInfo model
				clusterOrderInfoModel := new(vmwareasaserviceapiv1.ClusterOrderInfo)
				Expect(clusterOrderInfoModel).ToNot(BeNil())
				clusterOrderInfoModel.Name = core.StringPtr("testString")
				clusterOrderInfoModel.StorageType = core.StringPtr("nfs")
				clusterOrderInfoModel.HostCount = core.Int64Ptr(int64(2))
				clusterOrderInfoModel.FileShares = fileSharesModel
				clusterOrderInfoModel.HostProfile = core.StringPtr("testString")
				Expect(clusterOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterOrderInfoModel.StorageType).To(Equal(core.StringPtr("nfs")))
				Expect(clusterOrderInfoModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterOrderInfoModel.FileShares).To(Equal(fileSharesModel))
				Expect(clusterOrderInfoModel.HostProfile).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the PVDCOrderInfo model
				pvdcOrderInfoModel := new(vmwareasaserviceapiv1.PVDCOrderInfo)
				Expect(pvdcOrderInfoModel).ToNot(BeNil())
				pvdcOrderInfoModel.Name = core.StringPtr("testString")
				pvdcOrderInfoModel.DataCenter = core.StringPtr("testString")
				pvdcOrderInfoModel.Clusters = []vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}
				Expect(pvdcOrderInfoModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(pvdcOrderInfoModel.DataCenter).To(Equal(core.StringPtr("testString")))
				Expect(pvdcOrderInfoModel.Clusters).To(Equal([]vmwareasaserviceapiv1.ClusterOrderInfo{*clusterOrderInfoModel}))

				// Construct an instance of the GetVcddPriceOptions model
				getVcddPriceOptionsName := "testString"
				getVcddPriceOptionsResourceGroup := "testString"
				getVcddPriceOptionsPvdcs := []vmwareasaserviceapiv1.PVDCOrderInfo{}
				getVcddPriceOptionsModel := vMwareAsAServiceApiService.NewGetVcddPriceOptions(getVcddPriceOptionsName, getVcddPriceOptionsResourceGroup, getVcddPriceOptionsPvdcs)
				getVcddPriceOptionsModel.SetName("testString")
				getVcddPriceOptionsModel.SetResourceGroup("testString")
				getVcddPriceOptionsModel.SetPvdcs([]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel})
				getVcddPriceOptionsModel.SetAcceptLanguage("testString")
				getVcddPriceOptionsModel.SetXGlobalTransactionID("testString")
				getVcddPriceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVcddPriceOptionsModel).ToNot(BeNil())
				Expect(getVcddPriceOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.ResourceGroup).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.Pvdcs).To(Equal([]vmwareasaserviceapiv1.PVDCOrderInfo{*pvdcOrderInfoModel}))
				Expect(getVcddPriceOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getVcddPriceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVdcOptions successfully`, func() {
				// Construct an instance of the GetVdcOptions model
				vdcID := "testString"
				getVdcOptionsModel := vMwareAsAServiceApiService.NewGetVdcOptions(vdcID)
				getVdcOptionsModel.SetVdcID("testString")
				getVdcOptionsModel.SetAcceptLanguage("testString")
				getVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVdcOptionsModel).ToNot(BeNil())
				Expect(getVdcOptionsModel.VdcID).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewJSONPatchOperation successfully`, func() {
				op := "add"
				path := "testString"
				_model, err := vMwareAsAServiceApiService.NewJSONPatchOperation(op, path)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewListDirectorSiteHostProfilesOptions successfully`, func() {
				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := vMwareAsAServiceApiService.NewListDirectorSiteHostProfilesOptions()
				listDirectorSiteHostProfilesOptionsModel.SetAcceptLanguage("testString")
				listDirectorSiteHostProfilesOptionsModel.SetXGlobalTransactionID("testString")
				listDirectorSiteHostProfilesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDirectorSiteHostProfilesOptionsModel).ToNot(BeNil())
				Expect(listDirectorSiteHostProfilesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSiteHostProfilesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDirectorSiteRegionsOptions successfully`, func() {
				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := vMwareAsAServiceApiService.NewListDirectorSiteRegionsOptions()
				listDirectorSiteRegionsOptionsModel.SetAcceptLanguage("testString")
				listDirectorSiteRegionsOptionsModel.SetXGlobalTransactionID("testString")
				listDirectorSiteRegionsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDirectorSiteRegionsOptionsModel).ToNot(BeNil())
				Expect(listDirectorSiteRegionsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSiteRegionsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSiteRegionsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDirectorSitesOptions successfully`, func() {
				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := vMwareAsAServiceApiService.NewListDirectorSitesOptions()
				listDirectorSitesOptionsModel.SetAcceptLanguage("testString")
				listDirectorSitesOptionsModel.SetXGlobalTransactionID("testString")
				listDirectorSitesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDirectorSitesOptionsModel).ToNot(BeNil())
				Expect(listDirectorSitesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDirectorSitesPvdcsClustersOptions successfully`, func() {
				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				siteID := "testString"
				pvdcID := "testString"
				listDirectorSitesPvdcsClustersOptionsModel := vMwareAsAServiceApiService.NewListDirectorSitesPvdcsClustersOptions(siteID, pvdcID)
				listDirectorSitesPvdcsClustersOptionsModel.SetSiteID("testString")
				listDirectorSitesPvdcsClustersOptionsModel.SetPvdcID("testString")
				listDirectorSitesPvdcsClustersOptionsModel.SetAcceptLanguage("testString")
				listDirectorSitesPvdcsClustersOptionsModel.SetXGlobalTransactionID("testString")
				listDirectorSitesPvdcsClustersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDirectorSitesPvdcsClustersOptionsModel).ToNot(BeNil())
				Expect(listDirectorSitesPvdcsClustersOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsClustersOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsClustersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDirectorSitesPvdcsOptions successfully`, func() {
				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				siteID := "testString"
				listDirectorSitesPvdcsOptionsModel := vMwareAsAServiceApiService.NewListDirectorSitesPvdcsOptions(siteID)
				listDirectorSitesPvdcsOptionsModel.SetSiteID("testString")
				listDirectorSitesPvdcsOptionsModel.SetAcceptLanguage("testString")
				listDirectorSitesPvdcsOptionsModel.SetXGlobalTransactionID("testString")
				listDirectorSitesPvdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listDirectorSitesPvdcsOptionsModel).ToNot(BeNil())
				Expect(listDirectorSitesPvdcsOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listDirectorSitesPvdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListPricesOptions successfully`, func() {
				// Construct an instance of the ListPricesOptions model
				listPricesOptionsModel := vMwareAsAServiceApiService.NewListPricesOptions()
				listPricesOptionsModel.SetAcceptLanguage("testString")
				listPricesOptionsModel.SetXGlobalTransactionID("testString")
				listPricesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listPricesOptionsModel).ToNot(BeNil())
				Expect(listPricesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listPricesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listPricesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVdcsOptions successfully`, func() {
				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := vMwareAsAServiceApiService.NewListVdcsOptions()
				listVdcsOptionsModel.SetAcceptLanguage("testString")
				listVdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVdcsOptionsModel).ToNot(BeNil())
				Expect(listVdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listVdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewNewVDCDirectorSite successfully`, func() {
				id := "testString"
				var cluster *vmwareasaserviceapiv1.VDCDirectorSiteCluster = nil
				_, err := vMwareAsAServiceApiService.NewNewVDCDirectorSite(id, cluster)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewNewVDCEdge successfully`, func() {
				typeVar := "dedicated"
				_model, err := vMwareAsAServiceApiService.NewNewVDCEdge(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewNewVDCResourceGroup successfully`, func() {
				id := "testString"
				_model, err := vMwareAsAServiceApiService.NewNewVDCResourceGroup(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewPVDCOrderInfo successfully`, func() {
				name := "testString"
				dataCenter := "testString"
				clusters := []vmwareasaserviceapiv1.ClusterOrderInfo{}
				_model, err := vMwareAsAServiceApiService.NewPVDCOrderInfo(name, dataCenter, clusters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewReplaceOrgAdminPasswordOptions successfully`, func() {
				// Construct an instance of the ReplaceOrgAdminPasswordOptions model
				siteID := "testString"
				replaceOrgAdminPasswordOptionsModel := vMwareAsAServiceApiService.NewReplaceOrgAdminPasswordOptions(siteID)
				replaceOrgAdminPasswordOptionsModel.SetSiteID("testString")
				replaceOrgAdminPasswordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(replaceOrgAdminPasswordOptionsModel).ToNot(BeNil())
				Expect(replaceOrgAdminPasswordOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(replaceOrgAdminPasswordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateClusterResponsePatch successfully`, func() {
				// Construct an instance of the UpdateClusterResponse model
				updateClusterResponse := new(vmwareasaserviceapiv1.UpdateClusterResponse)
				updateClusterResponse.Message = core.StringPtr("The request has been accepted.")

				updateClusterResponsePatch := vMwareAsAServiceApiService.NewUpdateClusterResponsePatch(updateClusterResponse)
				Expect(updateClusterResponsePatch).ToNot(BeNil())

				_path := func(op interface{}) string {
					return *op.(vmwareasaserviceapiv1.JSONPatchOperation).Path
				}
				Expect(updateClusterResponsePatch).To(MatchAllElements(_path, Elements{
				"/message": MatchAllFields(Fields{
					"Op": PointTo(Equal(vmwareasaserviceapiv1.JSONPatchOperation_Op_Add)),
					"Path": PointTo(Equal("/message")),
					"From": BeNil(),
					"Value": Equal(updateClusterResponse.Message),
					}),
				}))
			})
			It(`Invoke NewUpdateDirectorSitesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the JSONPatchOperation model
				jsonPatchOperationModel := new(vmwareasaserviceapiv1.JSONPatchOperation)
				Expect(jsonPatchOperationModel).ToNot(BeNil())
				jsonPatchOperationModel.Op = core.StringPtr("add")
				jsonPatchOperationModel.Path = core.StringPtr("testString")
				jsonPatchOperationModel.From = core.StringPtr("testString")
				jsonPatchOperationModel.Value = core.StringPtr("testString")
				Expect(jsonPatchOperationModel.Op).To(Equal(core.StringPtr("add")))
				Expect(jsonPatchOperationModel.Path).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.From).To(Equal(core.StringPtr("testString")))
				Expect(jsonPatchOperationModel.Value).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				siteID := "testString"
				clusterID := "testString"
				pvdcID := "testString"
				body := []vmwareasaserviceapiv1.JSONPatchOperation{}
				updateDirectorSitesPvdcsClusterOptionsModel := vMwareAsAServiceApiService.NewUpdateDirectorSitesPvdcsClusterOptions(siteID, clusterID, pvdcID, body)
				updateDirectorSitesPvdcsClusterOptionsModel.SetSiteID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetClusterID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetPvdcID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetBody([]vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel})
				updateDirectorSitesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDirectorSitesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.ClusterID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.Body).To(Equal([]vmwareasaserviceapiv1.JSONPatchOperation{*jsonPatchOperationModel}))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVDCDirectorSiteCluster successfully`, func() {
				id := "testString"
				_model, err := vMwareAsAServiceApiService.NewVDCDirectorSiteCluster(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return io.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
