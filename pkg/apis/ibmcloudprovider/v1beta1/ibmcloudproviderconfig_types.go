/*
Copyright 2021.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// +genclient
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object

// IBMCloudMachineProviderSpec is the type that will be embedded in a Machine.Spec.ProviderSpec field
// for an IBM Cloud VPC virtual machine. It is used by the IBM Cloud machine actuator to create a single Machine.
// +k8s:openapi-gen=true
type IBMCloudMachineProviderSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// vpc type Instance struct
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	// VPC name where the instance will be created
	VPC string `json:"vpc"`

	// Actuator will apply these tags to an virtual server instance if not present in additon
	// to default tags applied by the actuator
	Tags []TagSpecs `json:"tags,omitempty"`

	// TODO: Add labels to the virtual server
	// Labels map[string]string `json:"labels,omitempty"`

	// Image is the id of the custom OS image in VPC
	// Example: rchos-4-4-7 (Image name)
	Image string `json:"image"`

	// Profile indicates the flavor of instance.
	// Example: bx2-8x32 (8 vCPUs, 32 GB RAM)
	Profile string `json:"profile"`

	// Region of the virtual machine
	Region string `json:"region"`

	// Zone where the virtual server instance will be created
	Zone string `json:"zone"`

	// PrimaryNetworkInterface is required to specify subnet
	PrimaryNetworkInterface NetworkInterface `json:"primaryNetworkInterface,omitempty"`

	// TODO: Probably not needed for the worker machines
	// SSHKeys is the SSH pub keys that will be used to access virtual service instance
	// SSHKeys []*string `json:"sshKeys,omitempty"`

	// UserDataSecret holds reference to a secret which containes Instance Ignition data (User Data)
	UserDataSecret *corev1.LocalObjectReference `json:"userDataSecret,omitempty"`

	// CredentialsSecret is a reference to the secret with IBM Cloud credentials.
	CredentialsSecret *corev1.LocalObjectReference `json:"credentialsSecret,omitempty"`
}

// NetworkInterface struct
type NetworkInterface struct {
	// Subnet ID of the network interface
	Subnet string `json:"subnet,omitempty"`
}

// TagSpecs is the name:value pair for a tag
type TagSpecs struct {
	// Name and Value of the tag
	Name  string `json:"name"`
	Value string `json:"value"`
}

// TODO: want to configure Disk/Block Device Mapping for VPC instances

// // IBMCloudMetadata describes metadata for IBM Cloud.
// type IBMCloudMetadata struct {
// 	Key   string  `json:"key"`
// 	Value *string `json:"value"`
// }

// TODO: IBMCloudLoadBalancerReferece - register an instance with the LoadBalancer

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
func init() {
	SchemeBuilder.Register(&IBMCloudMachineProviderSpec{})
}