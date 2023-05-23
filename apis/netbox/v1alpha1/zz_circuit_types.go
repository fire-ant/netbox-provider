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

type CircuitObservation struct {
	Cid *string `json:"cid,omitempty" tf:"cid,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	ProviderID *float64 `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`
}

type CircuitParameters struct {

	// +kubebuilder:validation:Optional
	Cid *string `json:"cid,omitempty" tf:"cid,omitempty"`

	// +kubebuilder:validation:Optional
	ProviderID *float64 `json:"providerId,omitempty" tf:"provider_id,omitempty"`

	// +kubebuilder:validation:Optional
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// +kubebuilder:validation:Optional
	TenantID *float64 `json:"tenantId,omitempty" tf:"tenant_id,omitempty"`

	// +kubebuilder:validation:Optional
	TypeID *float64 `json:"typeId,omitempty" tf:"type_id,omitempty"`
}

// CircuitSpec defines the desired state of Circuit
type CircuitSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     CircuitParameters `json:"forProvider"`
}

// CircuitStatus defines the observed state of Circuit.
type CircuitStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        CircuitObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Circuit is the Schema for the Circuits API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Circuit struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.cid)",message="cid is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.providerId)",message="providerId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.status)",message="status is a required parameter"
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.typeId)",message="typeId is a required parameter"
	Spec   CircuitSpec   `json:"spec"`
	Status CircuitStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CircuitList contains a list of Circuits
type CircuitList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Circuit `json:"items"`
}

// Repository type metadata.
var (
	Circuit_Kind             = "Circuit"
	Circuit_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Circuit_Kind}.String()
	Circuit_KindAPIVersion   = Circuit_Kind + "." + CRDGroupVersion.String()
	Circuit_GroupVersionKind = CRDGroupVersion.WithKind(Circuit_Kind)
)

func init() {
	SchemeBuilder.Register(&Circuit{}, &CircuitList{})
}
