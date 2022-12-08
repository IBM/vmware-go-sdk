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

/*
 * IBM OpenAPI SDK Code Generator Version: 3.54.2-6c0e29d4-20220824-204545
 */

// Package vmwareasaserviceapiv1 : Operations and models for the VMwareAsAServiceApiV1 service
package vmwareasaserviceapiv1

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/IBM/vmware-go-sdk/common"
	"github.com/go-openapi/strfmt"
)

// VMwareAsAServiceApiV1 : IBM Cloud for VMware as a Service API
//
// API Version: 1.0
type VMwareAsAServiceApiV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://v-mware-as-a-service-api.cloud.ibm.com/v1"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "v_mware_as_a_service_api"

// VMwareAsAServiceApiV1Options : Service options
type VMwareAsAServiceApiV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewVMwareAsAServiceApiV1UsingExternalConfig : constructs an instance of VMwareAsAServiceApiV1 with passed in options and external configuration.
func NewVMwareAsAServiceApiV1UsingExternalConfig(options *VMwareAsAServiceApiV1Options) (vMwareAsAServiceApi *VMwareAsAServiceApiV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	if options.Authenticator == nil {
		options.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	vMwareAsAServiceApi, err = NewVMwareAsAServiceApiV1(options)
	if err != nil {
		return
	}

	err = vMwareAsAServiceApi.Service.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = vMwareAsAServiceApi.Service.SetServiceURL(options.URL)
	}
	return
}

// NewVMwareAsAServiceApiV1 : constructs an instance of VMwareAsAServiceApiV1 with passed in options.
func NewVMwareAsAServiceApiV1(options *VMwareAsAServiceApiV1Options) (service *VMwareAsAServiceApiV1, err error) {
	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &VMwareAsAServiceApiV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "vMwareAsAServiceApi" suitable for processing requests.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) Clone() *VMwareAsAServiceApiV1 {
	if core.IsNil(vMwareAsAServiceApi) {
		return nil
	}
	clone := *vMwareAsAServiceApi
	clone.Service = vMwareAsAServiceApi.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) SetServiceURL(url string) error {
	return vMwareAsAServiceApi.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetServiceURL() string {
	return vMwareAsAServiceApi.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) SetDefaultHeaders(headers http.Header) {
	vMwareAsAServiceApi.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) SetEnableGzipCompression(enableGzip bool) {
	vMwareAsAServiceApi.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetEnableGzipCompression() bool {
	return vMwareAsAServiceApi.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	vMwareAsAServiceApi.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DisableRetries() {
	vMwareAsAServiceApi.Service.DisableRetries()
}

// CreateDirectorSites : Create a director site instance
// Create a new instance of a director site with specified configurations. The director site instance is the
// infrastructure and associated VMware software stack consisting of vCenter, NSX-T, and VMware Cloud Director. VMware
// platform management and operations are performed with VMware Cloud Director. The minimum initial order size is 2
// hosts (2-Socket 32 Cores, 192 GB RAM) with 24 TB of 2.0 IOPS/GB storage.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateDirectorSites(createDirectorSitesOptions *CreateDirectorSitesOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.CreateDirectorSitesWithContext(context.Background(), createDirectorSitesOptions)
}

// CreateDirectorSitesWithContext is an alternate form of the CreateDirectorSites method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateDirectorSitesWithContext(ctx context.Context, createDirectorSitesOptions *CreateDirectorSitesOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDirectorSitesOptions, "createDirectorSitesOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDirectorSitesOptions, "createDirectorSitesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDirectorSitesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "CreateDirectorSites")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDirectorSitesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createDirectorSitesOptions.AcceptLanguage))
	}
	if createDirectorSitesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*createDirectorSitesOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if createDirectorSitesOptions.Name != nil {
		body["name"] = createDirectorSitesOptions.Name
	}
	if createDirectorSitesOptions.ResourceGroup != nil {
		body["resource_group"] = createDirectorSitesOptions.ResourceGroup
	}
	if createDirectorSitesOptions.Pvdcs != nil {
		body["pvdcs"] = createDirectorSitesOptions.Pvdcs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDirectorSites : List director site instances
// List all VMware director site instances that the user can access in the cloud account.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSites(listDirectorSitesOptions *ListDirectorSitesOptions) (result *ListDirectorSites, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListDirectorSitesWithContext(context.Background(), listDirectorSitesOptions)
}

// ListDirectorSitesWithContext is an alternate form of the ListDirectorSites method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSitesWithContext(ctx context.Context, listDirectorSitesOptions *ListDirectorSitesOptions) (result *ListDirectorSites, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDirectorSitesOptions, "listDirectorSitesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDirectorSitesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListDirectorSites")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDirectorSitesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listDirectorSitesOptions.AcceptLanguage))
	}
	if listDirectorSitesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listDirectorSitesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListDirectorSites)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDirectorSite : Get a director site instance
// Get a director site instance by specifying the instance ID.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorSite(getDirectorSiteOptions *GetDirectorSiteOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.GetDirectorSiteWithContext(context.Background(), getDirectorSiteOptions)
}

// GetDirectorSiteWithContext is an alternate form of the GetDirectorSite method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorSiteWithContext(ctx context.Context, getDirectorSiteOptions *GetDirectorSiteOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDirectorSiteOptions, "getDirectorSiteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDirectorSiteOptions, "getDirectorSiteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *getDirectorSiteOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDirectorSiteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "GetDirectorSite")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDirectorSiteOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getDirectorSiteOptions.AcceptLanguage))
	}
	if getDirectorSiteOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getDirectorSiteOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDirectorSite : Delete a director site instance
// Delete a director site instance by specifying the instance ID.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteDirectorSite(deleteDirectorSiteOptions *DeleteDirectorSiteOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.DeleteDirectorSiteWithContext(context.Background(), deleteDirectorSiteOptions)
}

// DeleteDirectorSiteWithContext is an alternate form of the DeleteDirectorSite method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteDirectorSiteWithContext(ctx context.Context, deleteDirectorSiteOptions *DeleteDirectorSiteOptions) (result *DirectorSite, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDirectorSiteOptions, "deleteDirectorSiteOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDirectorSiteOptions, "deleteDirectorSiteOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *deleteDirectorSiteOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDirectorSiteOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "DeleteDirectorSite")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteDirectorSiteOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*deleteDirectorSiteOptions.AcceptLanguage))
	}
	if deleteDirectorSiteOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*deleteDirectorSiteOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSite)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDirectorSitesPvdcs : List the provider virtual data centers in a director site instance
// List the provider virtual data centers in a specified director site.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSitesPvdcs(listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions) (result *ListPVDCs, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListDirectorSitesPvdcsWithContext(context.Background(), listDirectorSitesPvdcsOptions)
}

// ListDirectorSitesPvdcsWithContext is an alternate form of the ListDirectorSitesPvdcs method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSitesPvdcsWithContext(ctx context.Context, listDirectorSitesPvdcsOptions *ListDirectorSitesPvdcsOptions) (result *ListPVDCs, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDirectorSitesPvdcsOptions, "listDirectorSitesPvdcsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDirectorSitesPvdcsOptions, "listDirectorSitesPvdcsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *listDirectorSitesPvdcsOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDirectorSitesPvdcsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListDirectorSitesPvdcs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDirectorSitesPvdcsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listDirectorSitesPvdcsOptions.AcceptLanguage))
	}
	if listDirectorSitesPvdcsOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listDirectorSitesPvdcsOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListPVDCs)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateDirectorSitesPvdcs : Create a provider virtual data center instance in a specified director site
// Create a new instance of a provider virtual datacentre with specified configurations. The director site instance is
// the infrastructure and associated VMware software stack consisting of vCenter, NSX-T, and VMware Cloud Director.
// VMware platform management and operations are performed with VMware Cloud Director. The minimum initial order size is
// 2 hosts (2-Socket 32 Cores, 192 GB RAM) with 24 TB of 2.0 IOPS/GB storage.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateDirectorSitesPvdcs(createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions) (result *PVDCResponse, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.CreateDirectorSitesPvdcsWithContext(context.Background(), createDirectorSitesPvdcsOptions)
}

// CreateDirectorSitesPvdcsWithContext is an alternate form of the CreateDirectorSitesPvdcs method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateDirectorSitesPvdcsWithContext(ctx context.Context, createDirectorSitesPvdcsOptions *CreateDirectorSitesPvdcsOptions) (result *PVDCResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createDirectorSitesPvdcsOptions, "createDirectorSitesPvdcsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createDirectorSitesPvdcsOptions, "createDirectorSitesPvdcsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *createDirectorSitesPvdcsOptions.SiteID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range createDirectorSitesPvdcsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "CreateDirectorSitesPvdcs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createDirectorSitesPvdcsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createDirectorSitesPvdcsOptions.AcceptLanguage))
	}
	if createDirectorSitesPvdcsOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*createDirectorSitesPvdcsOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if createDirectorSitesPvdcsOptions.Name != nil {
		body["name"] = createDirectorSitesPvdcsOptions.Name
	}
	if createDirectorSitesPvdcsOptions.DataCenter != nil {
		body["data_center"] = createDirectorSitesPvdcsOptions.DataCenter
	}
	if createDirectorSitesPvdcsOptions.Clusters != nil {
		body["clusters"] = createDirectorSitesPvdcsOptions.Clusters
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPVDCResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDirectorSitesPvdcs : Get the specified provider virtual data center in a director site instance
// Get the specified provider virtual data centers in a specified director site.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorSitesPvdcs(getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions) (result *PVDCSummary, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.GetDirectorSitesPvdcsWithContext(context.Background(), getDirectorSitesPvdcsOptions)
}

// GetDirectorSitesPvdcsWithContext is an alternate form of the GetDirectorSitesPvdcs method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorSitesPvdcsWithContext(ctx context.Context, getDirectorSitesPvdcsOptions *GetDirectorSitesPvdcsOptions) (result *PVDCSummary, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDirectorSitesPvdcsOptions, "getDirectorSitesPvdcsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDirectorSitesPvdcsOptions, "getDirectorSitesPvdcsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *getDirectorSitesPvdcsOptions.SiteID,
		"pvdc_id": *getDirectorSitesPvdcsOptions.PvdcID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs/{pvdc_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDirectorSitesPvdcsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "GetDirectorSitesPvdcs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDirectorSitesPvdcsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getDirectorSitesPvdcsOptions.AcceptLanguage))
	}
	if getDirectorSitesPvdcsOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getDirectorSitesPvdcsOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPVDCSummary)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDirectorSitesPvdcsClusters : List clusters
// List all VMware clusters of a director site instance by specifying the ID of the instance.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSitesPvdcsClusters(listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions) (result *ListClusters, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListDirectorSitesPvdcsClustersWithContext(context.Background(), listDirectorSitesPvdcsClustersOptions)
}

