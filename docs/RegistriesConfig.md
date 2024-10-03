# RegistriesConfig

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

### NewRegistriesConfig

`func NewRegistriesConfig() *RegistriesConfig`

NewRegistriesConfig instantiates a new RegistriesConfig object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistriesConfigWithDefaults

`func NewRegistriesConfigWithDefaults() *RegistriesConfig`

NewRegistriesConfigWithDefaults instantiates a new RegistriesConfig object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetArgsEscaped

`func (o *RegistriesConfig) GetArgsEscaped() bool`

GetArgsEscaped returns the ArgsEscaped field if non-nil, zero value otherwise.

### GetArgsEscapedOk

`func (o *RegistriesConfig) GetArgsEscapedOk() (*bool, bool)`

GetArgsEscapedOk returns a tuple with the ArgsEscaped field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetArgsEscaped

`func (o *RegistriesConfig) SetArgsEscaped(v bool)`

SetArgsEscaped sets ArgsEscaped field to given value.

### HasArgsEscaped

`func (o *RegistriesConfig) HasArgsEscaped() bool`

HasArgsEscaped returns a boolean if a field has been set.

### GetAttachStderr

`func (o *RegistriesConfig) GetAttachStderr() bool`

GetAttachStderr returns the AttachStderr field if non-nil, zero value otherwise.

### GetAttachStderrOk

`func (o *RegistriesConfig) GetAttachStderrOk() (*bool, bool)`

GetAttachStderrOk returns a tuple with the AttachStderr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStderr

`func (o *RegistriesConfig) SetAttachStderr(v bool)`

SetAttachStderr sets AttachStderr field to given value.

### HasAttachStderr

`func (o *RegistriesConfig) HasAttachStderr() bool`

HasAttachStderr returns a boolean if a field has been set.

### GetAttachStdin

`func (o *RegistriesConfig) GetAttachStdin() bool`

GetAttachStdin returns the AttachStdin field if non-nil, zero value otherwise.

### GetAttachStdinOk

`func (o *RegistriesConfig) GetAttachStdinOk() (*bool, bool)`

GetAttachStdinOk returns a tuple with the AttachStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdin

`func (o *RegistriesConfig) SetAttachStdin(v bool)`

SetAttachStdin sets AttachStdin field to given value.

### HasAttachStdin

`func (o *RegistriesConfig) HasAttachStdin() bool`

HasAttachStdin returns a boolean if a field has been set.

### GetAttachStdout

`func (o *RegistriesConfig) GetAttachStdout() bool`

GetAttachStdout returns the AttachStdout field if non-nil, zero value otherwise.

### GetAttachStdoutOk

`func (o *RegistriesConfig) GetAttachStdoutOk() (*bool, bool)`

GetAttachStdoutOk returns a tuple with the AttachStdout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAttachStdout

`func (o *RegistriesConfig) SetAttachStdout(v bool)`

SetAttachStdout sets AttachStdout field to given value.

### HasAttachStdout

`func (o *RegistriesConfig) HasAttachStdout() bool`

HasAttachStdout returns a boolean if a field has been set.

### GetCmd

`func (o *RegistriesConfig) GetCmd() []string`

GetCmd returns the Cmd field if non-nil, zero value otherwise.

### GetCmdOk

`func (o *RegistriesConfig) GetCmdOk() (*[]string, bool)`

GetCmdOk returns a tuple with the Cmd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCmd

`func (o *RegistriesConfig) SetCmd(v []string)`

SetCmd sets Cmd field to given value.

### HasCmd

`func (o *RegistriesConfig) HasCmd() bool`

HasCmd returns a boolean if a field has been set.

### GetDomainname

`func (o *RegistriesConfig) GetDomainname() string`

GetDomainname returns the Domainname field if non-nil, zero value otherwise.

### GetDomainnameOk

`func (o *RegistriesConfig) GetDomainnameOk() (*string, bool)`

GetDomainnameOk returns a tuple with the Domainname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDomainname

`func (o *RegistriesConfig) SetDomainname(v string)`

SetDomainname sets Domainname field to given value.

### HasDomainname

