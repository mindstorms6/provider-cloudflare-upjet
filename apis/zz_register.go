// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by upjet. DO NOT EDIT.

// Package apis contains Kubernetes API for the provider.
package apis

import (
	"k8s.io/apimachinery/pkg/runtime"

	v1alpha1 "github.com/mindstorms6/provider-cloudflare-upjet/apis/access/v1alpha1"
	v1alpha1account "github.com/mindstorms6/provider-cloudflare-upjet/apis/account/v1alpha1"
	v1alpha1api "github.com/mindstorms6/provider-cloudflare-upjet/apis/api/v1alpha1"
	v1alpha1argo "github.com/mindstorms6/provider-cloudflare-upjet/apis/argo/v1alpha1"
	v1alpha1authenticated "github.com/mindstorms6/provider-cloudflare-upjet/apis/authenticated/v1alpha1"
	v1alpha1bot "github.com/mindstorms6/provider-cloudflare-upjet/apis/bot/v1alpha1"
	v1alpha1byo "github.com/mindstorms6/provider-cloudflare-upjet/apis/byo/v1alpha1"
	v1alpha1calls "github.com/mindstorms6/provider-cloudflare-upjet/apis/calls/v1alpha1"
	v1alpha1certificate "github.com/mindstorms6/provider-cloudflare-upjet/apis/certificate/v1alpha1"
	v1alpha1cloud "github.com/mindstorms6/provider-cloudflare-upjet/apis/cloud/v1alpha1"
	v1alpha1cloudflare "github.com/mindstorms6/provider-cloudflare-upjet/apis/cloudflare/v1alpha1"
	v1alpha1cloudforce "github.com/mindstorms6/provider-cloudflare-upjet/apis/cloudforce/v1alpha1"
	v1alpha1content "github.com/mindstorms6/provider-cloudflare-upjet/apis/content/v1alpha1"
	v1alpha1custom "github.com/mindstorms6/provider-cloudflare-upjet/apis/custom/v1alpha1"
	v1alpha1d1 "github.com/mindstorms6/provider-cloudflare-upjet/apis/d1/v1alpha1"
	v1alpha1dns "github.com/mindstorms6/provider-cloudflare-upjet/apis/dns/v1alpha1"
	v1alpha1email "github.com/mindstorms6/provider-cloudflare-upjet/apis/email/v1alpha1"
	v1alpha1firewall "github.com/mindstorms6/provider-cloudflare-upjet/apis/firewall/v1alpha1"
	v1alpha1hostname "github.com/mindstorms6/provider-cloudflare-upjet/apis/hostname/v1alpha1"
	v1alpha1hyperdrive "github.com/mindstorms6/provider-cloudflare-upjet/apis/hyperdrive/v1alpha1"
	v1alpha1image "github.com/mindstorms6/provider-cloudflare-upjet/apis/image/v1alpha1"
	v1alpha1keyless "github.com/mindstorms6/provider-cloudflare-upjet/apis/keyless/v1alpha1"
	v1alpha1leaked "github.com/mindstorms6/provider-cloudflare-upjet/apis/leaked/v1alpha1"
	v1alpha1list "github.com/mindstorms6/provider-cloudflare-upjet/apis/list/v1alpha1"
	v1alpha1load "github.com/mindstorms6/provider-cloudflare-upjet/apis/load/v1alpha1"
	v1alpha1logpull "github.com/mindstorms6/provider-cloudflare-upjet/apis/logpull/v1alpha1"
	v1alpha1logpush "github.com/mindstorms6/provider-cloudflare-upjet/apis/logpush/v1alpha1"
	v1alpha1magic "github.com/mindstorms6/provider-cloudflare-upjet/apis/magic/v1alpha1"
	v1alpha1managed "github.com/mindstorms6/provider-cloudflare-upjet/apis/managed/v1alpha1"
	v1alpha1mtls "github.com/mindstorms6/provider-cloudflare-upjet/apis/mtls/v1alpha1"
	v1alpha1notification "github.com/mindstorms6/provider-cloudflare-upjet/apis/notification/v1alpha1"
	v1alpha1observatory "github.com/mindstorms6/provider-cloudflare-upjet/apis/observatory/v1alpha1"
	v1alpha1origin "github.com/mindstorms6/provider-cloudflare-upjet/apis/origin/v1alpha1"
	v1alpha1page "github.com/mindstorms6/provider-cloudflare-upjet/apis/page/v1alpha1"
	v1alpha1pages "github.com/mindstorms6/provider-cloudflare-upjet/apis/pages/v1alpha1"
	v1alpha1queue "github.com/mindstorms6/provider-cloudflare-upjet/apis/queue/v1alpha1"
	v1alpha1r2 "github.com/mindstorms6/provider-cloudflare-upjet/apis/r2/v1alpha1"
	v1alpha1rate "github.com/mindstorms6/provider-cloudflare-upjet/apis/rate/v1alpha1"
	v1alpha1regional "github.com/mindstorms6/provider-cloudflare-upjet/apis/regional/v1alpha1"
	v1alpha1registrar "github.com/mindstorms6/provider-cloudflare-upjet/apis/registrar/v1alpha1"
	v1alpha1snippet "github.com/mindstorms6/provider-cloudflare-upjet/apis/snippet/v1alpha1"
	v1alpha1spectrum "github.com/mindstorms6/provider-cloudflare-upjet/apis/spectrum/v1alpha1"
	v1alpha1stream "github.com/mindstorms6/provider-cloudflare-upjet/apis/stream/v1alpha1"
	v1alpha1tiered "github.com/mindstorms6/provider-cloudflare-upjet/apis/tiered/v1alpha1"
	v1alpha1total "github.com/mindstorms6/provider-cloudflare-upjet/apis/total/v1alpha1"
	v1alpha1turnstile "github.com/mindstorms6/provider-cloudflare-upjet/apis/turnstile/v1alpha1"
	v1alpha1url "github.com/mindstorms6/provider-cloudflare-upjet/apis/url/v1alpha1"
	v1alpha1user "github.com/mindstorms6/provider-cloudflare-upjet/apis/user/v1alpha1"
	v1alpha1apis "github.com/mindstorms6/provider-cloudflare-upjet/apis/v1alpha1"
	v1beta1 "github.com/mindstorms6/provider-cloudflare-upjet/apis/v1beta1"
	v1alpha1waiting "github.com/mindstorms6/provider-cloudflare-upjet/apis/waiting/v1alpha1"
	v1alpha1web "github.com/mindstorms6/provider-cloudflare-upjet/apis/web/v1alpha1"
	v1alpha1web3 "github.com/mindstorms6/provider-cloudflare-upjet/apis/web3/v1alpha1"
	v1alpha1workers "github.com/mindstorms6/provider-cloudflare-upjet/apis/workers/v1alpha1"
	v1alpha1zero "github.com/mindstorms6/provider-cloudflare-upjet/apis/zero/v1alpha1"
	v1alpha1zone "github.com/mindstorms6/provider-cloudflare-upjet/apis/zone/v1alpha1"
)