// ListDirectorSitesPvdcsClustersWithContext is an alternate form of the ListDirectorSitesPvdcsClusters method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSitesPvdcsClustersWithContext(ctx context.Context, listDirectorSitesPvdcsClustersOptions *ListDirectorSitesPvdcsClustersOptions) (result *ListClusters, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listDirectorSitesPvdcsClustersOptions, "listDirectorSitesPvdcsClustersOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listDirectorSitesPvdcsClustersOptions, "listDirectorSitesPvdcsClustersOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *listDirectorSitesPvdcsClustersOptions.SiteID,
		"pvdc_id": *listDirectorSitesPvdcsClustersOptions.PvdcID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs/{pvdc_id}/clusters`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDirectorSitesPvdcsClustersOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListDirectorSitesPvdcsClusters")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDirectorSitesPvdcsClustersOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listDirectorSitesPvdcsClustersOptions.AcceptLanguage))
	}
	if listDirectorSitesPvdcsClustersOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listDirectorSitesPvdcsClustersOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListClusters)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetDirectorInstancesPvdcsCluster : Get a cluster
// Get a specific VMware cluster from the provider virtual data center in a director site instance.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorInstancesPvdcsCluster(getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.GetDirectorInstancesPvdcsClusterWithContext(context.Background(), getDirectorInstancesPvdcsClusterOptions)
}

// GetDirectorInstancesPvdcsClusterWithContext is an alternate form of the GetDirectorInstancesPvdcsCluster method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetDirectorInstancesPvdcsClusterWithContext(ctx context.Context, getDirectorInstancesPvdcsClusterOptions *GetDirectorInstancesPvdcsClusterOptions) (result *Cluster, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getDirectorInstancesPvdcsClusterOptions, "getDirectorInstancesPvdcsClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getDirectorInstancesPvdcsClusterOptions, "getDirectorInstancesPvdcsClusterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *getDirectorInstancesPvdcsClusterOptions.SiteID,
		"cluster_id": *getDirectorInstancesPvdcsClusterOptions.ClusterID,
		"pvdc_id": *getDirectorInstancesPvdcsClusterOptions.PvdcID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs/{pvdc_id}/clusters/{cluster_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getDirectorInstancesPvdcsClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "GetDirectorInstancesPvdcsCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getDirectorInstancesPvdcsClusterOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getDirectorInstancesPvdcsClusterOptions.AcceptLanguage))
	}
	if getDirectorInstancesPvdcsClusterOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getDirectorInstancesPvdcsClusterOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCluster)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteDirectorSitesPvdcsCluster : Delete a cluster
// Delete a cluster from an existing provider virtual data center in director site instance by specifying the instance
// ID.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteDirectorSitesPvdcsCluster(deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions) (result *PVDCResponse, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.DeleteDirectorSitesPvdcsClusterWithContext(context.Background(), deleteDirectorSitesPvdcsClusterOptions)
}

// DeleteDirectorSitesPvdcsClusterWithContext is an alternate form of the DeleteDirectorSitesPvdcsCluster method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteDirectorSitesPvdcsClusterWithContext(ctx context.Context, deleteDirectorSitesPvdcsClusterOptions *DeleteDirectorSitesPvdcsClusterOptions) (result *PVDCResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteDirectorSitesPvdcsClusterOptions, "deleteDirectorSitesPvdcsClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteDirectorSitesPvdcsClusterOptions, "deleteDirectorSitesPvdcsClusterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *deleteDirectorSitesPvdcsClusterOptions.SiteID,
		"cluster_id": *deleteDirectorSitesPvdcsClusterOptions.ClusterID,
		"pvdc_id": *deleteDirectorSitesPvdcsClusterOptions.PvdcID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs/{pvdc_id}/clusters/{cluster_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteDirectorSitesPvdcsClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "DeleteDirectorSitesPvdcsCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteDirectorSitesPvdcsClusterOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*deleteDirectorSitesPvdcsClusterOptions.AcceptLanguage))
	}
	if deleteDirectorSitesPvdcsClusterOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*deleteDirectorSitesPvdcsClusterOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPVDCResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateDirectorSitesPvdcsCluster : Update a cluster
// Update the number of hosts or file storage shares of a specific cluster in a specific director site instance. VMware
// clusters must have between [2-25] hosts.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) UpdateDirectorSitesPvdcsCluster(updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions) (result *UpdateClusterResponse, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.UpdateDirectorSitesPvdcsClusterWithContext(context.Background(), updateDirectorSitesPvdcsClusterOptions)
}

// UpdateDirectorSitesPvdcsClusterWithContext is an alternate form of the UpdateDirectorSitesPvdcsCluster method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) UpdateDirectorSitesPvdcsClusterWithContext(ctx context.Context, updateDirectorSitesPvdcsClusterOptions *UpdateDirectorSitesPvdcsClusterOptions) (result *UpdateClusterResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateDirectorSitesPvdcsClusterOptions, "updateDirectorSitesPvdcsClusterOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateDirectorSitesPvdcsClusterOptions, "updateDirectorSitesPvdcsClusterOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"site_id": *updateDirectorSitesPvdcsClusterOptions.SiteID,
		"cluster_id": *updateDirectorSitesPvdcsClusterOptions.ClusterID,
		"pvdc_id": *updateDirectorSitesPvdcsClusterOptions.PvdcID,
	}

	builder := core.NewRequestBuilder(core.PATCH)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_sites/{site_id}/pvdcs/{pvdc_id}/clusters/{cluster_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateDirectorSitesPvdcsClusterOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "UpdateDirectorSitesPvdcsCluster")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json-patch+json")
	if updateDirectorSitesPvdcsClusterOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*updateDirectorSitesPvdcsClusterOptions.AcceptLanguage))
	}
	if updateDirectorSitesPvdcsClusterOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*updateDirectorSitesPvdcsClusterOptions.XGlobalTransactionID))
	}

	_, err = builder.SetBodyContentJSON(updateDirectorSitesPvdcsClusterOptions.Body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalUpdateClusterResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDirectorSiteRegions : List regions
// List all IBM Cloud regions enabled for users to create a new director site instance.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSiteRegions(listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions) (result *DirectorSiteRegions, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListDirectorSiteRegionsWithContext(context.Background(), listDirectorSiteRegionsOptions)
}

// ListDirectorSiteRegionsWithContext is an alternate form of the ListDirectorSiteRegions method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSiteRegionsWithContext(ctx context.Context, listDirectorSiteRegionsOptions *ListDirectorSiteRegionsOptions) (result *DirectorSiteRegions, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDirectorSiteRegionsOptions, "listDirectorSiteRegionsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_site_regions`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDirectorSiteRegionsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListDirectorSiteRegions")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDirectorSiteRegionsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listDirectorSiteRegionsOptions.AcceptLanguage))
	}
	if listDirectorSiteRegionsOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listDirectorSiteRegionsOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSiteRegions)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListDirectorSiteHostProfiles : List host profiles
// List available host profiles that could be used when creating a director site instance. IBM Cloud offers several
// different host types. Typically, the host type is selected based on the properties of the workload to be run in the
// VMware cluster.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSiteHostProfiles(listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions) (result *ListHostProfiles, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListDirectorSiteHostProfilesWithContext(context.Background(), listDirectorSiteHostProfilesOptions)
}

// ListDirectorSiteHostProfilesWithContext is an alternate form of the ListDirectorSiteHostProfiles method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListDirectorSiteHostProfilesWithContext(ctx context.Context, listDirectorSiteHostProfilesOptions *ListDirectorSiteHostProfilesOptions) (result *ListHostProfiles, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listDirectorSiteHostProfilesOptions, "listDirectorSiteHostProfilesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_site_host_profiles`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listDirectorSiteHostProfilesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListDirectorSiteHostProfiles")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listDirectorSiteHostProfilesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listDirectorSiteHostProfilesOptions.AcceptLanguage))
	}
	if listDirectorSiteHostProfilesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listDirectorSiteHostProfilesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListHostProfiles)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ReplaceOrgAdminPassword : Replace the password of VMware Cloud Director tenant portal
// Replace the admin password used to log on to the VMware Cloud Director tenant portal and return the new value. VMware
// Cloud Director has its own authentication and authorization model. The first time that you access the VMware Cloud
// Director console you must set the admin credentials to generate an initial, complex, and random password. After the
// first admin password is generated, the VMware Cloud Director console option is enabled on the VDC details page. IBM
// Cloud does not capture the password. If the password is lost it needs to be reset.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ReplaceOrgAdminPassword(replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) (result *NewPassword, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ReplaceOrgAdminPasswordWithContext(context.Background(), replaceOrgAdminPasswordOptions)
}

// ReplaceOrgAdminPasswordWithContext is an alternate form of the ReplaceOrgAdminPassword method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ReplaceOrgAdminPasswordWithContext(ctx context.Context, replaceOrgAdminPasswordOptions *ReplaceOrgAdminPasswordOptions) (result *NewPassword, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(replaceOrgAdminPasswordOptions, "replaceOrgAdminPasswordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(replaceOrgAdminPasswordOptions, "replaceOrgAdminPasswordOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_site_password`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range replaceOrgAdminPasswordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ReplaceOrgAdminPassword")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("site_id", fmt.Sprint(*replaceOrgAdminPasswordOptions.SiteID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalNewPassword)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListPrices : List billing metrics
// List all billing metrics and associated prices.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListPrices(listPricesOptions *ListPricesOptions) (result *DirectorSitePricingInfo, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListPricesWithContext(context.Background(), listPricesOptions)
}

// ListPricesWithContext is an alternate form of the ListPrices method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListPricesWithContext(ctx context.Context, listPricesOptions *ListPricesOptions) (result *DirectorSitePricingInfo, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listPricesOptions, "listPricesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_site_pricing`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listPricesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListPrices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listPricesOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listPricesOptions.AcceptLanguage))
	}
	if listPricesOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*listPricesOptions.XGlobalTransactionID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSitePricingInfo)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVcddPrice : Quote price
// Quote price for a specific director site instance configuration.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetVcddPrice(getVcddPriceOptions *GetVcddPriceOptions) (result *DirectorSitePriceQuoteResponse, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.GetVcddPriceWithContext(context.Background(), getVcddPriceOptions)
}

// GetVcddPriceWithContext is an alternate form of the GetVcddPrice method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetVcddPriceWithContext(ctx context.Context, getVcddPriceOptions *GetVcddPriceOptions) (result *DirectorSitePriceQuoteResponse, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVcddPriceOptions, "getVcddPriceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVcddPriceOptions, "getVcddPriceOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/director_site_price_quote`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVcddPriceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "GetVcddPrice")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if getVcddPriceOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getVcddPriceOptions.AcceptLanguage))
	}
	if getVcddPriceOptions.XGlobalTransactionID != nil {
		builder.AddHeader("X-Global-Transaction-ID", fmt.Sprint(*getVcddPriceOptions.XGlobalTransactionID))
	}

	body := make(map[string]interface{})
	if getVcddPriceOptions.Name != nil {
		body["name"] = getVcddPriceOptions.Name
	}
	if getVcddPriceOptions.ResourceGroup != nil {
		body["resource_group"] = getVcddPriceOptions.ResourceGroup
	}
	if getVcddPriceOptions.Pvdcs != nil {
		body["pvdcs"] = getVcddPriceOptions.Pvdcs
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalDirectorSitePriceQuoteResponse)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListVdcs : List Virtual Data Centers
// List all Virtual Data Centers that user has access to in the cloud account.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListVdcs(listVdcsOptions *ListVdcsOptions) (result *ListVDCs, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.ListVdcsWithContext(context.Background(), listVdcsOptions)
}

