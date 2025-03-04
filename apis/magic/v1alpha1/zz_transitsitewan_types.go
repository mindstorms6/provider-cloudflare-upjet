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

type TransitSiteWanInitParameters struct {

	// (String) Identifier
	// Identifier
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Number)
	Physport *float64 `json:"physport,omitempty" tf:"physport,omitempty"`

	// (Number)
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Identifier
	// Identifier
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (Attributes)  if omitted, use DHCP. Submit secondary_address when site is in high availability mode. (see below for nested schema)
	StaticAddressing *TransitSiteWanStaticAddressingInitParameters `json:"staticAddressing,omitempty" tf:"static_addressing,omitempty"`

	// (Number) VLAN port number.
	// VLAN port number.
	VlanTag *float64 `json:"vlanTag,omitempty" tf:"vlan_tag,omitempty"`
}

type TransitSiteWanObservation struct {

	// (String) Identifier
	// Identifier
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (String) Magic WAN health check rate for tunnels created on this link. The default value is mid.
	// Magic WAN health check rate for tunnels created on this link. The default value is `mid`.
	HealthCheckRate *string `json:"healthCheckRate,omitempty" tf:"health_check_rate,omitempty"`

	// (String) Identifier
	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// (Number)
	Physport *float64 `json:"physport,omitempty" tf:"physport,omitempty"`

	// (Number)
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Identifier
	// Identifier
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (Attributes)  if omitted, use DHCP. Submit secondary_address when site is in high availability mode. (see below for nested schema)
	StaticAddressing *TransitSiteWanStaticAddressingObservation `json:"staticAddressing,omitempty" tf:"static_addressing,omitempty"`

	// (Number) VLAN port number.
	// VLAN port number.
	VlanTag *float64 `json:"vlanTag,omitempty" tf:"vlan_tag,omitempty"`
}

type TransitSiteWanParameters struct {

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Physport *float64 `json:"physport,omitempty" tf:"physport,omitempty"`

	// (Number)
	// +kubebuilder:validation:Optional
	Priority *float64 `json:"priority,omitempty" tf:"priority,omitempty"`

	// (String) Identifier
	// Identifier
	// +kubebuilder:validation:Optional
	SiteID *string `json:"siteId,omitempty" tf:"site_id,omitempty"`

	// (Attributes)  if omitted, use DHCP. Submit secondary_address when site is in high availability mode. (see below for nested schema)
	// +kubebuilder:validation:Optional
	StaticAddressing *TransitSiteWanStaticAddressingParameters `json:"staticAddressing,omitempty" tf:"static_addressing,omitempty"`

	// (Number) VLAN port number.
	// VLAN port number.
	// +kubebuilder:validation:Optional
	VlanTag *float64 `json:"vlanTag,omitempty" tf:"vlan_tag,omitempty"`
}

type TransitSiteWanStaticAddressingInitParameters struct {

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (String) A valid IPv4 address.
	// A valid IPv4 address.
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	SecondaryAddress *string `json:"secondaryAddress,omitempty" tf:"secondary_address,omitempty"`
}

type TransitSiteWanStaticAddressingObservation struct {

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	Address *string `json:"address,omitempty" tf:"address,omitempty"`

	// (String) A valid IPv4 address.
	// A valid IPv4 address.
	GatewayAddress *string `json:"gatewayAddress,omitempty" tf:"gateway_address,omitempty"`

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	SecondaryAddress *string `json:"secondaryAddress,omitempty" tf:"secondary_address,omitempty"`
}

type TransitSiteWanStaticAddressingParameters struct {

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	// +kubebuilder:validation:Optional
	Address *string `json:"address" tf:"address,omitempty"`

	// (String) A valid IPv4 address.
	// A valid IPv4 address.
	// +kubebuilder:validation:Optional
	GatewayAddress *string `json:"gatewayAddress" tf:"gateway_address,omitempty"`

	// (String) A valid CIDR notation representing an IP range.
	// A valid CIDR notation representing an IP range.
	// +kubebuilder:validation:Optional
	SecondaryAddress *string `json:"secondaryAddress,omitempty" tf:"secondary_address,omitempty"`
}

// TransitSiteWanSpec defines the desired state of TransitSiteWan
type TransitSiteWanSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TransitSiteWanParameters `json:"forProvider"`
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
	InitProvider TransitSiteWanInitParameters `json:"initProvider,omitempty"`
}

// TransitSiteWanStatus defines the observed state of TransitSiteWan.
type TransitSiteWanStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TransitSiteWanObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TransitSiteWan is the Schema for the TransitSiteWans API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare-upjet}
type TransitSiteWan struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.accountId) || (has(self.initProvider) && has(self.initProvider.accountId))",message="spec.forProvider.accountId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.physport) || (has(self.initProvider) && has(self.initProvider.physport))",message="spec.forProvider.physport is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.siteId) || (has(self.initProvider) && has(self.initProvider.siteId))",message="spec.forProvider.siteId is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.staticAddressing) || (has(self.initProvider) && has(self.initProvider.staticAddressing))",message="spec.forProvider.staticAddressing is a required parameter"
	// +kubebuilder:validation:XValidation:rule="!('*' in self.managementPolicies || 'Create' in self.managementPolicies || 'Update' in self.managementPolicies) || has(self.forProvider.vlanTag) || (has(self.initProvider) && has(self.initProvider.vlanTag))",message="spec.forProvider.vlanTag is a required parameter"
	Spec   TransitSiteWanSpec   `json:"spec"`
	Status TransitSiteWanStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TransitSiteWanList contains a list of TransitSiteWans
type TransitSiteWanList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TransitSiteWan `json:"items"`
}

// Repository type metadata.
var (
	TransitSiteWan_Kind             = "TransitSiteWan"
	TransitSiteWan_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TransitSiteWan_Kind}.String()
	TransitSiteWan_KindAPIVersion   = TransitSiteWan_Kind + "." + CRDGroupVersion.String()
	TransitSiteWan_GroupVersionKind = CRDGroupVersion.WithKind(TransitSiteWan_Kind)
)

func init() {
	SchemeBuilder.Register(&TransitSiteWan{}, &TransitSiteWanList{})
}
