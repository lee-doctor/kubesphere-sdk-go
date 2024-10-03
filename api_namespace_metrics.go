/*
KubeSphere

KubeSphere OpenAPI

API version: v3.1.0
Contact: info@kubesphere.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
	"strings"
)


// NamespaceMetricsAPIService NamespaceMetricsAPI service
type NamespaceMetricsAPIService service

type ApiHandleAllNamespaceMetricsQueryRequest struct {
	ctx context.Context
	ApiService *NamespaceMetricsAPIService
	metricsFilter *string
	resourcesFilter *string
	start *string
	end *string
	step *string
	time *string
	sortMetric *string
	sortType *string
	page *int32
	limit *int32
}

// The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/).
func (r ApiHandleAllNamespaceMetricsQueryRequest) MetricsFilter(metricsFilter string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.metricsFilter = &metricsFilter
	return r
}

// The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: &#x60;test|kube-system&#x60;.
func (r ApiHandleAllNamespaceMetricsQueryRequest) ResourcesFilter(resourcesFilter string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.resourcesFilter = &resourcesFilter
	return r
}

// Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleAllNamespaceMetricsQueryRequest) Start(start string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.start = &start
	return r
}

// End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleAllNamespaceMetricsQueryRequest) End(end string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.end = &end
	return r
}

// Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).
func (r ApiHandleAllNamespaceMetricsQueryRequest) Step(step string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.step = &step
	return r
}

// A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.
func (r ApiHandleAllNamespaceMetricsQueryRequest) Time(time string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.time = &time
	return r
}

// Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided.
func (r ApiHandleAllNamespaceMetricsQueryRequest) SortMetric(sortMetric string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.sortMetric = &sortMetric
	return r
}

// Sort order. One of asc, desc.
func (r ApiHandleAllNamespaceMetricsQueryRequest) SortType(sortType string) ApiHandleAllNamespaceMetricsQueryRequest {
	r.sortType = &sortType
	return r
}

// The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data.
func (r ApiHandleAllNamespaceMetricsQueryRequest) Page(page int32) ApiHandleAllNamespaceMetricsQueryRequest {
	r.page = &page
	return r
}

// Page size, the maximum number of results in a single page. Defaults to 5.
func (r ApiHandleAllNamespaceMetricsQueryRequest) Limit(limit int32) ApiHandleAllNamespaceMetricsQueryRequest {
	r.limit = &limit
	return r
}

func (r ApiHandleAllNamespaceMetricsQueryRequest) Execute() (*MonitoringMetrics, *http.Response, error) {
	return r.ApiService.HandleAllNamespaceMetricsQueryExecute(r)
}

/*
HandleAllNamespaceMetricsQuery Get namespace-level metric data of all namespaces.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHandleAllNamespaceMetricsQueryRequest
*/
func (a *NamespaceMetricsAPIService) HandleAllNamespaceMetricsQuery(ctx context.Context) ApiHandleAllNamespaceMetricsQueryRequest {
	return ApiHandleAllNamespaceMetricsQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MonitoringMetrics
func (a *NamespaceMetricsAPIService) HandleAllNamespaceMetricsQueryExecute(r ApiHandleAllNamespaceMetricsQueryRequest) (*MonitoringMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceMetricsAPIService.HandleAllNamespaceMetricsQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metricsFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics_filter", r.metricsFilter, "form", "")
	}
	if r.resourcesFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resources_filter", r.resourcesFilter, "form", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
	}
	if r.step != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "step", r.step, "form", "")
	} else {
		var defaultValue string = "10m"
		r.step = &defaultValue
	}
	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	if r.sortMetric != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_metric", r.sortMetric, "form", "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_type", r.sortType, "form", "")
	} else {
		var defaultValue string = "desc."
		r.sortType = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 5
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["jwt"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHandleNamespaceMetricsQueryRequest struct {
	ctx context.Context
	ApiService *NamespaceMetricsAPIService
	namespace string
	metricsFilter *string
	start *string
	end *string
	step *string
	time *string
}

// The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/).
func (r ApiHandleNamespaceMetricsQueryRequest) MetricsFilter(metricsFilter string) ApiHandleNamespaceMetricsQueryRequest {
	r.metricsFilter = &metricsFilter
	return r
}

// Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleNamespaceMetricsQueryRequest) Start(start string) ApiHandleNamespaceMetricsQueryRequest {
	r.start = &start
	return r
}

// End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleNamespaceMetricsQueryRequest) End(end string) ApiHandleNamespaceMetricsQueryRequest {
	r.end = &end
	return r
}

// Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).
func (r ApiHandleNamespaceMetricsQueryRequest) Step(step string) ApiHandleNamespaceMetricsQueryRequest {
	r.step = &step
	return r
}

// A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.
func (r ApiHandleNamespaceMetricsQueryRequest) Time(time string) ApiHandleNamespaceMetricsQueryRequest {
	r.time = &time
	return r
}

func (r ApiHandleNamespaceMetricsQueryRequest) Execute() (*MonitoringMetrics, *http.Response, error) {
	return r.ApiService.HandleNamespaceMetricsQueryExecute(r)
}

