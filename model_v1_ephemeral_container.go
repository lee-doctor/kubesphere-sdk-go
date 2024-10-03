/*
KubeSphere

KubeSphere OpenAPI

API version: v3.1.0
Contact: info@kubesphere.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"bytes"
	"fmt"
)

// checks if the V1EphemeralContainer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &V1EphemeralContainer{}

// V1EphemeralContainer An EphemeralContainer is a container that may be added temporarily to an existing pod for user-initiated activities such as debugging. Ephemeral containers have no resource or scheduling guarantees, and they will not be restarted when they exit or when a pod is removed or restarted. If an ephemeral container causes a pod to exceed its resource allocation, the pod may be evicted. Ephemeral containers may not be added by directly updating the pod spec. They must be added via the pod's ephemeralcontainers subresource, and they will appear in the pod spec once added. This is an alpha feature enabled by the EphemeralContainers feature flag.
type V1EphemeralContainer struct {
	// Arguments to the entrypoint. The docker image's CMD is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Args []string `json:"args,omitempty"`
	// Entrypoint array. Not executed within a shell. The docker image's ENTRYPOINT is used if this is not provided. Variable references $(VAR_NAME) are expanded using the container's environment. If a variable cannot be resolved, the reference in the input string will be unchanged. The $(VAR_NAME) syntax can be escaped with a double $$, ie: $$(VAR_NAME). Escaped references will never be expanded, regardless of whether the variable exists or not. Cannot be updated. More info: https://kubernetes.io/docs/tasks/inject-data-application/define-command-argument-container/#running-a-command-in-a-shell
	Command []string `json:"command,omitempty"`
	// List of environment variables to set in the container. Cannot be updated.
	Env []V1EnvVar `json:"env,omitempty"`
	// List of sources to populate environment variables in the container. The keys defined within a source must be a C_IDENTIFIER. All invalid keys will be reported as an event when the container is starting. When a key exists in multiple sources, the value associated with the last source will take precedence. Values defined by an Env with a duplicate key will take precedence. Cannot be updated.
	EnvFrom []V1EnvFromSource `json:"envFrom,omitempty"`
	// Docker image name. More info: https://kubernetes.io/docs/concepts/containers/images
	Image *string `json:"image,omitempty"`
	// Image pull policy. One of Always, Never, IfNotPresent. Defaults to Always if :latest tag is specified, or IfNotPresent otherwise. Cannot be updated. More info: https://kubernetes.io/docs/concepts/containers/images#updating-images
	ImagePullPolicy *string `json:"imagePullPolicy,omitempty"`
	Lifecycle *V1Lifecycle `json:"lifecycle,omitempty"`
	LivenessProbe *V1Probe `json:"livenessProbe,omitempty"`
	// Name of the ephemeral container specified as a DNS_LABEL. This name must be unique among all containers, init containers and ephemeral containers.
	Name string `json:"name"`
	// Ports are not allowed for ephemeral containers.
	Ports []V1ContainerPort `json:"ports,omitempty"`
	ReadinessProbe *V1Probe `json:"readinessProbe,omitempty"`
	Resources *V1ResourceRequirements `json:"resources,omitempty"`
	SecurityContext *V1SecurityContext `json:"securityContext,omitempty"`
	StartupProbe *V1Probe `json:"startupProbe,omitempty"`
	// Whether this container should allocate a buffer for stdin in the container runtime. If this is not set, reads from stdin in the container will always result in EOF. Default is false.
	Stdin *bool `json:"stdin,omitempty"`
	// Whether the container runtime should close the stdin channel after it has been opened by a single attach. When stdin is true the stdin stream will remain open across multiple attach sessions. If stdinOnce is set to true, stdin is opened on container start, is empty until the first client attaches to stdin, and then remains open and accepts data until the client disconnects, at which time stdin is closed and remains closed until the container is restarted. If this flag is false, a container processes that reads from stdin will never receive an EOF. Default is false
	StdinOnce *bool `json:"stdinOnce,omitempty"`
	// If set, the name of the container from PodSpec that this ephemeral container targets. The ephemeral container will be run in the namespaces (IPC, PID, etc) of this container. If not set then the ephemeral container is run in whatever namespaces are shared for the pod. Note that the container runtime must support this feature.
	TargetContainerName *string `json:"targetContainerName,omitempty"`
	// Optional: Path at which the file to which the container's termination message will be written is mounted into the container's filesystem. Message written is intended to be brief final status, such as an assertion failure message. Will be truncated by the node if greater than 4096 bytes. The total message length across all containers will be limited to 12kb. Defaults to /dev/termination-log. Cannot be updated.
	TerminationMessagePath *string `json:"terminationMessagePath,omitempty"`
	// Indicate how the termination message should be populated. File will use the contents of terminationMessagePath to populate the container status message on both success and failure. FallbackToLogsOnError will use the last chunk of container log output if the termination message file is empty and the container exited with an error. The log output is limited to 2048 bytes or 80 lines, whichever is smaller. Defaults to File. Cannot be updated.
	TerminationMessagePolicy *string `json:"terminationMessagePolicy,omitempty"`
	// Whether this container should allocate a TTY for itself, also requires 'stdin' to be true. Default is false.
	Tty *bool `json:"tty,omitempty"`
	// volumeDevices is the list of block devices to be used by the container.
	VolumeDevices []V1VolumeDevice `json:"volumeDevices,omitempty"`
	// Pod volumes to mount into the container's filesystem. Cannot be updated.
	VolumeMounts []V1VolumeMount `json:"volumeMounts,omitempty"`
	// Container's working directory. If not specified, the container runtime's default will be used, which might be configured in the container image. Cannot be updated.
	WorkingDir *string `json:"workingDir,omitempty"`
}

type _V1EphemeralContainer V1EphemeralContainer

// NewV1EphemeralContainer instantiates a new V1EphemeralContainer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewV1EphemeralContainer(name string) *V1EphemeralContainer {
	this := V1EphemeralContainer{}
	this.Name = name
	return &this
}

// NewV1EphemeralContainerWithDefaults instantiates a new V1EphemeralContainer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewV1EphemeralContainerWithDefaults() *V1EphemeralContainer {
	this := V1EphemeralContainer{}
	return &this
}

// GetArgs returns the Args field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetArgs() []string {
	if o == nil || IsNil(o.Args) {
		var ret []string
		return ret
	}
	return o.Args
}

// GetArgsOk returns a tuple with the Args field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetArgsOk() ([]string, bool) {
	if o == nil || IsNil(o.Args) {
		return nil, false
	}
	return o.Args, true
}

// HasArgs returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasArgs() bool {
	if o != nil && !IsNil(o.Args) {
		return true
	}

	return false
}

// SetArgs gets a reference to the given []string and assigns it to the Args field.
func (o *V1EphemeralContainer) SetArgs(v []string) {
	o.Args = v
}

// GetCommand returns the Command field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetCommand() []string {
	if o == nil || IsNil(o.Command) {
		var ret []string
		return ret
	}
	return o.Command
}

// GetCommandOk returns a tuple with the Command field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetCommandOk() ([]string, bool) {
	if o == nil || IsNil(o.Command) {
		return nil, false
	}
	return o.Command, true
}

// HasCommand returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasCommand() bool {
	if o != nil && !IsNil(o.Command) {
		return true
	}

	return false
}

// SetCommand gets a reference to the given []string and assigns it to the Command field.
func (o *V1EphemeralContainer) SetCommand(v []string) {
	o.Command = v
}

// GetEnv returns the Env field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetEnv() []V1EnvVar {
	if o == nil || IsNil(o.Env) {
		var ret []V1EnvVar
		return ret
	}
	return o.Env
}

// GetEnvOk returns a tuple with the Env field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetEnvOk() ([]V1EnvVar, bool) {
	if o == nil || IsNil(o.Env) {
		return nil, false
	}
	return o.Env, true
}

// HasEnv returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasEnv() bool {
	if o != nil && !IsNil(o.Env) {
		return true
	}

	return false
}

// SetEnv gets a reference to the given []V1EnvVar and assigns it to the Env field.
func (o *V1EphemeralContainer) SetEnv(v []V1EnvVar) {
	o.Env = v
}

// GetEnvFrom returns the EnvFrom field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetEnvFrom() []V1EnvFromSource {
	if o == nil || IsNil(o.EnvFrom) {
		var ret []V1EnvFromSource
		return ret
	}
	return o.EnvFrom
}

// GetEnvFromOk returns a tuple with the EnvFrom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetEnvFromOk() ([]V1EnvFromSource, bool) {
	if o == nil || IsNil(o.EnvFrom) {
		return nil, false
	}
	return o.EnvFrom, true
}

// HasEnvFrom returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasEnvFrom() bool {
	if o != nil && !IsNil(o.EnvFrom) {
		return true
	}

	return false
}

// SetEnvFrom gets a reference to the given []V1EnvFromSource and assigns it to the EnvFrom field.
func (o *V1EphemeralContainer) SetEnvFrom(v []V1EnvFromSource) {
	o.EnvFrom = v
}

// GetImage returns the Image field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetImage() string {
	if o == nil || IsNil(o.Image) {
		var ret string
		return ret
	}
	return *o.Image
}

// GetImageOk returns a tuple with the Image field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetImageOk() (*string, bool) {
	if o == nil || IsNil(o.Image) {
		return nil, false
	}
	return o.Image, true
}

// HasImage returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasImage() bool {
	if o != nil && !IsNil(o.Image) {
		return true
	}

	return false
}

// SetImage gets a reference to the given string and assigns it to the Image field.
func (o *V1EphemeralContainer) SetImage(v string) {
	o.Image = &v
}

// GetImagePullPolicy returns the ImagePullPolicy field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetImagePullPolicy() string {
	if o == nil || IsNil(o.ImagePullPolicy) {
		var ret string
		return ret
	}
	return *o.ImagePullPolicy
}

// GetImagePullPolicyOk returns a tuple with the ImagePullPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetImagePullPolicyOk() (*string, bool) {
	if o == nil || IsNil(o.ImagePullPolicy) {
		return nil, false
	}
	return o.ImagePullPolicy, true
}

// HasImagePullPolicy returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasImagePullPolicy() bool {
	if o != nil && !IsNil(o.ImagePullPolicy) {
		return true
	}

	return false
}

// SetImagePullPolicy gets a reference to the given string and assigns it to the ImagePullPolicy field.
func (o *V1EphemeralContainer) SetImagePullPolicy(v string) {
	o.ImagePullPolicy = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetLifecycle() V1Lifecycle {
	if o == nil || IsNil(o.Lifecycle) {
		var ret V1Lifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetLifecycleOk() (*V1Lifecycle, bool) {
	if o == nil || IsNil(o.Lifecycle) {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasLifecycle() bool {
	if o != nil && !IsNil(o.Lifecycle) {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given V1Lifecycle and assigns it to the Lifecycle field.
func (o *V1EphemeralContainer) SetLifecycle(v V1Lifecycle) {
	o.Lifecycle = &v
}

// GetLivenessProbe returns the LivenessProbe field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetLivenessProbe() V1Probe {
	if o == nil || IsNil(o.LivenessProbe) {
		var ret V1Probe
		return ret
	}
	return *o.LivenessProbe
}

// GetLivenessProbeOk returns a tuple with the LivenessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetLivenessProbeOk() (*V1Probe, bool) {
	if o == nil || IsNil(o.LivenessProbe) {
		return nil, false
	}
	return o.LivenessProbe, true
}

// HasLivenessProbe returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasLivenessProbe() bool {
	if o != nil && !IsNil(o.LivenessProbe) {
		return true
	}

	return false
}

// SetLivenessProbe gets a reference to the given V1Probe and assigns it to the LivenessProbe field.
func (o *V1EphemeralContainer) SetLivenessProbe(v V1Probe) {
	o.LivenessProbe = &v
}

// GetName returns the Name field value
func (o *V1EphemeralContainer) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *V1EphemeralContainer) SetName(v string) {
	o.Name = v
}

// GetPorts returns the Ports field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetPorts() []V1ContainerPort {
	if o == nil || IsNil(o.Ports) {
		var ret []V1ContainerPort
		return ret
	}
	return o.Ports
}

// GetPortsOk returns a tuple with the Ports field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetPortsOk() ([]V1ContainerPort, bool) {
	if o == nil || IsNil(o.Ports) {
		return nil, false
	}
	return o.Ports, true
}

// HasPorts returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasPorts() bool {
	if o != nil && !IsNil(o.Ports) {
		return true
	}

	return false
}

// SetPorts gets a reference to the given []V1ContainerPort and assigns it to the Ports field.
func (o *V1EphemeralContainer) SetPorts(v []V1ContainerPort) {
	o.Ports = v
}

// GetReadinessProbe returns the ReadinessProbe field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetReadinessProbe() V1Probe {
	if o == nil || IsNil(o.ReadinessProbe) {
		var ret V1Probe
		return ret
	}
	return *o.ReadinessProbe
}

// GetReadinessProbeOk returns a tuple with the ReadinessProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetReadinessProbeOk() (*V1Probe, bool) {
	if o == nil || IsNil(o.ReadinessProbe) {
		return nil, false
	}
	return o.ReadinessProbe, true
}

// HasReadinessProbe returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasReadinessProbe() bool {
	if o != nil && !IsNil(o.ReadinessProbe) {
		return true
	}

	return false
}

// SetReadinessProbe gets a reference to the given V1Probe and assigns it to the ReadinessProbe field.
func (o *V1EphemeralContainer) SetReadinessProbe(v V1Probe) {
	o.ReadinessProbe = &v
}

// GetResources returns the Resources field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetResources() V1ResourceRequirements {
	if o == nil || IsNil(o.Resources) {
		var ret V1ResourceRequirements
		return ret
	}
	return *o.Resources
}

// GetResourcesOk returns a tuple with the Resources field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetResourcesOk() (*V1ResourceRequirements, bool) {
	if o == nil || IsNil(o.Resources) {
		return nil, false
	}
	return o.Resources, true
}

// HasResources returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasResources() bool {
	if o != nil && !IsNil(o.Resources) {
		return true
	}

	return false
}

// SetResources gets a reference to the given V1ResourceRequirements and assigns it to the Resources field.
func (o *V1EphemeralContainer) SetResources(v V1ResourceRequirements) {
	o.Resources = &v
}

// GetSecurityContext returns the SecurityContext field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetSecurityContext() V1SecurityContext {
	if o == nil || IsNil(o.SecurityContext) {
		var ret V1SecurityContext
		return ret
	}
	return *o.SecurityContext
}

// GetSecurityContextOk returns a tuple with the SecurityContext field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetSecurityContextOk() (*V1SecurityContext, bool) {
	if o == nil || IsNil(o.SecurityContext) {
		return nil, false
	}
	return o.SecurityContext, true
}

// HasSecurityContext returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasSecurityContext() bool {
	if o != nil && !IsNil(o.SecurityContext) {
		return true
	}

	return false
}

// SetSecurityContext gets a reference to the given V1SecurityContext and assigns it to the SecurityContext field.
func (o *V1EphemeralContainer) SetSecurityContext(v V1SecurityContext) {
	o.SecurityContext = &v
}

// GetStartupProbe returns the StartupProbe field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetStartupProbe() V1Probe {
	if o == nil || IsNil(o.StartupProbe) {
		var ret V1Probe
		return ret
	}
	return *o.StartupProbe
}

// GetStartupProbeOk returns a tuple with the StartupProbe field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetStartupProbeOk() (*V1Probe, bool) {
	if o == nil || IsNil(o.StartupProbe) {
		return nil, false
	}
	return o.StartupProbe, true
}

// HasStartupProbe returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasStartupProbe() bool {
	if o != nil && !IsNil(o.StartupProbe) {
		return true
	}

	return false
}

// SetStartupProbe gets a reference to the given V1Probe and assigns it to the StartupProbe field.
func (o *V1EphemeralContainer) SetStartupProbe(v V1Probe) {
	o.StartupProbe = &v
}

// GetStdin returns the Stdin field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetStdin() bool {
	if o == nil || IsNil(o.Stdin) {
		var ret bool
		return ret
	}
	return *o.Stdin
}

// GetStdinOk returns a tuple with the Stdin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetStdinOk() (*bool, bool) {
	if o == nil || IsNil(o.Stdin) {
		return nil, false
	}
	return o.Stdin, true
}

// HasStdin returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasStdin() bool {
	if o != nil && !IsNil(o.Stdin) {
		return true
	}

	return false
}

// SetStdin gets a reference to the given bool and assigns it to the Stdin field.
func (o *V1EphemeralContainer) SetStdin(v bool) {
	o.Stdin = &v
}

// GetStdinOnce returns the StdinOnce field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetStdinOnce() bool {
	if o == nil || IsNil(o.StdinOnce) {
		var ret bool
		return ret
	}
	return *o.StdinOnce
}

// GetStdinOnceOk returns a tuple with the StdinOnce field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetStdinOnceOk() (*bool, bool) {
	if o == nil || IsNil(o.StdinOnce) {
		return nil, false
	}
	return o.StdinOnce, true
}

// HasStdinOnce returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasStdinOnce() bool {
	if o != nil && !IsNil(o.StdinOnce) {
		return true
	}

	return false
}

// SetStdinOnce gets a reference to the given bool and assigns it to the StdinOnce field.
func (o *V1EphemeralContainer) SetStdinOnce(v bool) {
	o.StdinOnce = &v
}

// GetTargetContainerName returns the TargetContainerName field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetTargetContainerName() string {
	if o == nil || IsNil(o.TargetContainerName) {
		var ret string
		return ret
	}
	return *o.TargetContainerName
}

// GetTargetContainerNameOk returns a tuple with the TargetContainerName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetTargetContainerNameOk() (*string, bool) {
	if o == nil || IsNil(o.TargetContainerName) {
		return nil, false
	}
	return o.TargetContainerName, true
}

// HasTargetContainerName returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasTargetContainerName() bool {
	if o != nil && !IsNil(o.TargetContainerName) {
		return true
	}

	return false
}

// SetTargetContainerName gets a reference to the given string and assigns it to the TargetContainerName field.
func (o *V1EphemeralContainer) SetTargetContainerName(v string) {
	o.TargetContainerName = &v
}

// GetTerminationMessagePath returns the TerminationMessagePath field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetTerminationMessagePath() string {
	if o == nil || IsNil(o.TerminationMessagePath) {
		var ret string
		return ret
	}
	return *o.TerminationMessagePath
}

// GetTerminationMessagePathOk returns a tuple with the TerminationMessagePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetTerminationMessagePathOk() (*string, bool) {
	if o == nil || IsNil(o.TerminationMessagePath) {
		return nil, false
	}
	return o.TerminationMessagePath, true
}

// HasTerminationMessagePath returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasTerminationMessagePath() bool {
	if o != nil && !IsNil(o.TerminationMessagePath) {
		return true
	}

	return false
}

// SetTerminationMessagePath gets a reference to the given string and assigns it to the TerminationMessagePath field.
func (o *V1EphemeralContainer) SetTerminationMessagePath(v string) {
	o.TerminationMessagePath = &v
}

// GetTerminationMessagePolicy returns the TerminationMessagePolicy field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetTerminationMessagePolicy() string {
	if o == nil || IsNil(o.TerminationMessagePolicy) {
		var ret string
		return ret
	}
	return *o.TerminationMessagePolicy
}

// GetTerminationMessagePolicyOk returns a tuple with the TerminationMessagePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetTerminationMessagePolicyOk() (*string, bool) {
	if o == nil || IsNil(o.TerminationMessagePolicy) {
		return nil, false
	}
	return o.TerminationMessagePolicy, true
}

// HasTerminationMessagePolicy returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasTerminationMessagePolicy() bool {
	if o != nil && !IsNil(o.TerminationMessagePolicy) {
		return true
	}

	return false
}

// SetTerminationMessagePolicy gets a reference to the given string and assigns it to the TerminationMessagePolicy field.
func (o *V1EphemeralContainer) SetTerminationMessagePolicy(v string) {
	o.TerminationMessagePolicy = &v
}

// GetTty returns the Tty field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetTty() bool {
	if o == nil || IsNil(o.Tty) {
		var ret bool
		return ret
	}
	return *o.Tty
}

// GetTtyOk returns a tuple with the Tty field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetTtyOk() (*bool, bool) {
	if o == nil || IsNil(o.Tty) {
		return nil, false
	}
	return o.Tty, true
}

// HasTty returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasTty() bool {
	if o != nil && !IsNil(o.Tty) {
		return true
	}

	return false
}

// SetTty gets a reference to the given bool and assigns it to the Tty field.
func (o *V1EphemeralContainer) SetTty(v bool) {
	o.Tty = &v
}

// GetVolumeDevices returns the VolumeDevices field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetVolumeDevices() []V1VolumeDevice {
	if o == nil || IsNil(o.VolumeDevices) {
		var ret []V1VolumeDevice
		return ret
	}
	return o.VolumeDevices
}

// GetVolumeDevicesOk returns a tuple with the VolumeDevices field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetVolumeDevicesOk() ([]V1VolumeDevice, bool) {
	if o == nil || IsNil(o.VolumeDevices) {
		return nil, false
	}
	return o.VolumeDevices, true
}

// HasVolumeDevices returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasVolumeDevices() bool {
	if o != nil && !IsNil(o.VolumeDevices) {
		return true
	}

	return false
}

// SetVolumeDevices gets a reference to the given []V1VolumeDevice and assigns it to the VolumeDevices field.
func (o *V1EphemeralContainer) SetVolumeDevices(v []V1VolumeDevice) {
	o.VolumeDevices = v
}

// GetVolumeMounts returns the VolumeMounts field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetVolumeMounts() []V1VolumeMount {
	if o == nil || IsNil(o.VolumeMounts) {
		var ret []V1VolumeMount
		return ret
	}
	return o.VolumeMounts
}

// GetVolumeMountsOk returns a tuple with the VolumeMounts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetVolumeMountsOk() ([]V1VolumeMount, bool) {
	if o == nil || IsNil(o.VolumeMounts) {
		return nil, false
	}
	return o.VolumeMounts, true
}

// HasVolumeMounts returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasVolumeMounts() bool {
	if o != nil && !IsNil(o.VolumeMounts) {
		return true
	}

	return false
}

// SetVolumeMounts gets a reference to the given []V1VolumeMount and assigns it to the VolumeMounts field.
func (o *V1EphemeralContainer) SetVolumeMounts(v []V1VolumeMount) {
	o.VolumeMounts = v
}

// GetWorkingDir returns the WorkingDir field value if set, zero value otherwise.
func (o *V1EphemeralContainer) GetWorkingDir() string {
	if o == nil || IsNil(o.WorkingDir) {
		var ret string
		return ret
	}
	return *o.WorkingDir
}

// GetWorkingDirOk returns a tuple with the WorkingDir field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *V1EphemeralContainer) GetWorkingDirOk() (*string, bool) {
	if o == nil || IsNil(o.WorkingDir) {
		return nil, false
	}
	return o.WorkingDir, true
}

// HasWorkingDir returns a boolean if a field has been set.
func (o *V1EphemeralContainer) HasWorkingDir() bool {
	if o != nil && !IsNil(o.WorkingDir) {
		return true
	}

	return false
}

// SetWorkingDir gets a reference to the given string and assigns it to the WorkingDir field.
func (o *V1EphemeralContainer) SetWorkingDir(v string) {
	o.WorkingDir = &v
}

func (o V1EphemeralContainer) MarshalJSON() ([]byte, error) {
	toSerialize,err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o V1EphemeralContainer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Args) {
		toSerialize["args"] = o.Args
	}
	if !IsNil(o.Command) {
		toSerialize["command"] = o.Command
	}
	if !IsNil(o.Env) {
		toSerialize["env"] = o.Env
	}
	if !IsNil(o.EnvFrom) {
		toSerialize["envFrom"] = o.EnvFrom
	}
	if !IsNil(o.Image) {
		toSerialize["image"] = o.Image
	}
	if !IsNil(o.ImagePullPolicy) {
		toSerialize["imagePullPolicy"] = o.ImagePullPolicy
	}
	if !IsNil(o.Lifecycle) {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if !IsNil(o.LivenessProbe) {
		toSerialize["livenessProbe"] = o.LivenessProbe
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Ports) {
		toSerialize["ports"] = o.Ports
	}
	if !IsNil(o.ReadinessProbe) {
		toSerialize["readinessProbe"] = o.ReadinessProbe
	}
	if !IsNil(o.Resources) {
		toSerialize["resources"] = o.Resources
	}
	if !IsNil(o.SecurityContext) {
		toSerialize["securityContext"] = o.SecurityContext
	}
	if !IsNil(o.StartupProbe) {
		toSerialize["startupProbe"] = o.StartupProbe
	}
	if !IsNil(o.Stdin) {
		toSerialize["stdin"] = o.Stdin
	}
	if !IsNil(o.StdinOnce) {
		toSerialize["stdinOnce"] = o.StdinOnce
	}
	if !IsNil(o.TargetContainerName) {
		toSerialize["targetContainerName"] = o.TargetContainerName
	}
	if !IsNil(o.TerminationMessagePath) {
		toSerialize["terminationMessagePath"] = o.TerminationMessagePath
	}
	if !IsNil(o.TerminationMessagePolicy) {
		toSerialize["terminationMessagePolicy"] = o.TerminationMessagePolicy
	}
	if !IsNil(o.Tty) {
		toSerialize["tty"] = o.Tty
	}
	if !IsNil(o.VolumeDevices) {
		toSerialize["volumeDevices"] = o.VolumeDevices
	}
	if !IsNil(o.VolumeMounts) {
		toSerialize["volumeMounts"] = o.VolumeMounts
	}
	if !IsNil(o.WorkingDir) {
		toSerialize["workingDir"] = o.WorkingDir
	}
	return toSerialize, nil
}

func (o *V1EphemeralContainer) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err;
	}

	for _, requiredProperty := range(requiredProperties) {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varV1EphemeralContainer := _V1EphemeralContainer{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varV1EphemeralContainer)

	if err != nil {
		return err
	}

	*o = V1EphemeralContainer(varV1EphemeralContainer)

	return err
}

type NullableV1EphemeralContainer struct {
	value *V1EphemeralContainer
	isSet bool
}

func (v NullableV1EphemeralContainer) Get() *V1EphemeralContainer {
	return v.value
}

func (v *NullableV1EphemeralContainer) Set(val *V1EphemeralContainer) {
	v.value = val
	v.isSet = true
}

func (v NullableV1EphemeralContainer) IsSet() bool {
	return v.isSet
}

func (v *NullableV1EphemeralContainer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableV1EphemeralContainer(val *V1EphemeralContainer) *NullableV1EphemeralContainer {
	return &NullableV1EphemeralContainer{value: val, isSet: true}
}

func (v NullableV1EphemeralContainer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableV1EphemeralContainer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

