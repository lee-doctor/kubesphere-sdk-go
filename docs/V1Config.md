# V1Config

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgsEscaped** | Pointer to **bool** |  | [optional] 
**AttachStderr** | Pointer to **bool** |  | [optional] 
**AttachStdin** | Pointer to **bool** |  | [optional] 
**AttachStdout** | Pointer to **bool** |  | [optional] 
**Cmd** | Pointer to **[]string** |  | [optional] 
**Domainname** | Pointer to **string** |  | [optional] 
**Entrypoint** | Pointer to **[]string** |  | [optional] 
**Env** | Pointer to **[]string** |  | [optional] 
**ExposedPorts** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**Healthcheck** | Pointer to [**V1HealthConfig**](V1HealthConfig.md) |  | [optional] 
**Hostname** | Pointer to **string** |  | [optional] 
**Image** | Pointer to **string** |  | [optional] 
**Labels** | Pointer to **map[string]string** |  | [optional] 
**MacAddress** | Pointer to **string** |  | [optional] 
**NetworkDisabled** | Pointer to **bool** |  | [optional] 
**OnBuild** | Pointer to **[]string** |  | [optional] 
**OpenStdin** | Pointer to **bool** |  | [optional] 
**Shell** | Pointer to **[]string** |  | [optional] 
**StdinOnce** | Pointer to **bool** |  | [optional] 
**StopSignal** | Pointer to **string** |  | [optional] 
**Tty** | Pointer to **bool** |  | [optional] 
**User** | Pointer to **string** |  | [optional] 
**Volumes** | Pointer to **map[string]map[string]interface{}** |  | [optional] 
**WorkingDir** | Pointer to **string** |  | [optional] 

## Methods

### NewV1Config

`func NewV1Config() *V1Config`

NewV1Config instantiates a new V1Config object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewV1ConfigWithDefaults

`func NewV1ConfigWithDefaults() *V1Config`

NewV1ConfigWithDefaults instantiates a new V1Config object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgsEscaped

`func (o *V1Config) GetArgsEscaped() bool`

GetArgsEscaped returns the ArgsEscaped field if non-nil, zero value otherwise.

### GetArgsEscapedOk

`func (o *V1Config) GetArgsEscapedOk() (*bool, bool)`

GetArgsEscapedOk returns a tuple with the ArgsEscaped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsEscaped

`func (o *V1Config) SetArgsEscaped(v bool)`

SetArgsEscaped sets ArgsEscaped field to given value.

### HasArgsEscaped

`func (o *V1Config) HasArgsEscaped() bool`

HasArgsEscaped returns a boolean if a field has been set.

### GetAttachStderr

`func (o *V1Config) GetAttachStderr() bool`

GetAttachStderr returns the AttachStderr field if non-nil, zero value otherwise.

### GetAttachStderrOk

`func (o *V1Config) GetAttachStderrOk() (*bool, bool)`

GetAttachStderrOk returns a tuple with the AttachStderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStderr

`func (o *V1Config) SetAttachStderr(v bool)`

SetAttachStderr sets AttachStderr field to given value.

### HasAttachStderr

`func (o *V1Config) HasAttachStderr() bool`

HasAttachStderr returns a boolean if a field has been set.

### GetAttachStdin

`func (o *V1Config) GetAttachStdin() bool`

GetAttachStdin returns the AttachStdin field if non-nil, zero value otherwise.

### GetAttachStdinOk

`func (o *V1Config) GetAttachStdinOk() (*bool, bool)`

GetAttachStdinOk returns a tuple with the AttachStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdin

`func (o *V1Config) SetAttachStdin(v bool)`

SetAttachStdin sets AttachStdin field to given value.

### HasAttachStdin

`func (o *V1Config) HasAttachStdin() bool`

HasAttachStdin returns a boolean if a field has been set.

### GetAttachStdout

`func (o *V1Config) GetAttachStdout() bool`

GetAttachStdout returns the AttachStdout field if non-nil, zero value otherwise.

### GetAttachStdoutOk

`func (o *V1Config) GetAttachStdoutOk() (*bool, bool)`

GetAttachStdoutOk returns a tuple with the AttachStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdout

`func (o *V1Config) SetAttachStdout(v bool)`

SetAttachStdout sets AttachStdout field to given value.

### HasAttachStdout

`func (o *V1Config) HasAttachStdout() bool`

HasAttachStdout returns a boolean if a field has been set.

### GetCmd

`func (o *V1Config) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *V1Config) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *V1Config) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *V1Config) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetDomainname

`func (o *V1Config) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *V1Config) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *V1Config) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *V1Config) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetEntrypoint

`func (o *V1Config) GetEntrypoint() []string`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *V1Config) GetEntrypointOk() (*[]string, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *V1Config) SetEntrypoint(v []string)`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *V1Config) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *V1Config) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *V1Config) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *V1Config) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *V1Config) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExposedPorts

`func (o *V1Config) GetExposedPorts() map[string]map[string]interface{}`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *V1Config) GetExposedPortsOk() (*map[string]map[string]interface{}, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *V1Config) SetExposedPorts(v map[string]map[string]interface{})`

SetExposedPorts sets ExposedPorts field to given value.

