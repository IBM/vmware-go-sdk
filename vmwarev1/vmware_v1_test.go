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
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.ibm.com/VMWSolutions/vmware-go-sdk/vmwarev1"
)

var _ = Describe(`VmwareV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(vmwareService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(vmwareService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL: "https://vmwarev1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(vmwareService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_URL":       "https://vmwarev1/api",
				"VMWARE_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{})
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
					URL: "https://testService/api",
				})
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{})
				err := vmwareService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := vmwareService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != vmwareService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(vmwareService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(vmwareService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_URL":       "https://vmwarev1/api",
				"VMWARE_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(vmwareService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"VMWARE_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			vmwareService, serviceErr := vmwarev1.NewVmwareV1UsingExternalConfig(&vmwarev1.VmwareV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(vmwareService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = vmwarev1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`Parameterized URL tests`, func() {
		It(`Format parameterized URL with all default values`, func() {
			constructedURL, err := vmwarev1.ConstructServiceURL(nil)
			Expect(constructedURL).To(Equal("https://api.us-south.vmware.cloud.ibm.com/v1"))
			Expect(constructedURL).ToNot(BeNil())
			Expect(err).To(BeNil())
		})
		It(`Return an error if a provided variable name is invalid`, func() {
			var providedUrlVariables = map[string]string{
				"invalid_variable_name": "value",
			}
			constructedURL, err := vmwarev1.ConstructServiceURL(providedUrlVariables)
			Expect(constructedURL).To(Equal(""))
			Expect(err).ToNot(BeNil())
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				serviceIdentityModel.Name = core.StringPtr("veeam")

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwarev1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwarev1.PVDCPrototype{*pvdcPrototypeModel}
				createDirectorSitesOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createDirectorSitesOptionsModel.Services = []vmwarev1.ServiceIdentity{*serviceIdentityModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke CreateDirectorSites successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				serviceIdentityModel.Name = core.StringPtr("veeam")

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwarev1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwarev1.PVDCPrototype{*pvdcPrototypeModel}
				createDirectorSitesOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createDirectorSitesOptionsModel.Services = []vmwarev1.ServiceIdentity{*serviceIdentityModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateDirectorSitesWithContext(ctx, createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateDirectorSitesWithContext(ctx, createDirectorSitesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke CreateDirectorSites successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateDirectorSites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				serviceIdentityModel.Name = core.StringPtr("veeam")

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwarev1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwarev1.PVDCPrototype{*pvdcPrototypeModel}
				createDirectorSitesOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createDirectorSitesOptionsModel.Services = []vmwarev1.ServiceIdentity{*serviceIdentityModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDirectorSites with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				serviceIdentityModel.Name = core.StringPtr("veeam")

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwarev1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwarev1.PVDCPrototype{*pvdcPrototypeModel}
				createDirectorSitesOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createDirectorSitesOptionsModel.Services = []vmwarev1.ServiceIdentity{*serviceIdentityModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDirectorSitesOptions model with no property values
				createDirectorSitesOptionsModelNew := new(vmwarev1.CreateDirectorSitesOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateDirectorSites(createDirectorSitesOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				serviceIdentityModel.Name = core.StringPtr("veeam")

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsModel := new(vmwarev1.CreateDirectorSitesOptions)
				createDirectorSitesOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Pvdcs = []vmwarev1.PVDCPrototype{*pvdcPrototypeModel}
				createDirectorSitesOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createDirectorSitesOptionsModel.Services = []vmwarev1.ServiceIdentity{*serviceIdentityModel}
				createDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateDirectorSites(createDirectorSitesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwarev1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"director_sites": [{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}]}`)
				}))
			})
			It(`Invoke ListDirectorSites successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwarev1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListDirectorSitesWithContext(ctx, listDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListDirectorSitesWithContext(ctx, listDirectorSitesOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"director_sites": [{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}]}`)
				}))
			})
			It(`Invoke ListDirectorSites successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListDirectorSites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwarev1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSites with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwarev1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesOptions model
				listDirectorSitesOptionsModel := new(vmwarev1.ListDirectorSitesOptions)
				listDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListDirectorSites(listDirectorSitesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwarev1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke GetDirectorSite successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwarev1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetDirectorSiteWithContext(ctx, getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetDirectorSiteWithContext(ctx, getDirectorSiteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke GetDirectorSite successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetDirectorSite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwarev1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorSite with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwarev1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorSiteOptions model with no property values
				getDirectorSiteOptionsModelNew := new(vmwarev1.GetDirectorSiteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetDirectorSite(getDirectorSiteOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSiteOptions model
				getDirectorSiteOptionsModel := new(vmwarev1.GetDirectorSiteOptions)
				getDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetDirectorSite(getDirectorSiteOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwarev1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke DeleteDirectorSite successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwarev1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.DeleteDirectorSiteWithContext(ctx, deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.DeleteDirectorSiteWithContext(ctx, deleteDirectorSiteOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"crn": "Crn", "href": "Href", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "name": "Name", "status": "creating", "resource_group": {"id": "ID", "name": "Name", "crn": "Crn"}, "pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}], "type": "single_tenant", "services": [{"name": "veeam", "id": "ID", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "status": "creating", "console_url": "ConsoleURL"}], "rhel_vm_activation_key": "RhelVmActivationKey"}`)
				}))
			})
			It(`Invoke DeleteDirectorSite successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.DeleteDirectorSite(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwarev1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDirectorSite with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwarev1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDirectorSiteOptions model with no property values
				deleteDirectorSiteOptionsModelNew := new(vmwarev1.DeleteDirectorSiteOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSiteOptions model
				deleteDirectorSiteOptionsModel := new(vmwarev1.DeleteDirectorSiteOptions)
				deleteDirectorSiteOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSiteOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.DeleteDirectorSite(deleteDirectorSiteOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListDirectorSitesPvdcsWithContext(ctx, listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListDirectorSitesPvdcsWithContext(ctx, listDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"pvdcs": [{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDirectorSitesPvdcsOptions model with no property values
				listDirectorSitesPvdcsOptionsModelNew := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsOptions model
				listDirectorSitesPvdcsOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsOptions)
				listDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenterName = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenterName = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateDirectorSitesPvdcsWithContext(ctx, createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateDirectorSitesPvdcsWithContext(ctx, createDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenterName = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenterName = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDirectorSitesPvdcsOptions model with no property values
				createDirectorSitesPvdcsOptionsModelNew := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				createDirectorSitesPvdcsOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsOptions)
				createDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.DataCenterName = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				createDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.ID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}`)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.ID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetDirectorSitesPvdcsWithContext(ctx, getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetDirectorSitesPvdcsWithContext(ctx, getDirectorSitesPvdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "data_center_name": "DataCenterName", "id": "ID", "href": "Href", "clusters": [{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}], "status": "creating", "provider_types": [{"name": "on_demand"}]}`)
				}))
			})
			It(`Invoke GetDirectorSitesPvdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetDirectorSitesPvdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.ID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorSitesPvdcs with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.ID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorSitesPvdcsOptions model with no property values
				getDirectorSitesPvdcsOptionsModelNew := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				getDirectorSitesPvdcsOptionsModel := new(vmwarev1.GetDirectorSitesPvdcsOptions)
				getDirectorSitesPvdcsOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.ID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorSitesPvdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListDirectorSitesPvdcsClustersWithContext(ctx, listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListDirectorSitesPvdcsClustersWithContext(ctx, listDirectorSitesPvdcsClustersOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"clusters": [{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}]}`)
				}))
			})
			It(`Invoke ListDirectorSitesPvdcsClusters successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcsClusters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSitesPvdcsClusters with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListDirectorSitesPvdcsClustersOptions model with no property values
				listDirectorSitesPvdcsClustersOptionsModelNew := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSitesPvdcsClustersOptions model
				listDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.ListDirectorSitesPvdcsClustersOptions)
				listDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptionsModel)
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
	Describe(`CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions *CreateDirectorSitesPvdcsClustersOptions) - Operation response error`, func() {
		createDirectorSitesPvdcsClustersPath := "/director_sites/testString/pvdcs/testString/clusters"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsClustersPath))
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
			It(`Invoke CreateDirectorSitesPvdcsClusters with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				createDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				createDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.HostCount = core.Int64Ptr(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.HostProfile = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.FileShares = fileSharesPrototypeModel
				createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptions *CreateDirectorSitesPvdcsClustersOptions)`, func() {
		createDirectorSitesPvdcsClustersPath := "/director_sites/testString/pvdcs/testString/clusters"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsClustersPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcsClusters successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				createDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				createDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.HostCount = core.Int64Ptr(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.HostProfile = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.FileShares = fileSharesPrototypeModel
				createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateDirectorSitesPvdcsClustersWithContext(ctx, createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateDirectorSitesPvdcsClustersWithContext(ctx, createDirectorSitesPvdcsClustersOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(createDirectorSitesPvdcsClustersPath))
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke CreateDirectorSitesPvdcsClusters successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcsClusters(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				createDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				createDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.HostCount = core.Int64Ptr(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.HostProfile = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.FileShares = fileSharesPrototypeModel
				createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateDirectorSitesPvdcsClusters with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				createDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				createDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.HostCount = core.Int64Ptr(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.HostProfile = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.FileShares = fileSharesPrototypeModel
				createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateDirectorSitesPvdcsClustersOptions model with no property values
				createDirectorSitesPvdcsClustersOptionsModelNew := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModelNew)
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
			It(`Invoke CreateDirectorSitesPvdcsClusters successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				createDirectorSitesPvdcsClustersOptionsModel := new(vmwarev1.CreateDirectorSitesPvdcsClustersOptions)
				createDirectorSitesPvdcsClustersOptionsModel.SiteID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.PvdcID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Name = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.HostCount = core.Int64Ptr(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.HostProfile = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.FileShares = fileSharesPrototypeModel
				createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				createDirectorSitesPvdcsClustersOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateDirectorSitesPvdcsClusters(createDirectorSitesPvdcsClustersOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetDirectorInstancesPvdcsClusterWithContext(ctx, getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetDirectorInstancesPvdcsClusterWithContext(ctx, getDirectorInstancesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke GetDirectorInstancesPvdcsCluster successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetDirectorInstancesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetDirectorInstancesPvdcsCluster with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetDirectorInstancesPvdcsClusterOptions model with no property values
				getDirectorInstancesPvdcsClusterOptionsModelNew := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				getDirectorInstancesPvdcsClusterOptionsModel := new(vmwarev1.GetDirectorInstancesPvdcsClusterOptions)
				getDirectorInstancesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.DeleteDirectorSitesPvdcsClusterWithContext(ctx, deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.DeleteDirectorSitesPvdcsClusterWithContext(ctx, deleteDirectorSitesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"name": "Name", "host_count": 2, "host_profile": "HostProfile", "id": "ID", "data_center_name": "DataCenterName", "status": "Status", "href": "Href", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}}`)
				}))
			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.DeleteDirectorSitesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteDirectorSitesPvdcsCluster with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteDirectorSitesPvdcsClusterOptions model with no property values
				deleteDirectorSitesPvdcsClusterOptionsModelNew := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				deleteDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.DeleteDirectorSitesPvdcsClusterOptions)
				deleteDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPatch model
				clusterPatchModel := new(vmwarev1.ClusterPatch)
				clusterPatchModel.FileShares = fileSharesPrototypeModel
				clusterPatchModel.HostCount = core.Int64Ptr(int64(2))
				clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = clusterPatchModelAsPatch
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "message": "The request has been accepted.", "operation_id": "OperationID"}`)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPatch model
				clusterPatchModel := new(vmwarev1.ClusterPatch)
				clusterPatchModel.FileShares = fileSharesPrototypeModel
				clusterPatchModel.HostCount = core.Int64Ptr(int64(2))
				clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = clusterPatchModelAsPatch
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.UpdateDirectorSitesPvdcsClusterWithContext(ctx, updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.UpdateDirectorSitesPvdcsClusterWithContext(ctx, updateDirectorSitesPvdcsClusterOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"id": "ID", "name": "Name", "href": "Href", "ordered_at": "2019-01-01T12:00:00.000Z", "provisioned_at": "2019-01-01T12:00:00.000Z", "host_count": 9, "status": "Status", "data_center_name": "DataCenterName", "director_site": {"crn": "Crn", "href": "Href", "id": "ID"}, "host_profile": "HostProfile", "storage_type": "nfs", "billing_plan": "monthly", "file_shares": {"STORAGE_POINT_TWO_FIVE_IOPS_GB": 0, "STORAGE_TWO_IOPS_GB": 0, "STORAGE_FOUR_IOPS_GB": 0, "STORAGE_TEN_IOPS_GB": 0}, "message": "The request has been accepted.", "operation_id": "OperationID"}`)
				}))
			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.UpdateDirectorSitesPvdcsCluster(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPatch model
				clusterPatchModel := new(vmwarev1.ClusterPatch)
				clusterPatchModel.FileShares = fileSharesPrototypeModel
				clusterPatchModel.HostCount = core.Int64Ptr(int64(2))
				clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = clusterPatchModelAsPatch
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateDirectorSitesPvdcsCluster with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPatch model
				clusterPatchModel := new(vmwarev1.ClusterPatch)
				clusterPatchModel.FileShares = fileSharesPrototypeModel
				clusterPatchModel.HostCount = core.Int64Ptr(int64(2))
				clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = clusterPatchModelAsPatch
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateDirectorSitesPvdcsClusterOptions model with no property values
				updateDirectorSitesPvdcsClusterOptionsModelNew := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))

				// Construct an instance of the ClusterPatch model
				clusterPatchModel := new(vmwarev1.ClusterPatch)
				clusterPatchModel.FileShares = fileSharesPrototypeModel
				clusterPatchModel.HostCount = core.Int64Ptr(int64(2))
				clusterPatchModelAsPatch, asPatchErr := clusterPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				updateDirectorSitesPvdcsClusterOptionsModel := new(vmwarev1.UpdateDirectorSitesPvdcsClusterOptions)
				updateDirectorSitesPvdcsClusterOptionsModel.SiteID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.ID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.PvdcID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Body = clusterPatchModelAsPatch
				updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwarev1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"director_site_regions": [{"name": "Name", "data_centers": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}]}`)
				}))
			})
			It(`Invoke ListDirectorSiteRegions successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwarev1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListDirectorSiteRegionsWithContext(ctx, listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListDirectorSiteRegionsWithContext(ctx, listDirectorSiteRegionsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"director_site_regions": [{"name": "Name", "data_centers": [{"display_name": "DisplayName", "name": "Name", "uplink_speed": "UplinkSpeed"}], "endpoint": "Endpoint"}]}`)
				}))
			})
			It(`Invoke ListDirectorSiteRegions successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListDirectorSiteRegions(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwarev1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSiteRegions with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwarev1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteRegionsOptions model
				listDirectorSiteRegionsOptionsModel := new(vmwarev1.ListDirectorSiteRegionsOptions)
				listDirectorSiteRegionsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteRegionsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListDirectorSiteRegions(listDirectorSiteRegionsOptionsModel)
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
	Describe(`ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions *ListMultitenantDirectorSitesOptions) - Operation response error`, func() {
		listMultitenantDirectorSitesPath := "/multitenant_director_sites"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMultitenantDirectorSitesPath))
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
			It(`Invoke ListMultitenantDirectorSites with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := new(vmwarev1.ListMultitenantDirectorSitesOptions)
				listMultitenantDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListMultitenantDirectorSites(listMultitenantDirectorSitesOptions *ListMultitenantDirectorSitesOptions)`, func() {
		listMultitenantDirectorSitesPath := "/multitenant_director_sites"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listMultitenantDirectorSitesPath))
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
					fmt.Fprintf(res, "%s", `{"multitenant_director_sites": [{"name": "Name", "display_name": "DisplayName", "id": "ID", "region": "Region", "pvdcs": [{"name": "Name", "id": "ID", "data_center_name": "DataCenterName", "provider_types": [{"name": "on_demand"}]}], "services": ["veeam"]}]}`)
				}))
			})
			It(`Invoke ListMultitenantDirectorSites successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := new(vmwarev1.ListMultitenantDirectorSitesOptions)
				listMultitenantDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListMultitenantDirectorSitesWithContext(ctx, listMultitenantDirectorSitesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListMultitenantDirectorSitesWithContext(ctx, listMultitenantDirectorSitesOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(listMultitenantDirectorSitesPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					Expect(req.Header["X-Global-Transaction-Id"]).ToNot(BeNil())
					Expect(req.Header["X-Global-Transaction-Id"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"multitenant_director_sites": [{"name": "Name", "display_name": "DisplayName", "id": "ID", "region": "Region", "pvdcs": [{"name": "Name", "id": "ID", "data_center_name": "DataCenterName", "provider_types": [{"name": "on_demand"}]}], "services": ["veeam"]}]}`)
				}))
			})
			It(`Invoke ListMultitenantDirectorSites successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListMultitenantDirectorSites(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := new(vmwarev1.ListMultitenantDirectorSitesOptions)
				listMultitenantDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListMultitenantDirectorSites with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := new(vmwarev1.ListMultitenantDirectorSitesOptions)
				listMultitenantDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
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
			It(`Invoke ListMultitenantDirectorSites successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := new(vmwarev1.ListMultitenantDirectorSitesOptions)
				listMultitenantDirectorSitesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listMultitenantDirectorSitesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListMultitenantDirectorSites(listMultitenantDirectorSitesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwarev1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwarev1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListDirectorSiteHostProfilesWithContext(ctx, listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListDirectorSiteHostProfilesWithContext(ctx, listDirectorSiteHostProfilesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListDirectorSiteHostProfiles(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwarev1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListDirectorSiteHostProfiles with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwarev1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := new(vmwarev1.ListDirectorSiteHostProfilesOptions)
				listDirectorSiteHostProfilesOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.XGlobalTransactionID = core.StringPtr("testString")
				listDirectorSiteHostProfilesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.ListVdcs(listVdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"vdcs": [{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.ListVdcsWithContext(ctx, listVdcsOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"vdcs": [{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}]}`)
				}))
			})
			It(`Invoke ListVdcs successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.ListVdcs(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.ListVdcs(listVdcsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVdcs with error: Operation request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := new(vmwarev1.ListVdcsOptions)
				listVdcsOptionsModel.AcceptLanguage = core.StringPtr("testString")
				listVdcsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.ListVdcs(listVdcsOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				vdcProviderTypeModel.Name = core.StringPtr("paygo")

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = vdcDirectorSitePrototypeModel
				createVdcOptionsModel.Edge = vdcEdgePrototypeModel
				createVdcOptionsModel.FastProvisioningEnabled = core.BoolPtr(true)
				createVdcOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createVdcOptionsModel.Cpu = core.Int64Ptr(int64(0))
				createVdcOptionsModel.Ram = core.Int64Ptr(int64(0))
				createVdcOptionsModel.RhelByol = core.BoolPtr(false)
				createVdcOptionsModel.WindowsByol = core.BoolPtr(false)
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke CreateVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				vdcProviderTypeModel.Name = core.StringPtr("paygo")

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = vdcDirectorSitePrototypeModel
				createVdcOptionsModel.Edge = vdcEdgePrototypeModel
				createVdcOptionsModel.FastProvisioningEnabled = core.BoolPtr(true)
				createVdcOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createVdcOptionsModel.Cpu = core.Int64Ptr(int64(0))
				createVdcOptionsModel.Ram = core.Int64Ptr(int64(0))
				createVdcOptionsModel.RhelByol = core.BoolPtr(false)
				createVdcOptionsModel.WindowsByol = core.BoolPtr(false)
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.CreateVdcWithContext(ctx, createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.CreateVdcWithContext(ctx, createVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke CreateVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.CreateVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				vdcProviderTypeModel.Name = core.StringPtr("paygo")

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = vdcDirectorSitePrototypeModel
				createVdcOptionsModel.Edge = vdcEdgePrototypeModel
				createVdcOptionsModel.FastProvisioningEnabled = core.BoolPtr(true)
				createVdcOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createVdcOptionsModel.Cpu = core.Int64Ptr(int64(0))
				createVdcOptionsModel.Ram = core.Int64Ptr(int64(0))
				createVdcOptionsModel.RhelByol = core.BoolPtr(false)
				createVdcOptionsModel.WindowsByol = core.BoolPtr(false)
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				vdcProviderTypeModel.Name = core.StringPtr("paygo")

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = vdcDirectorSitePrototypeModel
				createVdcOptionsModel.Edge = vdcEdgePrototypeModel
				createVdcOptionsModel.FastProvisioningEnabled = core.BoolPtr(true)
				createVdcOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createVdcOptionsModel.Cpu = core.Int64Ptr(int64(0))
				createVdcOptionsModel.Ram = core.Int64Ptr(int64(0))
				createVdcOptionsModel.RhelByol = core.BoolPtr(false)
				createVdcOptionsModel.WindowsByol = core.BoolPtr(false)
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateVdcOptions model with no property values
				createVdcOptionsModelNew := new(vmwarev1.CreateVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.CreateVdc(createVdcOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				vdcProviderTypeModel.Name = core.StringPtr("paygo")

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				resourceGroupIdentityModel.ID = core.StringPtr("testString")

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsModel := new(vmwarev1.CreateVdcOptions)
				createVdcOptionsModel.Name = core.StringPtr("testString")
				createVdcOptionsModel.DirectorSite = vdcDirectorSitePrototypeModel
				createVdcOptionsModel.Edge = vdcEdgePrototypeModel
				createVdcOptionsModel.FastProvisioningEnabled = core.BoolPtr(true)
				createVdcOptionsModel.ResourceGroup = resourceGroupIdentityModel
				createVdcOptionsModel.Cpu = core.Int64Ptr(int64(0))
				createVdcOptionsModel.Ram = core.Int64Ptr(int64(0))
				createVdcOptionsModel.RhelByol = core.BoolPtr(false)
				createVdcOptionsModel.WindowsByol = core.BoolPtr(false)
				createVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				createVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.CreateVdc(createVdcOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.ID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke GetVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.ID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.GetVdcWithContext(ctx, getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.GetVdcWithContext(ctx, getVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke GetVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.GetVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.ID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.ID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVdcOptions model with no property values
				getVdcOptionsModelNew := new(vmwarev1.GetVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.GetVdc(getVdcOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the GetVdcOptions model
				getVdcOptionsModel := new(vmwarev1.GetVdcOptions)
				getVdcOptionsModel.ID = core.StringPtr("testString")
				getVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				getVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.GetVdc(getVdcOptionsModel)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.ID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke DeleteVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.ID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.DeleteVdcWithContext(ctx, deleteVdcOptionsModel)
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
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke DeleteVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.DeleteVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.ID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke DeleteVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.ID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the DeleteVdcOptions model with no property values
				deleteVdcOptionsModelNew := new(vmwarev1.DeleteVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.DeleteVdc(deleteVdcOptionsModelNew)
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
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the DeleteVdcOptions model
				deleteVdcOptionsModel := new(vmwarev1.DeleteVdcOptions)
				deleteVdcOptionsModel.ID = core.StringPtr("testString")
				deleteVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				deleteVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.DeleteVdc(deleteVdcOptionsModel)
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
	Describe(`UpdateVdc(updateVdcOptions *UpdateVdcOptions) - Operation response error`, func() {
		updateVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVdcPath))
					Expect(req.Method).To(Equal("PATCH"))
					Expect(req.Header["Accept-Language"]).ToNot(BeNil())
					Expect(req.Header["Accept-Language"][0]).To(Equal(fmt.Sprintf("%v", "testString")))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke UpdateVdc with error: Operation response processing error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCPatch model
				vdcPatchModel := new(vmwarev1.VDCPatch)
				vdcPatchModel.Cpu = core.Int64Ptr(int64(0))
				vdcPatchModel.FastProvisioningEnabled = core.BoolPtr(true)
				vdcPatchModel.Ram = core.Int64Ptr(int64(0))
				vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVdcOptions model
				updateVdcOptionsModel := new(vmwarev1.UpdateVdcOptions)
				updateVdcOptionsModel.ID = core.StringPtr("testString")
				updateVdcOptionsModel.VDCPatch = vdcPatchModelAsPatch
				updateVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := vmwareService.UpdateVdc(updateVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				vmwareService.EnableRetries(0, 0)
				result, response, operationErr = vmwareService.UpdateVdc(updateVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateVdc(updateVdcOptions *UpdateVdcOptions)`, func() {
		updateVdcPath := "/vdcs/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateVdcPath))
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
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke UpdateVdc successfully with retries`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())
				vmwareService.EnableRetries(0, 0)

				// Construct an instance of the VDCPatch model
				vdcPatchModel := new(vmwarev1.VDCPatch)
				vdcPatchModel.Cpu = core.Int64Ptr(int64(0))
				vdcPatchModel.FastProvisioningEnabled = core.BoolPtr(true)
				vdcPatchModel.Ram = core.Int64Ptr(int64(0))
				vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVdcOptions model
				updateVdcOptionsModel := new(vmwarev1.UpdateVdcOptions)
				updateVdcOptionsModel.ID = core.StringPtr("testString")
				updateVdcOptionsModel.VDCPatch = vdcPatchModelAsPatch
				updateVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := vmwareService.UpdateVdcWithContext(ctx, updateVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				vmwareService.DisableRetries()
				result, response, operationErr := vmwareService.UpdateVdc(updateVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = vmwareService.UpdateVdcWithContext(ctx, updateVdcOptionsModel)
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
					Expect(req.URL.EscapedPath()).To(Equal(updateVdcPath))
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
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(202)
					fmt.Fprintf(res, "%s", `{"href": "Href", "id": "ID", "provisioned_at": "2019-01-01T12:00:00.000Z", "cpu": 0, "crn": "Crn", "deleted_at": "2019-01-01T12:00:00.000Z", "director_site": {"id": "ID", "pvdc": {"id": "ID", "provider_type": {"name": "paygo"}}, "url": "URL"}, "edges": [{"id": "ID", "public_ips": ["PublicIps"], "size": "medium", "status": "creating", "type": "performance"}], "status_reasons": [{"code": "insufficent_cpu", "message": "Message", "more_info": "MoreInfo"}], "name": "Name", "ordered_at": "2019-01-01T12:00:00.000Z", "org_name": "OrgName", "ram": 0, "status": "creating", "type": "single_tenant", "fast_provisioning_enabled": false, "rhel_byol": true, "windows_byol": false}`)
				}))
			})
			It(`Invoke UpdateVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := vmwareService.UpdateVdc(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the VDCPatch model
				vdcPatchModel := new(vmwarev1.VDCPatch)
				vdcPatchModel.Cpu = core.Int64Ptr(int64(0))
				vdcPatchModel.FastProvisioningEnabled = core.BoolPtr(true)
				vdcPatchModel.Ram = core.Int64Ptr(int64(0))
				vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVdcOptions model
				updateVdcOptionsModel := new(vmwarev1.UpdateVdcOptions)
				updateVdcOptionsModel.ID = core.StringPtr("testString")
				updateVdcOptionsModel.VDCPatch = vdcPatchModelAsPatch
				updateVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = vmwareService.UpdateVdc(updateVdcOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke UpdateVdc with error: Operation validation and request error`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCPatch model
				vdcPatchModel := new(vmwarev1.VDCPatch)
				vdcPatchModel.Cpu = core.Int64Ptr(int64(0))
				vdcPatchModel.FastProvisioningEnabled = core.BoolPtr(true)
				vdcPatchModel.Ram = core.Int64Ptr(int64(0))
				vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVdcOptions model
				updateVdcOptionsModel := new(vmwarev1.UpdateVdcOptions)
				updateVdcOptionsModel.ID = core.StringPtr("testString")
				updateVdcOptionsModel.VDCPatch = vdcPatchModelAsPatch
				updateVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := vmwareService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := vmwareService.UpdateVdc(updateVdcOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the UpdateVdcOptions model with no property values
				updateVdcOptionsModelNew := new(vmwarev1.UpdateVdcOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = vmwareService.UpdateVdc(updateVdcOptionsModelNew)
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
			It(`Invoke UpdateVdc successfully`, func() {
				vmwareService, serviceErr := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(vmwareService).ToNot(BeNil())

				// Construct an instance of the VDCPatch model
				vdcPatchModel := new(vmwarev1.VDCPatch)
				vdcPatchModel.Cpu = core.Int64Ptr(int64(0))
				vdcPatchModel.FastProvisioningEnabled = core.BoolPtr(true)
				vdcPatchModel.Ram = core.Int64Ptr(int64(0))
				vdcPatchModelAsPatch, asPatchErr := vdcPatchModel.AsPatch()
				Expect(asPatchErr).To(BeNil())

				// Construct an instance of the UpdateVdcOptions model
				updateVdcOptionsModel := new(vmwarev1.UpdateVdcOptions)
				updateVdcOptionsModel.ID = core.StringPtr("testString")
				updateVdcOptionsModel.VDCPatch = vdcPatchModelAsPatch
				updateVdcOptionsModel.AcceptLanguage = core.StringPtr("testString")
				updateVdcOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := vmwareService.UpdateVdc(updateVdcOptionsModel)
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
			vmwareService, _ := vmwarev1.NewVmwareV1(&vmwarev1.VmwareV1Options{
				URL:           "http://vmwarev1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewClusterPrototype successfully`, func() {
				name := "testString"
				hostCount := int64(2)
				hostProfile := "testString"
				var fileShares *vmwarev1.FileSharesPrototype = nil
				_, err := vmwareService.NewClusterPrototype(name, hostCount, hostProfile, fileShares)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewCreateDirectorSitesOptions successfully`, func() {
				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				Expect(fileSharesPrototypeModel).ToNot(BeNil())
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				Expect(clusterPrototypeModel).ToNot(BeNil())
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel
				Expect(clusterPrototypeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterPrototypeModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterPrototypeModel.HostProfile).To(Equal(core.StringPtr("testString")))
				Expect(clusterPrototypeModel.FileShares).To(Equal(fileSharesPrototypeModel))

				// Construct an instance of the PVDCPrototype model
				pvdcPrototypeModel := new(vmwarev1.PVDCPrototype)
				Expect(pvdcPrototypeModel).ToNot(BeNil())
				pvdcPrototypeModel.Name = core.StringPtr("testString")
				pvdcPrototypeModel.DataCenterName = core.StringPtr("testString")
				pvdcPrototypeModel.Clusters = []vmwarev1.ClusterPrototype{*clusterPrototypeModel}
				Expect(pvdcPrototypeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(pvdcPrototypeModel.DataCenterName).To(Equal(core.StringPtr("testString")))
				Expect(pvdcPrototypeModel.Clusters).To(Equal([]vmwarev1.ClusterPrototype{*clusterPrototypeModel}))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("testString")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the ServiceIdentity model
				serviceIdentityModel := new(vmwarev1.ServiceIdentity)
				Expect(serviceIdentityModel).ToNot(BeNil())
				serviceIdentityModel.Name = core.StringPtr("veeam")
				Expect(serviceIdentityModel.Name).To(Equal(core.StringPtr("veeam")))

				// Construct an instance of the CreateDirectorSitesOptions model
				createDirectorSitesOptionsName := "testString"
				createDirectorSitesOptionsPvdcs := []vmwarev1.PVDCPrototype{}
				createDirectorSitesOptionsModel := vmwareService.NewCreateDirectorSitesOptions(createDirectorSitesOptionsName, createDirectorSitesOptionsPvdcs)
				createDirectorSitesOptionsModel.SetName("testString")
				createDirectorSitesOptionsModel.SetPvdcs([]vmwarev1.PVDCPrototype{*pvdcPrototypeModel})
				createDirectorSitesOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createDirectorSitesOptionsModel.SetServices([]vmwarev1.ServiceIdentity{*serviceIdentityModel})
				createDirectorSitesOptionsModel.SetAcceptLanguage("testString")
				createDirectorSitesOptionsModel.SetXGlobalTransactionID("testString")
				createDirectorSitesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDirectorSitesOptionsModel).ToNot(BeNil())
				Expect(createDirectorSitesOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.Pvdcs).To(Equal([]vmwarev1.PVDCPrototype{*pvdcPrototypeModel}))
				Expect(createDirectorSitesOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createDirectorSitesOptionsModel.Services).To(Equal([]vmwarev1.ServiceIdentity{*serviceIdentityModel}))
				Expect(createDirectorSitesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDirectorSitesPvdcsClustersOptions successfully`, func() {
				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				Expect(fileSharesPrototypeModel).ToNot(BeNil())
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the CreateDirectorSitesPvdcsClustersOptions model
				siteID := "testString"
				pvdcID := "testString"
				createDirectorSitesPvdcsClustersOptionsName := "testString"
				createDirectorSitesPvdcsClustersOptionsHostCount := int64(2)
				createDirectorSitesPvdcsClustersOptionsHostProfile := "testString"
				var createDirectorSitesPvdcsClustersOptionsFileShares *vmwarev1.FileSharesPrototype = nil
				createDirectorSitesPvdcsClustersOptionsModel := vmwareService.NewCreateDirectorSitesPvdcsClustersOptions(siteID, pvdcID, createDirectorSitesPvdcsClustersOptionsName, createDirectorSitesPvdcsClustersOptionsHostCount, createDirectorSitesPvdcsClustersOptionsHostProfile, createDirectorSitesPvdcsClustersOptionsFileShares)
				createDirectorSitesPvdcsClustersOptionsModel.SetSiteID("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetPvdcID("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetName("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetHostCount(int64(2))
				createDirectorSitesPvdcsClustersOptionsModel.SetHostProfile("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetFileShares(fileSharesPrototypeModel)
				createDirectorSitesPvdcsClustersOptionsModel.SetAcceptLanguage("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetXGlobalTransactionID("testString")
				createDirectorSitesPvdcsClustersOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDirectorSitesPvdcsClustersOptionsModel).ToNot(BeNil())
				Expect(createDirectorSitesPvdcsClustersOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.HostProfile).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.FileShares).To(Equal(fileSharesPrototypeModel))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsClustersOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateDirectorSitesPvdcsOptions successfully`, func() {
				// Construct an instance of the FileSharesPrototype model
				fileSharesPrototypeModel := new(vmwarev1.FileSharesPrototype)
				Expect(fileSharesPrototypeModel).ToNot(BeNil())
				fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETWOIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGEFOURIOPSGB = core.Int64Ptr(int64(0))
				fileSharesPrototypeModel.STORAGETENIOPSGB = core.Int64Ptr(int64(0))
				Expect(fileSharesPrototypeModel.STORAGEPOINTTWOFIVEIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETWOIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGEFOURIOPSGB).To(Equal(core.Int64Ptr(int64(0))))
				Expect(fileSharesPrototypeModel.STORAGETENIOPSGB).To(Equal(core.Int64Ptr(int64(0))))

				// Construct an instance of the ClusterPrototype model
				clusterPrototypeModel := new(vmwarev1.ClusterPrototype)
				Expect(clusterPrototypeModel).ToNot(BeNil())
				clusterPrototypeModel.Name = core.StringPtr("testString")
				clusterPrototypeModel.HostCount = core.Int64Ptr(int64(2))
				clusterPrototypeModel.HostProfile = core.StringPtr("testString")
				clusterPrototypeModel.FileShares = fileSharesPrototypeModel
				Expect(clusterPrototypeModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(clusterPrototypeModel.HostCount).To(Equal(core.Int64Ptr(int64(2))))
				Expect(clusterPrototypeModel.HostProfile).To(Equal(core.StringPtr("testString")))
				Expect(clusterPrototypeModel.FileShares).To(Equal(fileSharesPrototypeModel))

				// Construct an instance of the CreateDirectorSitesPvdcsOptions model
				siteID := "testString"
				createDirectorSitesPvdcsOptionsName := "testString"
				createDirectorSitesPvdcsOptionsDataCenterName := "testString"
				createDirectorSitesPvdcsOptionsClusters := []vmwarev1.ClusterPrototype{}
				createDirectorSitesPvdcsOptionsModel := vmwareService.NewCreateDirectorSitesPvdcsOptions(siteID, createDirectorSitesPvdcsOptionsName, createDirectorSitesPvdcsOptionsDataCenterName, createDirectorSitesPvdcsOptionsClusters)
				createDirectorSitesPvdcsOptionsModel.SetSiteID("testString")
				createDirectorSitesPvdcsOptionsModel.SetName("testString")
				createDirectorSitesPvdcsOptionsModel.SetDataCenterName("testString")
				createDirectorSitesPvdcsOptionsModel.SetClusters([]vmwarev1.ClusterPrototype{*clusterPrototypeModel})
				createDirectorSitesPvdcsOptionsModel.SetAcceptLanguage("testString")
				createDirectorSitesPvdcsOptionsModel.SetXGlobalTransactionID("testString")
				createDirectorSitesPvdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createDirectorSitesPvdcsOptionsModel).ToNot(BeNil())
				Expect(createDirectorSitesPvdcsOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.DataCenterName).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Clusters).To(Equal([]vmwarev1.ClusterPrototype{*clusterPrototypeModel}))
				Expect(createDirectorSitesPvdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(createDirectorSitesPvdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateVdcOptions successfully`, func() {
				// Construct an instance of the VDCProviderType model
				vdcProviderTypeModel := new(vmwarev1.VDCProviderType)
				Expect(vdcProviderTypeModel).ToNot(BeNil())
				vdcProviderTypeModel.Name = core.StringPtr("paygo")
				Expect(vdcProviderTypeModel.Name).To(Equal(core.StringPtr("paygo")))

				// Construct an instance of the DirectorSitePVDC model
				directorSitePvdcModel := new(vmwarev1.DirectorSitePVDC)
				Expect(directorSitePvdcModel).ToNot(BeNil())
				directorSitePvdcModel.ID = core.StringPtr("testString")
				directorSitePvdcModel.ProviderType = vdcProviderTypeModel
				Expect(directorSitePvdcModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(directorSitePvdcModel.ProviderType).To(Equal(vdcProviderTypeModel))

				// Construct an instance of the VDCDirectorSitePrototype model
				vdcDirectorSitePrototypeModel := new(vmwarev1.VDCDirectorSitePrototype)
				Expect(vdcDirectorSitePrototypeModel).ToNot(BeNil())
				vdcDirectorSitePrototypeModel.ID = core.StringPtr("testString")
				vdcDirectorSitePrototypeModel.Pvdc = directorSitePvdcModel
				Expect(vdcDirectorSitePrototypeModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(vdcDirectorSitePrototypeModel.Pvdc).To(Equal(directorSitePvdcModel))

				// Construct an instance of the VDCEdgePrototype model
				vdcEdgePrototypeModel := new(vmwarev1.VDCEdgePrototype)
				Expect(vdcEdgePrototypeModel).ToNot(BeNil())
				vdcEdgePrototypeModel.Size = core.StringPtr("medium")
				vdcEdgePrototypeModel.Type = core.StringPtr("performance")
				Expect(vdcEdgePrototypeModel.Size).To(Equal(core.StringPtr("medium")))
				Expect(vdcEdgePrototypeModel.Type).To(Equal(core.StringPtr("performance")))

				// Construct an instance of the ResourceGroupIdentity model
				resourceGroupIdentityModel := new(vmwarev1.ResourceGroupIdentity)
				Expect(resourceGroupIdentityModel).ToNot(BeNil())
				resourceGroupIdentityModel.ID = core.StringPtr("testString")
				Expect(resourceGroupIdentityModel.ID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the CreateVdcOptions model
				createVdcOptionsName := "testString"
				var createVdcOptionsDirectorSite *vmwarev1.VDCDirectorSitePrototype = nil
				createVdcOptionsModel := vmwareService.NewCreateVdcOptions(createVdcOptionsName, createVdcOptionsDirectorSite)
				createVdcOptionsModel.SetName("testString")
				createVdcOptionsModel.SetDirectorSite(vdcDirectorSitePrototypeModel)
				createVdcOptionsModel.SetEdge(vdcEdgePrototypeModel)
				createVdcOptionsModel.SetFastProvisioningEnabled(true)
				createVdcOptionsModel.SetResourceGroup(resourceGroupIdentityModel)
				createVdcOptionsModel.SetCpu(int64(0))
				createVdcOptionsModel.SetRam(int64(0))
				createVdcOptionsModel.SetRhelByol(false)
				createVdcOptionsModel.SetWindowsByol(false)
				createVdcOptionsModel.SetAcceptLanguage("testString")
				createVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createVdcOptionsModel).ToNot(BeNil())
				Expect(createVdcOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.DirectorSite).To(Equal(vdcDirectorSitePrototypeModel))
				Expect(createVdcOptionsModel.Edge).To(Equal(vdcEdgePrototypeModel))
				Expect(createVdcOptionsModel.FastProvisioningEnabled).To(Equal(core.BoolPtr(true)))
				Expect(createVdcOptionsModel.ResourceGroup).To(Equal(resourceGroupIdentityModel))
				Expect(createVdcOptionsModel.Cpu).To(Equal(core.Int64Ptr(int64(0))))
				Expect(createVdcOptionsModel.Ram).To(Equal(core.Int64Ptr(int64(0))))
				Expect(createVdcOptionsModel.RhelByol).To(Equal(core.BoolPtr(false)))
				Expect(createVdcOptionsModel.WindowsByol).To(Equal(core.BoolPtr(false)))
				Expect(createVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(createVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDirectorSiteOptions successfully`, func() {
				// Construct an instance of the DeleteDirectorSiteOptions model
				id := "testString"
				deleteDirectorSiteOptionsModel := vmwareService.NewDeleteDirectorSiteOptions(id)
				deleteDirectorSiteOptionsModel.SetID("testString")
				deleteDirectorSiteOptionsModel.SetAcceptLanguage("testString")
				deleteDirectorSiteOptionsModel.SetXGlobalTransactionID("testString")
				deleteDirectorSiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDirectorSiteOptionsModel).ToNot(BeNil())
				Expect(deleteDirectorSiteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteDirectorSitesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the DeleteDirectorSitesPvdcsClusterOptions model
				siteID := "testString"
				id := "testString"
				pvdcID := "testString"
				deleteDirectorSitesPvdcsClusterOptionsModel := vmwareService.NewDeleteDirectorSitesPvdcsClusterOptions(siteID, id, pvdcID)
				deleteDirectorSitesPvdcsClusterOptionsModel.SetSiteID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetPvdcID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				deleteDirectorSitesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(deleteDirectorSitesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteVdcOptions successfully`, func() {
				// Construct an instance of the DeleteVdcOptions model
				id := "testString"
				deleteVdcOptionsModel := vmwareService.NewDeleteVdcOptions(id)
				deleteVdcOptionsModel.SetID("testString")
				deleteVdcOptionsModel.SetAcceptLanguage("testString")
				deleteVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteVdcOptionsModel).ToNot(BeNil())
				Expect(deleteVdcOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(deleteVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDirectorSitePVDC successfully`, func() {
				id := "testString"
				_model, err := vmwareService.NewDirectorSitePVDC(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewGetDirectorInstancesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the GetDirectorInstancesPvdcsClusterOptions model
				siteID := "testString"
				id := "testString"
				pvdcID := "testString"
				getDirectorInstancesPvdcsClusterOptionsModel := vmwareService.NewGetDirectorInstancesPvdcsClusterOptions(siteID, id, pvdcID)
				getDirectorInstancesPvdcsClusterOptionsModel.SetSiteID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetPvdcID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorInstancesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorInstancesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorInstancesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDirectorSiteOptions successfully`, func() {
				// Construct an instance of the GetDirectorSiteOptions model
				id := "testString"
				getDirectorSiteOptionsModel := vmwareService.NewGetDirectorSiteOptions(id)
				getDirectorSiteOptionsModel.SetID("testString")
				getDirectorSiteOptionsModel.SetAcceptLanguage("testString")
				getDirectorSiteOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorSiteOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorSiteOptionsModel).ToNot(BeNil())
				Expect(getDirectorSiteOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSiteOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetDirectorSitesPvdcsOptions successfully`, func() {
				// Construct an instance of the GetDirectorSitesPvdcsOptions model
				siteID := "testString"
				id := "testString"
				getDirectorSitesPvdcsOptionsModel := vmwareService.NewGetDirectorSitesPvdcsOptions(siteID, id)
				getDirectorSitesPvdcsOptionsModel.SetSiteID("testString")
				getDirectorSitesPvdcsOptionsModel.SetID("testString")
				getDirectorSitesPvdcsOptionsModel.SetAcceptLanguage("testString")
				getDirectorSitesPvdcsOptionsModel.SetXGlobalTransactionID("testString")
				getDirectorSitesPvdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getDirectorSitesPvdcsOptionsModel).ToNot(BeNil())
				Expect(getDirectorSitesPvdcsOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(getDirectorSitesPvdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVdcOptions successfully`, func() {
				// Construct an instance of the GetVdcOptions model
				id := "testString"
				getVdcOptionsModel := vmwareService.NewGetVdcOptions(id)
				getVdcOptionsModel.SetID("testString")
				getVdcOptionsModel.SetAcceptLanguage("testString")
				getVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVdcOptionsModel).ToNot(BeNil())
				Expect(getVdcOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(getVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListDirectorSiteHostProfilesOptions successfully`, func() {
				// Construct an instance of the ListDirectorSiteHostProfilesOptions model
				listDirectorSiteHostProfilesOptionsModel := vmwareService.NewListDirectorSiteHostProfilesOptions()
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
				listDirectorSiteRegionsOptionsModel := vmwareService.NewListDirectorSiteRegionsOptions()
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
				listDirectorSitesOptionsModel := vmwareService.NewListDirectorSitesOptions()
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
				listDirectorSitesPvdcsClustersOptionsModel := vmwareService.NewListDirectorSitesPvdcsClustersOptions(siteID, pvdcID)
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
				listDirectorSitesPvdcsOptionsModel := vmwareService.NewListDirectorSitesPvdcsOptions(siteID)
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
			It(`Invoke NewListMultitenantDirectorSitesOptions successfully`, func() {
				// Construct an instance of the ListMultitenantDirectorSitesOptions model
				listMultitenantDirectorSitesOptionsModel := vmwareService.NewListMultitenantDirectorSitesOptions()
				listMultitenantDirectorSitesOptionsModel.SetAcceptLanguage("testString")
				listMultitenantDirectorSitesOptionsModel.SetXGlobalTransactionID("testString")
				listMultitenantDirectorSitesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listMultitenantDirectorSitesOptionsModel).ToNot(BeNil())
				Expect(listMultitenantDirectorSitesOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listMultitenantDirectorSitesOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(listMultitenantDirectorSitesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVdcsOptions successfully`, func() {
				// Construct an instance of the ListVdcsOptions model
				listVdcsOptionsModel := vmwareService.NewListVdcsOptions()
				listVdcsOptionsModel.SetAcceptLanguage("testString")
				listVdcsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVdcsOptionsModel).ToNot(BeNil())
				Expect(listVdcsOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(listVdcsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPVDCPrototype successfully`, func() {
				name := "testString"
				dataCenterName := "testString"
				clusters := []vmwarev1.ClusterPrototype{}
				_model, err := vmwareService.NewPVDCPrototype(name, dataCenterName, clusters)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewResourceGroupIdentity successfully`, func() {
				id := "testString"
				_model, err := vmwareService.NewResourceGroupIdentity(id)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewServiceIdentity successfully`, func() {
				name := "veeam"
				_model, err := vmwareService.NewServiceIdentity(name)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateDirectorSitesPvdcsClusterOptions successfully`, func() {
				// Construct an instance of the UpdateDirectorSitesPvdcsClusterOptions model
				siteID := "testString"
				id := "testString"
				pvdcID := "testString"
				body := map[string]interface{}{"anyKey": "anyValue"}
				updateDirectorSitesPvdcsClusterOptionsModel := vmwareService.NewUpdateDirectorSitesPvdcsClusterOptions(siteID, id, pvdcID, body)
				updateDirectorSitesPvdcsClusterOptionsModel.SetSiteID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetPvdcID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetBody(map[string]interface{}{"anyKey": "anyValue"})
				updateDirectorSitesPvdcsClusterOptionsModel.SetAcceptLanguage("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetXGlobalTransactionID("testString")
				updateDirectorSitesPvdcsClusterOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateDirectorSitesPvdcsClusterOptionsModel).ToNot(BeNil())
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.SiteID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.PvdcID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.Body).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.XGlobalTransactionID).To(Equal(core.StringPtr("testString")))
				Expect(updateDirectorSitesPvdcsClusterOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewUpdateVdcOptions successfully`, func() {
				// Construct an instance of the UpdateVdcOptions model
				id := "testString"
				vdcPatch := map[string]interface{}{"anyKey": "anyValue"}
				updateVdcOptionsModel := vmwareService.NewUpdateVdcOptions(id, vdcPatch)
				updateVdcOptionsModel.SetID("testString")
				updateVdcOptionsModel.SetVDCPatch(map[string]interface{}{"anyKey": "anyValue"})
				updateVdcOptionsModel.SetAcceptLanguage("testString")
				updateVdcOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateVdcOptionsModel).ToNot(BeNil())
				Expect(updateVdcOptionsModel.ID).To(Equal(core.StringPtr("testString")))
				Expect(updateVdcOptionsModel.VDCPatch).To(Equal(map[string]interface{}{"anyKey": "anyValue"}))
				Expect(updateVdcOptionsModel.AcceptLanguage).To(Equal(core.StringPtr("testString")))
				Expect(updateVdcOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewVDCDirectorSitePrototype successfully`, func() {
				id := "testString"
				var pvdc *vmwarev1.DirectorSitePVDC = nil
				_, err := vmwareService.NewVDCDirectorSitePrototype(id, pvdc)
				Expect(err).ToNot(BeNil())
			})
			It(`Invoke NewVDCEdgePrototype successfully`, func() {
				typeVar := "performance"
				_model, err := vmwareService.NewVDCEdgePrototype(typeVar)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewVDCProviderType successfully`, func() {
				name := "paygo"
				_model, err := vmwareService.NewVDCProviderType(name)
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