func init() {
	// Register the types with the Scheme so the components can map objects to GroupVersionKinds and back
	AddToSchemes = append(AddToSchemes,
		v1alpha1.SchemeBuilder.AddToScheme,
		v1alpha1account.SchemeBuilder.AddToScheme,
		v1alpha1api.SchemeBuilder.AddToScheme,
		v1alpha1argo.SchemeBuilder.AddToScheme,
		v1alpha1authenticated.SchemeBuilder.AddToScheme,
		v1alpha1bot.SchemeBuilder.AddToScheme,
		v1alpha1byo.SchemeBuilder.AddToScheme,
		v1alpha1calls.SchemeBuilder.AddToScheme,
		v1alpha1certificate.SchemeBuilder.AddToScheme,
		v1alpha1cloud.SchemeBuilder.AddToScheme,
		v1alpha1cloudflare.SchemeBuilder.AddToScheme,
		v1alpha1cloudforce.SchemeBuilder.AddToScheme,
		v1alpha1content.SchemeBuilder.AddToScheme,
		v1alpha1custom.SchemeBuilder.AddToScheme,
		v1alpha1d1.SchemeBuilder.AddToScheme,
		v1alpha1dns.SchemeBuilder.AddToScheme,
		v1alpha1email.SchemeBuilder.AddToScheme,
		v1alpha1firewall.SchemeBuilder.AddToScheme,
		v1alpha1hostname.SchemeBuilder.AddToScheme,
		v1alpha1hyperdrive.SchemeBuilder.AddToScheme,
		v1alpha1image.SchemeBuilder.AddToScheme,
		v1alpha1keyless.SchemeBuilder.AddToScheme,
		v1alpha1leaked.SchemeBuilder.AddToScheme,
		v1alpha1list.SchemeBuilder.AddToScheme,
		v1alpha1load.SchemeBuilder.AddToScheme,
		v1alpha1logpull.SchemeBuilder.AddToScheme,
		v1alpha1logpush.SchemeBuilder.AddToScheme,
		v1alpha1magic.SchemeBuilder.AddToScheme,
		v1alpha1managed.SchemeBuilder.AddToScheme,
		v1alpha1mtls.SchemeBuilder.AddToScheme,
		v1alpha1notification.SchemeBuilder.AddToScheme,
		v1alpha1observatory.SchemeBuilder.AddToScheme,
		v1alpha1origin.SchemeBuilder.AddToScheme,
		v1alpha1page.SchemeBuilder.AddToScheme,
		v1alpha1pages.SchemeBuilder.AddToScheme,
		v1alpha1queue.SchemeBuilder.AddToScheme,
		v1alpha1r2.SchemeBuilder.AddToScheme,
		v1alpha1rate.SchemeBuilder.AddToScheme,
		v1alpha1regional.SchemeBuilder.AddToScheme,
		v1alpha1registrar.SchemeBuilder.AddToScheme,
		v1alpha1snippet.SchemeBuilder.AddToScheme,
		v1alpha1spectrum.SchemeBuilder.AddToScheme,
		v1alpha1stream.SchemeBuilder.AddToScheme,
		v1alpha1tiered.SchemeBuilder.AddToScheme,
		v1alpha1total.SchemeBuilder.AddToScheme,
		v1alpha1turnstile.SchemeBuilder.AddToScheme,
		v1alpha1url.SchemeBuilder.AddToScheme,
		v1alpha1user.SchemeBuilder.AddToScheme,
		v1alpha1apis.SchemeBuilder.AddToScheme,
		v1beta1.SchemeBuilder.AddToScheme,
		v1alpha1waiting.SchemeBuilder.AddToScheme,
		v1alpha1web.SchemeBuilder.AddToScheme,
		v1alpha1web3.SchemeBuilder.AddToScheme,
		v1alpha1workers.SchemeBuilder.AddToScheme,
		v1alpha1zero.SchemeBuilder.AddToScheme,
		v1alpha1zone.SchemeBuilder.AddToScheme,
	)
}

// AddToSchemes may be used to add all resources defined in the project to a Scheme
var AddToSchemes runtime.SchemeBuilder

// AddToScheme adds all Resources to the Scheme
func AddToScheme(s *runtime.Scheme) error {
	return AddToSchemes.AddToScheme(s)
}