### HasExposedPorts

`func (o *V1Config) HasExposedPorts() bool`

HasExposedPorts returns a boolean if a field has been set.

### GetHealthcheck

`func (o *V1Config) GetHealthcheck() V1HealthConfig`

GetHealthcheck returns the Healthcheck field if non-nil, zero value otherwise.

### GetHealthcheckOk

`func (o *V1Config) GetHealthcheckOk() (*V1HealthConfig, bool)`

GetHealthcheckOk returns a tuple with the Healthcheck field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHealthcheck

`func (o *V1Config) SetHealthcheck(v V1HealthConfig)`

SetHealthcheck sets Healthcheck field to given value.

### HasHealthcheck

`func (o *V1Config) HasHealthcheck() bool`

HasHealthcheck returns a boolean if a field has been set.

### GetHostname

`func (o *V1Config) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *V1Config) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *V1Config) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *V1Config) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetImage

`func (o *V1Config) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *V1Config) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *V1Config) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *V1Config) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLabels

`func (o *V1Config) GetLabels() map[string]string`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *V1Config) GetLabelsOk() (*map[string]string, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *V1Config) SetLabels(v map[string]string)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *V1Config) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetMacAddress

`func (o *V1Config) GetMacAddress() string`

GetMacAddress returns the MacAddress field if non-nil, zero value otherwise.

### GetMacAddressOk

`func (o *V1Config) GetMacAddressOk() (*string, bool)`

GetMacAddressOk returns a tuple with the MacAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMacAddress

`func (o *V1Config) SetMacAddress(v string)`

SetMacAddress sets MacAddress field to given value.

### HasMacAddress

`func (o *V1Config) HasMacAddress() bool`

HasMacAddress returns a boolean if a field has been set.

### GetNetworkDisabled

`func (o *V1Config) GetNetworkDisabled() bool`

GetNetworkDisabled returns the NetworkDisabled field if non-nil, zero value otherwise.

### GetNetworkDisabledOk

`func (o *V1Config) GetNetworkDisabledOk() (*bool, bool)`

GetNetworkDisabledOk returns a tuple with the NetworkDisabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNetworkDisabled

`func (o *V1Config) SetNetworkDisabled(v bool)`

SetNetworkDisabled sets NetworkDisabled field to given value.

### HasNetworkDisabled

`func (o *V1Config) HasNetworkDisabled() bool`

HasNetworkDisabled returns a boolean if a field has been set.

### GetOnBuild

`func (o *V1Config) GetOnBuild() []string`

GetOnBuild returns the OnBuild field if non-nil, zero value otherwise.

### GetOnBuildOk

`func (o *V1Config) GetOnBuildOk() (*[]string, bool)`

GetOnBuildOk returns a tuple with the OnBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBuild

`func (o *V1Config) SetOnBuild(v []string)`

SetOnBuild sets OnBuild field to given value.

### HasOnBuild

`func (o *V1Config) HasOnBuild() bool`

HasOnBuild returns a boolean if a field has been set.

### GetOpenStdin

`func (o *V1Config) GetOpenStdin() bool`

GetOpenStdin returns the OpenStdin field if non-nil, zero value otherwise.

### GetOpenStdinOk

`func (o *V1Config) GetOpenStdinOk() (*bool, bool)`

GetOpenStdinOk returns a tuple with the OpenStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStdin

`func (o *V1Config) SetOpenStdin(v bool)`

SetOpenStdin sets OpenStdin field to given value.

### HasOpenStdin

`func (o *V1Config) HasOpenStdin() bool`

HasOpenStdin returns a boolean if a field has been set.

### GetShell

`func (o *V1Config) GetShell() []string`

GetShell returns the Shell field if non-nil, zero value otherwise.

### GetShellOk

`func (o *V1Config) GetShellOk() (*[]string, bool)`

GetShellOk returns a tuple with the Shell field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShell

`func (o *V1Config) SetShell(v []string)`

SetShell sets Shell field to given value.

### HasShell

`func (o *V1Config) HasShell() bool`

HasShell returns a boolean if a field has been set.

### GetStdinOnce

`func (o *V1Config) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *V1Config) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *V1Config) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *V1Config) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetStopSignal

`func (o *V1Config) GetStopSignal() string`

GetStopSignal returns the StopSignal field if non-nil, zero value otherwise.

### GetStopSignalOk

`func (o *V1Config) GetStopSignalOk() (*string, bool)`

GetStopSignalOk returns a tuple with the StopSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSignal

`func (o *V1Config) SetStopSignal(v string)`

SetStopSignal sets StopSignal field to given value.

### HasStopSignal

`func (o *V1Config) HasStopSignal() bool`

HasStopSignal returns a boolean if a field has been set.

### GetTty

`func (o *V1Config) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *V1Config) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *V1Config) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *V1Config) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUser

`func (o *V1Config) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *V1Config) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *V1Config) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *V1Config) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVolumes

`func (o *V1Config) GetVolumes() map[string]map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *V1Config) GetVolumesOk() (*map[string]map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *V1Config) SetVolumes(v map[string]map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *V1Config) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetWorkingDir

`func (o *V1Config) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *V1Config) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *V1Config) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *V1Config) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


