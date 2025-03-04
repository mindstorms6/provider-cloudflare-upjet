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

type RoutingSettingsInitParameters struct {

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RoutingSettingsObservation struct {

	// (String) The date and time the settings have been created.
	// The date and time the settings have been created.
	Created *string `json:"created,omitempty" tf:"created,omitempty"`

	// (Boolean) State of the zone settings for Email Routing.
	// State of the zone settings for Email Routing.
	Enabled *bool `json:"enabled,omitempty" tf:"enabled,omitempty"`

	// (String) Email Routing settings identifier.
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (String) The date and time the settings have been modified.
	// The date and time the settings have been modified.
	Modified *string `json:"modified,omitempty" tf:"modified,omitempty"`

	// (Boolean) Flag to check if the user skipped the configuration wizard.
	// Flag to check if the user skipped the configuration wizard.
	SkipWizard *bool `json:"skipWizard,omitempty" tf:"skip_wizard,omitempty"`

	// (String) Show the state of your account, and the type or configuration error.
	// Show the state of your account, and the type or configuration error.
	Status *string `json:"status,omitempty" tf:"status,omitempty"`

	// (String) Email Routing settings tag. (Deprecated, replaced by Email Routing settings identifier)
	// Email Routing settings tag. (Deprecated, replaced by Email Routing settings identifier)
	Tag *string `json:"tag,omitempty" tf:"tag,omitempty"`

	// (String) Identifier
	// Identifier
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type RoutingSettingsParameters struct {

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// RoutingSettingsSpec defines the desired state of RoutingSettings
type RoutingSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     RoutingSettingsParameters `json:"forProvider"`
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
	InitProvider RoutingSettingsInitParameters `json:"initProvider,omitempty"`
}

// RoutingSettingsStatus defines the observed state of RoutingSettings.
type RoutingSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        RoutingSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// RoutingSettings is the Schema for the RoutingSettingss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare-upjet}
type RoutingSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.zoneId) || (has(self.initProvider) && has(self.initProvider.zoneId))",message="spec.forProvider.zoneId is a required parameter"
	Spec   RoutingSettingsSpec   `json:"spec"`
	Status RoutingSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// RoutingSettingsList contains a list of RoutingSettingss
type RoutingSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []RoutingSettings `json:"items"`
}

// Repository type metadata.
var (
	RoutingSettings_Kind             = "RoutingSettings"
	RoutingSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: RoutingSettings_Kind}.String()
	RoutingSettings_KindAPIVersion   = RoutingSettings_Kind + "." + CRDGroupVersion.String()
	RoutingSettings_GroupVersionKind = CRDGroupVersion.WithKind(RoutingSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&RoutingSettings{}, &RoutingSettingsList{})
}