/*
HandleNamespaceMetricsQuery Get namespace-level metric data of the specific namespace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param namespace The name of the namespace.
 @return ApiHandleNamespaceMetricsQueryRequest
*/
func (a *NamespaceMetricsAPIService) HandleNamespaceMetricsQuery(ctx context.Context, namespace string) ApiHandleNamespaceMetricsQueryRequest {
	return ApiHandleNamespaceMetricsQueryRequest{
		ApiService: a,
		ctx: ctx,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return MonitoringMetrics
func (a *NamespaceMetricsAPIService) HandleNamespaceMetricsQueryExecute(r ApiHandleNamespaceMetricsQueryRequest) (*MonitoringMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceMetricsAPIService.HandleNamespaceMetricsQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metricsFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics_filter", r.metricsFilter, "form", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
	}
	if r.step != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "step", r.step, "form", "")
	} else {
		var defaultValue string = "10m"
		r.step = &defaultValue
	}
	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["jwt"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiHandleWorkspaceNamespaceMetricsQueryRequest struct {
	ctx context.Context
	ApiService *NamespaceMetricsAPIService
	workspace string
	metricsFilter *string
	resourcesFilter *string
	start *string
	end *string
	step *string
	time *string
	sortMetric *string
	sortType *string
	page *int32
	limit *int32
}

// The metric name filter consists of a regexp pattern. It specifies which metric data to return. For example, the following filter matches both namespace CPU usage and memory usage: &#x60;namespace_cpu_usage|namespace_memory_usage&#x60;. View available metrics at [kubesphere.io](https://v2-0.docs.kubesphere.io/docs/api-reference/monitoring-metrics/).
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) MetricsFilter(metricsFilter string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.metricsFilter = &metricsFilter
	return r
}

// The namespace filter consists of a regexp pattern. It specifies which namespace data to return. For example, the following filter matches both namespace test and kube-system: &#x60;test|kube-system&#x60;.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) ResourcesFilter(resourcesFilter string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.resourcesFilter = &resourcesFilter
	return r
}

// Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Start(start string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.start = &start
	return r
}

// End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) End(end string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.end = &end
	return r
}

// Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Step(step string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.step = &step
	return r
}

// A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Time(time string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.time = &time
	return r
}

// Sort namespaces by the specified metric. Not applicable if **start** and **end** are provided.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) SortMetric(sortMetric string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.sortMetric = &sortMetric
	return r
}

// Sort order. One of asc, desc.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) SortType(sortType string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.sortType = &sortType
	return r
}

// The page number. This field paginates result data of each metric, then returns a specific page. For example, setting **page** to 2 returns the second page. It only applies to sorted metric data.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Page(page int32) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.page = &page
	return r
}

// Page size, the maximum number of results in a single page. Defaults to 5.
func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Limit(limit int32) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	r.limit = &limit
	return r
}

func (r ApiHandleWorkspaceNamespaceMetricsQueryRequest) Execute() (*MonitoringMetrics, *http.Response, error) {
	return r.ApiService.HandleWorkspaceNamespaceMetricsQueryExecute(r)
}

/*
HandleWorkspaceNamespaceMetricsQuery Get namespace-level metric data of a specific workspace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param workspace Workspace name.
 @return ApiHandleWorkspaceNamespaceMetricsQueryRequest
*/
func (a *NamespaceMetricsAPIService) HandleWorkspaceNamespaceMetricsQuery(ctx context.Context, workspace string) ApiHandleWorkspaceNamespaceMetricsQueryRequest {
	return ApiHandleWorkspaceNamespaceMetricsQueryRequest{
		ApiService: a,
		ctx: ctx,
		workspace: workspace,
	}
}

// Execute executes the request
//  @return MonitoringMetrics
func (a *NamespaceMetricsAPIService) HandleWorkspaceNamespaceMetricsQueryExecute(r ApiHandleWorkspaceNamespaceMetricsQueryRequest) (*MonitoringMetrics, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetrics
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "NamespaceMetricsAPIService.HandleWorkspaceNamespaceMetricsQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/workspaces/{workspace}/namespaces"
	localVarPath = strings.Replace(localVarPath, "{"+"workspace"+"}", url.PathEscape(parameterValueToString(r.workspace, "workspace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.metricsFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "metrics_filter", r.metricsFilter, "form", "")
	}
	if r.resourcesFilter != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "resources_filter", r.resourcesFilter, "form", "")
	}
	if r.start != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	}
	if r.end != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
	}
	if r.step != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "step", r.step, "form", "")
	} else {
		var defaultValue string = "10m"
		r.step = &defaultValue
	}
	if r.time != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "time", r.time, "form", "")
	}
	if r.sortMetric != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_metric", r.sortMetric, "form", "")
	}
	if r.sortType != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sort_type", r.sortType, "form", "")
	} else {
		var defaultValue string = "desc."
		r.sortType = &defaultValue
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	} else {
		var defaultValue int32 = 5
		r.limit = &defaultValue
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["jwt"]; ok {
				var key string
				if apiKey.Prefix != "" {
					key = apiKey.Prefix + " " + apiKey.Key
				} else {
					key = apiKey.Key
				}
				localVarHeaderParams["Authorization"] = key
			}
		}
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
