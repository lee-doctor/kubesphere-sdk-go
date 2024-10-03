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


// CustomMetricsAPIService CustomMetricsAPI service
type CustomMetricsAPIService service

type ApiHandleClusterAdhocQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
	expr *string
	start *string
	end *string
	step *string
	time *string
}

// The expression to be evaluated.
func (r ApiHandleClusterAdhocQueryRequest) Expr(expr string) ApiHandleClusterAdhocQueryRequest {
	r.expr = &expr
	return r
}

// Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleClusterAdhocQueryRequest) Start(start string) ApiHandleClusterAdhocQueryRequest {
	r.start = &start
	return r
}

// End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleClusterAdhocQueryRequest) End(end string) ApiHandleClusterAdhocQueryRequest {
	r.end = &end
	return r
}

// Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).
func (r ApiHandleClusterAdhocQueryRequest) Step(step string) ApiHandleClusterAdhocQueryRequest {
	r.step = &step
	return r
}

// A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.
func (r ApiHandleClusterAdhocQueryRequest) Time(time string) ApiHandleClusterAdhocQueryRequest {
	r.time = &time
	return r
}

func (r ApiHandleClusterAdhocQueryRequest) Execute() (*MonitoringMetric, *http.Response, error) {
	return r.ApiService.HandleClusterAdhocQueryExecute(r)
}

/*
HandleClusterAdhocQuery Make an ad-hoc query in the whole cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHandleClusterAdhocQueryRequest
*/
func (a *CustomMetricsAPIService) HandleClusterAdhocQuery(ctx context.Context) ApiHandleClusterAdhocQueryRequest {
	return ApiHandleClusterAdhocQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MonitoringMetric
func (a *CustomMetricsAPIService) HandleClusterAdhocQueryExecute(r ApiHandleClusterAdhocQueryRequest) (*MonitoringMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleClusterAdhocQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/targets/query"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expr != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expr", r.expr, "form", "")
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

type ApiHandleClusterMetadataQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
}

func (r ApiHandleClusterMetadataQueryRequest) Execute() (*MonitoringMetadata, *http.Response, error) {
	return r.ApiService.HandleClusterMetadataQueryExecute(r)
}

/*
HandleClusterMetadataQuery Get metadata of metrics in the whole cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHandleClusterMetadataQueryRequest
*/
func (a *CustomMetricsAPIService) HandleClusterMetadataQuery(ctx context.Context) ApiHandleClusterMetadataQueryRequest {
	return ApiHandleClusterMetadataQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MonitoringMetadata
func (a *CustomMetricsAPIService) HandleClusterMetadataQueryExecute(r ApiHandleClusterMetadataQueryRequest) (*MonitoringMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleClusterMetadataQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/targets/metadata"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiHandleClusterMetricLabelSetQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
	metric *string
	start *string
	end *string
}

// The name of the metric
func (r ApiHandleClusterMetricLabelSetQueryRequest) Metric(metric string) ApiHandleClusterMetricLabelSetQueryRequest {
	r.metric = &metric
	return r
}

// Start time of query. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleClusterMetricLabelSetQueryRequest) Start(start string) ApiHandleClusterMetricLabelSetQueryRequest {
	r.start = &start
	return r
}

// End time of query. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleClusterMetricLabelSetQueryRequest) End(end string) ApiHandleClusterMetricLabelSetQueryRequest {
	r.end = &end
	return r
}

func (r ApiHandleClusterMetricLabelSetQueryRequest) Execute() (*MonitoringMetricLabelSet, *http.Response, error) {
	return r.ApiService.HandleClusterMetricLabelSetQueryExecute(r)
}