// ListVdcsWithContext is an alternate form of the ListVdcs method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) ListVdcsWithContext(ctx context.Context, listVdcsOptions *ListVdcsOptions) (result *ListVDCs, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVdcsOptions, "listVdcsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/vdcs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVdcsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "ListVdcs")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if listVdcsOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*listVdcsOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalListVDCs)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateVdc : Create a Virtual Data Center
// Create a new Virtual Data Center with specified configurations.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateVdc(createVdcOptions *CreateVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.CreateVdcWithContext(context.Background(), createVdcOptions)
}

// CreateVdcWithContext is an alternate form of the CreateVdc method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) CreateVdcWithContext(ctx context.Context, createVdcOptions *CreateVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createVdcOptions, "createVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createVdcOptions, "createVdcOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/vdcs`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "CreateVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")
	if createVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*createVdcOptions.AcceptLanguage))
	}

	body := make(map[string]interface{})
	if createVdcOptions.Name != nil {
		body["name"] = createVdcOptions.Name
	}
	if createVdcOptions.DirectorSite != nil {
		body["director_site"] = createVdcOptions.DirectorSite
	}
	if createVdcOptions.Edge != nil {
		body["edge"] = createVdcOptions.Edge
	}
	if createVdcOptions.ResourceGroup != nil {
		body["resource_group"] = createVdcOptions.ResourceGroup
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVdc : Get a Virtual Data Center
// Get details about a Virtual Data Center by specifying the VDC ID.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetVdc(getVdcOptions *GetVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.GetVdcWithContext(context.Background(), getVdcOptions)
}

// GetVdcWithContext is an alternate form of the GetVdc method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) GetVdcWithContext(ctx context.Context, getVdcOptions *GetVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVdcOptions, "getVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVdcOptions, "getVdcOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"vdc_id": *getVdcOptions.VdcID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/vdcs/{vdc_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "GetVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if getVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*getVdcOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteVdc : Delete a Virtual Data Center
// Delete a Virtual Data Center by specifying the VDC ID.
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteVdc(deleteVdcOptions *DeleteVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	return vMwareAsAServiceApi.DeleteVdcWithContext(context.Background(), deleteVdcOptions)
}

// DeleteVdcWithContext is an alternate form of the DeleteVdc method which supports a Context parameter
func (vMwareAsAServiceApi *VMwareAsAServiceApiV1) DeleteVdcWithContext(ctx context.Context, deleteVdcOptions *DeleteVdcOptions) (result *VDC, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteVdcOptions, "deleteVdcOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteVdcOptions, "deleteVdcOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"vdc_id": *deleteVdcOptions.VdcID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = vMwareAsAServiceApi.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(vMwareAsAServiceApi.Service.Options.URL, `/vdcs/{vdc_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteVdcOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("v_mware_as_a_service_api", "V1", "DeleteVdc")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	if deleteVdcOptions.AcceptLanguage != nil {
		builder.AddHeader("Accept-Language", fmt.Sprint(*deleteVdcOptions.AcceptLanguage))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = vMwareAsAServiceApi.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVDC)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Cluster : A cluster resource.
type Cluster struct {
	// The cluster ID.
	ID *string `json:"id,omitempty"`

	// The cluster name.
	Name *string `json:"name,omitempty"`

	// The hyperlink of the cluster resource.
	Href *string `json:"href,omitempty"`

	// The time that the instance is ordered.
	InstanceOrdered *strfmt.DateTime `json:"instance_ordered,omitempty"`

	// The time that the instance is created.
	InstanceCreated *strfmt.DateTime `json:"instance_created,omitempty"`

	// The number of hosts in the cluster.
	HostCount *int64 `json:"host_count,omitempty"`

	// The status of the director site cluster.
	Status *string `json:"status,omitempty"`

	// The ID of the provider virtual data center.
	PvdcID *string `json:"pvdc_id,omitempty"`

	// The ID of the director site.
	DirectorSite *string `json:"director_site,omitempty"`

	// The name of the host profile.
	HostProfile *string `json:"host_profile,omitempty"`

	// The storage type of the cluster.
	StorageType *string `json:"storage_type,omitempty"`

	// The billing plan for the cluster.
	BillingPlan *string `json:"billing_plan,omitempty"`

	// The chosen storage policies and their sizes.
	FileShares interface{} `json:"file_shares,omitempty"`
}

// Constants associated with the Cluster.StorageType property.
// The storage type of the cluster.
const (
	Cluster_StorageType_Nfs = "nfs"
)

// Constants associated with the Cluster.BillingPlan property.
// The billing plan for the cluster.
const (
	Cluster_BillingPlan_Monthly = "monthly"
)

// UnmarshalCluster unmarshals an instance of Cluster from the specified map of raw messages.
func UnmarshalCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Cluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_ordered", &obj.InstanceOrdered)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_created", &obj.InstanceCreated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "pvdc_id", &obj.PvdcID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "director_site", &obj.DirectorSite)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_type", &obj.StorageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "billing_plan", &obj.BillingPlan)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "file_shares", &obj.FileShares)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterOrderInfo : VMware Cluster order information. Clusters form VMware workload availibility boundaries.
type ClusterOrderInfo struct {
	// Name of the VMware cluster. Cluster names must be unique per director site instance. Cluster names cannot be changed
	// after creation.
	Name *string `json:"name" validate:"required"`

	// The storage type of the cluster.
	StorageType *string `json:"storage_type" validate:"required"`

	// Number of hosts in the VMware cluster.
	HostCount *int64 `json:"host_count" validate:"required"`

	// Chosen storage policies and their sizes.
	FileShares *FileShares `json:"file_shares" validate:"required"`

	// The host type. IBM Cloud offers several different host types. Typically, the host type is selected based on the
	// properties of the workload to be run in the VMware cluster.
	HostProfile *string `json:"host_profile" validate:"required"`
}

// Constants associated with the ClusterOrderInfo.StorageType property.
// The storage type of the cluster.
const (
	ClusterOrderInfo_StorageType_Nfs = "nfs"
)

// NewClusterOrderInfo : Instantiate ClusterOrderInfo (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewClusterOrderInfo(name string, storageType string, hostCount int64, fileShares *FileShares, hostProfile string) (_model *ClusterOrderInfo, err error) {
	_model = &ClusterOrderInfo{
		Name: core.StringPtr(name),
		StorageType: core.StringPtr(storageType),
		HostCount: core.Int64Ptr(hostCount),
		FileShares: fileShares,
		HostProfile: core.StringPtr(hostProfile),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalClusterOrderInfo unmarshals an instance of ClusterOrderInfo from the specified map of raw messages.
func UnmarshalClusterOrderInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterOrderInfo)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_type", &obj.StorageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "file_shares", &obj.FileShares, UnmarshalFileShares)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ClusterSummary : VMware Cluster order information. Clusters form VMware workload availibility boundaries.
type ClusterSummary struct {
	// Name of the VMware cluster. Cluster names must be unique per director site instance. Cluster names cannot be changed
	// after creation.
	Name *string `json:"name" validate:"required"`

	// The storage type of the cluster.
	StorageType *string `json:"storage_type" validate:"required"`

	// Number of hosts in the VMware cluster.
	HostCount *int64 `json:"host_count" validate:"required"`

	// Chosen storage policies and their sizes.
	FileShares *FileShares `json:"file_shares" validate:"required"`

	// The host type. IBM Cloud offers several different host types. Typically, the host type is selected based on the
	// properties of the workload to be run in the VMware cluster.
	HostProfile *string `json:"host_profile" validate:"required"`

	// The cluster ID.
	ID *string `json:"id,omitempty"`

	// The hyperlink of the cluster resource.
	Href *string `json:"href,omitempty"`
}

// Constants associated with the ClusterSummary.StorageType property.
// The storage type of the cluster.
const (
	ClusterSummary_StorageType_Nfs = "nfs"
)

// UnmarshalClusterSummary unmarshals an instance of ClusterSummary from the specified map of raw messages.
func UnmarshalClusterSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ClusterSummary)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "storage_type", &obj.StorageType)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_count", &obj.HostCount)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "file_shares", &obj.FileShares, UnmarshalFileShares)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "host_profile", &obj.HostProfile)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CreateDirectorSitesOptions : The CreateDirectorSites options.
type CreateDirectorSitesOptions struct {
	// Name of the director site instance. Use a name that is unique to your region and meaningful. Names cannot be changed
	// after initial creation.
	Name *string `json:"name" validate:"required"`

	// The name or ID of the IBM resource group where the instance is deployed.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// List of VMware provider virtual data centers to deploy on the instance.
	Pvdcs []PVDCOrderInfo `json:"pvdcs" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDirectorSitesOptions : Instantiate CreateDirectorSitesOptions
func (*VMwareAsAServiceApiV1) NewCreateDirectorSitesOptions(name string, resourceGroup string, pvdcs []PVDCOrderInfo) *CreateDirectorSitesOptions {
	return &CreateDirectorSitesOptions{
		Name: core.StringPtr(name),
		ResourceGroup: core.StringPtr(resourceGroup),
		Pvdcs: pvdcs,
	}
}

// SetName : Allow user to set Name
func (_options *CreateDirectorSitesOptions) SetName(name string) *CreateDirectorSitesOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateDirectorSitesOptions) SetResourceGroup(resourceGroup string) *CreateDirectorSitesOptions {
	_options.ResourceGroup = core.StringPtr(resourceGroup)
	return _options
}

// SetPvdcs : Allow user to set Pvdcs
func (_options *CreateDirectorSitesOptions) SetPvdcs(pvdcs []PVDCOrderInfo) *CreateDirectorSitesOptions {
	_options.Pvdcs = pvdcs
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateDirectorSitesOptions) SetAcceptLanguage(acceptLanguage string) *CreateDirectorSitesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *CreateDirectorSitesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *CreateDirectorSitesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDirectorSitesOptions) SetHeaders(param map[string]string) *CreateDirectorSitesOptions {
	options.Headers = param
	return options
}

