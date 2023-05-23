/*
Copyright 2022 Upbound Inc.
*/

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type DeviceInterfaceObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	DeviceID *float64 `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Defaults to `true`.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	Mgmtonly *bool `json:"mgmtonly,omitempty" tf:"mgmtonly,omitempty"`

	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	TaggedVlans []*float64 `json:"taggedVlans,omitempty" tf:"tagged_vlans,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	UntaggedVlan *float64 `json:"untaggedVlan,omitempty" tf:"untagged_vlan,omitempty"`
}

type DeviceInterfaceParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +kubebuilder:validation:Optional
	DeviceID *float64 `json:"deviceId,omitempty" tf:"device_id,omitempty"`

	// Defaults to `true`.
	// +kubebuilder:validation:Optional
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// +kubebuilder:validation:Optional
	MacAddress *string `json:"macAddress,omitempty" tf:"mac_address,omitempty"`

	// +kubebuilder:validation:Optional
	Mgmtonly *bool `json:"mgmtonly,omitempty" tf:"mgmtonly,omitempty"`

	// +kubebuilder:validation:Optional
	Mode *string `json:"mode,omitempty" tf:"mode,omitempty"`

	// +kubebuilder:validation:Optional
	Mtu *float64 `json:"mtu,omitempty" tf:"mtu,omitempty"`

	// +kubebuilder:validation:Optional
	TaggedVlans []*float64 `json:"taggedVlans,omitempty" tf:"tagged_vlans,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`

	// +kubebuilder:validation:Optional
	Type *string `json:"type,omitempty" tf:"type,omitempty"`

	// +kubebuilder:validation:Optional
	UntaggedVlan *float64 `json:"untaggedVlan,omitempty" tf:"untagged_vlan,omitempty"`
}

// DeviceInterfaceSpec defines the desired state of DeviceInterface
type DeviceInterfaceSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     DeviceInterfaceParameters `json:"forProvider"`
}

// DeviceInterfaceStatus defines the observed state of DeviceInterface.
type DeviceInterfaceStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        DeviceInterfaceObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceInterface is the Schema for the DeviceInterfaces API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type DeviceInterface struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.deviceId)",message="deviceId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.type)",message="type is a required parameter"
	Spec   DeviceInterfaceSpec   `json:"spec"`
	Status DeviceInterfaceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// DeviceInterfaceList contains a list of DeviceInterfaces
type DeviceInterfaceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []DeviceInterface `json:"items"`
}

// Repository type metadata.
var (
	DeviceInterface_Kind             = "DeviceInterface"
	DeviceInterface_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: DeviceInterface_Kind}.String()
	DeviceInterface_KindAPIVersion   = DeviceInterface_Kind + "." + CRDGroupVersion.String()
	DeviceInterface_GroupVersionKind = CRDGroupVersion.WithKind(DeviceInterface_Kind)
)

func init() {
	SchemeBuilder.Register(&DeviceInterface{}, &DeviceInterfaceList{})
}
