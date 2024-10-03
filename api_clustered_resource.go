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


// ClusteredResourceAPIService ClusteredResourceAPI service
type ClusteredResourceAPIService service

type ApiHandleGetResourcesRequest struct {
	ctx context.Context
	ApiService *ClusteredResourceAPIService
	resources string
	name string
}

func (r ApiHandleGetResourcesRequest) Execute() (*http.Response, error) {
	return r.ApiService.HandleGetResourcesExecute(r)
}

/*
HandleGetResources Cluster level resource

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resources cluster level resource type, e.g. pods,jobs,configmaps,services.
 @param name the name of the clustered resources
 @return ApiHandleGetResourcesRequest
*/
func (a *ClusteredResourceAPIService) HandleGetResources(ctx context.Context, resources string, name string) ApiHandleGetResourcesRequest {
	return ApiHandleGetResourcesRequest{
		ApiService: a,
		ctx: ctx,
		resources: resources,
		name: name,
	}
}

// Execute executes the request
func (a *ClusteredResourceAPIService) HandleGetResourcesExecute(r ApiHandleGetResourcesRequest) (*http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusteredResourceAPIService.HandleGetResources")
	if err != nil {
		return nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/resources.kubesphere.io/v1alpha3/{resources}/{name}"
	localVarPath = strings.Replace(localVarPath, "{"+"resources"+"}", url.PathEscape(parameterValueToString(r.resources, "resources")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"name"+"}", url.PathEscape(parameterValueToString(r.name, "name")), -1)

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
	localVarHTTPHeaderAccepts := []string{}

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
		return nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarHTTPResponse, newErr
	}

	return localVarHTTPResponse, nil
}

type ApiHandleListResourcesV3Request struct {
	ctx context.Context
	ApiService *ClusteredResourceAPIService
	resources string
	name *string
	page *string
	limit *string
	ascending *string
	sortBy *string
}

// name used to do filtering
func (r ApiHandleListResourcesV3Request) Name(name string) ApiHandleListResourcesV3Request {
	r.name = &name
	return r
}

// page
func (r ApiHandleListResourcesV3Request) Page(page string) ApiHandleListResourcesV3Request {
	r.page = &page
	return r
}

// limit
func (r ApiHandleListResourcesV3Request) Limit(limit string) ApiHandleListResourcesV3Request {
	r.limit = &limit
	return r
}

// sort parameters, e.g. reverse&#x3D;true
func (r ApiHandleListResourcesV3Request) Ascending(ascending string) ApiHandleListResourcesV3Request {
	r.ascending = &ascending
	return r
}

// sort parameters, e.g. orderBy&#x3D;createTime
func (r ApiHandleListResourcesV3Request) SortBy(sortBy string) ApiHandleListResourcesV3Request {
	r.sortBy = &sortBy
	return r
}

func (r ApiHandleListResourcesV3Request) Execute() (*ApiListResult, *http.Response, error) {
	return r.ApiService.HandleListResourcesV3Execute(r)
}

/*
HandleListResourcesV3 Cluster level resources

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param resources cluster level resource type, e.g. pods,jobs,configmaps,services.
 @return ApiHandleListResourcesV3Request
*/
func (a *ClusteredResourceAPIService) HandleListResourcesV3(ctx context.Context, resources string) ApiHandleListResourcesV3Request {
	return ApiHandleListResourcesV3Request{
		ApiService: a,
		ctx: ctx,
		resources: resources,
	}
}

// Execute executes the request
//  @return ApiListResult
func (a *ClusteredResourceAPIService) HandleListResourcesV3Execute(r ApiHandleListResourcesV3Request) (*ApiListResult, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ApiListResult
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ClusteredResourceAPIService.HandleListResourcesV3")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/resources.kubesphere.io/v1alpha3/{resources}"
	localVarPath = strings.Replace(localVarPath, "{"+"resources"+"}", url.PathEscape(parameterValueToString(r.resources, "resources")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	if r.name != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "name", r.name, "form", "")
	}
	if r.page != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "page", r.page, "form", "")
	} else {
		var defaultValue string = "page=1"
		r.page = &defaultValue
	}
	if r.limit != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "limit", r.limit, "form", "")
	}
	if r.ascending != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "ascending", r.ascending, "form", "")
	} else {
		var defaultValue string = "ascending=false"
		r.ascending = &defaultValue
	}
	if r.sortBy != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "sortBy", r.sortBy, "form", "")
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