// CreateDirectorSitesPvdcsOptions : The CreateDirectorSitesPvdcs options.
type CreateDirectorSitesPvdcsOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Name of the provider virtual data center. Provider virtual data center names must be unique per director site
	// instance. Provider virtual data center names cannot be changed after creation.
	Name *string `json:"name" validate:"required"`

	// Data center location to deploy the cluster. See `GET /director_site_regions` for supported data center locations.
	DataCenter *string `json:"data_center" validate:"required"`

	// List of VMware clusters to deploy on the instance. Clusters form VMware workload availibility boundaries.
	Clusters []ClusterOrderInfo `json:"clusters" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateDirectorSitesPvdcsOptions : Instantiate CreateDirectorSitesPvdcsOptions
func (*VMwareAsAServiceApiV1) NewCreateDirectorSitesPvdcsOptions(siteID string, name string, dataCenter string, clusters []ClusterOrderInfo) *CreateDirectorSitesPvdcsOptions {
	return &CreateDirectorSitesPvdcsOptions{
		SiteID: core.StringPtr(siteID),
		Name: core.StringPtr(name),
		DataCenter: core.StringPtr(dataCenter),
		Clusters: clusters,
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *CreateDirectorSitesPvdcsOptions) SetSiteID(siteID string) *CreateDirectorSitesPvdcsOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetName : Allow user to set Name
func (_options *CreateDirectorSitesPvdcsOptions) SetName(name string) *CreateDirectorSitesPvdcsOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDataCenter : Allow user to set DataCenter
func (_options *CreateDirectorSitesPvdcsOptions) SetDataCenter(dataCenter string) *CreateDirectorSitesPvdcsOptions {
	_options.DataCenter = core.StringPtr(dataCenter)
	return _options
}

// SetClusters : Allow user to set Clusters
func (_options *CreateDirectorSitesPvdcsOptions) SetClusters(clusters []ClusterOrderInfo) *CreateDirectorSitesPvdcsOptions {
	_options.Clusters = clusters
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateDirectorSitesPvdcsOptions) SetAcceptLanguage(acceptLanguage string) *CreateDirectorSitesPvdcsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *CreateDirectorSitesPvdcsOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *CreateDirectorSitesPvdcsOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateDirectorSitesPvdcsOptions) SetHeaders(param map[string]string) *CreateDirectorSitesPvdcsOptions {
	options.Headers = param
	return options
}

// CreateVdcOptions : The CreateVdc options.
type CreateVdcOptions struct {
	// A human readable identifier for the Virtual Data Center. Use a name that is unique to your region.
	Name *string `json:"name" validate:"required"`

	// The director site in which to deploy the Virtual Data Center.
	DirectorSite *NewVDCDirectorSite `json:"director_site" validate:"required"`

	// The networking Edge to be deployed on the Virtual Data Center.
	Edge *NewVDCEdge `json:"edge,omitempty"`

	// The resource group to associate with the Virtual Data Center.
	// If not specified, the default resource group in the account is used.
	ResourceGroup *NewVDCResourceGroup `json:"resource_group,omitempty"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateVdcOptions : Instantiate CreateVdcOptions
func (*VMwareAsAServiceApiV1) NewCreateVdcOptions(name string, directorSite *NewVDCDirectorSite) *CreateVdcOptions {
	return &CreateVdcOptions{
		Name: core.StringPtr(name),
		DirectorSite: directorSite,
	}
}

// SetName : Allow user to set Name
func (_options *CreateVdcOptions) SetName(name string) *CreateVdcOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDirectorSite : Allow user to set DirectorSite
func (_options *CreateVdcOptions) SetDirectorSite(directorSite *NewVDCDirectorSite) *CreateVdcOptions {
	_options.DirectorSite = directorSite
	return _options
}

// SetEdge : Allow user to set Edge
func (_options *CreateVdcOptions) SetEdge(edge *NewVDCEdge) *CreateVdcOptions {
	_options.Edge = edge
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *CreateVdcOptions) SetResourceGroup(resourceGroup *NewVDCResourceGroup) *CreateVdcOptions {
	_options.ResourceGroup = resourceGroup
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *CreateVdcOptions) SetAcceptLanguage(acceptLanguage string) *CreateVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateVdcOptions) SetHeaders(param map[string]string) *CreateVdcOptions {
	options.Headers = param
	return options
}

// DataCenterInfo : Details of the data center.
type DataCenterInfo struct {
	// The display name of the data center.
	DisplayName *string `json:"display_name,omitempty"`

	// The name of the data center.
	Name *string `json:"name,omitempty"`

	// The speed available per data center.
	UplinkSpeed *string `json:"uplink_speed,omitempty"`
}

// UnmarshalDataCenterInfo unmarshals an instance of DataCenterInfo from the specified map of raw messages.
func UnmarshalDataCenterInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DataCenterInfo)
	err = core.UnmarshalPrimitive(m, "display_name", &obj.DisplayName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "uplink_speed", &obj.UplinkSpeed)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteDirectorSiteOptions : The DeleteDirectorSite options.
type DeleteDirectorSiteOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDirectorSiteOptions : Instantiate DeleteDirectorSiteOptions
func (*VMwareAsAServiceApiV1) NewDeleteDirectorSiteOptions(siteID string) *DeleteDirectorSiteOptions {
	return &DeleteDirectorSiteOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *DeleteDirectorSiteOptions) SetSiteID(siteID string) *DeleteDirectorSiteOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *DeleteDirectorSiteOptions) SetAcceptLanguage(acceptLanguage string) *DeleteDirectorSiteOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *DeleteDirectorSiteOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *DeleteDirectorSiteOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDirectorSiteOptions) SetHeaders(param map[string]string) *DeleteDirectorSiteOptions {
	options.Headers = param
	return options
}

// DeleteDirectorSitesPvdcsClusterOptions : The DeleteDirectorSitesPvdcsCluster options.
type DeleteDirectorSitesPvdcsClusterOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// A unique identifier for the provider virtual data center in a director site.
	PvdcID *string `json:"pvdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteDirectorSitesPvdcsClusterOptions : Instantiate DeleteDirectorSitesPvdcsClusterOptions
func (*VMwareAsAServiceApiV1) NewDeleteDirectorSitesPvdcsClusterOptions(siteID string, clusterID string, pvdcID string) *DeleteDirectorSitesPvdcsClusterOptions {
	return &DeleteDirectorSitesPvdcsClusterOptions{
		SiteID: core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
		PvdcID: core.StringPtr(pvdcID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *DeleteDirectorSitesPvdcsClusterOptions) SetSiteID(siteID string) *DeleteDirectorSitesPvdcsClusterOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *DeleteDirectorSitesPvdcsClusterOptions) SetClusterID(clusterID string) *DeleteDirectorSitesPvdcsClusterOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetPvdcID : Allow user to set PvdcID
func (_options *DeleteDirectorSitesPvdcsClusterOptions) SetPvdcID(pvdcID string) *DeleteDirectorSitesPvdcsClusterOptions {
	_options.PvdcID = core.StringPtr(pvdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *DeleteDirectorSitesPvdcsClusterOptions) SetAcceptLanguage(acceptLanguage string) *DeleteDirectorSitesPvdcsClusterOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *DeleteDirectorSitesPvdcsClusterOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *DeleteDirectorSitesPvdcsClusterOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteDirectorSitesPvdcsClusterOptions) SetHeaders(param map[string]string) *DeleteDirectorSitesPvdcsClusterOptions {
	options.Headers = param
	return options
}

// DeleteVdcOptions : The DeleteVdc options.
type DeleteVdcOptions struct {
	// A unique identifier for a given Virtual Data Center.
	VdcID *string `json:"vdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteVdcOptions : Instantiate DeleteVdcOptions
func (*VMwareAsAServiceApiV1) NewDeleteVdcOptions(vdcID string) *DeleteVdcOptions {
	return &DeleteVdcOptions{
		VdcID: core.StringPtr(vdcID),
	}
}

// SetVdcID : Allow user to set VdcID
func (_options *DeleteVdcOptions) SetVdcID(vdcID string) *DeleteVdcOptions {
	_options.VdcID = core.StringPtr(vdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *DeleteVdcOptions) SetAcceptLanguage(acceptLanguage string) *DeleteVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteVdcOptions) SetHeaders(param map[string]string) *DeleteVdcOptions {
	options.Headers = param
	return options
}

// DirectorSite : A director site resource. The director site instance is the infrastructure and associated VMware software stack
// consisting of vCenter, NSX-T and VMware Cloud Director.
type DirectorSite struct {
	// A unique identifier for the director site in IBM Cloud.
	Crn *string `json:"crn,omitempty"`

	// The hyperlink of the director site resource.
	Href *string `json:"href,omitempty"`

	// ID of the director site.
	ID *string `json:"id,omitempty"`

	// The time that the instance is ordered.
	InstanceOrdered *strfmt.DateTime `json:"instance_ordered,omitempty"`

	// The time that the instance is created and available to use.
	InstanceCreated *strfmt.DateTime `json:"instance_created,omitempty"`

	// The name of director site. The name of the director site cannot be changed after creation.
	Name *string `json:"name,omitempty"`

	// The status of director site.
	Status *string `json:"status,omitempty"`

	// The name of the IBM resource group.
	ResourceGroup *string `json:"resource_group,omitempty"`

	// The email identity of the user that ordered the VMware as a Service director site instance.
	Creator *string `json:"creator,omitempty"`

	// The ID of the resource group.
	ResourceGroupID *string `json:"resource_group_id,omitempty"`

	// The CRN of the resource group.
	ResourceGroupCrn *string `json:"resource_group_crn,omitempty"`

	// List of VMware provider virtual data centers to deploy on the instance.
	Pvdcs []PVDCSummary `json:"pvdcs,omitempty"`
}

// Constants associated with the DirectorSite.Status property.
// The status of director site.
const (
	DirectorSite_Status_Creating = "Creating"
	DirectorSite_Status_Deleted = "Deleted"
	DirectorSite_Status_Deleting = "Deleting"
	DirectorSite_Status_Readytouse = "ReadyToUse"
	DirectorSite_Status_Updating = "Updating"
)

// UnmarshalDirectorSite unmarshals an instance of DirectorSite from the specified map of raw messages.
func UnmarshalDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSite)
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_ordered", &obj.InstanceOrdered)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "instance_created", &obj.InstanceCreated)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group", &obj.ResourceGroup)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "creator", &obj.Creator)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_id", &obj.ResourceGroupID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "resource_group_crn", &obj.ResourceGroupCrn)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "pvdcs", &obj.Pvdcs, UnmarshalPVDCSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceItem : sub items for a metric and associated prices.
type DirectorSitePriceItem struct {
	// The price for the metric.
	Price *float64 `json:"price,omitempty"`

	// Quantity tier.
	QuantityTier *int64 `json:"quantity_tier,omitempty"`
}