/*
HandleClusterMetricLabelSetQuery List all available labels and values of a metric within a specific time span in the whole cluster.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiHandleClusterMetricLabelSetQueryRequest
*/
func (a *CustomMetricsAPIService) HandleClusterMetricLabelSetQuery(ctx context.Context) ApiHandleClusterMetricLabelSetQueryRequest {
	return ApiHandleClusterMetricLabelSetQueryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return MonitoringMetricLabelSet
func (a *CustomMetricsAPIService) HandleClusterMetricLabelSetQueryExecute(r ApiHandleClusterMetricLabelSetQueryRequest) (*MonitoringMetricLabelSet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetricLabelSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleClusterMetricLabelSetQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/targets/labelsets"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metric == nil {
		return localVarReturnValue, nil, reportError("metric is required and must be specified")
	}
	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}
	if r.end == nil {
		return localVarReturnValue, nil, reportError("end is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "metric", r.metric, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
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

type ApiHandleMetadataQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
	namespace string
}

func (r ApiHandleMetadataQueryRequest) Execute() (*MonitoringMetadata, *http.Response, error) {
	return r.ApiService.HandleMetadataQueryExecute(r)
}

/*
HandleMetadataQuery Get metadata of metrics for the specific namespace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param namespace The name of the namespace.
 @return ApiHandleMetadataQueryRequest
*/
func (a *CustomMetricsAPIService) HandleMetadataQuery(ctx context.Context, namespace string) ApiHandleMetadataQueryRequest {
	return ApiHandleMetadataQueryRequest{
		ApiService: a,
		ctx: ctx,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return MonitoringMetadata
func (a *CustomMetricsAPIService) HandleMetadataQueryExecute(r ApiHandleMetadataQueryRequest) (*MonitoringMetadata, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetadata
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleMetadataQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/metadata"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

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

type ApiHandleMetricLabelSetQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
	namespace string
	metric *string
	start *string
	end *string
}

// The name of the metric
func (r ApiHandleMetricLabelSetQueryRequest) Metric(metric string) ApiHandleMetricLabelSetQueryRequest {
	r.metric = &metric
	return r
}

// Start time of query. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleMetricLabelSetQueryRequest) Start(start string) ApiHandleMetricLabelSetQueryRequest {
	r.start = &start
	return r
}

// End time of query. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleMetricLabelSetQueryRequest) End(end string) ApiHandleMetricLabelSetQueryRequest {
	r.end = &end
	return r
}

func (r ApiHandleMetricLabelSetQueryRequest) Execute() (*MonitoringMetricLabelSet, *http.Response, error) {
	return r.ApiService.HandleMetricLabelSetQueryExecute(r)
}

/*
HandleMetricLabelSetQuery List all available labels and values of a metric within a specific time span.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param namespace The name of the namespace.
 @return ApiHandleMetricLabelSetQueryRequest
*/
func (a *CustomMetricsAPIService) HandleMetricLabelSetQuery(ctx context.Context, namespace string) ApiHandleMetricLabelSetQueryRequest {
	return ApiHandleMetricLabelSetQueryRequest{
		ApiService: a,
		ctx: ctx,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return MonitoringMetricLabelSet
func (a *CustomMetricsAPIService) HandleMetricLabelSetQueryExecute(r ApiHandleMetricLabelSetQueryRequest) (*MonitoringMetricLabelSet, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetricLabelSet
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleMetricLabelSetQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/labelsets"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.metric == nil {
		return localVarReturnValue, nil, reportError("metric is required and must be specified")
	}
	if r.start == nil {
		return localVarReturnValue, nil, reportError("start is required and must be specified")
	}
	if r.end == nil {
		return localVarReturnValue, nil, reportError("end is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "metric", r.metric, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "start", r.start, "form", "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "end", r.end, "form", "")
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

type ApiHandleNamespaceAdhocQueryRequest struct {
	ctx context.Context
	ApiService *CustomMetricsAPIService
	namespace string
	expr *string
	start *string
	end *string
	step *string
	time *string
}

// The expression to be evaluated.
func (r ApiHandleNamespaceAdhocQueryRequest) Expr(expr string) ApiHandleNamespaceAdhocQueryRequest {
	r.expr = &expr
	return r
}

// Start time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1559347200. 
func (r ApiHandleNamespaceAdhocQueryRequest) Start(start string) ApiHandleNamespaceAdhocQueryRequest {
	r.start = &start
	return r
}

// End time of query. Use **start** and **end** to retrieve metric data over a time span. It is a string with Unix time format, eg. 1561939200. 
func (r ApiHandleNamespaceAdhocQueryRequest) End(end string) ApiHandleNamespaceAdhocQueryRequest {
	r.end = &end
	return r
}

// Time interval. Retrieve metric data at a fixed interval within the time range of start and end. It requires both **start** and **end** are provided. The format is [0-9]+[smhdwy]. Defaults to 10m (i.e. 10 min).
func (r ApiHandleNamespaceAdhocQueryRequest) Step(step string) ApiHandleNamespaceAdhocQueryRequest {
	r.step = &step
	return r
}

// A timestamp in Unix time format. Retrieve metric data at a single point in time. Defaults to now. Time and the combination of start, end, step are mutually exclusive.
func (r ApiHandleNamespaceAdhocQueryRequest) Time(time string) ApiHandleNamespaceAdhocQueryRequest {
	r.time = &time
	return r
}

func (r ApiHandleNamespaceAdhocQueryRequest) Execute() (*MonitoringMetric, *http.Response, error) {
	return r.ApiService.HandleNamespaceAdhocQueryExecute(r)
}

/*
HandleNamespaceAdhocQuery Make an ad-hoc query in the specific namespace.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param namespace The name of the namespace.
 @return ApiHandleNamespaceAdhocQueryRequest
*/
func (a *CustomMetricsAPIService) HandleNamespaceAdhocQuery(ctx context.Context, namespace string) ApiHandleNamespaceAdhocQueryRequest {
	return ApiHandleNamespaceAdhocQueryRequest{
		ApiService: a,
		ctx: ctx,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return MonitoringMetric
func (a *CustomMetricsAPIService) HandleNamespaceAdhocQueryExecute(r ApiHandleNamespaceAdhocQueryRequest) (*MonitoringMetric, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *MonitoringMetric
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "CustomMetricsAPIService.HandleNamespaceAdhocQuery")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/monitoring.kubesphere.io/v1alpha3/namespaces/{namespace}/targets/query"
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.expr != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "expr", r.expr, "form", "")
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