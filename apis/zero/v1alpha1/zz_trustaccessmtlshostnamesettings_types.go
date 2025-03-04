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

type SettingsInitParameters struct {

	// (Boolean) Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	// Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	ChinaNetwork *bool `json:"chinaNetwork,omitempty" tf:"china_network,omitempty"`

	// (Boolean) Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	// Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	ClientCertificateForwarding *bool `json:"clientCertificateForwarding,omitempty" tf:"client_certificate_forwarding,omitempty"`

	// (String) The hostname that these settings apply to.
	// The hostname that these settings apply to.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`
}

type SettingsObservation struct {

	// (Boolean) Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	// Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	ChinaNetwork *bool `json:"chinaNetwork,omitempty" tf:"china_network,omitempty"`

	// (Boolean) Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	// Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	ClientCertificateForwarding *bool `json:"clientCertificateForwarding,omitempty" tf:"client_certificate_forwarding,omitempty"`

	// (String) The hostname that these settings apply to.
	// The hostname that these settings apply to.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`
}

type SettingsParameters struct {

	// (Boolean) Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	// Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	// +kubebuilder:validation:Optional
	ChinaNetwork *bool `json:"chinaNetwork" tf:"china_network,omitempty"`

	// (Boolean) Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	// Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	// +kubebuilder:validation:Optional
	ClientCertificateForwarding *bool `json:"clientCertificateForwarding" tf:"client_certificate_forwarding,omitempty"`

	// (String) The hostname that these settings apply to.
	// The hostname that these settings apply to.
	// +kubebuilder:validation:Optional
	Hostname *string `json:"hostname" tf:"hostname,omitempty"`
}

type TrustAccessMtlsHostnameSettingsInitParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Attributes List) (see below for nested schema)
	Settings []SettingsInitParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TrustAccessMtlsHostnameSettingsObservation struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	// Request client certificates for this hostname in China. Can only be set to true if this zone is china network enabled.
	ChinaNetwork *bool `json:"chinaNetwork,omitempty" tf:"china_network,omitempty"`

	// (Boolean) Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	// Client Certificate Forwarding is a feature that takes the client cert provided by the eyeball to the edge, and forwards it to the origin as a HTTP header to allow logging on the origin.
	ClientCertificateForwarding *bool `json:"clientCertificateForwarding,omitempty" tf:"client_certificate_forwarding,omitempty"`

	// (String) The hostname that these settings apply to.
	// The hostname that these settings apply to.
	Hostname *string `json:"hostname,omitempty" tf:"hostname,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Attributes List) (see below for nested schema)
	Settings []SettingsObservation `json:"settings,omitempty" tf:"settings,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TrustAccessMtlsHostnameSettingsParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Attributes List) (see below for nested schema)
	// +kubebuilder:validation:Optional
	Settings []SettingsParameters `json:"settings,omitempty" tf:"settings,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// TrustAccessMtlsHostnameSettingsSpec defines the desired state of TrustAccessMtlsHostnameSettings
type TrustAccessMtlsHostnameSettingsSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustAccessMtlsHostnameSettingsParameters `json:"forProvider"`
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
	InitProvider TrustAccessMtlsHostnameSettingsInitParameters `json:"initProvider,omitempty"`
}

// TrustAccessMtlsHostnameSettingsStatus defines the observed state of TrustAccessMtlsHostnameSettings.
type TrustAccessMtlsHostnameSettingsStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustAccessMtlsHostnameSettingsObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustAccessMtlsHostnameSettings is the Schema for the TrustAccessMtlsHostnameSettingss API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare-upjet}
type TrustAccessMtlsHostnameSettings struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.settings) || (has(self.initProvider) && has(self.initProvider.settings))",message="spec.forProvider.settings is a required parameter"
	Spec   TrustAccessMtlsHostnameSettingsSpec   `json:"spec"`
	Status TrustAccessMtlsHostnameSettingsStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustAccessMtlsHostnameSettingsList contains a list of TrustAccessMtlsHostnameSettingss
type TrustAccessMtlsHostnameSettingsList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustAccessMtlsHostnameSettings `json:"items"`
}

// Repository type metadata.
var (
	TrustAccessMtlsHostnameSettings_Kind             = "TrustAccessMtlsHostnameSettings"
	TrustAccessMtlsHostnameSettings_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustAccessMtlsHostnameSettings_Kind}.String()
	TrustAccessMtlsHostnameSettings_KindAPIVersion   = TrustAccessMtlsHostnameSettings_Kind + "." + CRDGroupVersion.String()
	TrustAccessMtlsHostnameSettings_GroupVersionKind = CRDGroupVersion.WithKind(TrustAccessMtlsHostnameSettings_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustAccessMtlsHostnameSettings{}, &TrustAccessMtlsHostnameSettingsList{})
}
