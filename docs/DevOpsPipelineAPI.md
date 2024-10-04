# \DevOpsPipelineAPI

All URIs are relative to *http://localhost*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckCron**](DevOpsPipelineAPI.md#CheckCron) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/checkCron | Check cron script compile.
[**CheckScriptCompile**](DevOpsPipelineAPI.md#CheckScriptCompile) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/checkScriptCompile | Check pipeline script compile.
[**DeleteCredential**](DevOpsPipelineAPI.md#DeleteCredential) | **Delete** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/credentials/{credential} | delete the credential of the specified devops for the current user
[**DeletePipeline**](DevOpsPipelineAPI.md#DeletePipeline) | **Delete** /kapis/devops.kubesphere.io/v1alpha3/devops/{devops}/pipelines/{pipeline} | delete the pipeline of the specified devops for the current user
[**GetArtifacts**](DevOpsPipelineAPI.md#GetArtifacts) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/artifacts | Get all artifacts in the specified pipeline.
[**GetBranchArtifacts**](DevOpsPipelineAPI.md#GetBranchArtifacts) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/artifacts | (MultiBranchesPipeline) Get all artifacts generated from the specified run of the pipeline branch.
[**GetBranchNodeSteps**](DevOpsPipelineAPI.md#GetBranchNodeSteps) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/nodes/{node}/steps | (MultiBranchesPipeline) Get all steps in the specified node.
[**GetBranchNodesDetail**](DevOpsPipelineAPI.md#GetBranchNodesDetail) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/nodesdetail | (MultiBranchesPipeline) Get steps details in an activity node. For a node, the steps which is defined inside the node.
[**GetBranchPipeline**](DevOpsPipelineAPI.md#GetBranchPipeline) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch} | (MultiBranchesPipeline) Get the specified branch pipeline of the DevOps project
[**GetBranchPipelineRun**](DevOpsPipelineAPI.md#GetBranchPipelineRun) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run} | (MultiBranchesPipeline) Get details in the specified pipeline activity.
[**GetBranchPipelineRunNodes**](DevOpsPipelineAPI.md#GetBranchPipelineRunNodes) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/nodes | (MultiBranchesPipeline) Get run nodes.
[**GetBranchRunLog**](DevOpsPipelineAPI.md#GetBranchRunLog) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/log | (MultiBranchesPipeline) Get run logs of the specified pipeline activity.
[**GetBranchStepLog**](DevOpsPipelineAPI.md#GetBranchStepLog) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/nodes/{node}/steps/{step}/log | (MultiBranchesPipeline) Get the step logs in the specified pipeline activity.
[**GetConsoleLog**](DevOpsPipelineAPI.md#GetConsoleLog) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/consolelog | Get scan reponsitory logs in the specified pipeline.
[**GetCrumb**](DevOpsPipelineAPI.md#GetCrumb) | **Get** /kapis/devops.kubesphere.io/v1alpha2/crumbissuer | Get crumb issuer. A CrumbIssuer represents an algorithm to generate a nonce value, known as a crumb, to counter cross site request forgery exploits. Crumbs are typically hashes incorporating information that uniquely identifies an agent that sends a request, along with a guarded secret so that the crumb value cannot be forged by a third party.
[**GetNodeSteps**](DevOpsPipelineAPI.md#GetNodeSteps) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/nodes/{node}/steps | Get all steps in the specified node.
[**GetNodesDetail**](DevOpsPipelineAPI.md#GetNodesDetail) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/nodesdetail | Get steps details inside a activity node. For a node, the steps which defined inside the node.
[**GetPipeline**](DevOpsPipelineAPI.md#GetPipeline) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline} | Get the specified pipeline of the DevOps project
[**GetPipelineBranch**](DevOpsPipelineAPI.md#GetPipelineBranch) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches | (MultiBranchesPipeline) Get all branches in the specified pipeline.
[**GetPipelineRun**](DevOpsPipelineAPI.md#GetPipelineRun) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run} | Get details in the specified pipeline activity.
[**GetPipelineRunNodes**](DevOpsPipelineAPI.md#GetPipelineRunNodes) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/nodes | Get all nodes in the specified activity. node is the stage in the pipeline task
[**GetRunLog**](DevOpsPipelineAPI.md#GetRunLog) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/log | Get run logs of the specified pipeline activity.
[**GetStepLog**](DevOpsPipelineAPI.md#GetStepLog) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/nodes/{node}/steps/{step}/log | Get pipelines step log.
[**ListPipelineRuns**](DevOpsPipelineAPI.md#ListPipelineRuns) | **Get** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs | Get all runs of the specified pipeline
[**ListPipelines**](DevOpsPipelineAPI.md#ListPipelines) | **Get** /kapis/devops.kubesphere.io/v1alpha2/search | Search DevOps resource. More info: https://github.com/jenkinsci/blueocean-plugin/tree/master/blueocean-rest#get-pipelines-across-organization
[**ReplayBranchPipeline**](DevOpsPipelineAPI.md#ReplayBranchPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/replay | (MultiBranchesPipeline) Replay the specified pipeline of the DevOps project
[**ReplayPipeline**](DevOpsPipelineAPI.md#ReplayPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/replay | Replay pipeline
[**RunBranchPipeline**](DevOpsPipelineAPI.md#RunBranchPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs | (MultiBranchesPipeline) Run the specified pipeline of the DevOps project.
[**RunPipeline**](DevOpsPipelineAPI.md#RunPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs | Run pipeline.
[**ScanBranch**](DevOpsPipelineAPI.md#ScanBranch) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/scan | Scan remote Repository, Start a build if have new branch.
[**StopBranchPipeline**](DevOpsPipelineAPI.md#StopBranchPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/stop | (MultiBranchesPipeline) Stop the specified pipeline of the DevOps project.
[**StopPipeline**](DevOpsPipelineAPI.md#StopPipeline) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/stop | Stop pipeline
[**SubmitBranchInputStep**](DevOpsPipelineAPI.md#SubmitBranchInputStep) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/branches/{branch}/runs/{run}/nodes/{node}/steps/{step} | (MultiBranchesPipeline) Proceed or Break the paused pipeline which waiting for user input.
[**SubmitInputStep**](DevOpsPipelineAPI.md#SubmitInputStep) | **Post** /kapis/devops.kubesphere.io/v1alpha2/devops/{devops}/pipelines/{pipeline}/runs/{run}/nodes/{node}/steps/{step} | Proceed or Break the paused pipeline which is waiting for user input.



## CheckCron

> DevopsCheckCronRes CheckCron(ctx, devops).Body(body).Execute()

Check cron script compile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	body := *openapiclient.NewDevopsCronData("Cron_example") // DevopsCronData | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.CheckCron(context.Background(), devops).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.CheckCron``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckCron`: DevopsCheckCronRes
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.CheckCron`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckCronRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**DevopsCronData**](DevopsCronData.md) |  | 

### Return type

[**DevopsCheckCronRes**](DevopsCheckCronRes.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CheckScriptCompile

> DevopsCheckScript CheckScriptCompile(ctx, devops, pipeline).Value(value).Execute()

Check pipeline script compile.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	value := "value_example" // string | Pipeline script data (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.CheckScriptCompile(context.Background(), devops, pipeline).Value(value).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.CheckScriptCompile``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CheckScriptCompile`: DevopsCheckScript
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.CheckScriptCompile`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiCheckScriptCompileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **value** | **string** | Pipeline script data | 

### Return type

[**DevopsCheckScript**](DevopsCheckScript.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: application/x-www-form-urlencoded, charset=utf-8
- **Accept**: application/json, charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeleteCredential

> []V1Secret DeleteCredential(ctx, devops, credential).Execute()

delete the credential of the specified devops for the current user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | project name
	credential := "credential_example" // string | credential name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.DeleteCredential(context.Background(), devops, credential).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.DeleteCredential``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteCredential`: []V1Secret
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.DeleteCredential`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**credential** | **string** | credential name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCredentialRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1Secret**](V1Secret.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePipeline

> []V1alpha3Pipeline DeletePipeline(ctx, devops, pipeline).Execute()

delete the pipeline of the specified devops for the current user

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | project name
	pipeline := "pipeline_example" // string | pipeline name

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.DeletePipeline(context.Background(), devops, pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.DeletePipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeletePipeline`: []V1alpha3Pipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.DeletePipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | project name | 
**pipeline** | **string** | pipeline name | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]V1alpha3Pipeline**](V1alpha3Pipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetArtifacts

> []DevopsArtifacts GetArtifacts(ctx, devops, pipeline, run).Start(start).Limit(limit).Execute()

Get all artifacts in the specified pipeline.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	start := "start_example" // string | the item number that the search starts from. (optional)
	limit := "limit_example" // string | the limit item count of the search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetArtifacts(context.Background(), devops, pipeline, run).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetArtifacts`: []DevopsArtifacts
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **string** | the item number that the search starts from. | 
 **limit** | **string** | the limit item count of the search. | 

### Return type

[**[]DevopsArtifacts**](DevopsArtifacts.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchArtifacts

> []DevopsArtifacts GetBranchArtifacts(ctx, devops, pipeline, branch, run).Start(start).Limit(limit).Execute()

(MultiBranchesPipeline) Get all artifacts generated from the specified run of the pipeline branch.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	start := "start_example" // string | the item number that the search starts from. (optional)
	limit := "limit_example" // string | the limit item count of the search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchArtifacts(context.Background(), devops, pipeline, branch, run).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchArtifacts``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchArtifacts`: []DevopsArtifacts
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchArtifacts`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchArtifactsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **string** | the item number that the search starts from. | 
 **limit** | **string** | the limit item count of the search. | 

### Return type

[**[]DevopsArtifacts**](DevopsArtifacts.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchNodeSteps

> []DevopsNodeSteps GetBranchNodeSteps(ctx, devops, pipeline, branch, run, node).Execute()

(MultiBranchesPipeline) Get all steps in the specified node.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | the name of devops project
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	node := "node_example" // string | pipeline node ID, the stage in pipeline.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchNodeSteps(context.Background(), devops, pipeline, branch, run, node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchNodeSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchNodeSteps`: []DevopsNodeSteps
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchNodeSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | the name of devops project | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 
**node** | **string** | pipeline node ID, the stage in pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchNodeStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






### Return type

[**[]DevopsNodeSteps**](DevopsNodeSteps.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchNodesDetail

> []DevopsNodesDetail GetBranchNodesDetail(ctx, devops, pipeline, branch, run).Execute()

(MultiBranchesPipeline) Get steps details in an activity node. For a node, the steps which is defined inside the node.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchNodesDetail(context.Background(), devops, pipeline, branch, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchNodesDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchNodesDetail`: []DevopsNodesDetail
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchNodesDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchNodesDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]DevopsNodesDetail**](DevopsNodesDetail.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchPipeline

> DevopsBranchPipeline GetBranchPipeline(ctx, devops, pipeline, branch).Execute()

(MultiBranchesPipeline) Get the specified branch pipeline of the DevOps project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | the name of devops project
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchPipeline(context.Background(), devops, pipeline, branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchPipeline`: DevopsBranchPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | the name of devops project | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DevopsBranchPipeline**](DevopsBranchPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchPipelineRun

> DevopsPipelineRun GetBranchPipelineRun(ctx, devops, pipeline, branch, run).Execute()

(MultiBranchesPipeline) Get details in the specified pipeline activity.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run id, the unique id for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchPipelineRun(context.Background(), devops, pipeline, branch, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchPipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchPipelineRun`: DevopsPipelineRun
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchPipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run id, the unique id for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DevopsPipelineRun**](DevopsPipelineRun.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchPipelineRunNodes

> []DevopsBranchPipelineRunNodes GetBranchPipelineRunNodes(ctx, devops, pipeline, branch, run).Limit(limit).Execute()

(MultiBranchesPipeline) Get run nodes.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run id, the unique id for a pipeline once build.
	limit := "limit_example" // string | the limit item count of the search. (optional) (default to "limit=10000")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetBranchPipelineRunNodes(context.Background(), devops, pipeline, branch, run).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchPipelineRunNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetBranchPipelineRunNodes`: []DevopsBranchPipelineRunNodes
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetBranchPipelineRunNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run id, the unique id for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchPipelineRunNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **limit** | **string** | the limit item count of the search. | [default to &quot;limit&#x3D;10000&quot;]

### Return type

[**[]DevopsBranchPipelineRunNodes**](DevopsBranchPipelineRunNodes.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchRunLog

> GetBranchRunLog(ctx, devops, pipeline, branch, run).Start(start).Execute()

(MultiBranchesPipeline) Get run logs of the specified pipeline activity.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	start := "start_example" // string | the item number that the search starts from. (optional) (default to "start=0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.GetBranchRunLog(context.Background(), devops, pipeline, branch, run).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchRunLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchRunLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **start** | **string** | the item number that the search starts from. | [default to &quot;start&#x3D;0&quot;]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetBranchStepLog

> GetBranchStepLog(ctx, devops, pipeline, branch, run, node, step).Start(start).Execute()

(MultiBranchesPipeline) Get the step logs in the specified pipeline activity.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run id, the unique id for a pipeline once build.
	node := "node_example" // string | pipeline node id, the stage in pipeline.
	step := "step_example" // string | pipeline step id, the step in pipeline.
	start := "start_example" // string | the item number that the search starts from. (optional) (default to "start=0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.GetBranchStepLog(context.Background(), devops, pipeline, branch, run, node, step).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetBranchStepLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run id, the unique id for a pipeline once build. | 
**node** | **string** | pipeline node id, the stage in pipeline. | 
**step** | **string** | pipeline step id, the step in pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetBranchStepLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **start** | **string** | the item number that the search starts from. | [default to &quot;start&#x3D;0&quot;]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetConsoleLog

> GetConsoleLog(ctx, devops, pipeline).Execute()

Get scan reponsitory logs in the specified pipeline.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.GetConsoleLog(context.Background(), devops, pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetConsoleLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetConsoleLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetCrumb

> DevopsCrumb GetCrumb(ctx).Execute()

Get crumb issuer. A CrumbIssuer represents an algorithm to generate a nonce value, known as a crumb, to counter cross site request forgery exploits. Crumbs are typically hashes incorporating information that uniquely identifies an agent that sends a request, along with a guarded secret so that the crumb value cannot be forged by a third party.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetCrumb(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetCrumb``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetCrumb`: DevopsCrumb
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetCrumb`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCrumbRequest struct via the builder pattern


### Return type

[**DevopsCrumb**](DevopsCrumb.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodeSteps

> []DevopsNodeSteps GetNodeSteps(ctx, devops, pipeline, run, node).Execute()

Get all steps in the specified node.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | the name of devops project
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build
	node := "node_example" // string | pipeline node ID, the stage in pipeline.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetNodeSteps(context.Background(), devops, pipeline, run, node).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetNodeSteps``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodeSteps`: []DevopsNodeSteps
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetNodeSteps`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | the name of devops project | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build | 
**node** | **string** | pipeline node ID, the stage in pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodeStepsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**[]DevopsNodeSteps**](DevopsNodeSteps.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetNodesDetail

> []DevopsNodesDetail GetNodesDetail(ctx, devops, pipeline, run).Execute()

Get steps details inside a activity node. For a node, the steps which defined inside the node.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetNodesDetail(context.Background(), devops, pipeline, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetNodesDetail``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNodesDetail`: []DevopsNodesDetail
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetNodesDetail`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetNodesDetailRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]DevopsNodesDetail**](DevopsNodesDetail.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipeline

> DevopsPipeline GetPipeline(ctx, devops, pipeline).Execute()

Get the specified pipeline of the DevOps project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetPipeline(context.Background(), devops, pipeline).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipeline`: DevopsPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**DevopsPipeline**](DevopsPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineBranch

> []map[string]interface{} GetPipelineBranch(ctx, devops, pipeline).Filter(filter).Start(start).Limit(limit).Execute()

(MultiBranchesPipeline) Get all branches in the specified pipeline.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	filter := "filter_example" // string | filter remote scm. e.g. origin (optional)
	start := "start_example" // string | the count of branches start. (optional) (default to "start=0")
	limit := "limit_example" // string | the count of branches limit. (optional) (default to "limit=100")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetPipelineBranch(context.Background(), devops, pipeline).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetPipelineBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineBranch`: []map[string]interface{}
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetPipelineBranch`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **filter** | **string** | filter remote scm. e.g. origin | 
 **start** | **string** | the count of branches start. | [default to &quot;start&#x3D;0&quot;]
 **limit** | **string** | the count of branches limit. | [default to &quot;limit&#x3D;100&quot;]

### Return type

**[]map[string]interface{}**

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRun

> DevopsPipelineRun GetPipelineRun(ctx, devops, pipeline, run).Execute()

Get details in the specified pipeline activity.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | the name of devops project
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetPipelineRun(context.Background(), devops, pipeline, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetPipelineRun``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineRun`: DevopsPipelineRun
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetPipelineRun`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | the name of devops project | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRunRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DevopsPipelineRun**](DevopsPipelineRun.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetPipelineRunNodes

> []DevopsPipelineRunNodes GetPipelineRunNodes(ctx, devops, pipeline, run).Execute()

Get all nodes in the specified activity. node is the stage in the pipeline task

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | the name of devops project
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.GetPipelineRunNodes(context.Background(), devops, pipeline, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetPipelineRunNodes``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetPipelineRunNodes`: []DevopsPipelineRunNodes
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.GetPipelineRunNodes`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | the name of devops project | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetPipelineRunNodesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**[]DevopsPipelineRunNodes**](DevopsPipelineRunNodes.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRunLog

> GetRunLog(ctx, devops, pipeline, run).Start(start).Execute()

Get run logs of the specified pipeline activity.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	start := "start_example" // string | the item number that the search starts from. (optional) (default to "start=0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.GetRunLog(context.Background(), devops, pipeline, run).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetRunLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRunLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **start** | **string** | the item number that the search starts from. | [default to &quot;start&#x3D;0&quot;]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStepLog

> GetStepLog(ctx, devops, pipeline, run, node, step).Start(start).Execute()

Get pipelines step log.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	node := "node_example" // string | pipeline node ID, the stage in pipeline.
	step := "step_example" // string | pipeline step ID, the step in pipeline.
	start := "start_example" // string | the item number that the search starts from. (optional) (default to "start=0")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.GetStepLog(context.Background(), devops, pipeline, run, node, step).Start(start).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.GetStepLog``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 
**node** | **string** | pipeline node ID, the stage in pipeline. | 
**step** | **string** | pipeline step ID, the step in pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStepLogRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **start** | **string** | the item number that the search starts from. | [default to &quot;start&#x3D;0&quot;]

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelineRuns

> DevopsPipelineRunList ListPipelineRuns(ctx, pipeline, devops).Start(start).Limit(limit).Branch(branch).Execute()

Get all runs of the specified pipeline

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	start := "start_example" // string | the item number that the search starts from (optional)
	limit := "limit_example" // string | the limit item count of the search (optional)
	branch := "branch_example" // string | the name of branch, same as repository branch, will be filtered by branch. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.ListPipelineRuns(context.Background(), pipeline, devops).Start(start).Limit(limit).Branch(branch).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.ListPipelineRuns``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPipelineRuns`: DevopsPipelineRunList
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.ListPipelineRuns`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pipeline** | **string** | the name of the CI/CD pipeline | 
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 

### Other Parameters

Other parameters are passed through a pointer to a apiListPipelineRunsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **start** | **string** | the item number that the search starts from | 
 **limit** | **string** | the limit item count of the search | 
 **branch** | **string** | the name of branch, same as repository branch, will be filtered by branch. | 

### Return type

[**DevopsPipelineRunList**](DevopsPipelineRunList.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListPipelines

> DevopsPipelineList ListPipelines(ctx).Q(q).Filter(filter).Start(start).Limit(limit).Execute()

Search DevOps resource. More info: https://github.com/jenkinsci/blueocean-plugin/tree/master/blueocean-rest#get-pipelines-across-organization

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	q := "q_example" // string | query pipelines, condition for filtering.
	filter := "filter_example" // string | Filter some types of jobs. e.g. no-folder，will not get a job of type folder (optional)
	start := "start_example" // string | the item number that the search starts from. (optional)
	limit := "limit_example" // string | the limit item count of the search. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.ListPipelines(context.Background()).Q(q).Filter(filter).Start(start).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.ListPipelines``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListPipelines`: DevopsPipelineList
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.ListPipelines`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiListPipelinesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | query pipelines, condition for filtering. | 
 **filter** | **string** | Filter some types of jobs. e.g. no-folder，will not get a job of type folder | 
 **start** | **string** | the item number that the search starts from. | 
 **limit** | **string** | the limit item count of the search. | 

### Return type

[**DevopsPipelineList**](DevopsPipelineList.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayBranchPipeline

> DevopsReplayPipeline ReplayBranchPipeline(ctx, devops, pipeline, branch, run).Execute()

(MultiBranchesPipeline) Replay the specified pipeline of the DevOps project

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.ReplayBranchPipeline(context.Background(), devops, pipeline, branch, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.ReplayBranchPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayBranchPipeline`: DevopsReplayPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.ReplayBranchPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayBranchPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





### Return type

[**DevopsReplayPipeline**](DevopsReplayPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReplayPipeline

> DevopsReplayPipeline ReplayPipeline(ctx, devops, pipeline, run).Execute()

Replay pipeline

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.ReplayPipeline(context.Background(), devops, pipeline, run).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.ReplayPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReplayPipeline`: DevopsReplayPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.ReplayPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiReplayPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




### Return type

[**DevopsReplayPipeline**](DevopsReplayPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunBranchPipeline

> DevopsRunPipeline RunBranchPipeline(ctx, devops, pipeline, branch).Body(body).Execute()

(MultiBranchesPipeline) Run the specified pipeline of the DevOps project.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	body := *openapiclient.NewDevopsRunPayload() // DevopsRunPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.RunBranchPipeline(context.Background(), devops, pipeline, branch).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.RunBranchPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunBranchPipeline`: DevopsRunPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.RunBranchPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunBranchPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **body** | [**DevopsRunPayload**](DevopsRunPayload.md) |  | 

### Return type

[**DevopsRunPipeline**](DevopsRunPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## RunPipeline

> DevopsRunPipeline RunPipeline(ctx, devops, pipeline).Body(body).Execute()

Run pipeline.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	body := *openapiclient.NewDevopsRunPayload() // DevopsRunPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.RunPipeline(context.Background(), devops, pipeline).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.RunPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `RunPipeline`: DevopsRunPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.RunPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiRunPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **body** | [**DevopsRunPayload**](DevopsRunPayload.md) |  | 

### Return type

[**DevopsRunPipeline**](DevopsRunPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ScanBranch

> ScanBranch(ctx, devops, pipeline).Delay(delay).Execute()

Scan remote Repository, Start a build if have new branch.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	delay := "delay_example" // string | the delay time to scan (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.ScanBranch(context.Background(), devops, pipeline).Delay(delay).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.ScanBranch``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 

### Other Parameters

Other parameters are passed through a pointer to a apiScanBranchRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **delay** | **string** | the delay time to scan | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopBranchPipeline

> DevopsStopPipeline StopBranchPipeline(ctx, devops, pipeline, branch, run).Blocking(blocking).TimeOutInSecs(timeOutInSecs).Execute()

(MultiBranchesPipeline) Stop the specified pipeline of the DevOps project.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	blocking := "blocking_example" // string | stop and between each retries will sleep. (optional) (default to "blocking=false")
	timeOutInSecs := "timeOutInSecs_example" // string | the time of stop and between each retries sleep. (optional) (default to "timeOutInSecs=10")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.StopBranchPipeline(context.Background(), devops, pipeline, branch, run).Blocking(blocking).TimeOutInSecs(timeOutInSecs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.StopBranchPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopBranchPipeline`: DevopsStopPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.StopBranchPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopBranchPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------




 **blocking** | **string** | stop and between each retries will sleep. | [default to &quot;blocking&#x3D;false&quot;]
 **timeOutInSecs** | **string** | the time of stop and between each retries sleep. | [default to &quot;timeOutInSecs&#x3D;10&quot;]

### Return type

[**DevopsStopPipeline**](DevopsStopPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StopPipeline

> DevopsStopPipeline StopPipeline(ctx, devops, pipeline, run).Blocking(blocking).TimeOutInSecs(timeOutInSecs).Execute()

Stop pipeline

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	blocking := "blocking_example" // string | stop and between each retries will sleep. (optional) (default to "blocking=false")
	timeOutInSecs := "timeOutInSecs_example" // string | the time of stop and between each retries sleep. (optional) (default to "timeOutInSecs=10")

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.DevOpsPipelineAPI.StopPipeline(context.Background(), devops, pipeline, run).Blocking(blocking).TimeOutInSecs(timeOutInSecs).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.StopPipeline``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StopPipeline`: DevopsStopPipeline
	fmt.Fprintf(os.Stdout, "Response from `DevOpsPipelineAPI.StopPipeline`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 

### Other Parameters

Other parameters are passed through a pointer to a apiStopPipelineRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



 **blocking** | **string** | stop and between each retries will sleep. | [default to &quot;blocking&#x3D;false&quot;]
 **timeOutInSecs** | **string** | the time of stop and between each retries sleep. | [default to &quot;timeOutInSecs&#x3D;10&quot;]

### Return type

[**DevopsStopPipeline**](DevopsStopPipeline.md)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitBranchInputStep

> SubmitBranchInputStep(ctx, devops, pipeline, branch, run, node, step).Body(body).Execute()

(MultiBranchesPipeline) Proceed or Break the paused pipeline which waiting for user input.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	branch := "branch_example" // string | the name of branch, same as repository branch.
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	node := "node_example" // string | pipeline node ID, the stage in pipeline.
	step := "step_example" // string | pipeline step ID, the step in pipeline.
	body := *openapiclient.NewDevopsCheckPlayload() // DevopsCheckPlayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.SubmitBranchInputStep(context.Background(), devops, pipeline, branch, run, node, step).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.SubmitBranchInputStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**branch** | **string** | the name of branch, same as repository branch. | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 
**node** | **string** | pipeline node ID, the stage in pipeline. | 
**step** | **string** | pipeline step ID, the step in pipeline. | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitBranchInputStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------






 **body** | [**DevopsCheckPlayload**](DevopsCheckPlayload.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SubmitInputStep

> SubmitInputStep(ctx, devops, pipeline, run, node, step).Body(body).Execute()

Proceed or Break the paused pipeline which is waiting for user input.

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/lee-doctor/kubesphere-sdk-go"
)

func main() {
	devops := "devops_example" // string | DevOps project's ID, e.g. project-RRRRAzLBlLEm
	pipeline := "pipeline_example" // string | the name of the CI/CD pipeline
	run := "run_example" // string | pipeline run ID, the unique ID for a pipeline once build.
	node := "node_example" // string | pipeline node ID, the stage in pipeline.
	step := "step_example" // string | pipeline step ID
	body := *openapiclient.NewDevopsCheckPlayload() // DevopsCheckPlayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.DevOpsPipelineAPI.SubmitInputStep(context.Background(), devops, pipeline, run, node, step).Body(body).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `DevOpsPipelineAPI.SubmitInputStep``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**devops** | **string** | DevOps project&#39;s ID, e.g. project-RRRRAzLBlLEm | 
**pipeline** | **string** | the name of the CI/CD pipeline | 
**run** | **string** | pipeline run ID, the unique ID for a pipeline once build. | 
**node** | **string** | pipeline node ID, the stage in pipeline. | 
**step** | **string** | pipeline step ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiSubmitInputStepRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------





 **body** | [**DevopsCheckPlayload**](DevopsCheckPlayload.md) |  | 

### Return type

 (empty response body)

### Authorization

[jwt](../README.md#jwt)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

