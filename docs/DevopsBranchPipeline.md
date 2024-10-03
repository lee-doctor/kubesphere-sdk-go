# DevopsBranchPipeline

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Class** | Pointer to **string** | It’s a fully qualified name and is an identifier of the producer of this resource&#39;s capability. | [optional] 
**Links** | Pointer to [**Links**](Links.md) |  | [optional] 
**Actions** | Pointer to **[]map[string]interface{}** | the list of all actions | [optional] 
**Branch** | Pointer to [**Branch**](Branch.md) |  | [optional] 
**Disabled** | Pointer to **bool** | disable or not, if disabled, can not do any action | [optional] 
**DisplayName** | Pointer to **string** | display name | [optional] 
**EstimatedDurationInMillis** | Pointer to **int32** | estimated duration time in millis | [optional] 
**FullDisplayName** | Pointer to **string** | full display name | [optional] 
**FullName** | Pointer to **string** | full name | [optional] 
**LatestRun** | Pointer to [**LatestRun**](LatestRun.md) |  | [optional] 
**Name** | Pointer to **string** | name | [optional] 
**Organization** | Pointer to **string** | the name of organization | [optional] 
**Parameters** | Pointer to [**[]DevopsBranchPipelineParameters**](DevopsBranchPipelineParameters.md) |  | [optional] 
**Permissions** | Pointer to [**Permissions**](Permissions.md) |  | [optional] 
**WeatherScore** | Pointer to **int32** | the score to description the result of pipeline | [optional] 

## Methods

### NewDevopsBranchPipeline

`func NewDevopsBranchPipeline() *DevopsBranchPipeline`

NewDevopsBranchPipeline instantiates a new DevopsBranchPipeline object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDevopsBranchPipelineWithDefaults

`func NewDevopsBranchPipelineWithDefaults() *DevopsBranchPipeline`

NewDevopsBranchPipelineWithDefaults instantiates a new DevopsBranchPipeline object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClass

`func (o *DevopsBranchPipeline) GetClass() string`

GetClass returns the Class field if non-nil, zero value otherwise.

### GetClassOk

`func (o *DevopsBranchPipeline) GetClassOk() (*string, bool)`

GetClassOk returns a tuple with the Class field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClass

`func (o *DevopsBranchPipeline) SetClass(v string)`

SetClass sets Class field to given value.

### HasClass

`func (o *DevopsBranchPipeline) HasClass() bool`

HasClass returns a boolean if a field has been set.

### GetLinks

`func (o *DevopsBranchPipeline) GetLinks() Links`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *DevopsBranchPipeline) GetLinksOk() (*Links, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *DevopsBranchPipeline) SetLinks(v Links)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *DevopsBranchPipeline) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetActions

`func (o *DevopsBranchPipeline) GetActions() []map[string]interface{}`

GetActions returns the Actions field if non-nil, zero value otherwise.

### GetActionsOk

`func (o *DevopsBranchPipeline) GetActionsOk() (*[]map[string]interface{}, bool)`

GetActionsOk returns a tuple with the Actions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActions

`func (o *DevopsBranchPipeline) SetActions(v []map[string]interface{})`

SetActions sets Actions field to given value.

### HasActions

`func (o *DevopsBranchPipeline) HasActions() bool`

HasActions returns a boolean if a field has been set.

### GetBranch

`func (o *DevopsBranchPipeline) GetBranch() Branch`

GetBranch returns the Branch field if non-nil, zero value otherwise.

### GetBranchOk

`func (o *DevopsBranchPipeline) GetBranchOk() (*Branch, bool)`

GetBranchOk returns a tuple with the Branch field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBranch

`func (o *DevopsBranchPipeline) SetBranch(v Branch)`

SetBranch sets Branch field to given value.

### HasBranch

`func (o *DevopsBranchPipeline) HasBranch() bool`

HasBranch returns a boolean if a field has been set.

### GetDisabled

`func (o *DevopsBranchPipeline) GetDisabled() bool`

GetDisabled returns the Disabled field if non-nil, zero value otherwise.

### GetDisabledOk

`func (o *DevopsBranchPipeline) GetDisabledOk() (*bool, bool)`

GetDisabledOk returns a tuple with the Disabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisabled

`func (o *DevopsBranchPipeline) SetDisabled(v bool)`

SetDisabled sets Disabled field to given value.

### HasDisabled

`func (o *DevopsBranchPipeline) HasDisabled() bool`

HasDisabled returns a boolean if a field has been set.

### GetDisplayName

`func (o *DevopsBranchPipeline) GetDisplayName() string`

GetDisplayName returns the DisplayName field if non-nil, zero value otherwise.

### GetDisplayNameOk

`func (o *DevopsBranchPipeline) GetDisplayNameOk() (*string, bool)`

GetDisplayNameOk returns a tuple with the DisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayName

`func (o *DevopsBranchPipeline) SetDisplayName(v string)`

SetDisplayName sets DisplayName field to given value.

### HasDisplayName

`func (o *DevopsBranchPipeline) HasDisplayName() bool`

HasDisplayName returns a boolean if a field has been set.

### GetEstimatedDurationInMillis

`func (o *DevopsBranchPipeline) GetEstimatedDurationInMillis() int32`

GetEstimatedDurationInMillis returns the EstimatedDurationInMillis field if non-nil, zero value otherwise.