`func (o *RegistriesConfig) HasDomainname() bool`

HasDomainname returns a boolean if a field has been set.

### GetEntrypoint

`func (o *RegistriesConfig) GetEntrypoint() map[string]interface{}`

GetEntrypoint returns the Entrypoint field if non-nil, zero value otherwise.

### GetEntrypointOk

`func (o *RegistriesConfig) GetEntrypointOk() (*map[string]interface{}, bool)`

GetEntrypointOk returns a tuple with the Entrypoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEntrypoint

`func (o *RegistriesConfig) SetEntrypoint(v map[string]interface{})`

SetEntrypoint sets Entrypoint field to given value.

### HasEntrypoint

`func (o *RegistriesConfig) HasEntrypoint() bool`

HasEntrypoint returns a boolean if a field has been set.

### GetEnv

`func (o *RegistriesConfig) GetEnv() []string`

GetEnv returns the Env field if non-nil, zero value otherwise.

### GetEnvOk

`func (o *RegistriesConfig) GetEnvOk() (*[]string, bool)`

GetEnvOk returns a tuple with the Env field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnv

`func (o *RegistriesConfig) SetEnv(v []string)`

SetEnv sets Env field to given value.

### HasEnv

`func (o *RegistriesConfig) HasEnv() bool`

HasEnv returns a boolean if a field has been set.

### GetExposedPorts

`func (o *RegistriesConfig) GetExposedPorts() map[string]interface{}`

GetExposedPorts returns the ExposedPorts field if non-nil, zero value otherwise.

### GetExposedPortsOk

`func (o *RegistriesConfig) GetExposedPortsOk() (*map[string]interface{}, bool)`

GetExposedPortsOk returns a tuple with the ExposedPorts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExposedPorts

`func (o *RegistriesConfig) SetExposedPorts(v map[string]interface{})`

SetExposedPorts sets ExposedPorts field to given value.

### HasExposedPorts

`func (o *RegistriesConfig) HasExposedPorts() bool`

HasExposedPorts returns a boolean if a field has been set.

### GetHostname

`func (o *RegistriesConfig) GetHostname() string`

GetHostname returns the Hostname field if non-nil, zero value otherwise.

### GetHostnameOk

`func (o *RegistriesConfig) GetHostnameOk() (*string, bool)`

GetHostnameOk returns a tuple with the Hostname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHostname

`func (o *RegistriesConfig) SetHostname(v string)`

SetHostname sets Hostname field to given value.

### HasHostname

`func (o *RegistriesConfig) HasHostname() bool`

HasHostname returns a boolean if a field has been set.

### GetImage

`func (o *RegistriesConfig) GetImage() string`

GetImage returns the Image field if non-nil, zero value otherwise.

### GetImageOk

`func (o *RegistriesConfig) GetImageOk() (*string, bool)`

GetImageOk returns a tuple with the Image field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetImage

`func (o *RegistriesConfig) SetImage(v string)`

SetImage sets Image field to given value.

### HasImage

`func (o *RegistriesConfig) HasImage() bool`

HasImage returns a boolean if a field has been set.

### GetLabels

`func (o *RegistriesConfig) GetLabels() RegistriesLabels`

GetLabels returns the Labels field if non-nil, zero value otherwise.

### GetLabelsOk

`func (o *RegistriesConfig) GetLabelsOk() (*RegistriesLabels, bool)`

GetLabelsOk returns a tuple with the Labels field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabels

`func (o *RegistriesConfig) SetLabels(v RegistriesLabels)`

SetLabels sets Labels field to given value.

### HasLabels

`func (o *RegistriesConfig) HasLabels() bool`

HasLabels returns a boolean if a field has been set.

### GetOnBuild

`func (o *RegistriesConfig) GetOnBuild() map[string]interface{}`

GetOnBuild returns the OnBuild field if non-nil, zero value otherwise.

### GetOnBuildOk

`func (o *RegistriesConfig) GetOnBuildOk() (*map[string]interface{}, bool)`

GetOnBuildOk returns a tuple with the OnBuild field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnBuild

`func (o *RegistriesConfig) SetOnBuild(v map[string]interface{})`