// UnmarshalDirectorSitePriceItem unmarshals an instance of DirectorSitePriceItem from the specified map of raw messages.
func UnmarshalDirectorSitePriceItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceItem)
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "quantity_tier", &obj.QuantityTier)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceListItem : items for a metric and associated prices.
type DirectorSitePriceListItem struct {
	// The country for which this price applies.
	Country *string `json:"country,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// A list of prices.
	Prices []DirectorSitePriceItem `json:"prices,omitempty"`
}

// UnmarshalDirectorSitePriceListItem unmarshals an instance of DirectorSitePriceListItem from the specified map of raw messages.
func UnmarshalDirectorSitePriceListItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceListItem)
	err = core.UnmarshalPrimitive(m, "country", &obj.Country)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prices", &obj.Prices, UnmarshalDirectorSitePriceItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceMetric : A metric and associated prices.
type DirectorSitePriceMetric struct {
	// The metric name.
	Metric *string `json:"metric,omitempty"`

	// The metric description.
	Description *string `json:"description,omitempty"`

	// A list of prices for each country.
	PriceList []DirectorSitePriceListItem `json:"price_list,omitempty"`
}

// UnmarshalDirectorSitePriceMetric unmarshals an instance of DirectorSitePriceMetric from the specified map of raw messages.
func UnmarshalDirectorSitePriceMetric(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceMetric)
	err = core.UnmarshalPrimitive(m, "metric", &obj.Metric)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "price_list", &obj.PriceList, UnmarshalDirectorSitePriceListItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePriceQuoteResponse : Return price information for a VCDD instance.
type DirectorSitePriceQuoteResponse struct {
	// Details of the instance base charge.
	BaseCharge *PriceInfoBaseCharge `json:"base_charge,omitempty"`

	// A list of the clusters with price information.
	Clusters []PriceInfoClusterCharge `json:"clusters,omitempty"`

	// The currency unit for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for the instance.
	Total *float64 `json:"total,omitempty"`
}

// UnmarshalDirectorSitePriceQuoteResponse unmarshals an instance of DirectorSitePriceQuoteResponse from the specified map of raw messages.
func UnmarshalDirectorSitePriceQuoteResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePriceQuoteResponse)
	err = core.UnmarshalModel(m, "base_charge", &obj.BaseCharge, UnmarshalPriceInfoBaseCharge)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalPriceInfoClusterCharge)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "total", &obj.Total)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSitePricingInfo : Return all metrics with associate prices.
type DirectorSitePricingInfo struct {
	// A list of metrics and associated prices.
	DirectorSitePricing []DirectorSitePriceMetric `json:"director_site_pricing,omitempty"`
}

// UnmarshalDirectorSitePricingInfo unmarshals an instance of DirectorSitePricingInfo from the specified map of raw messages.
func UnmarshalDirectorSitePricingInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSitePricingInfo)
	err = core.UnmarshalModel(m, "director_site_pricing", &obj.DirectorSitePricing, UnmarshalDirectorSitePriceMetric)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DirectorSiteRegions : Success. The request was successfully processed.
type DirectorSiteRegions struct {
	// available region.
	DirectorSiteRegions map[string]RegionDetail `json:"director_site_regions,omitempty"`
}

// UnmarshalDirectorSiteRegions unmarshals an instance of DirectorSiteRegions from the specified map of raw messages.
func UnmarshalDirectorSiteRegions(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(DirectorSiteRegions)
	err = core.UnmarshalModel(m, "director_site_regions", &obj.DirectorSiteRegions, UnmarshalRegionDetail)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Edge : A networking Edge deployed on a Virtual Data Center. Networking edges are based on NSX-T and used for bridging
// virtualize networking to the physical public-internet and IBM private networking.
type Edge struct {
	// A unique identifier for the Edge.
	ID *string `json:"id" validate:"required"`

	// The public IP addresses assigned to the Edge.
	PublicIps []string `json:"public_ips" validate:"required"`

	// The size of the Edge.
	//
	// The size can only be specified for dedicated Edges. Larger sizes require more capacity from the director site in
	// which the Virtual Data Center was created to be deployed.
	Size *string `json:"size,omitempty"`

	// The type of Edge to be deployed.
	//
	// Shared Edges allow for multiple VDCs to share some Edge resources. Dedicated Edges do not share resources between
	// VDCs.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the Edge.Size property.
// The size of the Edge.
//
// The size can only be specified for dedicated Edges. Larger sizes require more capacity from the director site in
// which the Virtual Data Center was created to be deployed.
const (
	Edge_Size_ExtraLarge = "extra_large"
	Edge_Size_Large = "large"
	Edge_Size_Medium = "medium"
)

// Constants associated with the Edge.Type property.
// The type of Edge to be deployed.
//
// Shared Edges allow for multiple VDCs to share some Edge resources. Dedicated Edges do not share resources between
// VDCs.
const (
	Edge_Type_Dedicated = "dedicated"
	Edge_Type_Shared = "shared"
)

// UnmarshalEdge unmarshals an instance of Edge from the specified map of raw messages.
func UnmarshalEdge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Edge)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "public_ips", &obj.PublicIps)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Error : Information about why a request cannot be completed or why a resource could not be created.
type Error struct {
	// An error code specific to the error encountered.
	Code *string `json:"code" validate:"required"`

	// A message describing why the error ocurred.
	Message *string `json:"message" validate:"required"`

	// A URL that links to a page with more information about this error.
	MoreInfo *string `json:"more_info,omitempty"`
}

// UnmarshalError unmarshals an instance of Error from the specified map of raw messages.
func UnmarshalError(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Error)
	err = core.UnmarshalPrimitive(m, "code", &obj.Code)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "more_info", &obj.MoreInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// FileShares : Chosen storage policies and their sizes.
type FileShares struct {
	// The amount of 0.25 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEPOINTTWOFIVEIOPSGB *int64 `json:"STORAGE_POINT_TWO_FIVE_IOPS_GB,omitempty"`

	// The amount of 2 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETWOIOPSGB *int64 `json:"STORAGE_TWO_IOPS_GB,omitempty"`

	// The amount of 4 IOPS/GB storage in GB (1024^3 bytes).
	STORAGEFOURIOPSGB *int64 `json:"STORAGE_FOUR_IOPS_GB,omitempty"`

	// The amount of 10 IOPS/GB storage in GB (1024^3 bytes).
	STORAGETENIOPSGB *int64 `json:"STORAGE_TEN_IOPS_GB,omitempty"`
}

// UnmarshalFileShares unmarshals an instance of FileShares from the specified map of raw messages.
func UnmarshalFileShares(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(FileShares)
	err = core.UnmarshalPrimitive(m, "STORAGE_POINT_TWO_FIVE_IOPS_GB", &obj.STORAGEPOINTTWOFIVEIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_TWO_IOPS_GB", &obj.STORAGETWOIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_FOUR_IOPS_GB", &obj.STORAGEFOURIOPSGB)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "STORAGE_TEN_IOPS_GB", &obj.STORAGETENIOPSGB)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// GetDirectorInstancesPvdcsClusterOptions : The GetDirectorInstancesPvdcsCluster options.
type GetDirectorInstancesPvdcsClusterOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// A unique identifier for the provider virtual data center in a director site.
	PvdcID *string `json:"pvdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDirectorInstancesPvdcsClusterOptions : Instantiate GetDirectorInstancesPvdcsClusterOptions
func (*VMwareAsAServiceApiV1) NewGetDirectorInstancesPvdcsClusterOptions(siteID string, clusterID string, pvdcID string) *GetDirectorInstancesPvdcsClusterOptions {
	return &GetDirectorInstancesPvdcsClusterOptions{
		SiteID: core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
		PvdcID: core.StringPtr(pvdcID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *GetDirectorInstancesPvdcsClusterOptions) SetSiteID(siteID string) *GetDirectorInstancesPvdcsClusterOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *GetDirectorInstancesPvdcsClusterOptions) SetClusterID(clusterID string) *GetDirectorInstancesPvdcsClusterOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetPvdcID : Allow user to set PvdcID
func (_options *GetDirectorInstancesPvdcsClusterOptions) SetPvdcID(pvdcID string) *GetDirectorInstancesPvdcsClusterOptions {
	_options.PvdcID = core.StringPtr(pvdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetDirectorInstancesPvdcsClusterOptions) SetAcceptLanguage(acceptLanguage string) *GetDirectorInstancesPvdcsClusterOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetDirectorInstancesPvdcsClusterOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetDirectorInstancesPvdcsClusterOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDirectorInstancesPvdcsClusterOptions) SetHeaders(param map[string]string) *GetDirectorInstancesPvdcsClusterOptions {
	options.Headers = param
	return options
}

// GetDirectorSiteOptions : The GetDirectorSite options.
type GetDirectorSiteOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDirectorSiteOptions : Instantiate GetDirectorSiteOptions
func (*VMwareAsAServiceApiV1) NewGetDirectorSiteOptions(siteID string) *GetDirectorSiteOptions {
	return &GetDirectorSiteOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *GetDirectorSiteOptions) SetSiteID(siteID string) *GetDirectorSiteOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetDirectorSiteOptions) SetAcceptLanguage(acceptLanguage string) *GetDirectorSiteOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetDirectorSiteOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetDirectorSiteOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDirectorSiteOptions) SetHeaders(param map[string]string) *GetDirectorSiteOptions {
	options.Headers = param
	return options
}

// GetDirectorSitesPvdcsOptions : The GetDirectorSitesPvdcs options.
type GetDirectorSitesPvdcsOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// A unique identifier for the provider virtual data center in a director site.
	PvdcID *string `json:"pvdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetDirectorSitesPvdcsOptions : Instantiate GetDirectorSitesPvdcsOptions
func (*VMwareAsAServiceApiV1) NewGetDirectorSitesPvdcsOptions(siteID string, pvdcID string) *GetDirectorSitesPvdcsOptions {
	return &GetDirectorSitesPvdcsOptions{
		SiteID: core.StringPtr(siteID),
		PvdcID: core.StringPtr(pvdcID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *GetDirectorSitesPvdcsOptions) SetSiteID(siteID string) *GetDirectorSitesPvdcsOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetPvdcID : Allow user to set PvdcID
func (_options *GetDirectorSitesPvdcsOptions) SetPvdcID(pvdcID string) *GetDirectorSitesPvdcsOptions {
	_options.PvdcID = core.StringPtr(pvdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetDirectorSitesPvdcsOptions) SetAcceptLanguage(acceptLanguage string) *GetDirectorSitesPvdcsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetDirectorSitesPvdcsOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetDirectorSitesPvdcsOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetDirectorSitesPvdcsOptions) SetHeaders(param map[string]string) *GetDirectorSitesPvdcsOptions {
	options.Headers = param
	return options
}

// GetVcddPriceOptions : The GetVcddPrice options.
type GetVcddPriceOptions struct {
	// Name of the director site instance. Use a name that is unique to your region and meaningful. Names cannot be changed
	// after initial creation.
	Name *string `json:"name" validate:"required"`

	// The name or ID of the IBM resource group where the instance is deployed.
	ResourceGroup *string `json:"resource_group" validate:"required"`

	// List of VMware provider virtual data centers to deploy on the instance.
	Pvdcs []PVDCOrderInfo `json:"pvdcs" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVcddPriceOptions : Instantiate GetVcddPriceOptions
func (*VMwareAsAServiceApiV1) NewGetVcddPriceOptions(name string, resourceGroup string, pvdcs []PVDCOrderInfo) *GetVcddPriceOptions {
	return &GetVcddPriceOptions{
		Name: core.StringPtr(name),
		ResourceGroup: core.StringPtr(resourceGroup),
		Pvdcs: pvdcs,
	}
}

