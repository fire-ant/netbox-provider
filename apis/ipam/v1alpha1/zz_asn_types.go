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

type AsnObservation struct {
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	RirID *float64 `json:"rirId,omitempty" tf:"rir_id,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type AsnParameters struct {

	// +kubebuilder:validation:Optional
	Asn *float64 `json:"asn,omitempty" tf:"asn,omitempty"`

	// +crossplane:generate:reference:type=Rir
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	RirID *float64 `json:"rirId,omitempty" tf:"rir_id,omitempty"`

	// Reference to a Rir to populate rirId.
	// +kubebuilder:validation:Optional
	RirIDRef *v1.Reference `json:"rirIdRef,omitempty" tf:"-"`

	// Selector for a Rir to populate rirId.
	// +kubebuilder:validation:Optional
	RirIDSelector *v1.Selector `json:"rirIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// AsnSpec defines the desired state of Asn
type AsnSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     AsnParameters `json:"forProvider"`
}

// AsnStatus defines the observed state of Asn.
type AsnStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        AsnObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Asn is the Schema for the Asns API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Asn struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="self.managementPolicy == 'ObserveOnly' || has(self.forProvider.asn)",message="asn is a required parameter"
	Spec   AsnSpec   `json:"spec"`
	Status AsnStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AsnList contains a list of Asns
type AsnList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Asn `json:"items"`
}

// Repository type metadata.
var (
	Asn_Kind             = "Asn"
	Asn_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Asn_Kind}.String()
	Asn_KindAPIVersion   = Asn_Kind + "." + CRDGroupVersion.String()
	Asn_GroupVersionKind = CRDGroupVersion.WithKind(Asn_Kind)
)

func init() {
	SchemeBuilder.Register(&Asn{}, &AsnList{})
}
