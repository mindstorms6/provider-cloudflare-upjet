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

type CustomPagesInitParameters struct {

	// identity rule.
	// The uid of the custom page to use when a user is denied access after failing a non-identity rule.
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The uid of the custom page to use when a user is denied access.
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type CustomPagesObservation struct {

	// identity rule.
	// The uid of the custom page to use when a user is denied access after failing a non-identity rule.
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The uid of the custom page to use when a user is denied access.
	// The uid of the custom page to use when a user is denied access.
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type CustomPagesParameters struct {

	// identity rule.
	// The uid of the custom page to use when a user is denied access after failing a non-identity rule.
	// +kubebuilder:validation:Optional
	Forbidden *string `json:"forbidden,omitempty" tf:"forbidden,omitempty"`

	// (String) The uid of the custom page to use when a user is denied access.
	// The uid of the custom page to use when a user is denied access.
	// +kubebuilder:validation:Optional
	IdentityDenied *string `json:"identityDenied,omitempty" tf:"identity_denied,omitempty"`
}

type LoginDesignInitParameters struct {

	// (String) The background color on your login page.
	// The background color on your login page.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of your login page.
	// The text at the bottom of your login page.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of your login page.
	// The text at the top of your login page.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on your login page.
	// The URL of the logo on your login page.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on your login page.
	// The text color on your login page.
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type LoginDesignObservation struct {

	// (String) The background color on your login page.
	// The background color on your login page.
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of your login page.
	// The text at the bottom of your login page.
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of your login page.
	// The text at the top of your login page.
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on your login page.
	// The URL of the logo on your login page.
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on your login page.
	// The text color on your login page.
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type LoginDesignParameters struct {

	// (String) The background color on your login page.
	// The background color on your login page.
	// +kubebuilder:validation:Optional
	BackgroundColor *string `json:"backgroundColor,omitempty" tf:"background_color,omitempty"`

	// (String) The text at the bottom of your login page.
	// The text at the bottom of your login page.
	// +kubebuilder:validation:Optional
	FooterText *string `json:"footerText,omitempty" tf:"footer_text,omitempty"`

	// (String) The text at the top of your login page.
	// The text at the top of your login page.
	// +kubebuilder:validation:Optional
	HeaderText *string `json:"headerText,omitempty" tf:"header_text,omitempty"`

	// (String) The URL of the logo on your login page.
	// The URL of the logo on your login page.
	// +kubebuilder:validation:Optional
	LogoPath *string `json:"logoPath,omitempty" tf:"logo_path,omitempty"`

	// (String) The text color on your login page.
	// The text color on your login page.
	// +kubebuilder:validation:Optional
	TextColor *string `json:"textColor,omitempty" tf:"text_color,omitempty"`
}

type TrustOrganizationInitParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to `true`, users skip the identity provider selection step during login.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (Attributes) (see below for nested schema)
	CustomPages *CustomPagesInitParameters `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// Only in the Dashboard, regardless of user permission.
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Attributes) (see below for nested schema)
	LoginDesign *LoginDesignInitParameters `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TrustOrganizationObservation struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to `true`, users skip the identity provider selection step during login.
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (String)
	CreatedAt *string `json:"createdAt,omitempty" tf:"created_at,omitempty"`

	// (Attributes) (see below for nested schema)
	CustomPages *CustomPagesObservation `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	ID *string `json:"id,omitempty" tf:"id,omitempty"`

	// Only in the Dashboard, regardless of user permission.
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Attributes) (see below for nested schema)
	LoginDesign *LoginDesignObservation `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String)
	UpdatedAt *string `json:"updatedAt,omitempty" tf:"updated_at,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

type TrustOrganizationParameters struct {

	// (String) The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// The Account ID to use for this endpoint. Mutually exclusive with the Zone ID.
	// +kubebuilder:validation:Optional
	AccountID *string `json:"accountId,omitempty" tf:"account_id,omitempty"`

	// (Boolean) When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// When set to true, users can authenticate via WARP for any application in your organization. Application settings will take precedence over this value.
	// +kubebuilder:validation:Optional
	AllowAuthenticateViaWarp *bool `json:"allowAuthenticateViaWarp,omitempty" tf:"allow_authenticate_via_warp,omitempty"`

	// (String) The unique subdomain assigned to your Zero Trust organization.
	// The unique subdomain assigned to your Zero Trust organization.
	// +kubebuilder:validation:Optional
	AuthDomain *string `json:"authDomain,omitempty" tf:"auth_domain,omitempty"`

	// (Boolean) When set to true, users skip the identity provider selection step during login.
	// When set to `true`, users skip the identity provider selection step during login.
	// +kubebuilder:validation:Optional
	AutoRedirectToIdentity *bool `json:"autoRedirectToIdentity,omitempty" tf:"auto_redirect_to_identity,omitempty"`

	// (Attributes) (see below for nested schema)
	// +kubebuilder:validation:Optional
	CustomPages *CustomPagesParameters `json:"customPages,omitempty" tf:"custom_pages,omitempty"`

	// Only in the Dashboard, regardless of user permission.
	// Lock all settings as Read-Only in the Dashboard, regardless of user permission.
	// +kubebuilder:validation:Optional
	IsUIReadOnly *bool `json:"isUiReadOnly,omitempty" tf:"is_ui_read_only,omitempty"`

	// (Attributes) (see below for nested schema)
	// +kubebuilder:validation:Optional
	LoginDesign *LoginDesignParameters `json:"loginDesign,omitempty" tf:"login_design,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `300ms` or `2h45m`. Valid time units are: ns, us (or µs), ms, s, m, h.
	// +kubebuilder:validation:Optional
	SessionDuration *string `json:"sessionDuration,omitempty" tf:"session_duration,omitempty"`

	// (String) A description of the reason why the UI read only field is being toggled.
	// A description of the reason why the UI read only field is being toggled.
	// +kubebuilder:validation:Optional
	UIReadOnlyToggleReason *string `json:"uiReadOnlyToggleReason,omitempty" tf:"ui_read_only_toggle_reason,omitempty"`

	// (String) The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format 300ms or 2h45m. Valid time units are: ns, us (or µs), ms, s, m, h.
	// The amount of time a user seat is inactive before it expires. When the user seat exceeds the set time of inactivity, the user is removed as an active seat and no longer counts against your Teams seat count.  Minimum value for this setting is 1 month (730h). Must be in the format `300ms` or `2h45m`. Valid time units are: `ns`, `us` (or `µs`), `ms`, `s`, `m`, `h`.
	// +kubebuilder:validation:Optional
	UserSeatExpirationInactiveTime *string `json:"userSeatExpirationInactiveTime,omitempty" tf:"user_seat_expiration_inactive_time,omitempty"`

	// (String) The amount of time that tokens issued for applications will be valid. Must be in the format 30m or 2h45m. Valid time units are: m, h.
	// The amount of time that tokens issued for applications will be valid. Must be in the format `30m` or `2h45m`. Valid time units are: m, h.
	// +kubebuilder:validation:Optional
	WarpAuthSessionDuration *string `json:"warpAuthSessionDuration,omitempty" tf:"warp_auth_session_duration,omitempty"`

	// (String) The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// The Zone ID to use for this endpoint. Mutually exclusive with the Account ID.
	// +kubebuilder:validation:Optional
	ZoneID *string `json:"zoneId,omitempty" tf:"zone_id,omitempty"`
}

// TrustOrganizationSpec defines the desired state of TrustOrganization
type TrustOrganizationSpec struct {
	v1.ResourceSpec `json:",inline"`
	ForProvider     TrustOrganizationParameters `json:"forProvider"`
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
	InitProvider TrustOrganizationInitParameters `json:"initProvider,omitempty"`
}

// TrustOrganizationStatus defines the observed state of TrustOrganization.
type TrustOrganizationStatus struct {
	v1.ResourceStatus `json:",inline"`
	AtProvider        TrustOrganizationObservation `json:"atProvider,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:storageversion

// TrustOrganization is the Schema for the TrustOrganizations API.
// +kubebuilder:printcolumn:name="SYNCED",type="string",JSONPath=".status.conditions[?(@.type=='Synced')].status"
// +kubebuilder:printcolumn:name="READY",type="string",JSONPath=".status.conditions[?(@.type=='Ready')].status"
// +kubebuilder:printcolumn:name="EXTERNAL-NAME",type="string",JSONPath=".metadata.annotations.crossplane\\.io/external-name"
// +kubebuilder:printcolumn:name="AGE",type="date",JSONPath=".metadata.creationTimestamp"
// +kubebuilder:resource:scope=Cluster,categories={crossplane,managed,cloudflare-upjet}
type TrustOrganization struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec              TrustOrganizationSpec   `json:"spec"`
	Status            TrustOrganizationStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// TrustOrganizationList contains a list of TrustOrganizations
type TrustOrganizationList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []TrustOrganization `json:"items"`
}

// Repository type metadata.
var (
	TrustOrganization_Kind             = "TrustOrganization"
	TrustOrganization_GroupKind        = schema.GroupKind{Group: CRDGroup, Kind: TrustOrganization_Kind}.String()
	TrustOrganization_KindAPIVersion   = TrustOrganization_Kind + "." + CRDGroupVersion.String()
	TrustOrganization_GroupVersionKind = CRDGroupVersion.WithKind(TrustOrganization_Kind)
)

func init() {
	SchemeBuilder.Register(&TrustOrganization{}, &TrustOrganizationList{})
}