### GetEstimatedDurationInMillisOk

`func (o *DevopsBranchPipeline) GetEstimatedDurationInMillisOk() (*int32, bool)`

GetEstimatedDurationInMillisOk returns a tuple with the EstimatedDurationInMillis field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEstimatedDurationInMillis

`func (o *DevopsBranchPipeline) SetEstimatedDurationInMillis(v int32)`

SetEstimatedDurationInMillis sets EstimatedDurationInMillis field to given value.

### HasEstimatedDurationInMillis

`func (o *DevopsBranchPipeline) HasEstimatedDurationInMillis() bool`

HasEstimatedDurationInMillis returns a boolean if a field has been set.

### GetFullDisplayName

`func (o *DevopsBranchPipeline) GetFullDisplayName() string`

GetFullDisplayName returns the FullDisplayName field if non-nil, zero value otherwise.

### GetFullDisplayNameOk

`func (o *DevopsBranchPipeline) GetFullDisplayNameOk() (*string, bool)`

GetFullDisplayNameOk returns a tuple with the FullDisplayName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullDisplayName

`func (o *DevopsBranchPipeline) SetFullDisplayName(v string)`

SetFullDisplayName sets FullDisplayName field to given value.

### HasFullDisplayName

`func (o *DevopsBranchPipeline) HasFullDisplayName() bool`

HasFullDisplayName returns a boolean if a field has been set.

### GetFullName

`func (o *DevopsBranchPipeline) GetFullName() string`

GetFullName returns the FullName field if non-nil, zero value otherwise.

### GetFullNameOk

`func (o *DevopsBranchPipeline) GetFullNameOk() (*string, bool)`

GetFullNameOk returns a tuple with the FullName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFullName

`func (o *DevopsBranchPipeline) SetFullName(v string)`

SetFullName sets FullName field to given value.

### HasFullName

`func (o *DevopsBranchPipeline) HasFullName() bool`

HasFullName returns a boolean if a field has been set.

### GetLatestRun

`func (o *DevopsBranchPipeline) GetLatestRun() LatestRun`

GetLatestRun returns the LatestRun field if non-nil, zero value otherwise.

### GetLatestRunOk

`func (o *DevopsBranchPipeline) GetLatestRunOk() (*LatestRun, bool)`

GetLatestRunOk returns a tuple with the LatestRun field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLatestRun

`func (o *DevopsBranchPipeline) SetLatestRun(v LatestRun)`

SetLatestRun sets LatestRun field to given value.

### HasLatestRun

`func (o *DevopsBranchPipeline) HasLatestRun() bool`

HasLatestRun returns a boolean if a field has been set.

### GetName

`func (o *DevopsBranchPipeline) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *DevopsBranchPipeline) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *DevopsBranchPipeline) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *DevopsBranchPipeline) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOrganization

`func (o *DevopsBranchPipeline) GetOrganization() string`

GetOrganization returns the Organization field if non-nil, zero value otherwise.

### GetOrganizationOk

`func (o *DevopsBranchPipeline) GetOrganizationOk() (*string, bool)`

GetOrganizationOk returns a tuple with the Organization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOrganization

`func (o *DevopsBranchPipeline) SetOrganization(v string)`

SetOrganization sets Organization field to given value.

### HasOrganization

`func (o *DevopsBranchPipeline) HasOrganization() bool`

HasOrganization returns a boolean if a field has been set.

### GetParameters

`func (o *DevopsBranchPipeline) GetParameters() []DevopsBranchPipelineParameters`

GetParameters returns the Parameters field if non-nil, zero value otherwise.

### GetParametersOk

`func (o *DevopsBranchPipeline) GetParametersOk() (*[]DevopsBranchPipelineParameters, bool)`

GetParametersOk returns a tuple with the Parameters field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParameters

`func (o *DevopsBranchPipeline) SetParameters(v []DevopsBranchPipelineParameters)`

SetParameters sets Parameters field to given value.

### HasParameters

`func (o *DevopsBranchPipeline) HasParameters() bool`

HasParameters returns a boolean if a field has been set.

### GetPermissions

`func (o *DevopsBranchPipeline) GetPermissions() Permissions`

GetPermissions returns the Permissions field if non-nil, zero value otherwise.

### GetPermissionsOk

`func (o *DevopsBranchPipeline) GetPermissionsOk() (*Permissions, bool)`

GetPermissionsOk returns a tuple with the Permissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPermissions

`func (o *DevopsBranchPipeline) SetPermissions(v Permissions)`

SetPermissions sets Permissions field to given value.

### HasPermissions

`func (o *DevopsBranchPipeline) HasPermissions() bool`

HasPermissions returns a boolean if a field has been set.

### GetWeatherScore

`func (o *DevopsBranchPipeline) GetWeatherScore() int32`

GetWeatherScore returns the WeatherScore field if non-nil, zero value otherwise.

### GetWeatherScoreOk

`func (o *DevopsBranchPipeline) GetWeatherScoreOk() (*int32, bool)`

GetWeatherScoreOk returns a tuple with the WeatherScore field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeatherScore

`func (o *DevopsBranchPipeline) SetWeatherScore(v int32)`

SetWeatherScore sets WeatherScore field to given value.

### HasWeatherScore

`func (o *DevopsBranchPipeline) HasWeatherScore() bool`

HasWeatherScore returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


