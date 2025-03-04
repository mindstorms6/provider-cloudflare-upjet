// SPDX-FileCopyrightText: 2024 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	rule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/access/rule"
	member "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/member"
	subscription "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/subscription"
	token "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/account/token"
	shield "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shield"
	shielddiscoveryoperation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shielddiscoveryoperation"
	shieldoperation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldoperation"
	shieldoperationschemavalidationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldoperationschemavalidationsettings"
	shieldschema "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldschema"
	shieldschemavalidationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/shieldschemavalidationsettings"
	tokenapi "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/api/token"
	smartrouting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/argo/smartrouting"
	tieredcaching "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/argo/tieredcaching"
	originpulls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/authenticated/originpulls"
	originpullscertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/authenticated/originpullscertificate"
	management "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/bot/management"
	ipprefix "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/byo/ipprefix"
	sfuapp "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/calls/sfuapp"
	turnapp "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/calls/turnapp"
	pack "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/certificate/pack"
	connectorrules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloud/connectorrules"
	account "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/account"
	filter "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/filter"
	healthcheck "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/healthcheck"
	image "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/image"
	list "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/list"
	queue "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/queue"
	ruleset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/ruleset"
	snippets "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/snippets"
	stream "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/stream"
	user "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/user"
	zone "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudflare/zone"
	onerequest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequest"
	onerequestasset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestasset"
	onerequestmessage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestmessage"
	onerequestpriority "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/cloudforce/onerequestpriority"
	scanningexpression "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/content/scanningexpression"
	hostname "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/hostname"
	hostnamefallbackorigin "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/hostnamefallbackorigin"
	ssl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/custom/ssl"
	database "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/d1/database"
	firewall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/firewall"
	record "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/record"
	settings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/settings"
	settingsinternalview "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/settingsinternalview"
	zonetransfersacl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersacl"
	zonetransfersincoming "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersincoming"
	zonetransfersoutgoing "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransfersoutgoing"
	zonetransferspeer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransferspeer"
	zonetransferstsig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/dns/zonetransferstsig"
	routingaddress "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingaddress"
	routingcatchall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingcatchall"
	routingdns "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingdns"
	routingrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingrule"
	routingsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/routingsettings"
	securityblocksender "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securityblocksender"
	securityimpersonationregistry "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securityimpersonationregistry"
	securitytrusteddomains "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/email/securitytrusteddomains"
	rulefirewall "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/firewall/rule"
	tlssetting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/hostname/tlssetting"
	config "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/hyperdrive/config"
	variant "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/image/variant"
	certificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/keyless/certificate"
	credentialcheck "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/leaked/credentialcheck"
	credentialcheckrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/leaked/credentialcheckrule"
	item "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/list/item"
	balancer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancer"
	balancermonitor "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancermonitor"
	balancerpool "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/load/balancerpool"
	retention "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpull/retention"
	job "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpush/job"
	ownershipchallenge "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/logpush/ownershipchallenge"
	networkmonitoringconfiguration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/networkmonitoringconfiguration"
	networkmonitoringrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/networkmonitoringrule"
	transitconnector "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitconnector"
	transitsite "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsite"
	transitsiteacl "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsiteacl"
	transitsitelan "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsitelan"
	transitsitewan "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/transitsitewan"
	wangretunnel "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wangretunnel"
	wanipsectunnel "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wanipsectunnel"
	wanstaticroute "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/magic/wanstaticroute"
	transforms "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/managed/transforms"
	certificatemtls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/mtls/certificate"
	policy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/notification/policy"
	policywebhooks "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/notification/policywebhooks"
	scheduledtest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/observatory/scheduledtest"
	cacertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/origin/cacertificate"
	rulepage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/page/rule"
	shieldpolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/page/shieldpolicy"
	domain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/pages/domain"
	project "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/pages/project"
	providerconfig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/providerconfig"
	consumer "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/queue/consumer"
	bucket "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucket"
	bucketcors "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketcors"
	bucketeventnotification "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketeventnotification"
	bucketlifecycle "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketlifecycle"
	bucketlock "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketlock"
	bucketsippy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/bucketsippy"
	customdomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/customdomain"
	manageddomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/r2/manageddomain"
	limit "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/rate/limit"
	hostnameregional "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/regional/hostname"
	tieredcache "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/regional/tieredcache"
	domainregistrar "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/registrar/domain"
	rules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/snippet/rules"
	application "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/spectrum/application"
	audiotrack "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/audiotrack"
	captionlanguage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/captionlanguage"
	download "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/download"
	key "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/key"
	liveinput "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/liveinput"
	watermark "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/watermark"
	webhook "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/stream/webhook"
	cache "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/tiered/cache"
	tls "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/total/tls"
	widget "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/turnstile/widget"
	normalizationsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/url/normalizationsettings"
	agentblockingrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/user/agentblockingrule"
	room "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/room"
	roomevent "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomevent"
	roomrules "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomrules"
	roomsettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/waiting/roomsettings"
	analyticsrule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web/analyticsrule"
	analyticssite "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web/analyticssite"
	hostnameweb3 "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/web3/hostname"
	crontrigger "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/crontrigger"
	customdomainworkers "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/customdomain"
	deployment "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/deployment"
	forplatformsdispatchnamespace "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/forplatformsdispatchnamespace"
	kv "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/kv"
	kvnamespace "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/kvnamespace"
	route "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/route"
	script "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/script"
	scriptsubdomain "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/scriptsubdomain"
	secret "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/workers/secret"
	trustaccessapplication "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessapplication"
	trustaccesscustompage "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesscustompage"
	trustaccessgroup "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessgroup"
	trustaccessidentityprovider "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessidentityprovider"
	trustaccessinfrastructuretarget "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessinfrastructuretarget"
	trustaccesskeyconfiguration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesskeyconfiguration"
	trustaccessmtlscertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessmtlscertificate"
	trustaccessmtlshostnamesettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessmtlshostnamesettings"
	trustaccesspolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesspolicy"
	trustaccessservicetoken "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessservicetoken"
	trustaccessshortlivedcertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccessshortlivedcertificate"
	trustaccesstag "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustaccesstag"
	trustdevicecustomprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicecustomprofile"
	trustdevicecustomprofilelocaldomainfallback "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicecustomprofilelocaldomainfallback"
	trustdevicedefaultprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofile"
	trustdevicedefaultprofilecertificates "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofilecertificates"
	trustdevicedefaultprofilelocaldomainfallback "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicedefaultprofilelocaldomainfallback"
	trustdevicemanagednetworks "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicemanagednetworks"
	trustdevicepostureintegration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdevicepostureintegration"
	trustdeviceposturerule "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdeviceposturerule"
	trustdextest "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdextest"
	trustdlpcustomprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpcustomprofile"
	trustdlpdataset "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpdataset"
	trustdlpentry "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlpentry"
	trustdlppredefinedprofile "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdlppredefinedprofile"
	trustdnslocation "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustdnslocation"
	trustgatewaycertificate "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaycertificate"
	trustgatewaylogging "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaylogging"
	trustgatewaypolicy "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaypolicy"
	trustgatewayproxyendpoint "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewayproxyendpoint"
	trustgatewaysettings "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustgatewaysettings"
	trustlist "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustlist"
	trustorganization "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustorganization"
	trustriskbehavior "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustriskbehavior"
	trustriskscoringintegration "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trustriskscoringintegration"
	trusttunnelcloudflared "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflared"
	trusttunnelcloudflaredconfig "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredconfig"
	trusttunnelcloudflaredroute "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredroute"
	trusttunnelcloudflaredvirtualnetwork "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zero/trusttunnelcloudflaredvirtualnetwork"
	cachereserve "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/cachereserve"
	cachevariants "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/cachevariants"
	dnssec "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/dnssec"
	hold "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/hold"
	lockdown "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/lockdown"
	setting "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/setting"
	subscriptionzone "github.com/mindstorms6/provider-cloudflare-upjet/internal/controller/zone/subscription"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		rule.Setup,
		member.Setup,
		subscription.Setup,
		token.Setup,
		shield.Setup,
		shielddiscoveryoperation.Setup,
		shieldoperation.Setup,
		shieldoperationschemavalidationsettings.Setup,
		shieldschema.Setup,
		shieldschemavalidationsettings.Setup,
		tokenapi.Setup,
		smartrouting.Setup,
		tieredcaching.Setup,
		originpulls.Setup,
		originpullscertificate.Setup,
		management.Setup,
		ipprefix.Setup,
		sfuapp.Setup,
		turnapp.Setup,
		pack.Setup,
		connectorrules.Setup,
		account.Setup,
		filter.Setup,
		healthcheck.Setup,
		image.Setup,
		list.Setup,
		queue.Setup,
		ruleset.Setup,
		snippets.Setup,
		stream.Setup,
		user.Setup,
		zone.Setup,
		onerequest.Setup,
		onerequestasset.Setup,
		onerequestmessage.Setup,
		onerequestpriority.Setup,
		scanningexpression.Setup,
		hostname.Setup,
		hostnamefallbackorigin.Setup,
		ssl.Setup,
		database.Setup,
		firewall.Setup,
		record.Setup,
		settings.Setup,
		settingsinternalview.Setup,
		zonetransfersacl.Setup,
		zonetransfersincoming.Setup,
		zonetransfersoutgoing.Setup,
		zonetransferspeer.Setup,
		zonetransferstsig.Setup,
		routingaddress.Setup,
		routingcatchall.Setup,
		routingdns.Setup,
		routingrule.Setup,
		routingsettings.Setup,
		securityblocksender.Setup,
		securityimpersonationregistry.Setup,
		securitytrusteddomains.Setup,
		rulefirewall.Setup,
		tlssetting.Setup,
		config.Setup,
		variant.Setup,
		certificate.Setup,
		credentialcheck.Setup,
		credentialcheckrule.Setup,
		item.Setup,
		balancer.Setup,
		balancermonitor.Setup,
		balancerpool.Setup,
		retention.Setup,
		job.Setup,
		ownershipchallenge.Setup,
		networkmonitoringconfiguration.Setup,
		networkmonitoringrule.Setup,
		transitconnector.Setup,
		transitsite.Setup,
		transitsiteacl.Setup,
		transitsitelan.Setup,
		transitsitewan.Setup,
		wangretunnel.Setup,
		wanipsectunnel.Setup,
		wanstaticroute.Setup,
		transforms.Setup,
		certificatemtls.Setup,
		policy.Setup,
		policywebhooks.Setup,
		scheduledtest.Setup,
		cacertificate.Setup,
		rulepage.Setup,
		shieldpolicy.Setup,
		domain.Setup,
		project.Setup,
		providerconfig.Setup,
		consumer.Setup,
		bucket.Setup,
		bucketcors.Setup,
		bucketeventnotification.Setup,
		bucketlifecycle.Setup,
		bucketlock.Setup,
		bucketsippy.Setup,
		customdomain.Setup,
		manageddomain.Setup,
		limit.Setup,
		hostnameregional.Setup,
		tieredcache.Setup,
		domainregistrar.Setup,
		rules.Setup,
		application.Setup,
		audiotrack.Setup,
		captionlanguage.Setup,
		download.Setup,
		key.Setup,
		liveinput.Setup,
		watermark.Setup,
		webhook.Setup,
		cache.Setup,
		tls.Setup,
		widget.Setup,
		normalizationsettings.Setup,
		agentblockingrule.Setup,
		room.Setup,
		roomevent.Setup,
		roomrules.Setup,
		roomsettings.Setup,
		analyticsrule.Setup,
		analyticssite.Setup,
		hostnameweb3.Setup,
		crontrigger.Setup,
		customdomainworkers.Setup,
		deployment.Setup,
		forplatformsdispatchnamespace.Setup,
		kv.Setup,
		kvnamespace.Setup,
		route.Setup,
		script.Setup,
		scriptsubdomain.Setup,
		secret.Setup,
		trustaccessapplication.Setup,
		trustaccesscustompage.Setup,
		trustaccessgroup.Setup,
		trustaccessidentityprovider.Setup,
		trustaccessinfrastructuretarget.Setup,
		trustaccesskeyconfiguration.Setup,
		trustaccessmtlscertificate.Setup,
		trustaccessmtlshostnamesettings.Setup,
		trustaccesspolicy.Setup,
		trustaccessservicetoken.Setup,
		trustaccessshortlivedcertificate.Setup,
		trustaccesstag.Setup,
		trustdevicecustomprofile.Setup,
		trustdevicecustomprofilelocaldomainfallback.Setup,
		trustdevicedefaultprofile.Setup,
		trustdevicedefaultprofilecertificates.Setup,
		trustdevicedefaultprofilelocaldomainfallback.Setup,
		trustdevicemanagednetworks.Setup,
		trustdevicepostureintegration.Setup,
		trustdeviceposturerule.Setup,
		trustdextest.Setup,
		trustdlpcustomprofile.Setup,
		trustdlpdataset.Setup,
		trustdlpentry.Setup,
		trustdlppredefinedprofile.Setup,
		trustdnslocation.Setup,
		trustgatewaycertificate.Setup,
		trustgatewaylogging.Setup,
		trustgatewaypolicy.Setup,
		trustgatewayproxyendpoint.Setup,
		trustgatewaysettings.Setup,
		trustlist.Setup,
		trustorganization.Setup,
		trustriskbehavior.Setup,
		trustriskscoringintegration.Setup,
		trusttunnelcloudflared.Setup,
		trusttunnelcloudflaredconfig.Setup,
		trusttunnelcloudflaredroute.Setup,
		trusttunnelcloudflaredvirtualnetwork.Setup,
		cachereserve.Setup,
		cachevariants.Setup,
		dnssec.Setup,
		hold.Setup,
		lockdown.Setup,
		setting.Setup,
		subscriptionzone.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
