# RegistriesContainerConfig

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ArgsEscaped** | Pointer to **bool** | Command is already escaped (Windows only) | [optional] 
**AttachStderr** | Pointer to **bool** | Boolean value, attaches to stderr. | [optional] 
**AttachStdin** | Pointer to **bool** | Boolean value, attaches to stdin. | [optional] 
**AttachStdout** | Pointer to **bool** | Boolean value, attaches to stdout. | [optional] 
**Cmd** | Pointer to **[]string** | Command to run specified as a string or an array of strings. | [optional] 
**Domainname** | Pointer to **string** | A string value containing the domain name to use for the container. | [optional] 
**Entrypoint** | Pointer to **map[string]interface{}** |  | [optional] 
**Env** | Pointer to **[]string** | A list of environment variables in the form of [\&quot;VAR&#x3D;value\&quot;, ...] | [optional] 
**ExposedPorts** | Pointer to **map[string]interface{}** | An object mapping ports to an empty object in the form of: \&quot;ExposedPorts\&quot;: { \&quot;&lt;port&gt;/&lt;tcp|udp&gt;: {}\&quot; } | [optional] 
**Hostname** | Pointer to **string** | A string value containing the hostname to use for the container. | [optional] 
**Image** | Pointer to **string** | A string specifying the image name to use for the container. | [optional] 
**Labels** | Pointer to [**RegistriesLabels**](RegistriesLabels.md) |  | [optional] 
**OnBuild** | Pointer to **map[string]interface{}** |  | [optional] 
**OpenStdin** | Pointer to **bool** | Boolean value, opens stdin | [optional] 
**StdinOnce** | Pointer to **bool** | Boolean value, close stdin after the 1 attached client disconnects. | [optional] 
**StopSignal** | Pointer to **string** | Signal to stop a container as a string or unsigned integer. | [optional] 
**Tty** | Pointer to **bool** | Boolean value, Attach standard streams to a tty, including stdin if it is not closed. | [optional] 
**User** | Pointer to **string** | A string value specifying the user inside the container. | [optional] 
**Volumes** | Pointer to **map[string]interface{}** |  | [optional] 
**WorkingDir** | Pointer to **string** | A string specifying the working directory for commands to run in. | [optional] 

## Methods

### NewRegistriesContainerConfig

`func NewRegistriesContainerConfig() *RegistriesContainerConfig`

NewRegistriesContainerConfig instantiates a new RegistriesContainerConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesContainerConfigWithDefaults

`func NewRegistriesContainerConfigWithDefaults() *RegistriesContainerConfig`

NewRegistriesContainerConfigWithDefaults instantiates a new RegistriesContainerConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgsEscaped

`func (o *RegistriesContainerConfig) GetArgsEscaped() bool`

GetArgsEscaped returns the ArgsEscaped field if non-nil, zero value otherwise.

### GetArgsEscapedOk

`func (o *RegistriesContainerConfig) GetArgsEscapedOk() (*bool, bool)`

GetArgsEscapedOk returns a tuple with the ArgsEscaped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsEscaped

`func (o *RegistriesContainerConfig) SetArgsEscaped(v bool)`

SetArgsEscaped sets ArgsEscaped field to given value.

### HasArgsEscaped

`func (o *RegistriesContainerConfig) HasArgsEscaped() bool`

HasArgsEscaped returns a boolean if a field has been set.

### GetAttachStderr

`func (o *RegistriesContainerConfig) GetAttachStderr() bool`

GetAttachStderr returns the AttachStderr field if non-nil, zero value otherwise.

### GetAttachStderrOk

`func (o *RegistriesContainerConfig) GetAttachStderrOk() (*bool, bool)`

GetAttachStderrOk returns a tuple with the AttachStderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStderr

`func (o *RegistriesContainerConfig) SetAttachStderr(v bool)`

SetAttachStderr sets AttachStderr field to given value.

### HasAttachStderr

`func (o *RegistriesContainerConfig) HasAttachStderr() bool`

HasAttachStderr returns a boolean if a field has been set.

### GetAttachStdin

`func (o *RegistriesContainerConfig) GetAttachStdin() bool`

GetAttachStdin returns the AttachStdin field if non-nil, zero value otherwise.

### GetAttachStdinOk

`func (o *RegistriesContainerConfig) GetAttachStdinOk() (*bool, bool)`

GetAttachStdinOk returns a tuple with the AttachStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdin

`func (o *RegistriesContainerConfig) SetAttachStdin(v bool)`

SetAttachStdin sets AttachStdin field to given value.

### HasAttachStdin

`func (o *RegistriesContainerConfig) HasAttachStdin() bool`

HasAttachStdin returns a boolean if a field has been set.

### GetAttachStdout

`func (o *RegistriesContainerConfig) GetAttachStdout() bool`

GetAttachStdout returns the AttachStdout field if non-nil, zero value otherwise.

### GetAttachStdoutOk

`func (o *RegistriesContainerConfig) GetAttachStdoutOk() (*bool, bool)`

GetAttachStdoutOk returns a tuple with the AttachStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdout

`func (o *RegistriesContainerConfig) SetAttachStdout(v bool)`

SetAttachStdout sets AttachStdout field to given value.

### HasAttachStdout

`func (o *RegistriesContainerConfig) HasAttachStdout() bool`

HasAttachStdout returns a boolean if a field has been set.

### GetCmd

`func (o *RegistriesContainerConfig) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *RegistriesContainerConfig) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *RegistriesContainerConfig) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *RegistriesContainerConfig) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetDomainname

`func (o *RegistriesContainerConfig) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *RegistriesContainerConfig) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *RegistriesContainerConfig) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *RegistriesContainerConfig) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetEntrypoint

