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

type TenantObservation struct {
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

type TenantParameters struct {

	// +kubebuilder:validation:Optional
	Description *string `json:"description,omitempty" tf:"description,omitempty"`

	// +crossplane:generate:reference:type=TenantGroup
	// +crossplane:generate:reference:extractor=github.com/upbound/upjet/pkg/resource.ExtractResourceID()
	// +kubebuilder:validation:Optional
	GroupID *float64 `json:"groupId,omitempty" tf:"group_id,omitempty"`

	// Reference to a TenantGroup to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDRef *v1.Reference `json:"groupIdRef,omitempty" tf:"-"`

	// Selector for a TenantGroup to populate groupId.
	// +kubebuilder:validation:Optional
	GroupIDSelector *v1.Selector `json:"groupIdSelector,omitempty" tf:"-"`

	// +kubebuilder:validation:Optional
	Slug *string `json:"slug,omitempty" tf:"slug,omitempty"`

	// +kubebuilder:validation:Optional
	Tags []*string `json:"tags,omitempty" tf:"tags,omitempty"`
}

// TenantSpec defines the desired state of Tenant
type TenantSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TenantParameters `json:"forProvider"`
}

// TenantStatus defines the observed state of Tenant.
type TenantStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TenantObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true

// Tenant is the Schema for the Tenants API. <no value>
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,netbox}
type Tenant struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TenantSpec   `json:"spec"`
	Status            TenantStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TenantList contains a list of Tenants
type TenantList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Tenant `json:"items"`
}

// Repository type metadata.
var (
	Tenant_Kind             = "Tenant"
	Tenant_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: Tenant_Kind}.String()
	Tenant_KindAPIVersion   = Tenant_Kind + "." + CRDGroupVersion.String()
	Tenant_GroupVersionKind = CRDGroupVersion.WithKind(Tenant_Kind)
)

func init() {
	SchemeBuilder.Register(&Tenant{}, &TenantList{})
}