// SetName : Allow user to set Name
func (_options *GetVcddPriceOptions) SetName(name string) *GetVcddPriceOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetResourceGroup : Allow user to set ResourceGroup
func (_options *GetVcddPriceOptions) SetResourceGroup(resourceGroup string) *GetVcddPriceOptions {
	_options.ResourceGroup = core.StringPtr(resourceGroup)
	return _options
}

// SetPvdcs : Allow user to set Pvdcs
func (_options *GetVcddPriceOptions) SetPvdcs(pvdcs []PVDCOrderInfo) *GetVcddPriceOptions {
	_options.Pvdcs = pvdcs
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetVcddPriceOptions) SetAcceptLanguage(acceptLanguage string) *GetVcddPriceOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *GetVcddPriceOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *GetVcddPriceOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVcddPriceOptions) SetHeaders(param map[string]string) *GetVcddPriceOptions {
	options.Headers = param
	return options
}

// GetVdcOptions : The GetVdc options.
type GetVdcOptions struct {
	// A unique identifier for a given Virtual Data Center.
	VdcID *string `json:"vdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetVdcOptions : Instantiate GetVdcOptions
func (*VMwareAsAServiceApiV1) NewGetVdcOptions(vdcID string) *GetVdcOptions {
	return &GetVdcOptions{
		VdcID: core.StringPtr(vdcID),
	}
}

// SetVdcID : Allow user to set VdcID
func (_options *GetVdcOptions) SetVdcID(vdcID string) *GetVdcOptions {
	_options.VdcID = core.StringPtr(vdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *GetVdcOptions) SetAcceptLanguage(acceptLanguage string) *GetVdcOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVdcOptions) SetHeaders(param map[string]string) *GetVdcOptions {
	options.Headers = param
	return options
}

// HostProfile : Host profile template.
type HostProfile struct {
	// The id for this host profile.
	ID *string `json:"id,omitempty"`

	// The number CPU cores for this host profile.
	Cpu *int64 `json:"cpu,omitempty"`

	// The CPU family for this host profile.
	Family *string `json:"family,omitempty"`

	// The CPU type for this host profile.
	Processor *string `json:"processor,omitempty"`

	// The RAM for this host profile in GB (1024^3 bytes).
	Ram *int64 `json:"ram,omitempty"`

	// The number of CPU sockets available for this host profile.
	Socket *int64 `json:"socket,omitempty"`

	// The CPU clock speed.
	Speed *string `json:"speed,omitempty"`

	// The manufacturer for this host profile.
	Manufacturer *string `json:"manufacturer,omitempty"`

	// Additional features for this host profile.
	Features []string `json:"features,omitempty"`
}

// UnmarshalHostProfile unmarshals an instance of HostProfile from the specified map of raw messages.
func UnmarshalHostProfile(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(HostProfile)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "cpu", &obj.Cpu)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "family", &obj.Family)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "processor", &obj.Processor)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ram", &obj.Ram)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "socket", &obj.Socket)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speed", &obj.Speed)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "manufacturer", &obj.Manufacturer)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "features", &obj.Features)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// JSONPatchOperation : This model represents an individual patch operation to be performed on a JSON document, as defined by RFC 6902.
type JSONPatchOperation struct {
	// The operation to be performed.
	Op *string `json:"op" validate:"required"`

	// The JSON Pointer that identifies the field that is the target of the operation.
	Path *string `json:"path" validate:"required"`

	// The JSON Pointer that identifies the field that is the source of the operation.
	From *string `json:"from,omitempty"`

	// The value to be used within the operation.
	Value interface{} `json:"value,omitempty"`
}

// Constants associated with the JSONPatchOperation.Op property.
// The operation to be performed.
const (
	JSONPatchOperation_Op_Add = "add"
	JSONPatchOperation_Op_Copy = "copy"
	JSONPatchOperation_Op_Move = "move"
	JSONPatchOperation_Op_Remove = "remove"
	JSONPatchOperation_Op_Replace = "replace"
	JSONPatchOperation_Op_Test = "test"
)

// NewJSONPatchOperation : Instantiate JSONPatchOperation (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewJSONPatchOperation(op string, path string) (_model *JSONPatchOperation, err error) {
	_model = &JSONPatchOperation{
		Op: core.StringPtr(op),
		Path: core.StringPtr(path),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalJSONPatchOperation unmarshals an instance of JSONPatchOperation from the specified map of raw messages.
func UnmarshalJSONPatchOperation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(JSONPatchOperation)
	err = core.UnmarshalPrimitive(m, "op", &obj.Op)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "path", &obj.Path)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "from", &obj.From)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "value", &obj.Value)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListClusters : Return all clusters instances.
type ListClusters struct {
	// list of cluster objects.
	Clusters []Cluster `json:"clusters,omitempty"`
}

// UnmarshalListClusters unmarshals an instance of ListClusters from the specified map of raw messages.
func UnmarshalListClusters(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListClusters)
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalCluster)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDirectorSiteHostProfilesOptions : The ListDirectorSiteHostProfiles options.
type ListDirectorSiteHostProfilesOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDirectorSiteHostProfilesOptions : Instantiate ListDirectorSiteHostProfilesOptions
func (*VMwareAsAServiceApiV1) NewListDirectorSiteHostProfilesOptions() *ListDirectorSiteHostProfilesOptions {
	return &ListDirectorSiteHostProfilesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListDirectorSiteHostProfilesOptions) SetAcceptLanguage(acceptLanguage string) *ListDirectorSiteHostProfilesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListDirectorSiteHostProfilesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListDirectorSiteHostProfilesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDirectorSiteHostProfilesOptions) SetHeaders(param map[string]string) *ListDirectorSiteHostProfilesOptions {
	options.Headers = param
	return options
}

// ListDirectorSiteRegionsOptions : The ListDirectorSiteRegions options.
type ListDirectorSiteRegionsOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDirectorSiteRegionsOptions : Instantiate ListDirectorSiteRegionsOptions
func (*VMwareAsAServiceApiV1) NewListDirectorSiteRegionsOptions() *ListDirectorSiteRegionsOptions {
	return &ListDirectorSiteRegionsOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListDirectorSiteRegionsOptions) SetAcceptLanguage(acceptLanguage string) *ListDirectorSiteRegionsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListDirectorSiteRegionsOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListDirectorSiteRegionsOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDirectorSiteRegionsOptions) SetHeaders(param map[string]string) *ListDirectorSiteRegionsOptions {
	options.Headers = param
	return options
}

// ListDirectorSites : Return all director site instances.
type ListDirectorSites struct {
	// List of director site instances.
	DirectorSites []DirectorSite `json:"director_sites,omitempty"`
}

// UnmarshalListDirectorSites unmarshals an instance of ListDirectorSites from the specified map of raw messages.
func UnmarshalListDirectorSites(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListDirectorSites)
	err = core.UnmarshalModel(m, "director_sites", &obj.DirectorSites, UnmarshalDirectorSite)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListDirectorSitesOptions : The ListDirectorSites options.
type ListDirectorSitesOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDirectorSitesOptions : Instantiate ListDirectorSitesOptions
func (*VMwareAsAServiceApiV1) NewListDirectorSitesOptions() *ListDirectorSitesOptions {
	return &ListDirectorSitesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListDirectorSitesOptions) SetAcceptLanguage(acceptLanguage string) *ListDirectorSitesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListDirectorSitesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListDirectorSitesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDirectorSitesOptions) SetHeaders(param map[string]string) *ListDirectorSitesOptions {
	options.Headers = param
	return options
}

// ListDirectorSitesPvdcsClustersOptions : The ListDirectorSitesPvdcsClusters options.
type ListDirectorSitesPvdcsClustersOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// A unique identifier for the provider virtual data center in a director site.
	PvdcID *string `json:"pvdc_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDirectorSitesPvdcsClustersOptions : Instantiate ListDirectorSitesPvdcsClustersOptions
func (*VMwareAsAServiceApiV1) NewListDirectorSitesPvdcsClustersOptions(siteID string, pvdcID string) *ListDirectorSitesPvdcsClustersOptions {
	return &ListDirectorSitesPvdcsClustersOptions{
		SiteID: core.StringPtr(siteID),
		PvdcID: core.StringPtr(pvdcID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *ListDirectorSitesPvdcsClustersOptions) SetSiteID(siteID string) *ListDirectorSitesPvdcsClustersOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetPvdcID : Allow user to set PvdcID
func (_options *ListDirectorSitesPvdcsClustersOptions) SetPvdcID(pvdcID string) *ListDirectorSitesPvdcsClustersOptions {
	_options.PvdcID = core.StringPtr(pvdcID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListDirectorSitesPvdcsClustersOptions) SetAcceptLanguage(acceptLanguage string) *ListDirectorSitesPvdcsClustersOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListDirectorSitesPvdcsClustersOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListDirectorSitesPvdcsClustersOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDirectorSitesPvdcsClustersOptions) SetHeaders(param map[string]string) *ListDirectorSitesPvdcsClustersOptions {
	options.Headers = param
	return options
}

// ListDirectorSitesPvdcsOptions : The ListDirectorSitesPvdcs options.
type ListDirectorSitesPvdcsOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListDirectorSitesPvdcsOptions : Instantiate ListDirectorSitesPvdcsOptions
func (*VMwareAsAServiceApiV1) NewListDirectorSitesPvdcsOptions(siteID string) *ListDirectorSitesPvdcsOptions {
	return &ListDirectorSitesPvdcsOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *ListDirectorSitesPvdcsOptions) SetSiteID(siteID string) *ListDirectorSitesPvdcsOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListDirectorSitesPvdcsOptions) SetAcceptLanguage(acceptLanguage string) *ListDirectorSitesPvdcsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListDirectorSitesPvdcsOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListDirectorSitesPvdcsOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListDirectorSitesPvdcsOptions) SetHeaders(param map[string]string) *ListDirectorSitesPvdcsOptions {
	options.Headers = param
	return options
}

// ListHostProfiles : Success. The request was successfully processed.
type ListHostProfiles struct {
	// The list of available host profiles.
	DirectorSiteHostProfiles []HostProfile `json:"director_site_host_profiles,omitempty"`
}

// UnmarshalListHostProfiles unmarshals an instance of ListHostProfiles from the specified map of raw messages.
func UnmarshalListHostProfiles(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListHostProfiles)
	err = core.UnmarshalModel(m, "director_site_host_profiles", &obj.DirectorSiteHostProfiles, UnmarshalHostProfile)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPVDCs : Return all provider virtual data center instances.
type ListPVDCs struct {
	// List of provider virtual data center instances.
	Pvdcs []PVDCSummary `json:"pvdcs,omitempty"`
}

// UnmarshalListPVDCs unmarshals an instance of ListPVDCs from the specified map of raw messages.
func UnmarshalListPVDCs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListPVDCs)
	err = core.UnmarshalModel(m, "pvdcs", &obj.Pvdcs, UnmarshalPVDCSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListPricesOptions : The ListPrices options.
type ListPricesOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListPricesOptions : Instantiate ListPricesOptions
func (*VMwareAsAServiceApiV1) NewListPricesOptions() *ListPricesOptions {
	return &ListPricesOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListPricesOptions) SetAcceptLanguage(acceptLanguage string) *ListPricesOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *ListPricesOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *ListPricesOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListPricesOptions) SetHeaders(param map[string]string) *ListPricesOptions {
	options.Headers = param
	return options
}

