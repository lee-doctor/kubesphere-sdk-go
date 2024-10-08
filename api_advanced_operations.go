/*
KS API

KubeSphere OpenAPI

API version: v4.1.1
Contact: support@kubesphere.cloud
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


// AdvancedOperationsAPIService AdvancedOperationsAPI service
type AdvancedOperationsAPIService service

type ApiGetRegistryEntryRequest struct {
	ctx context.Context
	ApiService *AdvancedOperationsAPIService
	image *string
	namespace *string
	secret *string
	insecure *string
}

// query image, condition for filtering.
func (r ApiGetRegistryEntryRequest) Image(image string) ApiGetRegistryEntryRequest {
	r.image = &image
	return r
}

// namespace which secret in.
func (r ApiGetRegistryEntryRequest) Namespace(namespace string) ApiGetRegistryEntryRequest {
	r.namespace = &namespace
	return r
}

// secret name
func (r ApiGetRegistryEntryRequest) Secret(secret string) ApiGetRegistryEntryRequest {
	r.secret = &secret
	return r
}

// whether verify cert if using https repo
func (r ApiGetRegistryEntryRequest) Insecure(insecure string) ApiGetRegistryEntryRequest {
	r.insecure = &insecure
	return r
}

func (r ApiGetRegistryEntryRequest) Execute() (*RegistriesImageDetails, *http.Response, error) {
	return r.ApiService.GetRegistryEntryExecute(r)
}

/*
GetRegistryEntry Retrieve the blob from the registry

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiGetRegistryEntryRequest

Deprecated
*/
func (a *AdvancedOperationsAPIService) GetRegistryEntry(ctx context.Context) ApiGetRegistryEntryRequest {
	return ApiGetRegistryEntryRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return RegistriesImageDetails
// Deprecated
func (a *AdvancedOperationsAPIService) GetRegistryEntryExecute(r ApiGetRegistryEntryRequest) (*RegistriesImageDetails, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *RegistriesImageDetails
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedOperationsAPIService.GetRegistryEntry")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/resources.kubesphere.io/v1alpha2/registry/blob"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.image == nil {
		return localVarReturnValue, nil, reportError("image is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "image", r.image, "form", "")
	if r.namespace != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "namespace", r.namespace, "form", "")
	}
	if r.secret != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "secret", r.secret, "form", "")
	}
	if r.insecure != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "insecure", r.insecure, "form", "")
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
			if apiKey, ok := auth["BearerToken"]; ok {
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

type ApiJobReRunRequest struct {
	ctx context.Context
	ApiService *AdvancedOperationsAPIService
	job string
	namespace string
	resourceVersion *string
	action *string
}

// version of job, rerun when the version matches
func (r ApiJobReRunRequest) ResourceVersion(resourceVersion string) ApiJobReRunRequest {
	r.resourceVersion = &resourceVersion
	return r
}

// action must be \&quot;rerun\&quot;
func (r ApiJobReRunRequest) Action(action string) ApiJobReRunRequest {
	r.action = &action
	return r
}

func (r ApiJobReRunRequest) Execute() (*ErrorsError, *http.Response, error) {
	return r.ApiService.JobReRunExecute(r)
}

/*
JobReRun Job rerun

Rerun job whether the job is complete or not.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param job job name
 @param namespace The specified namespace.
 @return ApiJobReRunRequest

Deprecated
*/
func (a *AdvancedOperationsAPIService) JobReRun(ctx context.Context, job string, namespace string) ApiJobReRunRequest {
	return ApiJobReRunRequest{
		ApiService: a,
		ctx: ctx,
		job: job,
		namespace: namespace,
	}
}

// Execute executes the request
//  @return ErrorsError
// Deprecated
func (a *AdvancedOperationsAPIService) JobReRunExecute(r ApiJobReRunRequest) (*ErrorsError, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ErrorsError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedOperationsAPIService.JobReRun")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/operations.kubesphere.io/v1alpha2/namespaces/{namespace}/jobs/{job}"
	localVarPath = strings.Replace(localVarPath, "{"+"job"+"}", url.PathEscape(parameterValueToString(r.job, "job")), -1)
	localVarPath = strings.Replace(localVarPath, "{"+"namespace"+"}", url.PathEscape(parameterValueToString(r.namespace, "namespace")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.resourceVersion == nil {
		return localVarReturnValue, nil, reportError("resourceVersion is required and must be specified")
	}

	if r.action != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "action", r.action, "form", "")
	}
	parameterAddToHeaderOrQuery(localVarQueryParams, "resourceVersion", r.resourceVersion, "form", "")
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
			if apiKey, ok := auth["BearerToken"]; ok {
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

type ApiVerifyGitCredentialRequest struct {
	ctx context.Context
	ApiService *AdvancedOperationsAPIService
	body *GitAuthInfo
}

func (r ApiVerifyGitCredentialRequest) Body(body GitAuthInfo) ApiVerifyGitCredentialRequest {
	r.body = &body
	return r
}

func (r ApiVerifyGitCredentialRequest) Execute() (*ErrorsError, *http.Response, error) {
	return r.ApiService.VerifyGitCredentialExecute(r)
}

/*
VerifyGitCredential Verify the git credential

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyGitCredentialRequest

Deprecated
*/
func (a *AdvancedOperationsAPIService) VerifyGitCredential(ctx context.Context) ApiVerifyGitCredentialRequest {
	return ApiVerifyGitCredentialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ErrorsError
// Deprecated
func (a *AdvancedOperationsAPIService) VerifyGitCredentialExecute(r ApiVerifyGitCredentialRequest) (*ErrorsError, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ErrorsError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedOperationsAPIService.VerifyGitCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/resources.kubesphere.io/v1alpha2/git/verify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerToken"]; ok {
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

type ApiVerifyRegistryCredentialRequest struct {
	ctx context.Context
	ApiService *AdvancedOperationsAPIService
	body *ApiRegistryCredential
}

func (r ApiVerifyRegistryCredentialRequest) Body(body ApiRegistryCredential) ApiVerifyRegistryCredentialRequest {
	r.body = &body
	return r
}

func (r ApiVerifyRegistryCredentialRequest) Execute() (*ErrorsError, *http.Response, error) {
	return r.ApiService.VerifyRegistryCredentialExecute(r)
}

/*
VerifyRegistryCredential Verify registry credential

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiVerifyRegistryCredentialRequest

Deprecated
*/
func (a *AdvancedOperationsAPIService) VerifyRegistryCredential(ctx context.Context) ApiVerifyRegistryCredentialRequest {
	return ApiVerifyRegistryCredentialRequest{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ErrorsError
// Deprecated
func (a *AdvancedOperationsAPIService) VerifyRegistryCredentialExecute(r ApiVerifyRegistryCredentialRequest) (*ErrorsError, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ErrorsError
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdvancedOperationsAPIService.VerifyRegistryCredential")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/kapis/resources.kubesphere.io/v1alpha2/registry/verify"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.body == nil {
		return localVarReturnValue, nil, reportError("body is required and must be specified")
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
	// body params
	localVarPostBody = r.body
	if r.ctx != nil {
		// API Key Authentication
		if auth, ok := r.ctx.Value(ContextAPIKeys).(map[string]APIKey); ok {
			if apiKey, ok := auth["BearerToken"]; ok {
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