`func (o *RegistriesContainerConfig) GetEntrypoint() map[string]interface{}`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *RegistriesContainerConfig) GetEntrypointOk() (*map[string]interface{}, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *RegistriesContainerConfig) SetEntrypoint(v map[string]interface{})`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *RegistriesContainerConfig) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *RegistriesContainerConfig) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RegistriesContainerConfig) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RegistriesContainerConfig) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RegistriesContainerConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExposedPorts

`func (o *RegistriesContainerConfig) GetExposedPorts() map[string]interface{}`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *RegistriesContainerConfig) GetExposedPortsOk() (*map[string]interface{}, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *RegistriesContainerConfig) SetExposedPorts(v map[string]interface{})`

SetExposedPorts sets ExposedPorts field to given value.

### HasExposedPorts

`func (o *RegistriesContainerConfig) HasExposedPorts() bool`

HasExposedPorts returns a boolean if a field has been set.

### GetHostname

`func (o *RegistriesContainerConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RegistriesContainerConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RegistriesContainerConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RegistriesContainerConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetImage

`func (o *RegistriesContainerConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RegistriesContainerConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RegistriesContainerConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RegistriesContainerConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLabels

`func (o *RegistriesContainerConfig) GetLabels() RegistriesLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RegistriesContainerConfig) GetLabelsOk() (*RegistriesLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RegistriesContainerConfig) SetLabels(v RegistriesLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RegistriesContainerConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOnBuild

`func (o *RegistriesContainerConfig) GetOnBuild() map[string]interface{}`

GetOnBuild returns the OnBuild field if non-nil, zero value otherwise.

### GetOnBuildOk

`func (o *RegistriesContainerConfig) GetOnBuildOk() (*map[string]interface{}, bool)`

GetOnBuildOk returns a tuple with the OnBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBuild

`func (o *RegistriesContainerConfig) SetOnBuild(v map[string]interface{})`

SetOnBuild sets OnBuild field to given value.

### HasOnBuild

`func (o *RegistriesContainerConfig) HasOnBuild() bool`

HasOnBuild returns a boolean if a field has been set.

### GetOpenStdin

`func (o *RegistriesContainerConfig) GetOpenStdin() bool`

GetOpenStdin returns the OpenStdin field if non-nil, zero value otherwise.

### GetOpenStdinOk

`func (o *RegistriesContainerConfig) GetOpenStdinOk() (*bool, bool)`

GetOpenStdinOk returns a tuple with the OpenStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStdin

`func (o *RegistriesContainerConfig) SetOpenStdin(v bool)`

SetOpenStdin sets OpenStdin field to given value.

### HasOpenStdin

`func (o *RegistriesContainerConfig) HasOpenStdin() bool`

HasOpenStdin returns a boolean if a field has been set.

### GetStdinOnce

`func (o *RegistriesContainerConfig) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *RegistriesContainerConfig) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *RegistriesContainerConfig) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *RegistriesContainerConfig) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetStopSignal

`func (o *RegistriesContainerConfig) GetStopSignal() string`

GetStopSignal returns the StopSignal field if non-nil, zero value otherwise.

### GetStopSignalOk

`func (o *RegistriesContainerConfig) GetStopSignalOk() (*string, bool)`

GetStopSignalOk returns a tuple with the StopSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSignal

`func (o *RegistriesContainerConfig) SetStopSignal(v string)`

SetStopSignal sets StopSignal field to given value.

### HasStopSignal

`func (o *RegistriesContainerConfig) HasStopSignal() bool`

HasStopSignal returns a boolean if a field has been set.

### GetTty

`func (o *RegistriesContainerConfig) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *RegistriesContainerConfig) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *RegistriesContainerConfig) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *RegistriesContainerConfig) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUser

`func (o *RegistriesContainerConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RegistriesContainerConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RegistriesContainerConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RegistriesContainerConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVolumes

`func (o *RegistriesContainerConfig) GetVolumes() map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *RegistriesContainerConfig) GetVolumesOk() (*map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *RegistriesContainerConfig) SetVolumes(v map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *RegistriesContainerConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetWorkingDir

`func (o *RegistriesContainerConfig) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *RegistriesContainerConfig) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *RegistriesContainerConfig) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *RegistriesContainerConfig) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