// ListVDCs : A list of Virtual Data Centers.
type ListVDCs struct {
	// A List of Virtual Data Centers.
	Vdcs []VDC `json:"vdcs" validate:"required"`
}

// UnmarshalListVDCs unmarshals an instance of ListVDCs from the specified map of raw messages.
func UnmarshalListVDCs(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(ListVDCs)
	err = core.UnmarshalModel(m, "vdcs", &obj.Vdcs, UnmarshalVDC)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ListVdcsOptions : The ListVdcs options.
type ListVdcsOptions struct {
	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVdcsOptions : Instantiate ListVdcsOptions
func (*VMwareAsAServiceApiV1) NewListVdcsOptions() *ListVdcsOptions {
	return &ListVdcsOptions{}
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *ListVdcsOptions) SetAcceptLanguage(acceptLanguage string) *ListVdcsOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListVdcsOptions) SetHeaders(param map[string]string) *ListVdcsOptions {
	options.Headers = param
	return options
}

// NewPassword : The new admin password used to log in to the VMware Cloud Director tenant portal. VMware Cloud Director has its own
// internal authentication and authorization model. The previous Director admin password is reset to a newly generated
// random password.
type NewPassword struct {
	// The password used to log in to the VMware Cloud Director tenant portal.
	Password *string `json:"password" validate:"required"`
}

// UnmarshalNewPassword unmarshals an instance of NewPassword from the specified map of raw messages.
func UnmarshalNewPassword(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewPassword)
	err = core.UnmarshalPrimitive(m, "password", &obj.Password)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCDirectorSite : The director site in which to deploy the Virtual Data Center.
type NewVDCDirectorSite struct {
	// A unique identifier for the director site.
	ID *string `json:"id" validate:"required"`

	// The cluster within the director site in which to deploy the Virtual Data Center.
	Cluster *VDCDirectorSiteCluster `json:"cluster" validate:"required"`
}

// NewNewVDCDirectorSite : Instantiate NewVDCDirectorSite (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewNewVDCDirectorSite(id string, cluster *VDCDirectorSiteCluster) (_model *NewVDCDirectorSite, err error) {
	_model = &NewVDCDirectorSite{
		ID: core.StringPtr(id),
		Cluster: cluster,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCDirectorSite unmarshals an instance of NewVDCDirectorSite from the specified map of raw messages.
func UnmarshalNewVDCDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCDirectorSite)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cluster", &obj.Cluster, UnmarshalVDCDirectorSiteCluster)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCEdge : The networking Edge to be deployed on the Virtual Data Center.
type NewVDCEdge struct {
	// The size of the Edge. Only used for Edges of type dedicated.
	Size *string `json:"size,omitempty"`

	// The type of Edge to be deployed on the Virtual Data Center.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the NewVDCEdge.Size property.
// The size of the Edge. Only used for Edges of type dedicated.
const (
	NewVDCEdge_Size_ExtraLarge = "extra_large"
	NewVDCEdge_Size_Large = "large"
	NewVDCEdge_Size_Medium = "medium"
)

// Constants associated with the NewVDCEdge.Type property.
// The type of Edge to be deployed on the Virtual Data Center.
const (
	NewVDCEdge_Type_Dedicated = "dedicated"
	NewVDCEdge_Type_Shared = "shared"
)

// NewNewVDCEdge : Instantiate NewVDCEdge (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewNewVDCEdge(typeVar string) (_model *NewVDCEdge, err error) {
	_model = &NewVDCEdge{
		Type: core.StringPtr(typeVar),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCEdge unmarshals an instance of NewVDCEdge from the specified map of raw messages.
func UnmarshalNewVDCEdge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCEdge)
	err = core.UnmarshalPrimitive(m, "size", &obj.Size)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// NewVDCResourceGroup : The resource group to associate with the Virtual Data Center. If not specified, the default resource group in the
// account is used.
type NewVDCResourceGroup struct {
	// A unique identifier for the resource group.
	ID *string `json:"id" validate:"required"`
}

// NewNewVDCResourceGroup : Instantiate NewVDCResourceGroup (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewNewVDCResourceGroup(id string) (_model *NewVDCResourceGroup, err error) {
	_model = &NewVDCResourceGroup{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalNewVDCResourceGroup unmarshals an instance of NewVDCResourceGroup from the specified map of raw messages.
func UnmarshalNewVDCResourceGroup(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(NewVDCResourceGroup)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PVDCOrderInfo : VMware provider virtual data center order information.
type PVDCOrderInfo struct {
	// Name of the provider virtual data center. Provider virtual data center names must be unique per director site
	// instance. Provider virtual data center names cannot be changed after creation.
	Name *string `json:"name" validate:"required"`

	// Data center location to deploy the cluster. See `GET /director_site_regions` for supported data center locations.
	DataCenter *string `json:"data_center" validate:"required"`

	// List of VMware clusters to deploy on the instance. Clusters form VMware workload availibility boundaries.
	Clusters []ClusterOrderInfo `json:"clusters" validate:"required"`
}

// NewPVDCOrderInfo : Instantiate PVDCOrderInfo (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewPVDCOrderInfo(name string, dataCenter string, clusters []ClusterOrderInfo) (_model *PVDCOrderInfo, err error) {
	_model = &PVDCOrderInfo{
		Name: core.StringPtr(name),
		DataCenter: core.StringPtr(dataCenter),
		Clusters: clusters,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPVDCOrderInfo unmarshals an instance of PVDCOrderInfo from the specified map of raw messages.
func UnmarshalPVDCOrderInfo(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PVDCOrderInfo)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalClusterOrderInfo)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PVDCResponse : VMware provider virtual data center create information.
type PVDCResponse struct {
	// Name of the provider virtual data center. Provider virtual data center names must be unique per director site
	// instance. Provider virtual data center names cannot be changed after creation.
	Name *string `json:"name" validate:"required"`

	// Data center location to deploy the cluster. See `GET /director_site_regions` for supported data center locations.
	DataCenter *string `json:"data_center" validate:"required"`

	// The provider virtual data center ID.
	ID *string `json:"id" validate:"required"`

	// The hyperlink of the provider virtual data center resource.
	Href *string `json:"href" validate:"required"`

	// List of VMware clusters to deploy on the instance. Clusters form VMware workload availibility boundaries.
	Clusters []ClusterSummary `json:"clusters" validate:"required"`

	// The provider virtual data center ordering status.
	Status *string `json:"status" validate:"required"`
}

// Constants associated with the PVDCResponse.Status property.
// The provider virtual data center ordering status.
const (
	PVDCResponse_Status_Creating = "Creating"
	PVDCResponse_Status_Deleted = "Deleted"
	PVDCResponse_Status_Deleting = "Deleting"
	PVDCResponse_Status_Failed = "Failed"
	PVDCResponse_Status_Modifying = "Modifying"
	PVDCResponse_Status_Readytouse = "ReadyToUse"
)

// UnmarshalPVDCResponse unmarshals an instance of PVDCResponse from the specified map of raw messages.
func UnmarshalPVDCResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PVDCResponse)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalClusterSummary)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PVDCSummary : VMware provider virtual data center information.
type PVDCSummary struct {
	// Name of the provider virtual data center. Provider virtual data center names must be unique per director site
	// instance. Provider virtual data center names cannot be changed after creation.
	Name *string `json:"name" validate:"required"`

	// Data center location to deploy the cluster. See `GET /director_site_regions` for supported data center locations.
	DataCenter *string `json:"data_center" validate:"required"`

	// The provider virtual data center ID.
	ID *string `json:"id" validate:"required"`

	// The hyperlink of the provider virtual data center resource.
	Href *string `json:"href" validate:"required"`

	// List of VMware clusters to deploy on the instance. Clusters form VMware workload availibility boundaries.
	Clusters []ClusterSummary `json:"clusters" validate:"required"`
}

// UnmarshalPVDCSummary unmarshals an instance of PVDCSummary from the specified map of raw messages.
func UnmarshalPVDCSummary(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PVDCSummary)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "data_center", &obj.DataCenter)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "href", &obj.Href)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "clusters", &obj.Clusters, UnmarshalClusterSummary)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoBaseCharge : Details of the instance base charge.
type PriceInfoBaseCharge struct {
	// The name of the metric that is being charged.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this pric.
	Currency *string `json:"currency,omitempty"`

	// The price for this metric.
	Price *float64 `json:"price,omitempty"`
}

// UnmarshalPriceInfoBaseCharge unmarshals an instance of PriceInfoBaseCharge from the specified map of raw messages.
func UnmarshalPriceInfoBaseCharge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoBaseCharge)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterCharge : A cluster for the instance and its price information.
type PriceInfoClusterCharge struct {
	// The cluster name.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for this cluster.
	Price *float64 `json:"price,omitempty"`

	// A list of items that make up the cluster and their price information.
	Items []PriceInfoClusterItem `json:"items,omitempty"`
}

