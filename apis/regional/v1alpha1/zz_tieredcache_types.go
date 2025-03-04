// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime/schema"

	v1 "github.com/crossplane/crossplane-runtime/apis/common/v1"
)

type TieredCacheInitParameters struct {

	// (String) Value of the Regional Tiered Cache zone setting.
	// Value of the Regional Tiered Cache zone setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TieredCacheObservation struct {

	// (Boolean) Whether the setting is editable
	// Whether the setting is editable
	Editable *bool `json:"editable,omitempty" tf:"editable,omitempty"`

	// (String) Identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) Last time this setting was modified.
	// Last time this setting was modified.
	ModifiedOn *string `json:"modifiedOn,omitempty" tf:"modified_on,omitempty"`

	// (String) Value of the Regional Tiered Cache zone setting.
	// Value of the Regional Tiered Cache zone setting.
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TieredCacheParameters struct {

	// (String) Value of the Regional Tiered Cache zone setting.
	// Value of the Regional Tiered Cache zone setting.
	// +kubebuilder:validation:Optional
	Value *string `json:"value,omitempty" tf:"value,omitempty"`

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// TieredCacheSpec defines the desired state of TieredCache
type TieredCacheSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TieredCacheParameters `json:"forProvider"`
	// THIS IS A BETA FIELD. It will be honored
	// unless the Management Policies feature flag is disabled.
	// InitProvider holds the same fields as ForProvider, with the exception
	// of Identifier and other resource reference fields. The fields that are
	// in InitProvider are merged into ForProvider when the resource is created.
	// The same fields are also added to the terraform ignore_changes hook, to
	// avoid updating them after creation. This is useful for fields that are
	// required on creation, but we do not desire to update them after creation,
	// for example because of an external controller is managing them, like an
	// autoscaler.
	InitProvider TieredCacheInitParameters `json:"initProvider,omitempty"`
}

// TieredCacheStatus defines the observed state of TieredCache.
type TieredCacheStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TieredCacheObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TieredCache is the Schema for the TieredCaches API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare-upjet}
type TieredCache struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zoneId) || (has(self.initProvider) && has(self.initProvider.zoneId))",message="spec.forProvider.zoneId is a required parameter"
	Spec   TieredCacheSpec   `json:"spec"`
	Status TieredCacheStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TieredCacheList contains a list of TieredCaches
type TieredCacheList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TieredCache `json:"items"`
}

// Repository type metadata.
var (
	TieredCache_Kind             = "TieredCache"
	TieredCache_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TieredCache_Kind}.String()
	TieredCache_KindAPIVersion   = TieredCache_Kind + "." + CRDGroupVersion.String()
	TieredCache_GroupVersionKind = CRDGroupVersion.WithKind(TieredCache_Kind)
)

func init() {
	SchemeBuilder.Register(&TieredCache{}, &TieredCacheList{})
}