SetOnBuild sets OnBuild field to given value.

### HasOnBuild

`func (o *RegistriesConfig) HasOnBuild() bool`

HasOnBuild returns a boolean if a field has been set.

### GetOpenStdin

`func (o *RegistriesConfig) GetOpenStdin() bool`

GetOpenStdin returns the OpenStdin field if non-nil, zero value otherwise.

### GetOpenStdinOk

`func (o *RegistriesConfig) GetOpenStdinOk() (*bool, bool)`

GetOpenStdinOk returns a tuple with the OpenStdin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenStdin

`func (o *RegistriesConfig) SetOpenStdin(v bool)`

SetOpenStdin sets OpenStdin field to given value.

### HasOpenStdin

`func (o *RegistriesConfig) HasOpenStdin() bool`

HasOpenStdin returns a boolean if a field has been set.

### GetStdinOnce

`func (o *RegistriesConfig) GetStdinOnce() bool`

GetStdinOnce returns the StdinOnce field if non-nil, zero value otherwise.

### GetStdinOnceOk

`func (o *RegistriesConfig) GetStdinOnceOk() (*bool, bool)`

GetStdinOnceOk returns a tuple with the StdinOnce field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStdinOnce

`func (o *RegistriesConfig) SetStdinOnce(v bool)`

SetStdinOnce sets StdinOnce field to given value.

### HasStdinOnce

`func (o *RegistriesConfig) HasStdinOnce() bool`

HasStdinOnce returns a boolean if a field has been set.

### GetStopSignal

`func (o *RegistriesConfig) GetStopSignal() string`

GetStopSignal returns the StopSignal field if non-nil, zero value otherwise.

### GetStopSignalOk

`func (o *RegistriesConfig) GetStopSignalOk() (*string, bool)`

GetStopSignalOk returns a tuple with the StopSignal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStopSignal

`func (o *RegistriesConfig) SetStopSignal(v string)`

SetStopSignal sets StopSignal field to given value.

### HasStopSignal

`func (o *RegistriesConfig) HasStopSignal() bool`

HasStopSignal returns a boolean if a field has been set.

### GetTty

`func (o *RegistriesConfig) GetTty() bool`

GetTty returns the Tty field if non-nil, zero value otherwise.

### GetTtyOk

`func (o *RegistriesConfig) GetTtyOk() (*bool, bool)`

GetTtyOk returns a tuple with the Tty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTty

`func (o *RegistriesConfig) SetTty(v bool)`

SetTty sets Tty field to given value.

### HasTty

`func (o *RegistriesConfig) HasTty() bool`

HasTty returns a boolean if a field has been set.

### GetUser

`func (o *RegistriesConfig) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *RegistriesConfig) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *RegistriesConfig) SetUser(v string)`

SetUser sets User field to given value.

### HasUser

`func (o *RegistriesConfig) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVolumes

`func (o *RegistriesConfig) GetVolumes() map[string]interface{}`

GetVolumes returns the Volumes field if non-nil, zero value otherwise.

### GetVolumesOk

`func (o *RegistriesConfig) GetVolumesOk() (*map[string]interface{}, bool)`

GetVolumesOk returns a tuple with the Volumes field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVolumes

`func (o *RegistriesConfig) SetVolumes(v map[string]interface{})`

SetVolumes sets Volumes field to given value.

### HasVolumes

`func (o *RegistriesConfig) HasVolumes() bool`

HasVolumes returns a boolean if a field has been set.

### GetWorkingDir

`func (o *RegistriesConfig) GetWorkingDir() string`

GetWorkingDir returns the WorkingDir field if non-nil, zero value otherwise.

### GetWorkingDirOk

`func (o *RegistriesConfig) GetWorkingDirOk() (*string, bool)`

GetWorkingDirOk returns a tuple with the WorkingDir field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWorkingDir

`func (o *RegistriesConfig) SetWorkingDir(v string)`

SetWorkingDir sets WorkingDir field to given value.

### HasWorkingDir

`func (o *RegistriesConfig) HasWorkingDir() bool`

HasWorkingDir returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