// UnmarshalPriceInfoClusterCharge unmarshals an instance of PriceInfoClusterCharge from the specified map of raw messages.
func UnmarshalPriceInfoClusterCharge(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterCharge)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalPriceInfoClusterItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterItem : items for VCDD instance cluster price information.
type PriceInfoClusterItem struct {
	// The item name.
	Name *string `json:"name,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The total price for this item.
	Price *float64 `json:"price,omitempty"`

	// A list of subitems and their price information.
	Items []PriceInfoClusterSubItem `json:"items,omitempty"`
}

// UnmarshalPriceInfoClusterItem unmarshals an instance of PriceInfoClusterItem from the specified map of raw messages.
func UnmarshalPriceInfoClusterItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "items", &obj.Items, UnmarshalPriceInfoClusterSubItem)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PriceInfoClusterSubItem : sub items for VCDD instance cluster price information.
type PriceInfoClusterSubItem struct {
	// The metric that is being charged.
	Name *string `json:"name,omitempty"`

	// The number of items that this metric will be charged.
	Count *int64 `json:"count,omitempty"`

	// The unit of currency for this price.
	Currency *string `json:"currency,omitempty"`

	// The price for a single charge of this metric.
	Price *float64 `json:"price,omitempty"`
}

// UnmarshalPriceInfoClusterSubItem unmarshals an instance of PriceInfoClusterSubItem from the specified map of raw messages.
func UnmarshalPriceInfoClusterSubItem(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PriceInfoClusterSubItem)
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "count", &obj.Count)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "currency", &obj.Currency)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "price", &obj.Price)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// RegionDetail : The region details.
type RegionDetail struct {
	// The data center details.
	DataCenters []DataCenterInfo `json:"data_centers,omitempty"`

	// Accessable endpoint of the region.
	Endpoint *string `json:"endpoint,omitempty"`
}

// UnmarshalRegionDetail unmarshals an instance of RegionDetail from the specified map of raw messages.
func UnmarshalRegionDetail(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(RegionDetail)
	err = core.UnmarshalModel(m, "data_centers", &obj.DataCenters, UnmarshalDataCenterInfo)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "endpoint", &obj.Endpoint)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// ReplaceOrgAdminPasswordOptions : The ReplaceOrgAdminPassword options.
type ReplaceOrgAdminPasswordOptions struct {
	// A unique identifier for the director site.
	SiteID *string `json:"site_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewReplaceOrgAdminPasswordOptions : Instantiate ReplaceOrgAdminPasswordOptions
func (*VMwareAsAServiceApiV1) NewReplaceOrgAdminPasswordOptions(siteID string) *ReplaceOrgAdminPasswordOptions {
	return &ReplaceOrgAdminPasswordOptions{
		SiteID: core.StringPtr(siteID),
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *ReplaceOrgAdminPasswordOptions) SetSiteID(siteID string) *ReplaceOrgAdminPasswordOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ReplaceOrgAdminPasswordOptions) SetHeaders(param map[string]string) *ReplaceOrgAdminPasswordOptions {
	options.Headers = param
	return options
}

// UpdateClusterResponse : Response of cluster update.
type UpdateClusterResponse struct {
	// Information of request accepted.
	Message *string `json:"message,omitempty"`
}

// UnmarshalUpdateClusterResponse unmarshals an instance of UpdateClusterResponse from the specified map of raw messages.
func UnmarshalUpdateClusterResponse(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(UpdateClusterResponse)
	err = core.UnmarshalPrimitive(m, "message", &obj.Message)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

func (*VMwareAsAServiceApiV1) NewUpdateClusterResponsePatch(updateClusterResponse *UpdateClusterResponse) (_patch []JSONPatchOperation) {
	if (updateClusterResponse.Message != nil) {
		_patch = append(_patch, JSONPatchOperation{
			Op: core.StringPtr(JSONPatchOperation_Op_Add),
			Path: core.StringPtr("/message"),
			Value: updateClusterResponse.Message,
		})
	}
	return
}

// UpdateDirectorSitesPvdcsClusterOptions : The UpdateDirectorSitesPvdcsCluster options.
type UpdateDirectorSitesPvdcsClusterOptions struct {
	// A unique identifier for the director site in which the virtual data center was created.
	SiteID *string `json:"site_id" validate:"required,ne="`

	// The cluster to query.
	ClusterID *string `json:"cluster_id" validate:"required,ne="`

	// A unique identifier for the provider virtual data center in a director site.
	PvdcID *string `json:"pvdc_id" validate:"required,ne="`

	// Array of patch operations as defined in RFC 6902.
	Body []JSONPatchOperation `json:"body" validate:"required"`

	// Language.
	AcceptLanguage *string `json:"Accept-Language,omitempty"`

	// Transaction id.
	XGlobalTransactionID *string `json:"X-Global-Transaction-ID,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateDirectorSitesPvdcsClusterOptions : Instantiate UpdateDirectorSitesPvdcsClusterOptions
func (*VMwareAsAServiceApiV1) NewUpdateDirectorSitesPvdcsClusterOptions(siteID string, clusterID string, pvdcID string, body []JSONPatchOperation) *UpdateDirectorSitesPvdcsClusterOptions {
	return &UpdateDirectorSitesPvdcsClusterOptions{
		SiteID: core.StringPtr(siteID),
		ClusterID: core.StringPtr(clusterID),
		PvdcID: core.StringPtr(pvdcID),
		Body: body,
	}
}

// SetSiteID : Allow user to set SiteID
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetSiteID(siteID string) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.SiteID = core.StringPtr(siteID)
	return _options
}

// SetClusterID : Allow user to set ClusterID
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetClusterID(clusterID string) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.ClusterID = core.StringPtr(clusterID)
	return _options
}

// SetPvdcID : Allow user to set PvdcID
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetPvdcID(pvdcID string) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.PvdcID = core.StringPtr(pvdcID)
	return _options
}

// SetBody : Allow user to set Body
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetBody(body []JSONPatchOperation) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.Body = body
	return _options
}

// SetAcceptLanguage : Allow user to set AcceptLanguage
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetAcceptLanguage(acceptLanguage string) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.AcceptLanguage = core.StringPtr(acceptLanguage)
	return _options
}

// SetXGlobalTransactionID : Allow user to set XGlobalTransactionID
func (_options *UpdateDirectorSitesPvdcsClusterOptions) SetXGlobalTransactionID(xGlobalTransactionID string) *UpdateDirectorSitesPvdcsClusterOptions {
	_options.XGlobalTransactionID = core.StringPtr(xGlobalTransactionID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateDirectorSitesPvdcsClusterOptions) SetHeaders(param map[string]string) *UpdateDirectorSitesPvdcsClusterOptions {
	options.Headers = param
	return options
}

// VDC : A VMware Virtual Data Center (VDC). VMware VDCs are used to deploy and run VMware virtualized networking and run
// VMware workloads. VMware VDCs form loose boundaries of networking and workload where networking and workload can be
// shared or optionally isolated between VDCs. You can deploy one or more VDCs in an instance except when using the
// minimal instance configuration consisting of 2 hosts (2-Socket 32 Cores, 192 GB RAM). With the minimal instance
// configuration you can start with just one VDC and a performance network edge of medium size until additional hosts
// are added to the cluster.
type VDC struct {
	// A unique identifier for the Virtual Data Center.
	ID *string `json:"id" validate:"required"`

	// Determines how resources are made available to the Virtual Data Center. VMware as a Services uses the VMware Cloud
	// Director Pay-As-You-Go (paygo) allocation model. With paygo, resources are committed as they are allocated by VMware
	// vApps and VMs. IaaS resources are not reserved until vApps and VMs are specifically defined to VMware Cloud
	// Director. The paygo model supports an optimal use of resources where resources are allocated on-demand as needed
	// rather than prereserved without use.
	AllocationModel *string `json:"allocation_model" validate:"required"`

	// The time after which the Virtual Data Center is considered usable.
	CreatedTime *strfmt.DateTime `json:"created_time" validate:"required"`

	// A unique identifier for the Virtual Data Center in IBM Cloud.
	Crn *string `json:"crn" validate:"required"`

	// The time after which the Virtual Data Center is no longer considered usable.
	DeletedTime *strfmt.DateTime `json:"deleted_time" validate:"required"`

	// The director site in which to deploy the Virtual Data Center.
	DirectorSite *VDCDirectorSite `json:"director_site" validate:"required"`

	// The VMware NSX-T networking Edges deployed on the Virtual Data Center. NSX-T edges are used for bridging virtualize
	// networking to the physical public-internet and IBM private networking.
	Edges []Edge `json:"edges" validate:"required"`

	// Information about why the request to create the Virtual Data Center cannot be completed.
	Errors []Error `json:"errors" validate:"required"`

	// A human readable identifier for the Virtual Data Center.
	Name *string `json:"name" validate:"required"`

	// The time at which the request to create the Virtual Data Center was made.
	OrderedTime *strfmt.DateTime `json:"ordered_time" validate:"required"`

	// The name of the VMware Cloud Director organization containing this Virtual Data Center. VMware Cloud Director
	// organizations are used to create strong boundaries between virtual data centers. There is a complete isolation of
	// user administration, networking, workloads and VMware Cloud Director catalogs between different Director
	// organizations.
	OrgName *string `json:"org_name" validate:"required"`

	// Determines the state the Virtual Data Center is currently in.
	Status *string `json:"status" validate:"required"`

	// Determines if this Virtual Data Center is in a single-tenant or multi-tenant director site.
	Type *string `json:"type" validate:"required"`
}

// Constants associated with the VDC.AllocationModel property.
// Determines how resources are made available to the Virtual Data Center. VMware as a Services uses the VMware Cloud
// Director Pay-As-You-Go (paygo) allocation model. With paygo, resources are committed as they are allocated by VMware
// vApps and VMs. IaaS resources are not reserved until vApps and VMs are specifically defined to VMware Cloud Director.
// The paygo model supports an optimal use of resources where resources are allocated on-demand as needed rather than
// prereserved without use.
const (
	VDC_AllocationModel_Paygo = "paygo"
)

// Constants associated with the VDC.Status property.
// Determines the state the Virtual Data Center is currently in.
const (
	VDC_Status_Creating = "Creating"
	VDC_Status_Deleted = "Deleted"
	VDC_Status_Deleting = "Deleting"
	VDC_Status_Failed = "Failed"
	VDC_Status_Modifying = "Modifying"
	VDC_Status_Readytouse = "ReadyToUse"
)

// Constants associated with the VDC.Type property.
// Determines if this Virtual Data Center is in a single-tenant or multi-tenant director site.
const (
	VDC_Type_Dedicated = "dedicated"
)

// UnmarshalVDC unmarshals an instance of VDC from the specified map of raw messages.
func UnmarshalVDC(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDC)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "allocation_model", &obj.AllocationModel)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created_time", &obj.CreatedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "crn", &obj.Crn)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "deleted_time", &obj.DeletedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "director_site", &obj.DirectorSite, UnmarshalVDCDirectorSite)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "edges", &obj.Edges, UnmarshalEdge)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "errors", &obj.Errors, UnmarshalError)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "ordered_time", &obj.OrderedTime)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "org_name", &obj.OrgName)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "type", &obj.Type)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VDCDirectorSite : The director site in which to deploy the Virtual Data Center.
type VDCDirectorSite struct {
	// A unique identifier for the director site.
	ID *string `json:"id" validate:"required"`

	// The cluster within the director site in which to deploy the Virtual Data Center.
	Cluster *VDCDirectorSiteCluster `json:"cluster" validate:"required"`

	// The URL of the VMware Cloud Director tenant portal where this Virtual Data Center can be managed.
	URL *string `json:"url" validate:"required"`
}

// UnmarshalVDCDirectorSite unmarshals an instance of VDCDirectorSite from the specified map of raw messages.
func UnmarshalVDCDirectorSite(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDCDirectorSite)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "cluster", &obj.Cluster, UnmarshalVDCDirectorSiteCluster)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// VDCDirectorSiteCluster : The cluster within the director site in which to deploy the Virtual Data Center.
type VDCDirectorSiteCluster struct {
	// A unique identifier for the cluster.
	ID *string `json:"id" validate:"required"`
}

// NewVDCDirectorSiteCluster : Instantiate VDCDirectorSiteCluster (Generic Model Constructor)
func (*VMwareAsAServiceApiV1) NewVDCDirectorSiteCluster(id string) (_model *VDCDirectorSiteCluster, err error) {
	_model = &VDCDirectorSiteCluster{
		ID: core.StringPtr(id),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalVDCDirectorSiteCluster unmarshals an instance of VDCDirectorSiteCluster from the specified map of raw messages.
func UnmarshalVDCDirectorSiteCluster(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(VDCDirectorSiteCluster)
	err = core.UnmarshalPrimitive(m, "id", &obj.ID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
